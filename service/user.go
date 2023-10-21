package service

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"gorm.io/gorm"
	"io/ioutil"
	"log"
	"net/http"
	"xfd-backend/database/db"
	"xfd-backend/database/db/dao"
	"xfd-backend/database/db/model"
	"xfd-backend/pkg/consts"
	"xfd-backend/pkg/jwt"
	"xfd-backend/pkg/types"
	"xfd-backend/pkg/utils"
)

type UserService struct {
	userDao *dao.UserDao
}

func NewUserService() *UserService {
	return &UserService{userDao: dao.NewUserDB()}
}

func (s *UserService) Login(ctx context.Context, req *types.UserLoginReq) (*types.UserLoginResp, error) {
	resp, err := s.getOpenID(ctx, req.Code)
	if err != nil {
		return nil, err
	}

	// 开始事务
	tx := db.Get().Begin()
	user, err := s.loginOrRegister(tx, resp.OpenID, req.UserRole)
	if err != nil {
		tx.Rollback()
		return nil, err
	}

	// 提交事务
	if err = tx.Commit().Error; err != nil {
		return nil, err
	}

	// 使用userID+openID+role生成token
	info := &jwt.SubjectInfo{
		UserID:   user.UserID,
		UserName: user.Username,
		OpenID:   user.OpenID,
		Role:     user.UserRole,
	}
	token, err := jwt.Auth.GenerateToken(ctx, info)
	if err != nil {
		return nil, err
	}

	return &types.UserLoginResp{
		AccessToken: token.AccessToken,
		TokenType:   token.TokenType,
		ExpiresAt:   token.ExpiresAt,
	}, nil
}

func (s *UserService) loginOrRegister(tx *gorm.DB, openID string, userRole model.UserRole) (*model.User, error) {
	var (
		user *model.User
		err  error
	)
	user, err = s.userDao.GetByOpenIDAndRoleInTx(tx, openID, userRole)
	if err != nil {
		return nil, err
	}
	if user == nil {
		user = &model.User{
			UserID:   utils.GenUUID(),
			OpenID:   openID,
			UserRole: userRole,
		}
		if err = s.userDao.CreateInTx(tx, user); err != nil {
			return nil, err
		}
	}
	return user, nil
}

func (s *UserService) getOpenID(ctx context.Context, code string) (*types.Jscode2SessionResponse, error) {
	// 构建请求参数
	// https://developers.weixin.qq.com/miniprogram/dev/OpenApiDoc/user-login/code2Session.html
	url := "https://api.weixin.qq.com/sns/jscode2session"
	appID := consts.APP_ID
	secret := consts.APP_SECRET
	jsCode := code
	grantType := "authorization_code"

	// 构建完整的请求 URL
	fullURL := fmt.Sprintf("%s?appid=%s&secret=%s&js_code=%s&grant_type=%s", url, appID, secret, jsCode, grantType)

	// 发送 GET 请求
	resp, err := http.Get(fullURL)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	// 读取响应的内容
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	var result *types.Jscode2SessionResponse
	err = json.Unmarshal(body, &result)
	if err != nil {
		return nil, err
	}

	if result.ErrorCode != 0 {
		log.Println("[UserService] GetOpenID failed, err=", result.ErrorMsg)
		return nil, err
	}
	return result, nil
}

func (s *UserService) RefreshToken(ctx context.Context, req *types.UserLoginReq) (*types.UserLoginResp, error) {
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

	return &types.UserLoginResp{
		AccessToken: token.AccessToken,
		TokenType:   token.TokenType,
		ExpiresAt:   token.ExpiresAt,
	}, nil
}
