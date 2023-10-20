package service

import (
	"context"
	"xfd-backend/database/db"
	"xfd-backend/types"
)

type UserService struct {
	userDao *db.UserDao
}

func NewUserService() *UserService {
	return &UserService{userDao: db.NewUserDB()}
}

func (h *UserService) GetUserInfo(ctx context.Context, userID string) (*types.GetUserInfoResp, error) {
	userInfo, err := h.userDao.GetUserInfo(ctx, userID)
	if err != nil {
		return nil, err
	}
	return &types.GetUserInfoResp{
		UserName:       userInfo.UserName,
		Email:          userInfo.Email,
		PasswordDigest: userInfo.PasswordDigest,
		NickName:       userInfo.NickName,
		Status:         userInfo.Status,
		Avatar:         userInfo.Avatar,
		Money:          userInfo.Money,
	}, nil
}
