package db

import (
	"context"
	"gorm.io/gorm"
	"xfd-backend/database/model"
)

type UserDao struct {
	db *gorm.DB
}

func NewUserDB() *UserDao {
	return &UserDao{db: mySQL}
}

func (h *UserDao) GetUserInfo(ctx context.Context, userID string) (user *model.User, err error) {
	if err = h.db.Model(&model.User{}).Where("id = ?", userID).Find(&user).Error; err != nil {
		return nil, err
	}
	return user, nil
}
