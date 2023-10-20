package service

import (
	"context"
	"xfd-backend/database/db/dao"
	"xfd-backend/pkg/types"
)

type UserService struct {
	userDao *dao.UserDao
}

func NewUserService() *UserService {
	return &UserService{userDao: dao.NewUserDB()}
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
