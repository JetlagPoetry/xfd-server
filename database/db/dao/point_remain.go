package dao

import (
	"context"
	"github.com/shopspring/decimal"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
	"xfd-backend/database/db"
	"xfd-backend/database/db/model"
)

type PointRemainDao struct {
}

func NewPointRemainDao() *PointRemainDao {
	return &PointRemainDao{}
}

func (d *PointRemainDao) ListByUserID(tx *gorm.DB, userID string) (list []*model.PointRemain, err error) {
	err = tx.Model(&model.PointRemain{}).Where("user_id = ? AND point_remain > 0",
		userID).Order("id ASC").Find(&list).Error
	if err != nil {
		return nil, err
	}
	return list, nil
}

func (d *PointRemainDao) ListValidByUserIDCTX(ctx context.Context, userID string) (list []*model.PointRemain, err error) {
	err = db.GetRepo().GetDB(ctx).Model(&model.PointRemain{}).Where("user_id = ? AND point_remain > 0 AND start_time <= now() AND end_time >= now()",
		userID).Order("id ASC").Find(&list).Error
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

func (d *PointRemainDao) GetByIDForUpdate(tx *gorm.DB, id int) (record *model.PointRemain, err error) {
	err = tx.Model(&model.PointRemain{}).Clauses(clause.Locking{Strength: "UPDATE"}).Where("id = ?", id).First(&record).Error
	if err != nil {
		return nil, err
	}
	return record, nil
}

func (d *PointRemainDao) GetByIDForUpdateCTX(ctx context.Context, id int) (record *model.PointRemain, err error) {
	err = db.GetRepo().GetDB(ctx).Model(&model.PointRemain{}).Clauses(clause.Locking{Strength: "UPDATE"}).Where("id = ?", id).First(&record).Error
	if err != nil {
		return nil, err
	}
	return record, nil
}

func (d *PointRemainDao) GetByID(ctx context.Context, id int) (record *model.PointRemain, err error) {
	err = db.GetRepo().GetDB(ctx).Model(&model.PointRemain{}).Where("id = ?", id).First(&record).Error
	if err != nil {
		return nil, err
	}
	return record, nil
}

func (d *PointRemainDao) GetByUserID(ctx context.Context, userID string) (record *model.PointRemain, err error) {
	err = db.GetRepo().GetDB(ctx).Model(&model.PointRemain{}).Where("user_id = ?", userID).First(&record).Error
	if err != nil {
		return nil, err
	}
	return record, nil
}

func (d *PointRemainDao) Create(ctx context.Context, record *model.PointRemain) (err error) {
	err = db.GetRepo().GetDB(ctx).Model(&model.PointRemain{}).Create(record).Error
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
	updateResult := db.GetRepo().GetDB(ctx).Model(&model.PointRemain{}).Where("id =?", id).Updates(updateValue)
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

func (d *PointRemainDao) SumPointRemainByApplyID(ctx context.Context, applyID int) (sum decimal.Decimal, err error) {
	err = db.GetRepo().GetDB(ctx).Table("point_remain").Select("IFNULL(SUM(point_remain),0)").Where("point_application_id = ?", applyID).Row().Scan(&sum)
	if err != nil {
		return decimal.Decimal{}, err
	}
	return sum, nil
}

func (d *PointRemainDao) SumPointRemainByApplyIDInTx(tx *gorm.DB, applyID int) (sum decimal.Decimal, err error) {
	err = tx.Table("point_remain").Select("IFNULL(SUM(point_remain),0)").Where("point_application_id = ?", applyID).Row().Scan(&sum)
	if err != nil {
		return decimal.Decimal{}, err
	}
	return sum, nil
}

func (d *PointRemainDao) UpdateByIDInTxCTX(ctx context.Context, id int, updateValue *model.PointRemain) error {
	updateResult := db.GetRepo().GetDB(ctx).Model(&model.PointRemain{}).Where("id =?", id).Updates(updateValue)
	if err := updateResult.Error; err != nil {
		return err
	}
	return nil
}
