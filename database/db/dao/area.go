package dao

import (
	"context"
	"xfd-backend/database/db"
	"xfd-backend/database/db/model"
)

type AreaDao struct {
}

func NewAreaDao() *AreaDao {
	return &AreaDao{}
}

func (d *AreaDao) GetAreaInfo(ctx context.Context, code int) (result []model.AreaCode, err error) {
	if err = db.Get().Model(&model.AreaCode{}).
		Where("pcode = ? and status = 1", code).
		Find(&result).Error; err != nil {
		return nil, err
	}
	return result, nil
}
