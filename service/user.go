package service

import (
	"context"
	"errors"
	"fmt"
	goredis "github.com/go-redis/redis"
	"github.com/tencentyun/tls-sig-api-v2-golang/tencentyun"
	"gorm.io/gorm"
	"os"
	"sort"
	"strconv"
	"time"
	"xfd-backend/database/db"
	"xfd-backend/database/db/dao"
	"xfd-backend/database/db/model"
	"xfd-backend/database/redis"
	"xfd-backend/pkg/common"
	"xfd-backend/pkg/consts"
	"xfd-backend/pkg/im"
	"xfd-backend/pkg/jwt"
	"xfd-backend/pkg/types"
	"xfd-backend/pkg/utils"
	"xfd-backend/pkg/xerr"
)

type UserService struct {
	userDao         *dao.UserDao
	userVerifyDao   *dao.UserVerifyDao
	organizationDao *dao.OrganizationDao
	userAddressDao  *dao.UserAddressDao
}

func NewUserService() *UserService {
	return &UserService{
		userDao:         dao.NewUserDao(),
		userVerifyDao:   dao.NewUserVerifyDao(),
		organizationDao: dao.NewOrganizationDao(),
		userAddressDao:  dao.NewUserAddressDao(),
	}
}

func (s *UserService) SendCode(ctx context.Context, req types.UserSendCodeReq) (*types.UserSendCodeResp, xerr.XErr) {
	code := utils.GenSixDigitCode()

	err := redis.RedisClient.Set(fmt.Sprintf("user-login-code:phone:%s", req.Phone), code, 5*time.Minute).Err()
	if err != nil {
		return nil, xerr.WithCode(xerr.ErrorRedis, err)
	}

	err = utils.SendSms(req.Phone, code, 5)
	if err != nil {
		return nil, xerr.WithCode(xerr.ErrorCallApi, err)
	}
	return &types.UserSendCodeResp{}, nil
}

func (s *UserService) Login(ctx context.Context, req types.UserLoginReq) (*types.UserLoginResp, xerr.XErr) {
	//code, err := redis.RedisClient.Get(fmt.Sprintf("user-login-code:phone:%s", req.Phone)).Result()
	//if err == goredis.Nil {
	//	return nil, xerr.WithCode(xerr.ErrorRedis, errors.New("code invalid"))
	//} else if err != nil {
	//	return nil, xerr.WithCode(xerr.ErrorRedis, err)
	//}
	//
	//if req.Code != code {
	//	log.Println("[UserService] Login failed, code=", code, ", req.Code=", req.Code)
	//	return nil, xerr.WithCode(xerr.ErrorRedis, errors.New("code invalid"))
	//}

	resp := &types.UserLoginResp{}
	var (
		user *model.User
		err  error
	)
	if req.Source == types.SourceMiniApp {
		// 开始事务
		tx := db.Get().Begin()
		user, err = s.loginOrRegister(tx, req.Phone)
		if err != nil {
			tx.Rollback()
			return nil, xerr.WithCode(xerr.ErrorDatabase, err)
		}

		_, err = im.ImportAccount(user)
		if err != nil {
			return nil, xerr.WithCode(xerr.ErrorCallApi, err)
		}

		// 提交事务
		if err = tx.Commit().Error; err != nil {
			return nil, xerr.WithCode(xerr.ErrorDatabase, err)
		}

		resp = &types.UserLoginResp{
			UserRole: user.UserRole,
		}
	} else if req.Source == types.SourceCMS {
		user, err = s.userDao.GetByPhone(ctx, req.Phone)
		if err != nil {
			return nil, xerr.WithCode(xerr.ErrorDatabase, err)
		}
		if user == nil {
			return nil, xerr.WithCode(xerr.ErrorUserNotFound, errors.New("user not found"))
		}
		resp = &types.UserLoginResp{UserRole: user.UserRole}
	}

	// 使用userID+phone+role生成token
	info := &jwt.SubjectInfo{
		UserID: user.UserID,
		Phone:  user.Phone,
	}
	token, err := jwt.Auth.GenerateToken(ctx, info)
	if err != nil {
		return nil, xerr.WithCode(xerr.ErrorAuthToken, err)
	}

	redis.RedisClient.Del(fmt.Sprintf("user-login-code:phone:%s", req.Phone))

	resp.UserID = user.UserID
	resp.AccessToken = token.AccessToken
	resp.TokenType = token.TokenType
	resp.ExpiresAt = token.ExpiresAt
	return resp, nil
}

func (s *UserService) loginOrRegister(tx *gorm.DB, phone string) (*model.User, error) {
	var (
		user *model.User
		err  error
	)
	user, err = s.userDao.GetByPhoneInTx(tx, phone)
	if err != nil {
		return nil, err
	}
	if user == nil {
		user = &model.User{
			UserID:    utils.GenUUID(),
			Phone:     phone,
			Username:  utils.GenUsername(phone),
			AvatarURL: "https://xfd-t-1313159791.cos.ap-beijing.myqcloud.com/resources/common/aagj/WechatIMG1463.jpeg",
		}
		if err = s.userDao.CreateInTx(tx, user); err != nil {
			return nil, err
		}
		//// 注册时，在腾讯云im初始化
		//_, err = im.ImportAccount(user.UserID, phone)
		//if err != nil {
		//	return nil, err
		//}
	}

	return user, nil
}

func (s *UserService) SubmitRole(ctx context.Context, req types.UserSubmitRoleReq) (*types.UserSubmitRoleResp, xerr.XErr) {
	userID := common.GetUserID(ctx)
	user, err := s.userDao.GetByUserID(ctx, userID)
	if err != nil {
		return nil, xerr.WithCode(xerr.ErrorDatabase, err)
	}
	// 允许初次提交&重复提交认证
	if (user.UserRole == model.UserRoleSupplier || user.UserRole == model.UserRoleBuyer) && user.UserRole != req.Role {
		return nil, xerr.WithCode(xerr.ErrorOperationForbidden, errors.New("重复认证"))
	}

	// 开始事务
	tx := db.Get().Begin()
	xErr := s.updateRoleAndVerify(tx, userID, req)
	if xErr != nil {
		tx.Rollback()
		return nil, xerr.WithCode(xerr.ErrorDatabase, xErr)
	}
	// 提交事务
	if err = tx.Commit().Error; err != nil {
		return nil, xerr.WithCode(xerr.ErrorDatabase, err)
	}
	return &types.UserSubmitRoleResp{}, nil
}

func (s *UserService) updateRoleAndVerify(tx *gorm.DB, userID string, req types.UserSubmitRoleReq) xerr.XErr {
	err := s.userDao.UpdateByUserIDInTx(tx, userID, &model.User{UserRole: req.Role})
	if err != nil {
		return xerr.WithCode(xerr.ErrorDatabase, err)
	}
	if req.Role == model.UserRoleBuyer || req.Role == model.UserRoleSupplier {
		verify := &model.UserVerify{
			UserID:           userID,
			UserRole:         req.Role,
			Organization:     req.Organization,
			OrganizationCode: req.OrganizationCode,
			OrganizationURL:  req.OrganizationURL,
			IdentityURLA:     req.IdentityURLA,
			IdentityURLB:     req.IdentityURLB,
			RealName:         req.RealName,
			CertificateNo:    req.CertificateNo,
			Position:         req.Position,
			Phone:            req.Phone,
			Status:           model.UserVerifyStatusSubmitted,
			VerifyTime:       consts.TimeZeroValue,
		}
		err = s.userVerifyDao.CreateInTx(tx, verify)
		if err != nil {
			return xerr.WithCode(xerr.ErrorDatabase, err)
		}
	}

	return nil
}

//func (s *userService) getOpenID(ctx context.Context, code string) (*types.Jscode2SessionResponse, error) {
//	// 构建请求参数
//	// https://developers.weixin.qq.com/miniprogram/dev/OpenApiDoc/user-login/code2Session.html
//	url := "https://api.weixin.qq.com/sns/jscode2session"
//	appID := consts.APP_ID
//	secret := consts.APP_SECRET
//	jsCode := code
//	grantType := "authorization_code"
//
//	// 构建完整的请求 URL
//	fullURL := fmt.Sprintf("%s?appid=%s&secret=%s&js_code=%s&grant_type=%s", url, appID, secret, jsCode, grantType)
//
//	// 发送 GET 请求
//	resp, err := http.Get(fullURL)
//	if err != nil {
//		return nil, err
//	}
//	defer resp.Body.Close()
//
//	// 读取响应的内容
//	body, err := ioutil.ReadAll(resp.Body)
//	if err != nil {
//		return nil, err
//	}
//
//	var result *types.Jscode2SessionResponse
//	err = json.Unmarshal(body, &result)
//	if err != nil {
//		return nil, err
//	}
//
//	if result.ErrorCode != 0 {
//		log.Println("[userService] GetOpenID failed, err=", result.ErrorMsg)
//		return nil, err
//	}
//	return result, nil
//}

func (s *UserService) RefreshToken(ctx context.Context) (*types.UserRefreshTokenResp, xerr.XErr) {
	data := ctx.Value(consts.CONTEXT_HEADER_USER_AUTH_INFO)
	if data == nil {
		return nil, xerr.WithCode(xerr.ErrorUserAuthFailed, errors.New("token invalid"))
	}
	userInfo, ok := data.(*jwt.SubjectInfo)
	if !ok || userInfo == nil {
		return nil, xerr.WithCode(xerr.ErrorUserAuthFailed, errors.New("token invalid"))
	}
	token, err := jwt.Auth.GenerateToken(ctx, userInfo)
	if err != nil {
		return nil, xerr.WithCode(xerr.ErrorUserAuthFailed, errors.New("token invalid"))
	}

	return &types.UserRefreshTokenResp{
		AccessToken: token.AccessToken,
		TokenType:   token.TokenType,
		ExpiresAt:   token.ExpiresAt,
	}, nil
}

func (s *UserService) GetVerifyInfo(ctx context.Context) (*types.GetVerifyInfoResp, xerr.XErr) {
	userID := common.GetUserID(ctx)
	user, err := s.userDao.GetByUserID(ctx, userID)
	if err != nil {
		return nil, xerr.WithCode(xerr.ErrorDatabase, err)
	}

	// 返回用户信息+用户当前绑定的公司
	resp := &types.GetVerifyInfoResp{
		Username: user.Username,
		UserRole: user.UserRole,
	}

	verifyList, err := s.userVerifyDao.ListUserVerifyByUserID(ctx, user.UserID)
	if err != nil {
		return nil, xerr.WithCode(xerr.ErrorDatabase, err)
	}

	// 消费者及其他身份，直接返回
	if user.UserRole != model.UserRoleSupplier && user.UserRole != model.UserRoleBuyer {
		return resp, nil
	}

	// 没有认证信息，直接返回
	if len(verifyList) == 0 {
		return resp, nil
	}

	notify := false
	successHistory := false
	verifyHistory := verifyList[0] // 选取最后一次认证
	if verifyHistory.Status == model.UserVerifyStatusSuccess {
		_, err := redis.RedisClient.Get(fmt.Sprintf("user-verify-notify:verify_id:%d", verifyHistory.ID)).Result()
		if err == goredis.Nil {
			notify = true
			redis.RedisClient.Set(fmt.Sprintf("user-verify-notify:verify_id:%d", verifyHistory.ID), 1, 0)
		}
	}

	for _, verify := range verifyList {
		if verify.Status == model.UserVerifyStatusSuccess {
			successHistory = true
			break
		}
	}

	resp.NotifyVerify = notify
	resp.VerifyHistory = successHistory
	resp.VerifyComment = verifyHistory.Comment
	resp.VerifyStatus = verifyHistory.Status
	resp.Organization = verifyHistory.Organization

	return resp, nil
}

func (s *UserService) GetUserInfo(ctx context.Context) (*types.GetUserInfoResp, xerr.XErr) {
	userID := common.GetUserID(ctx)
	user, err := s.userDao.GetByUserID(ctx, userID)
	if err != nil {
		return nil, xerr.WithCode(xerr.ErrorDatabase, err)
	}

	// 返回用户信息+用户当前绑定的公司
	resp := &types.GetUserInfoResp{
		Username:       user.Username,
		AvatarURL:      user.AvatarURL,
		UserRole:       user.UserRole,
		Organization:   user.OrganizationName,
		OrganizationID: user.OrganizationID,
		Point:          user.Point.Round(2).String(),
	}

	return resp, nil
}

func (s *UserService) ModifyUserInfo(ctx context.Context, req types.UserModifyInfoReq) (*types.UserModifyInfoResp, xerr.XErr) {
	userID := common.GetUserID(ctx)
	err := s.userDao.UpdateByUserID(ctx, userID, &model.User{AvatarURL: req.AvatarURL, Username: req.Username})
	if err != nil {
		return nil, xerr.WithCode(xerr.ErrorDatabase, err)
	}
	return &types.UserModifyInfoResp{}, nil
}

func (s *UserService) AssignAdmin(ctx context.Context, req types.UserAssignAdminReq) (*types.UserAssignAdminResp, xerr.XErr) {
	userID := common.GetUserID(ctx)
	user, err := s.userDao.GetByUserID(ctx, userID)
	if err != nil {
		return nil, xerr.WithCode(xerr.ErrorDatabase, err)
	}
	if user.UserRole != model.UserRoleRoot {
		return nil, xerr.WithCode(xerr.ErrorOperationForbidden, errors.New("user is not root"))
	}

	tx := db.Get().Begin()
	_, xErr := s.assignOrRegisterAdmin(tx, req)
	if xErr != nil {
		tx.Rollback()
		return nil, xerr.WithCode(xerr.ErrorDatabase, xErr)
	}

	// 提交事务
	if err = tx.Commit().Error; err != nil {
		return nil, xerr.WithCode(xerr.ErrorDatabase, err)
	}

	return &types.UserAssignAdminResp{}, nil
}

func (s *UserService) assignOrRegisterAdmin(tx *gorm.DB, req types.UserAssignAdminReq) (*model.User, error) {
	var (
		user *model.User
		err  error
	)
	user, err = s.userDao.GetByPhoneInTx(tx, req.Phone)
	if err != nil {
		return nil, err
	}
	if user != nil && user.UserRole != model.UserRoleUnknown {
		return nil, errors.New("user exists")
	}
	if user == nil {
		user = &model.User{
			UserID:   utils.GenUUID(),
			Phone:    req.Phone,
			Username: req.Username,
		}
	}
	user.UserRole = req.Role
	if err = s.userDao.SaveInTx(tx, user); err != nil {
		return nil, err
	}

	return user, nil
}

func (s *UserService) GetAdmins(ctx context.Context, req types.UserGetAdminsReq) (*types.UserGetAdminsResp, xerr.XErr) {
	userID := common.GetUserID(ctx)
	user, err := s.userDao.GetByUserID(ctx, userID)
	if err != nil {
		return nil, xerr.WithCode(xerr.ErrorDatabase, err)
	}
	if user.UserRole != model.UserRoleRoot {
		return nil, xerr.WithCode(xerr.ErrorOperationForbidden, err)
	}

	userList, count, err := s.userDao.ListByStatus(ctx, req.PageRequest, []model.UserRole{model.UserRoleAdmin, model.UserRoleRoot})
	if err != nil {
		return nil, xerr.WithCode(xerr.ErrorDatabase, err)
	}

	_, rootCount, err := s.userDao.ListByStatus(ctx, req.PageRequest, []model.UserRole{model.UserRoleRoot})
	if err != nil {
		return nil, xerr.WithCode(xerr.ErrorDatabase, err)
	}

	list := make([]*types.UserAdmin, 0)
	for _, user := range userList {
		list = append(list, &types.UserAdmin{
			UserID:   user.UserID,
			Username: user.Username,
			Phone:    user.Phone,
			Role:     user.UserRole,
			RoleName: func() string {
				if user.UserRole == model.UserRoleAdmin {
					return "管理员"
				} else if user.UserRole == model.UserRoleRoot {
					return "超级管理员"
				}
				return ""
			}(),
			Comment: func() string {
				if user.UserRole == model.UserRoleAdmin {
					return "积分审核、身份审核、订单售后"
				} else if user.UserRole == model.UserRoleRoot {
					return "添加新管理员、积分审核、身份审核、订单售后"
				}
				return ""
			}(),
		})
	}
	return &types.UserGetAdminsResp{List: list, TotalNum: int(count), RootNum: int(rootCount)}, nil
}

func (s *UserService) DeleteUser(ctx context.Context, req types.UserDeleteUserReq) (*types.UserDeleteUserResp, xerr.XErr) {
	currentUserID := common.GetUserID(ctx)
	currentUser, err := s.userDao.GetByUserID(ctx, currentUserID)
	if err != nil {
		return nil, xerr.WithCode(xerr.ErrorDatabase, err)
	}
	if currentUser.UserRole != model.UserRoleRoot {
		return nil, xerr.WithCode(xerr.ErrorOperationForbidden, errors.New("user is not root"))
	}

	err = s.userDao.DeleteByUserID(ctx, req.UserID)
	if err != nil {
		return nil, xerr.WithCode(xerr.ErrorDatabase, err)
	}

	return &types.UserDeleteUserResp{}, nil
}

func (s *UserService) ImSig(ctx context.Context, req types.UserImSigReq) (*types.UserImSigResp, xerr.XErr) {
	currentUserID := common.GetUserID(ctx)
	currentUser, err := s.userDao.GetByUserID(ctx, currentUserID)
	if err != nil {
		return nil, xerr.WithCode(xerr.ErrorDatabase, err)
	}

	imAppID, err := strconv.Atoi(os.Getenv("IM_APP_ID"))
	if err != nil {
		return nil, xerr.WithCode(xerr.ErrorOperationForbidden, err)
	}
	imSecret := os.Getenv("IM_APP_SECRET")
	userSig, err := tencentyun.GenUserSig(imAppID, imSecret, currentUser.UserID, 30*86400)
	if err != nil {
		return nil, xerr.WithCode(xerr.ErrorOperationForbidden, err)
	}

	return &types.UserImSigResp{
		UserSig:   userSig,
		ExpiredAt: time.Now().Add(time.Second * 86400).Unix(),
	}, nil
}

func (s *UserService) GetAddressList(ctx context.Context, req types.UserGetAddressListReq) (*types.UserGetAddressListResp, xerr.XErr) {
	userID := common.GetUserID(ctx)

	addrList, err := s.userAddressDao.ListByUserID(ctx, userID)
	if err != nil {
		return nil, xerr.WithCode(xerr.ErrorDatabase, err)
	}
	defaultAddr := 0
	for _, addr := range addrList {
		if *addr.IsDefault == 1 {
			defaultAddr = int(addr.ID)
			break
		}
	}

	list := make([]*types.UserAddress, 0)
	for _, addr := range addrList {
		list = append(list, &types.UserAddress{
			ID:        int(addr.ID),
			Name:      addr.Name,
			Phone:     addr.Phone,
			Province:  addr.Province,
			City:      addr.City,
			Region:    addr.Region,
			Address:   addr.Address,
			IsDefault: int(addr.ID) == defaultAddr,
		})
	}

	sort.Slice(list, func(i, j int) bool {
		if list[i].IsDefault {
			return true
		} else if list[j].IsDefault {
			return false
		}
		return list[i].ID > list[j].ID
	})
	return &types.UserGetAddressListResp{List: list}, nil
}

func (s *UserService) GetDefaultAddress(ctx context.Context, req types.UserGetDefaultAddressReq) (*types.UserGetDefaultAddressResp, xerr.XErr) {
	userID := common.GetUserID(ctx)

	addr, err := s.userAddressDao.GetByUserIDAndDefault(ctx, userID)
	if err != nil {
		return nil, xerr.WithCode(xerr.ErrorDatabase, err)
	}
	if addr == nil {
		return &types.UserGetDefaultAddressResp{NoDefault: true}, nil
	}
	return &types.UserGetDefaultAddressResp{
		UserAddress: types.UserAddress{
			ID:        int(addr.ID),
			Name:      addr.Name,
			Phone:     addr.Phone,
			Province:  addr.Province,
			City:      addr.City,
			Region:    addr.Region,
			Address:   addr.Address,
			IsDefault: true,
		},
		NoDefault: false,
	}, nil
}

func (s *UserService) AddAddress(ctx context.Context, req types.UserAddAddressReq) (*types.UserAddAddressResp, xerr.XErr) {
	userID := common.GetUserID(ctx)

	tx := db.Get().Begin()
	xErr := s.addAddress(tx, userID, req)
	if xErr != nil {
		tx.Rollback()
		return nil, xerr.WithCode(xerr.ErrorDatabase, xErr)
	}

	// 提交事务
	if err := tx.Commit().Error; err != nil {
		return nil, xerr.WithCode(xerr.ErrorDatabase, err)
	}

	return &types.UserAddAddressResp{}, nil
}

func (s *UserService) addAddress(tx *gorm.DB, userID string, req types.UserAddAddressReq) xerr.XErr {
	if req.IsDefault {
		// 取消所有默认地址
		err := s.userAddressDao.UpdateByUserIDInTx(tx, userID, &model.UserAddress{IsDefault: utils.IntPtr(0)})
		if err != nil {
			return xerr.WithCode(xerr.ErrorDatabase, err)
		}
	}

	addr := &model.UserAddress{
		UserID:    userID,
		Name:      req.Name,
		Phone:     req.Phone,
		Province:  req.Province,
		City:      req.City,
		Region:    req.Region,
		Address:   req.Address,
		IsDefault: utils.IntPtr(utils.BoolToInt(req.IsDefault)),
	}
	err := s.userAddressDao.CreateInTx(tx, addr)
	if err != nil {
		return xerr.WithCode(xerr.ErrorDatabase, err)
	}

	return nil
}

func (s *UserService) ModifyAddress(ctx context.Context, req types.UserModifyAddressReq) (*types.UserModifyAddressResp, xerr.XErr) {
	userID := common.GetUserID(ctx)

	addr, err := s.userAddressDao.GetByID(ctx, req.ID)
	if err != nil {
		return nil, xerr.WithCode(xerr.ErrorDatabase, err)
	}
	if addr.UserID != userID {
		return nil, xerr.WithCode(xerr.ErrorDatabase, errors.New("address not belong to user"))
	}

	tx := db.Get().Begin()
	xErr := s.modifyAddress(tx, userID, req)
	if xErr != nil {
		tx.Rollback()
		return nil, xerr.WithCode(xerr.ErrorDatabase, xErr)
	}

	// 提交事务
	if err = tx.Commit().Error; err != nil {
		return nil, xerr.WithCode(xerr.ErrorDatabase, err)
	}

	return &types.UserModifyAddressResp{}, nil
}

func (s *UserService) modifyAddress(tx *gorm.DB, userID string, req types.UserModifyAddressReq) xerr.XErr {
	if req.IsDefault {
		// 取消所有默认地址
		err := s.userAddressDao.UpdateByUserIDInTx(tx, userID, &model.UserAddress{IsDefault: utils.IntPtr(0)})
		if err != nil {
			return xerr.WithCode(xerr.ErrorDatabase, err)
		}
	}

	addr := &model.UserAddress{}

	if len(req.Name) > 0 && len(req.Phone) > 0 && len(req.Province) > 0 && len(req.City) > 0 && len(req.Region) > 0 && len(req.Address) > 0 {
		addr.Name = req.Name
		addr.Phone = req.Phone
		addr.Province = req.Province
		addr.City = req.City
		addr.Region = req.Region
		addr.Address = req.Address
	}
	if req.IsDefault {
		addr.IsDefault = utils.IntPtr(utils.BoolToInt(req.IsDefault))
	}
	err := s.userAddressDao.UpdateByIDInTx(tx, req.ID, addr)
	if err != nil {
		return xerr.WithCode(xerr.ErrorDatabase, err)
	}
	return nil
}

func (s *UserService) DeleteAddress(ctx context.Context, req types.UserDeleteAddressReq) (*types.UserDeleteAddressResp, xerr.XErr) {
	userID := common.GetUserID(ctx)

	addr, err := s.userAddressDao.GetByID(ctx, req.ID)
	if err != nil {
		return nil, xerr.WithCode(xerr.ErrorDatabase, err)
	}
	if addr.UserID != userID {
		return nil, xerr.WithCode(xerr.ErrorDatabase, errors.New("address not belong to user"))
	}

	err = s.userAddressDao.Delete(ctx, req.ID)
	if err != nil {
		return nil, xerr.WithCode(xerr.ErrorDatabase, err)
	}

	return &types.UserDeleteAddressResp{}, nil
}
