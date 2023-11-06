package service

import (
	"context"
	"errors"
	"gorm.io/gorm"
	"xfd-backend/database/db"
	"xfd-backend/database/db/dao"
	"xfd-backend/database/db/model"
	"xfd-backend/pkg/common"
	"xfd-backend/pkg/consts"
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

func (s *UserService) Login(ctx context.Context, req types.UserLoginReq) (*types.UserLoginResp, xerr.XErr) {
	// 开始事务
	tx := db.Get().Begin()
	user, err := s.loginOrRegister(tx, req.Phone)
	if err != nil {
		tx.Rollback()
		return nil, xerr.WithCode(xerr.ErrorDatabase, err)
	}

	// 提交事务
	if err = tx.Commit().Error; err != nil {
		return nil, xerr.WithCode(xerr.ErrorDatabase, err)
	}

	verifyHistory := &model.UserVerify{}
	if user.UserRole == model.UserRoleSupplier || user.UserRole == model.UserRoleBuyer {
		verifyList, err := s.userVerifyDao.ListUserVerifyByUserID(ctx, user.UserID)
		if err != nil {
			return nil, xerr.WithCode(xerr.ErrorDatabase, err)
		}
		if len(verifyList) > 0 {
			verifyHistory = verifyList[0]
		}
	}

	// 使用userID+phone+role生成token
	info := &jwt.SubjectInfo{
		UserID: user.UserID,
		Phone:  user.Phone,
		Role:   user.UserRole,
	}
	token, err := jwt.Auth.GenerateToken(ctx, info)
	if err != nil {
		return nil, xerr.WithCode(xerr.ErrorAuthToken, err)
	}

	// todo 用verify.id 判断是否首次认证成功

	return &types.UserLoginResp{
		AccessToken:   token.AccessToken,
		TokenType:     token.TokenType,
		ExpiresAt:     token.ExpiresAt,
		UserRole:      user.UserRole,
		VerifyStatus:  verifyHistory.Status,
		VerifyComment: verifyHistory.Comment,
		NotifyVerify:  true, // todo
	}, nil
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
			UserRole: model.UserRoleCustomer,
			UserID:   utils.GenUUID(),
			Phone:    phone,
			Username: utils.GenUsername(phone),
		}
		if err = s.userDao.CreateInTx(tx, user); err != nil {
			return nil, err
		}
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
	if user.UserRole != model.UserRoleCustomer && user.UserRole != req.Role {
		return nil, xerr.WithCode(xerr.ErrorOperationForbidden, err)
	}

	// 开始事务
	tx := db.Get().Begin()
	err = s.updateRoleAndVerify(tx, userID, req)
	if err != nil {
		tx.Rollback()
		return nil, xerr.WithCode(xerr.ErrorDatabase, err)
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

func (s *UserService) GetUserInfo(ctx context.Context) (*types.GetUserInfoResp, xerr.XErr) {
	userID := common.GetUserID(ctx)
	user, err := s.userDao.GetByUserID(ctx, userID)
	if err != nil {
		return nil, xerr.WithCode(xerr.ErrorDatabase, err)
	}

	resp := &types.GetUserInfoResp{
		Username:     user.Username,
		AvatarURL:    user.AvatarURL,
		UserRole:     user.UserRole,
		VerifyStatus: types.UserVerifyStatusUnfinished,
		Point:        user.Point,
		NotifyVerify: false, // todo 用verify.id 判断是否首次认证成功
	}

	if user.UserRole == model.UserRoleSupplier || user.UserRole == model.UserRoleBuyer {
		verifyList, err := s.userVerifyDao.ListUserVerifyByUserID(ctx, userID)
		if err != nil {
			return nil, xerr.WithCode(xerr.ErrorDatabase, err)
		}
		if len(verifyList) > 0 {
			verify := verifyList[0]
			organization, err := s.organizationDao.GetByCode(ctx, verify.OrganizationCode)
			if err != nil {
				return nil, xerr.WithCode(xerr.ErrorDatabase, err)
			}

			resp.VerifyStatus = types.UserVerifyStatusDone
			resp.Organization = organization.Name
			resp.OrganizationID = int(organization.ID)
			resp.VerifyComment = verify.Comment
		}
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
		return nil, xerr.WithCode(xerr.ErrorOperationForbidden, err)
	}

	tx := db.Get().Begin()
	_, err = s.assignOrRegisterAdmin(tx, req.Phone)
	if err != nil {
		tx.Rollback()
		return nil, xerr.WithCode(xerr.ErrorDatabase, err)
	}

	// 提交事务
	if err = tx.Commit().Error; err != nil {
		return nil, xerr.WithCode(xerr.ErrorDatabase, err)
	}

	return &types.UserAssignAdminResp{}, nil
}

func (s *UserService) assignOrRegisterAdmin(tx *gorm.DB, phone string) (*model.User, error) {
	var (
		user *model.User
		err  error
	)
	user, err = s.userDao.GetByPhoneInTx(tx, phone)
	if err != nil {
		return nil, err
	}
	if user != nil && user.UserRole != model.UserRoleCustomer {
		return nil, errors.New("user exists")
	}
	user = &model.User{
		UserID:   utils.GenUUID(),
		Phone:    phone,
		UserRole: model.UserRoleAdmin,
		Username: utils.GenUsername(phone),
	}
	if err = s.userDao.CreateInTx(tx, user); err != nil {
		return nil, err
	}

	return user, nil
}

func (s *UserService) GetAddressList(ctx context.Context, req types.UserGetAddressListReq) (*types.UserGetAddressListResp, xerr.XErr) {
	userID := common.GetUserID(ctx)

	user, err := s.userDao.GetByUserID(ctx, userID)
	if err != nil {
		return nil, xerr.WithCode(xerr.ErrorDatabase, err)
	}
	defaultAddrID := user.AddressID

	addrList, err := s.userAddressDao.ListByUserID(ctx, userID)
	if err != nil {
		return nil, xerr.WithCode(xerr.ErrorDatabase, err)
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
			IsDefault: int(addr.ID) == defaultAddrID,
		})
	}
	return &types.UserGetAddressListResp{List: list}, nil
}

func (s *UserService) AddAddress(ctx context.Context, req types.UserAddAddressReq) (*types.UserAddAddressResp, xerr.XErr) {
	userID := common.GetUserID(ctx)

	tx := db.Get().Begin()
	cErr := s.addAddress(tx, userID, req)
	if cErr != nil {
		tx.Rollback()
		return nil, xerr.WithCode(xerr.ErrorDatabase, cErr)
	}

	// 提交事务
	if err := tx.Commit().Error; err != nil {
		return nil, xerr.WithCode(xerr.ErrorDatabase, err)
	}

	return &types.UserAddAddressResp{}, nil
}

func (s *UserService) addAddress(tx *gorm.DB, userID string, req types.UserAddAddressReq) xerr.XErr {
	addr := &model.UserAddress{
		UserID:   userID,
		Name:     req.Name,
		Phone:    req.Phone,
		Province: req.Province,
		City:     req.City,
		Region:   req.Region,
		Address:  req.Address,
	}
	err := s.userAddressDao.CreateInTx(tx, addr)
	if err != nil {
		return xerr.WithCode(xerr.ErrorDatabase, err)
	}

	if req.IsDefault == true {
		err = s.userDao.UpdateByUserIDInTx(tx, userID, &model.User{AddressID: int(addr.ID)})
		if err != nil {
			return xerr.WithCode(xerr.ErrorDatabase, err)
		}
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
	cErr := s.modifyAddress(tx, userID, req)
	if cErr != nil {
		tx.Rollback()
		return nil, xerr.WithCode(xerr.ErrorDatabase, cErr)
	}

	// 提交事务
	if err = tx.Commit().Error; err != nil {
		return nil, xerr.WithCode(xerr.ErrorDatabase, err)
	}

	return &types.UserModifyAddressResp{}, nil
}

func (s *UserService) modifyAddress(tx *gorm.DB, userID string, req types.UserModifyAddressReq) xerr.XErr {
	if len(req.Name) > 0 && len(req.Phone) > 0 && len(req.Province) > 0 && len(req.City) > 0 && len(req.Region) > 0 && len(req.Address) > 0 {
		addr := &model.UserAddress{
			Name:     req.Name,
			Phone:    req.Phone,
			Province: req.Province,
			City:     req.City,
			Region:   req.Region,
			Address:  req.Address,
		}
		err := s.userAddressDao.UpdateByIDInTx(tx, req.ID, addr)
		if err != nil {
			return xerr.WithCode(xerr.ErrorDatabase, err)
		}
	}
	if req.IsDefault == true {
		err := s.userDao.UpdateByUserIDInTx(tx, userID, &model.User{AddressID: req.ID})
		if err != nil {
			return xerr.WithCode(xerr.ErrorDatabase, err)
		}
	}
	return nil
}
