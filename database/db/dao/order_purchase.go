package dao

import (
	"context"
	"xfd-backend/database/db"
	"xfd-backend/database/db/model"
)

type OrderPurchaseDao struct {
}

func NewOrderPurchaseDao() *OrderPurchaseDao {
	return &OrderPurchaseDao{}
}

func (d *OrderPurchaseDao) Lists(ctx context.Context, limit, offset int) (UserList []*model.OrderPurchase, count int64, err error) {
	if err = db.Get().Model(&model.OrderPurchase{}).Find(&UserList).Limit(limit).Offset(offset).Error; err != nil {
		return nil, 0, err
	}

	if err = db.Get().Model(&model.OrderPurchase{}).Count(&count).Error; err != nil {
		return nil, 0, err
	}
	return UserList, count, nil
}

func (d *OrderPurchaseDao) GetByID(ctx context.Context, id int) (User *model.OrderPurchase, err error) {
	err = db.Get().Model(&model.OrderPurchase{}).Where("id = ?", id).First(&User).Error
	if err != nil {
		return nil, err
	}
	return User, nil
}

func (d *OrderPurchaseDao) GetByUserID(ctx context.Context, userID string) (User *model.OrderPurchase, err error) {
	err = db.Get().Model(&model.OrderPurchase{}).Where("user_id = ?", userID).First(&User).Error
	if err != nil {
		return nil, err
	}
	return User, nil
}

func (d *OrderPurchaseDao) Create(ctx context.Context, User *model.OrderPurchase) (err error) {
	err = db.Get().Model(&model.OrderPurchase{}).Create(User).Error
	if err != nil {
		return err
	}
	return nil
}
