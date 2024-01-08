package dao

import (
	"context"
	"gorm.io/gorm"
	"xfd-backend/database/db"
	"xfd-backend/database/db/model"
)

type PointApplicationTmpDao struct {
}

func NewPointApplicationTmpDao() *PointApplicationTmpDao {
	return &PointApplicationTmpDao{}
}

func (d *PointApplicationTmpDao) GetByID(ctx context.Context, id int) (order *model.PointApplicationTmp, err error) {
	err = db.GetRepo().GetDB(ctx).Model(&model.PointApplicationTmp{}).Where("id = ?", id).First(&order).Error
	if err != nil {
		return nil, err
	}
	return order, nil
}

func (d *PointApplicationTmpDao) ListByStatus(ctx context.Context, status model.PointApplicationTmpStatus) (apply []*model.PointApplicationTmp, err error) {
	err = db.GetRepo().GetDB(ctx).Model(&model.PointApplicationTmp{}).Where("status = ?", status).Limit(2000).Find(&apply).Error
	if err != nil {
		return nil, err
	}
	return apply, nil
}

func (d *PointApplicationTmpDao) GetByStatus(ctx context.Context, status model.PointApplicationTmpStatus) (apply *model.PointApplicationTmp, err error) {
	err = db.GetRepo().GetDB(ctx).Model(&model.PointApplicationTmp{}).Where("status = ?", status).First(&apply).Error
	if err != nil {
		return nil, err
	}
	return apply, nil
}

func (d *PointApplicationTmpDao) CountByStatus(ctx context.Context, status model.PointApplicationTmpStatus) (count int64, err error) {
	err = db.GetRepo().GetDB(ctx).Model(&model.PointApplicationTmp{}).Where("status = ?", status).Count(&count).Error
	if err != nil {
		return 0, err
	}
	return count, nil
}

func (d *PointApplicationTmpDao) BatchCreateInTx(tx *gorm.DB, list []*model.PointApplicationTmp) (err error) {
	err = tx.Model(&model.PointApplicationTmp{}).CreateInBatches(list, 100).Error
	if err != nil {
		return err
	}
	return nil
}

func (d *PointApplicationTmpDao) CreateInTx(tx *gorm.DB, record *model.PointApplicationTmp) (err error) {
	err = tx.Model(&model.PointApplicationTmp{}).Create(record).Error
	if err != nil {
		return err
	}
	return nil
}

func (d *PointApplicationTmpDao) UpdateByID(ctx context.Context, id int, updateValue *model.PointApplicationTmp) (err error) {
	updateResult := db.GetRepo().GetDB(ctx).Model(&model.PointApplicationTmp{}).Where("id =?", id).Updates(updateValue)
	if err = updateResult.Error; err != nil {
		return err
	}
	return nil
}

func (d *PointApplicationTmpDao) UpdateByIDInTx(tx *gorm.DB, id int, updateValue *model.PointApplicationTmp) (err error) {
	updateResult := tx.Model(&model.PointApplicationTmp{}).Where("id =?", id).Updates(updateValue)
	if err = updateResult.Error; err != nil {
		return err
	}
	return nil
}

func (d *PointApplicationTmpDao) UpdateByAppIDInTx(tx *gorm.DB, applyID int, updateValue *model.PointApplicationTmp) (err error) {
	updateResult := tx.Model(&model.PointApplicationTmp{}).Where("application_id =?", applyID).Updates(updateValue)
	if err = updateResult.Error; err != nil {
		return err
	}
	return nil
}
