package dao

import (
	"context"
	"xfd-backend/database/db"
	"xfd-backend/database/db/model"
	"xfd-backend/pkg/types"
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

func (d *OrderQuoteDao) ListByUserIDAndOrderID(ctx context.Context, userID string, orderID int) (list []*model.OrderQuote, err error) {
	err = db.Get().Model(&model.OrderQuote{}).Where("quote_user_id = ? AND purchase_order_id = ?", userID, orderID).Find(&list).Error
	if err != nil {
		return nil, err
	}
	return list, nil
}

func (d *OrderQuoteDao) ListByQuoteUserIDAndOrderIDs(ctx context.Context, userID string, orderIDs []int) (list []*model.OrderQuote, err error) {
	err = db.Get().Model(&model.OrderQuote{}).Where("quote_user_id = ? AND purchase_order_id IN (?)", userID, orderIDs).Find(&list).Error
	if err != nil {
		return nil, err
	}
	return list, nil
}

func (d *OrderQuoteDao) ListByOrderID(ctx context.Context, orderID int, page types.PageRequest) (list []*model.OrderQuote, count int64, err error) {
	sql := db.Get().Model(&model.OrderQuote{}).Where("purchase_order_id = ?", orderID)
	err = sql.Offset((page.PageNum - 1) * page.PageSize).Limit(page.PageSize).Find(&list).Error
	if err != nil {
		return nil, 0, err
	}
	if err = sql.Count(&count).Error; err != nil {
		return nil, 0, err
	}
	return list, 0, nil
}

func (d *OrderQuoteDao) ListByQuoteUserID(ctx context.Context, userID string) (list []*model.OrderQuote, err error) {
	err = db.Get().Model(&model.OrderQuote{}).Where("quote_user_id = ?", userID).Find(&list).Error
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

func (d *OrderQuoteDao) CountByOrderID(ctx context.Context, orderID int) (count int64, err error) {
	err = db.Get().Model(&model.OrderQuote{}).Where("purchase_order_id = ?", orderID).Count(&count).Error
	if err != nil {
		return 0, err
	}
	return count, nil
}

func (d *OrderQuoteDao) CountByPurchaseUserIDAndNotifyPurchase(ctx context.Context, userID string, notifyPurchase bool) (count int64, err error) {
	err = db.Get().Model(&model.OrderQuote{}).Where("purchase_user_id = ? AND notify_purchase = ?", userID, notifyPurchase).Count(&count).Error
	if err != nil {
		return 0, err
	}
	return count, nil
}

func (d *OrderQuoteDao) CountByOrderIDAndNotifyPurchase(ctx context.Context, orderID int, notifyPurchase bool) (count int64, err error) {
	err = db.Get().Model(&model.OrderQuote{}).Where("purchase_order_id = ? AND notify_purchase = ?", orderID, notifyPurchase).Count(&count).Error
	if err != nil {
		return 0, err
	}
	return count, nil
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

func (d *OrderQuoteDao) UpdateByOrderID(ctx context.Context, orderID int, updateValue *model.OrderQuote) (err error) {
	updateResult := db.Get().Model(&model.OrderQuote{}).Where("purchase_order_id =?", orderID).Updates(updateValue)
	if err = updateResult.Error; err != nil {
		return err
	}
	return nil
}

func (d *OrderQuoteDao) UpdateByPurchaseUserAndQuoteUser(ctx context.Context, purchaseUser string, quoteUser string, updateValue *model.OrderQuote) (err error) {
	updateResult := db.Get().Model(&model.OrderQuote{}).Where("purchase_user_id = ? AND quote_user_id = ?", purchaseUser, quoteUser).Updates(updateValue)
	if err = updateResult.Error; err != nil {
		return err
	}
	return nil
}
