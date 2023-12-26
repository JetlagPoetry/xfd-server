package dao

import (
	"context"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
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
	sql := db.GetRepo().GetDB(ctx).Model(&model.PointApplication{})

	if orgID > 0 {
		sql = sql.Where("organization_id = ? ", orgID)
	}
	if err = sql.Count(&count).Error; err != nil {
		return nil, 0, err
	}

	if err = sql.Order("id desc").Offset((page.PageNum - 1) * page.PageSize).Limit(page.PageSize).Find(&list).Error; err != nil {
		return nil, 0, err
	}

	return list, count, nil
}

func (d *PointApplicationDao) GetByID(ctx context.Context, id int) (order *model.PointApplication, err error) {
	err = db.GetRepo().GetDB(ctx).Model(&model.PointApplication{}).Where("id = ?", id).First(&order).Error
	if err != nil {
		return nil, err
	}
	return order, nil
}

func (d *PointApplicationDao) GetByIDForUpdate(tx *gorm.DB, id int) (order *model.PointApplication, err error) {
	err = tx.Model(&model.PointApplication{}).Where("id = ?", id).Clauses(clause.Locking{Strength: "UPDATE"}).First(&order).Error
	if err != nil {
		return nil, err
	}
	return order, nil
}

func (d *PointApplicationDao) ListByStatus(ctx context.Context, status model.PointApplicationStatus) (apply []*model.PointApplication, err error) {
	err = db.GetRepo().GetDB(ctx).Model(&model.PointApplication{}).Where("status = ?", status).Find(&apply).Error
	if err != nil {
		return nil, err
	}
	return apply, nil
}

func (d *PointApplicationDao) ListExpired(ctx context.Context) (apply []*model.PointApplication, err error) {
	err = db.GetRepo().GetDB(ctx).Model(&model.PointApplication{}).Where("end_time <= now() AND status = ? ", model.PointApplicationStatusFinish).Find(&apply).Error
	if err != nil {
		return nil, err
	}
	return apply, nil
}

func (d *PointApplicationDao) ListByOrgIDStatusInTx(tx *gorm.DB, orgID int, status model.PointApplicationStatus) (apply []*model.PointApplication, err error) {
	err = tx.Model(&model.PointApplication{}).Where("organization_id = ? AND status = ?", orgID, status).Find(&apply).Error
	if err != nil {
		return nil, err
	}
	return apply, nil
}

func (d *PointApplicationDao) GetByStatus(ctx context.Context, status model.PointApplicationStatus) (apply *model.PointApplication, err error) {
	err = db.GetRepo().GetDB(ctx).Model(&model.PointApplication{}).Where("status = ?", status).First(&apply).Error
	if err == gorm.ErrRecordNotFound {
		return nil, nil
	} else if err != nil {
		return nil, err
	}
	return apply, nil
}

func (d *PointApplicationDao) CountByStatus(ctx context.Context, status model.PointApplicationStatus) (count int64, err error) {
	err = db.GetRepo().GetDB(ctx).Model(&model.PointApplication{}).Where("status = ?", status).Count(&count).Error
	if err != nil {
		return 0, err
	}
	return count, nil
}

func (d *PointApplicationDao) CreateInTx(tx *gorm.DB, record *model.PointApplication) (err error) {
	err = tx.Model(&model.PointApplication{}).Create(record).Error
	if err != nil {
		return err
	}
	return nil
}

func (d *PointApplicationDao) UpdateByID(ctx context.Context, id int, updateValue *model.PointApplication) (err error) {
	updateResult := db.GetRepo().GetDB(ctx).Model(&model.PointApplication{}).Where("id =?", id).Updates(updateValue)
	if err = updateResult.Error; err != nil {
		return err
	}
	return nil
}

func (d *PointApplicationDao) UpdateByIDInTx(tx *gorm.DB, id int, updateValue *model.PointApplication) (err error) {
	updateResult := tx.Model(&model.PointApplication{}).Where("id =?", id).Updates(updateValue)
	if err = updateResult.Error; err != nil {
		return err
	}
	return nil
}

func (d *PointApplicationDao) UpdateByOrgIDInTx(tx *gorm.DB, orgID int, updateValue *model.PointApplication) (err error) {
	updateResult := tx.Model(&model.PointApplication{}).Where("organization_id =?", orgID).Updates(updateValue)
	if err = updateResult.Error; err != nil {
		return err
	}
	return nil
}
