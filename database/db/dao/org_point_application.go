package dao

import (
	"context"
	"xfd-backend/database/db"
	"xfd-backend/database/db/model"
	"xfd-backend/pkg/types"
)

type OrgPointApplicationDao struct {
}

func NewOrgPointApplicationDao() *OrgPointApplicationDao {
	return &OrgPointApplicationDao{}
}

func (d *OrgPointApplicationDao) Lists(ctx context.Context, page types.PageRequest, orgID int) (list []*model.OrgPointApplication, count int64, err error) {
	sql := db.Get().Model(&model.OrgPointApplication{})

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

func (d *OrgPointApplicationDao) GetByID(ctx context.Context, id int) (order *model.OrgPointApplication, err error) {
	err = db.Get().Model(&model.OrgPointApplication{}).Where("id = ?", id).First(&order).Error
	if err != nil {
		return nil, err
	}
	return order, nil
}

func (d *OrgPointApplicationDao) GetByStatus(ctx context.Context, status model.OrgPointApplicationStatus) (apply *model.OrgPointApplication, err error) {
	err = db.Get().Model(&model.OrgPointApplication{}).Where("id = ?", status).First(&apply).Error
	if err != nil {
		return nil, err
	}
	return apply, nil
}

func (d *OrgPointApplicationDao) CountByStatus(ctx context.Context, status model.OrgPointApplicationStatus) (count int64, err error) {
	err = db.Get().Model(&model.OrgPointApplication{}).Where("status = ?", status).Count(&count).Error
	if err != nil {
		return 0, err
	}
	return count, nil
}
