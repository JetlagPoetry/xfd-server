package dao

import (
	"context"
	"xfd-backend/database/db"
	"xfd-backend/database/db/model"
	"xfd-backend/pkg/types"
)

type PointApplicationDao struct {
}

func NewPointApplicationDao() *PointApplicationDao {
	return &PointApplicationDao{}
}

func (d *PointApplicationDao) Lists(ctx context.Context, page types.PageRequest, orgID int) (list []*model.PointApplication, count int64, err error) {
	sql := db.Get().Model(&model.PointApplication{})

	if orgID > 0 {
		sql = sql.Where("organization_id = ? ", orgID)
	}
	if err = sql.Offset((page.PageNum - 1) * page.PageSize).Limit(page.PageSize).Find(&list).Error; err != nil {
		return nil, 0, err
	}

	if err = sql.Count(&count).Error; err != nil {
		return nil, 0, err
	}
	return list, count, nil
}

func (d *PointApplicationDao) GetByID(ctx context.Context, id int) (order *model.PointApplication, err error) {
	err = db.Get().Model(&model.PointApplication{}).Where("id = ?", id).First(&order).Error
	if err != nil {
		return nil, err
	}
	return order, nil
}

func (d *PointApplicationDao) GetByStatus(ctx context.Context, status model.PointApplicationStatus) (apply *model.PointApplication, err error) {
	err = db.Get().Model(&model.PointApplication{}).Where("id = ?", status).First(&apply).Error
	if err != nil {
		return nil, err
	}
	return apply, nil
}

func (d *PointApplicationDao) CountByStatus(ctx context.Context, status model.PointApplicationStatus) (count int64, err error) {
	err = db.Get().Model(&model.PointApplication{}).Where("status = ?", status).Count(&count).Error
	if err != nil {
		return 0, err
	}
	return count, nil
}
