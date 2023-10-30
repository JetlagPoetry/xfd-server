package dao

import (
	"context"
	"xfd-backend/database/db"
	"xfd-backend/database/db/model"
	"xfd-backend/pkg/types"
)

type OrderPurchaseDao struct {
}

func NewOrderPurchaseDao() *OrderPurchaseDao {
	return &OrderPurchaseDao{}
}

func (d *OrderPurchaseDao) Lists(ctx context.Context, page types.PageRequest) (list []*model.OrderPurchase, count int64, err error) {
	if err = db.Get().Model(&model.OrderPurchase{}).Find(&list).Offset((page.PageNum - 1) * page.PageSize).Limit(page.PageSize).Error; err != nil {
		return nil, 0, err
	}

	if err = db.Get().Model(&model.OrderPurchase{}).Count(&count).Error; err != nil {
		return nil, 0, err
	}
	return list, count, nil
}

func (d *OrderPurchaseDao) GetByID(ctx context.Context, id int) (purchase *model.OrderPurchase, err error) {
	err = db.Get().Model(&model.OrderPurchase{}).Where("id = ?", id).First(&purchase).Error
	if err != nil {
		return nil, err
	}
	return purchase, nil
}

func (d *OrderPurchaseDao) GetByUserID(ctx context.Context, userID string) (purchase *model.OrderPurchase, err error) {
	err = db.Get().Model(&model.OrderPurchase{}).Where("user_id = ?", userID).First(&purchase).Error
	if err != nil {
		return nil, err
	}
	return purchase, nil
}

func (d *OrderPurchaseDao) Create(ctx context.Context, purchase *model.OrderPurchase) (err error) {
	err = db.Get().Model(&model.OrderPurchase{}).Create(purchase).Error
	if err != nil {
		return err
	}
	return nil
}

func (d *OrderPurchaseDao) UpdateByID(ctx context.Context, id int, updateValue *model.OrderPurchase) (err error) {
	updateResult := db.Get().Model(&model.OrderPurchase{}).Where("id =?", id).Updates(updateValue)
	if err = updateResult.Error; err != nil {
		return err
	}
	return nil
}
