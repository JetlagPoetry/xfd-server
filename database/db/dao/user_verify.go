package dao

import (
	"context"
	"gorm.io/gorm"
	"xfd-backend/database/db"
	"xfd-backend/database/db/model"
)

type UserVerifyDao struct {
}

func NewUserVerifyDao() *UserVerifyDao {
	return &UserVerifyDao{}
}

func (d *UserVerifyDao) ListUserVerifyByUserID(ctx context.Context, userID string) (userVerifyList []*model.UserVerify, err error) {
	if err = db.Get().Model(&model.UserVerify{}).Where("user_id = ? AND status = ?", userID, model.UserVerifyStatusSuccess).
		Order("created_at desc").Find(&userVerifyList).Error; err != nil {
		return nil, err
	}
	return userVerifyList, nil
}

func (d *UserVerifyDao) Create(ctx context.Context, User *model.UserVerify) (err error) {
	err = db.Get().Model(&model.UserVerify{}).Create(User).Error
	if err != nil {
		return err
	}
	return nil
}

func (d *UserVerifyDao) CreateInTx(tx *gorm.DB, User *model.UserVerify) (err error) {
	err = tx.Model(&model.UserVerify{}).Create(User).Error
	if err != nil {
		return err
	}
	return nil
}
