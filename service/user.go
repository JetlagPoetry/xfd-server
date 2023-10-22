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
)

type UserService struct {
	userDao       *dao.UserDao
	userVerifyDao *dao.UserVerifyDao
}

func NewUserService() *UserService {
	return &UserService{
		userDao:       dao.NewUserDao(),
		userVerifyDao: dao.NewUserVerifyDao(),
	}
}

func (s *UserService) SendCode(ctx context.Context, req *types.UserSendCodeReq) (*types.UserSendCodeResp, error) {
	// todo 发验证码

	return nil, nil
}

func (s *UserService) Login(ctx context.Context, req *types.UserLoginReq) (*types.UserLoginResp, error) {
	// todo 校验验证码

	// 开始事务
	tx := db.Get().Begin()
	user, err := s.loginOrRegister(tx, req.Phone)
	if err != nil {
		tx.Rollback()
		return nil, err
	}

	// 提交事务
	if err = tx.Commit().Error; err != nil {
		return nil, err
	}

	verifyHistory := &model.UserVerify{}
	if user.UserRole == model.UserRoleSupplier || user.UserRole == model.UserRoleBuyer {
		verifyList, err := s.userVerifyDao.ListUserVerifyByUserID(ctx, user.UserID)
		if err != nil {
			return nil, err
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
	// todo token过期错误码
	token, err := jwt.Auth.GenerateToken(ctx, info)
	if err != nil {
		return nil, err
	}

	return &types.UserLoginResp{
		AccessToken:   token.AccessToken,
		TokenType:     token.TokenType,
		ExpiresAt:     token.ExpiresAt,
		UserRole:      user.UserRole,
		VerifyStatus:  verifyHistory.Status,
		VerifyComment: verifyHistory.Comment,
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
			UserID: utils.GenUUID(),
			Phone:  phone,
		}
		if err = s.userDao.CreateInTx(tx, user); err != nil {
			return nil, err
		}
	}
	return user, nil
}

func (s *UserService) SubmitRole(ctx context.Context, req *types.UserSubmitRoleReq) (*types.UserSubmitRoleResp, error) {
	userID := common.GetUserID(ctx)
	user, err := s.userDao.GetByUserID(ctx, userID)
	if err != nil {
		return nil, err
	}
	// 允许初次提交&重复提交认证
	if user.UserRole != model.UserRoleUnknown && user.UserRole != req.Role {
		return nil, errors.New("request forbidden")
	}

	// 开始事务
	tx := db.Get().Begin()
	err = s.updateRoleAndVerify(tx, userID, req)
	if err != nil {
		tx.Rollback()
		return nil, err
	}
	// 提交事务
	if err = tx.Commit().Error; err != nil {
		return nil, err
	}
	return nil, nil
}

func (s *UserService) updateRoleAndVerify(tx *gorm.DB, userID string, req *types.UserSubmitRoleReq) error {
	err := s.userDao.UpdateByUserIDInTx(tx, userID, &model.User{UserRole: req.Role})
	if err != nil {
		return err
	}
	if req.Role == model.UserRoleBuyer || req.Role == model.UserRoleSupplier {
		verify := &model.UserVerify{
			UserID:           userID,
			Organization:     req.Organization,
			OrganizationCode: req.OrganizationCode,
			OrganizationURL:  req.OrganizationURL,
			CorporationURLA:  req.CorporationURLA,
			CorporationURLB:  req.CorporationURLB,
			RealName:         req.RealName,
			CertificateNo:    req.CertificateNo,
			Position:         req.Position,
			Phone:            req.Phone,
		}
		err = s.userVerifyDao.CreateInTx(tx, verify)
		if err != nil {
			return err
		}
	}

	return nil
}

//func (s *UserService) getOpenID(ctx context.Context, code string) (*types.Jscode2SessionResponse, error) {
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
//		log.Println("[UserService] GetOpenID failed, err=", result.ErrorMsg)
//		return nil, err
//	}
//	return result, nil
//}

func (s *UserService) RefreshToken(ctx context.Context) (*types.UserRefreshTokenResp, error) {
	data := ctx.Value(consts.CONTEXT_HEADER_USER_AUTH_INFO)
	if data == nil {
		return nil, errors.New("token invalid")
	}
	userInfo, ok := data.(*jwt.SubjectInfo)
	if !ok || userInfo == nil {
		return nil, errors.New("token invalid")
	}
	token, err := jwt.Auth.GenerateToken(ctx, userInfo)
	if err != nil {
		return nil, err
	}

	return &types.UserRefreshTokenResp{
		AccessToken: token.AccessToken,
		TokenType:   token.TokenType,
		ExpiresAt:   token.ExpiresAt,
	}, nil
}

func (s *UserService) GetUserInfo(ctx context.Context) (*types.GetUserInfoResp, error) {
	userID := common.GetUserID(ctx)
	user, err := s.userDao.GetByUserID(ctx, userID)
	if err != nil {
		return nil, err
	}

	resp := &types.GetUserInfoResp{
		Username:     user.Username,
		AvatarURL:    user.AvatarURL,
		UserRole:     user.UserRole,
		VerifyStatus: types.UserVerifyStatusUnfinished,
		//Point:        user.po, // todo jingyuan
	}

	if user.UserRole == model.UserRoleSupplier || user.UserRole == model.UserRoleBuyer {
		verifyList, err := s.userVerifyDao.ListUserVerifyByUserID(ctx, userID)
		if err != nil {
			return nil, err
		}
		if len(verifyList) > 0 {
			resp.VerifyStatus = types.UserVerifyStatusDone
			resp.Organization = verifyList[0].Organization
		}
	}

	return resp, nil
}

func (s *UserService) ModifyUserInfo(ctx context.Context, req *types.UserModifyInfoReq) (*types.UserModifyInfoResp, error) {
	userID := common.GetUserID(ctx)
	err := s.userDao.UpdateByUserID(ctx, userID, &model.User{AvatarURL: req.AvatarURL, Username: req.Username})
	if err != nil {
		return nil, err
	}
	return nil, nil
}
