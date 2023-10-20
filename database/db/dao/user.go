package dao

import (
	"context"
	"xfd-backend/database/db"
	"xfd-backend/database/db/model"
)

type UserDao struct {
}

func NewUserDB() *UserDao {
	return &UserDao{}
}

func (h *UserDao) GetUserInfo(ctx context.Context, userID string) (user *model.User, err error) {
	if err = db.Get().Model(&model.User{}).Where("id = ?", userID).Find(&user).Error; err != nil {
		return nil, err
	}
	return user, nil
}
