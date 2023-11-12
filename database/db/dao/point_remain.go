package dao

import (
	"context"
	"gorm.io/gorm"
	"xfd-backend/database/db"
	"xfd-backend/database/db/model"
	"xfd-backend/pkg/types"
)

type PointRemainDao struct {
}

func NewPointRemainDao() *PointRemainDao {
	return &PointRemainDao{}
}

func (d *PointRemainDao) List(ctx context.Context, page types.PageRequest, categoryA, categoryB, categoryC int, like string) (list []*model.PointRemain, count int64, err error) {
	sql := db.Get().Model(&model.PointRemain{})
	if len(like) > 0 {
		sql = sql.Where("category_name LIKE ? AND status = 1", "%"+like+"%")
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

func (d *PointRemainDao) ListByUserID(tx *gorm.DB, userID string) (list []*model.PointRemain, err error) {
	err = tx.Model(&model.PointRemain{}).Where("user_id = ? AND point_remain > 0", userID).Find(&list).Error
	if err != nil {
		return nil, err
	}
	return list, nil
}

func (d *PointRemainDao) ListByAppIDs(tx *gorm.DB, appIDs []int) (list []*model.PointRemain, err error) {
	err = tx.Model(&model.PointRemain{}).Where("point_application_id IN (?) AND point_remain > 0", appIDs).Find(&list).Error
	if err != nil {
		return nil, err
	}
	return list, nil
}

func (d *PointRemainDao) ListByAppIDInTx(tx *gorm.DB, appID int) (list []*model.PointRemain, err error) {
	err = tx.Model(&model.PointRemain{}).Where("point_application_id = ? AND point_remain > 0", appID).Find(&list).Error
	if err != nil {
		return nil, err
	}
	return list, nil
}

func (d *PointRemainDao) GetByID(ctx context.Context, id int) (record *model.PointRemain, err error) {
	err = db.Get().Model(&model.PointRemain{}).Where("id = ?", id).First(&record).Error
	if err != nil {
		return nil, err
	}
	return record, nil
}

func (d *PointRemainDao) GetByUserID(ctx context.Context, userID string) (record *model.PointRemain, err error) {
	err = db.Get().Model(&model.PointRemain{}).Where("user_id = ?", userID).First(&record).Error
	if err != nil {
		return nil, err
	}
	return record, nil
}

func (d *PointRemainDao) Create(ctx context.Context, record *model.PointRemain) (err error) {
	err = db.Get().Model(&model.PointRemain{}).Create(record).Error
	if err != nil {
		return err
	}
	return nil
}

func (d *PointRemainDao) CreateInTx(tx *gorm.DB, record *model.PointRemain) (err error) {
	err = tx.Model(&model.PointRemain{}).Create(record).Error
	if err != nil {
		return err
	}
	return nil
}

func (d *PointRemainDao) UpdateByID(ctx context.Context, id int, updateValue *model.PointRemain) (err error) {
	updateResult := db.Get().Model(&model.PointRemain{}).Where("id =?", id).Updates(updateValue)
	if err = updateResult.Error; err != nil {
		return err
	}
	return nil
}

func (d *PointRemainDao) UpdateByIDInTx(tx *gorm.DB, id int, updateValue *model.PointRemain) (err error) {
	updateResult := tx.Model(&model.PointRemain{}).Where("id =?", id).Updates(updateValue)
	if err = updateResult.Error; err != nil {
		return err
	}
	return nil
}

func (d *PointRemainDao) UpdateByOrgIDInTx(tx *gorm.DB, orgID int, updateValue *model.PointRemain) (err error) {
	updateResult := tx.Model(&model.PointRemain{}).Where("organization_id =?", orgID).Updates(updateValue)
	if err = updateResult.Error; err != nil {
		return err
	}
	return nil
}

func (d *PointRemainDao) UpdateByAppIDInTx(tx *gorm.DB, appID int, updateValue *model.PointRemain) (err error) {
	updateResult := tx.Model(&model.PointRemain{}).Where("point_application_id =?", appID).Updates(updateValue)
	if err = updateResult.Error; err != nil {
		return err
	}
	return nil
}

func (d *PointRemainDao) UpdateByUserIDInTx(tx *gorm.DB, userID string, updateValue *model.PointRemain) (err error) {
	updateResult := tx.Model(&model.PointRemain{}).Where("user_id =?", userID).Updates(updateValue)
	if err = updateResult.Error; err != nil {
		return err
	}
	return nil
}

func (d *PointRemainDao) SumPointRemainByApplyID(ctx context.Context, applyID int) (sum float64, err error) {
	err = db.Get().Table("point_remain").Select("sum(point_remain)").Where("point_application_id = ?", applyID).Row().Scan(&sum)
	if err != nil {
		return 0, err
	}
	return sum, nil
}

func (d *PointRemainDao) SumPointRemainByApplyIDInTx(tx *gorm.DB, applyID int) (sum float64, err error) {
	err = tx.Table("point_remain").Select("sum(point_remain)").Where("point_application_id = ?", applyID).Row().Scan(&sum)
	if err != nil {
		return 0, err
	}
	return sum, nil
}
