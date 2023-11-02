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

func (d *OrderPurchaseDao) List(ctx context.Context, page types.PageRequest, categoryA, categoryB, categoryC int, like string) (list []*model.OrderPurchase, count int64, err error) {
	sql := db.Get().Model(&model.OrderPurchase{})
	if len(like) > 0 {
		sql = sql.Where("category_name LIKE '%?%' AND status = 1", like) // todo debug
	} else {
		sql = sql.Where("category_a = ? AND category_b = ? AND status = 1", categoryA, categoryB)

		if categoryC > 0 {
			sql = sql.Where("category_c = ?", categoryC)
		}
	}
	if err = sql.Order("created_at desc").Offset((page.PageNum - 1) * page.PageSize).Limit(page.PageSize).Find(&list).Error; err != nil {
		return nil, 0, err
	}

	if err = sql.Count(&count).Error; err != nil {
		return nil, 0, err
	}
	return list, count, nil
}

func (d *OrderPurchaseDao) ListByOrderIDs(ctx context.Context, page types.PageRequest, orderIDs []int, status model.OrderPurchaseStatus) (list []*model.OrderPurchase, count int64, err error) {
	sql := db.Get().Model(&model.OrderPurchase{}).Where("id IN (?)", orderIDs)
	if status != 0 {
		sql = sql.Where("status = ?", status)
	}
	if err = sql.Order("created_at desc").Offset((page.PageNum - 1) * page.PageSize).Limit(page.PageSize).Find(&list).Error; err != nil {
		return nil, 0, err
	}

	if err = sql.Count(&count).Error; err != nil {
		return nil, 0, err
	}
	return list, count, nil
}

func (d *OrderPurchaseDao) ListByUser(ctx context.Context, page types.PageRequest, userID string, status model.OrderPurchaseStatus) (list []*model.OrderPurchase, count int64, err error) {
	sql := db.Get().Model(&model.OrderPurchase{}).Where("userID = ?", userID)
	if status != 0 {
		sql = sql.Where("status = ?", status)
	}
	if err = sql.Order("created_at desc").Offset((page.PageNum - 1) * page.PageSize).Limit(page.PageSize).Find(&list).Error; err != nil {
		return nil, 0, err
	}

	if err = sql.Count(&count).Error; err != nil {
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
