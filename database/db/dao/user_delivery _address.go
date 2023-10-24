package dao

import (
	"context"
	"xfd-backend/database/db"
	"xfd-backend/database/db/model"
)

type DeliveryAddressDao struct {
}

func NewDeliveryAddressDao() *DeliveryAddressDao {
	return &DeliveryAddressDao{}
}

// Create 在数据库中插入一条.
func (r *DeliveryAddressDao) Create(ctx context.Context, requestDO *model.UserDeliveryAddress) (deliveryID int, err error) {
	err = db.Get().Model(&model.UserDeliveryAddress{}).Create(requestDO).Error
	if err != nil {
		return 0, err
	}
	return int(requestDO.ID), err
}

// CheckExists 检查是否存在.
func (r *DeliveryAddressDao) CheckExists(ctx context.Context, userId string, addressType int) (exists bool, err error) {
	var count int64
	if err = db.Get().Model(&model.UserDeliveryAddress{}).
		Where("user_id = ? and address_type = ?", userId, addressType).Count(&count).Error; err != nil {
		return false, err
	}
	return count >= 1, nil
}
