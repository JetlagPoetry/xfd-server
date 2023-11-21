package dao

import (
	"context"
	"github.com/shopspring/decimal"
	"gorm.io/gorm"
	"xfd-backend/database/db"
	"xfd-backend/database/db/model"
	"xfd-backend/pkg/types"
)

type PointRecordDao struct {
}

func NewPointRecordDao() *PointRecordDao {
	return &PointRecordDao{}
}

func (d *PointRecordDao) List(ctx context.Context, page types.PageRequest, categoryA, categoryB, categoryC int, like string) (list []*model.PointRecord, count int64, err error) {
	sql := db.Get().Model(&model.PointRecord{})
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

func (d *PointRecordDao) ListByOrderIDs(ctx context.Context, page types.PageRequest, orderIDs []int, status model.PointRecordStatus) (list []*model.PointRecord, count int64, err error) {
	sql := db.Get().Model(&model.PointRecord{}).Where("id IN (?)", orderIDs)
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

func (d *PointRecordDao) ListByApplyID(ctx context.Context, page types.PageRequest, applyID int) (list []*model.PointRecord, count int64, err error) {
	sql := db.Get().Model(&model.PointRecord{}).Where("point_application_id = ? AND status = ?", applyID, model.PointRecordStatusConfirmed)
	if err = sql.Order("created_at desc").Offset((page.PageNum - 1) * page.PageSize).Limit(page.PageSize).Find(&list).Error; err != nil {
		return nil, 0, err
	}

	if err = sql.Count(&count).Error; err != nil {
		return nil, 0, err
	}
	return list, count, nil
}

func (d *PointRecordDao) ListByUserID(ctx context.Context, page types.PageRequest, userID string) (list []*model.PointRecord, count int64, err error) {
	sql := db.Get().Model(&model.PointRecord{}).Where("user_id = ? AND status = ?", userID, model.PointRecordStatusConfirmed)
	if err = sql.Order("created_at desc").Offset((page.PageNum - 1) * page.PageSize).Limit(page.PageSize).Find(&list).Error; err != nil {
		return nil, 0, err
	}

	if err = sql.Count(&count).Error; err != nil {
		return nil, 0, err
	}
	return list, count, nil
}

func (d *PointRecordDao) ListByOrgID(ctx context.Context, page types.PageRequest, orgID int) (list []*model.PointRecord, count int64, err error) {
	sql := db.Get().Model(&model.PointRecord{}).Where("organization_id = ? AND status = ?", orgID, model.PointRecordStatusConfirmed)
	if err = sql.Order("created_at desc").Offset((page.PageNum - 1) * page.PageSize).Limit(page.PageSize).Find(&list).Error; err != nil {
		return nil, 0, err
	}

	if err = sql.Count(&count).Error; err != nil {
		return nil, 0, err
	}
	return list, count, nil
}

func (d *PointRecordDao) ListByOrderIDInTx(tx *gorm.DB, userID string, orderID int) (list []*model.PointRecord, err error) {
	err = tx.Model(&model.PointRecord{}).Where("user_id = ? AND order_id = ? AND status = ?", userID,
		orderID, model.PointRecordStatusConfirmed).Find(&list).Error
	if err != nil {
		return nil, err
	}
	return list, nil
}

func (d *PointRecordDao) ListByOrderIDInTxCTX(ctx context.Context, userID string, orderID int) (list []*model.PointRecord, err error) {
	err = db.GetRepo().GetDB(ctx).Model(&model.PointRecord{}).Where("user_id = ? AND order_id = ? AND status = ?", userID,
		orderID, model.PointRecordStatusConfirmed).Find(&list).Error
	if err != nil {
		return nil, err
	}
	return list, nil
}
func (d *PointRecordDao) GetByID(ctx context.Context, id int) (record *model.PointRecord, err error) {
	err = db.Get().Model(&model.PointRecord{}).Where("id = ?", id).First(&record).Error
	if err != nil {
		return nil, err
	}
	return record, nil
}

func (d *PointRecordDao) GetByUserID(ctx context.Context, userID string) (record *model.PointRecord, err error) {
	err = db.Get().Model(&model.PointRecord{}).Where("user_id = ?", userID).First(&record).Error
	if err != nil {
		return nil, err
	}
	return record, nil
}

func (d *PointRecordDao) Create(ctx context.Context, record *model.PointRecord) (err error) {
	err = db.Get().Model(&model.PointRecord{}).Create(record).Error
	if err != nil {
		return err
	}
	return nil
}

func (d *PointRecordDao) CreateInTx(tx *gorm.DB, record *model.PointRecord) (err error) {
	err = tx.Model(&model.PointRecord{}).Create(record).Error
	if err != nil {
		return err
	}
	return nil
}

func (d *PointRecordDao) BatchCreateInTx(tx *gorm.DB, list []*model.PointRecord) (err error) {
	err = tx.Model(&model.PointRecord{}).CreateInBatches(list, 100).Error
	if err != nil {
		return err
	}
	return nil
}

func (d *PointRecordDao) UpdateByID(ctx context.Context, id int, updateValue *model.PointRecord) (err error) {
	updateResult := db.Get().Model(&model.PointRecord{}).Where("id =?", id).Updates(updateValue)
	if err = updateResult.Error; err != nil {
		return err
	}
	return nil
}

func (d *PointRecordDao) SumByAppIDInTx(ctx context.Context, applyID int, ty model.PointRecordType) (sum decimal.Decimal, err error) {
	err = db.Get().Table("point_record").Select("IFNULL(SUM(change_point),0)").Where("point_application_id = ? AND type = ?", applyID, ty).Row().Scan(&sum)
	if err != nil {
		return decimal.Decimal{}, err
	}
	return sum, nil
}

func (d *PointRecordDao) CreateInTxCTX(ctx context.Context, record *model.PointRecord) error {
	err := db.GetRepo().GetDB(ctx).Model(&model.PointRecord{}).Create(record).Error
	if err != nil {
		return err
	}
	return nil
}
