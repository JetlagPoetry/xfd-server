package dao

import (
	"context"
	"xfd-backend/database/db"
	"xfd-backend/database/db/model"
)

type OrderQuoteDao struct {
}

func NewOrderQuoteDao() *OrderQuoteDao {
	return &OrderQuoteDao{}
}

func (d *OrderQuoteDao) Lists(ctx context.Context, limit, offset int) (list []*model.OrderQuote, count int64, err error) {
	if err = db.Get().Model(&model.OrderQuote{}).Find(&list).Limit(limit).Offset(offset).Error; err != nil {
		return nil, 0, err
	}

	if err = db.Get().Model(&model.OrderQuote{}).Count(&count).Error; err != nil {
		return nil, 0, err
	}
	return list, count, nil
}

func (d *OrderQuoteDao) ListByUserID(ctx context.Context, userID string) (list []*model.OrderQuote, err error) {
	err = db.Get().Model(&model.OrderQuote{}).Where("user_id = ?", userID).Find(&list).Error
	if err != nil {
		return nil, err
	}
	return list, nil
}

func (d *OrderQuoteDao) GetByID(ctx context.Context, id int) (order *model.OrderQuote, err error) {
	err = db.Get().Model(&model.OrderQuote{}).Where("id = ?", id).First(&order).Error
	if err != nil {
		return nil, err
	}
	return order, nil
}

func (d *OrderQuoteDao) GetByUserID(ctx context.Context, userID string) (order *model.OrderQuote, err error) {
	err = db.Get().Model(&model.OrderQuote{}).Where("user_id = ?", userID).First(&order).Error
	if err != nil {
		return nil, err
	}
	return order, nil
}

func (d *OrderQuoteDao) Create(ctx context.Context, order *model.OrderQuote) (err error) {
	err = db.Get().Model(&model.OrderQuote{}).Create(order).Error
	if err != nil {
		return err
	}
	return nil
}

func (d *OrderQuoteDao) UpdateByID(ctx context.Context, id int, updateValue *model.OrderQuote) (err error) {
	updateResult := db.Get().Model(&model.OrderQuote{}).Where("id =?", id).Updates(updateValue)
	if err = updateResult.Error; err != nil {
		return err
	}
	return nil
}
