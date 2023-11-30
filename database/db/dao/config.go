package dao

import (
	"context"
	"xfd-backend/database/db"
	"xfd-backend/database/db/model"
)

type ConfigDao struct {
}

func NewConfigDao() *ConfigDao {
	return &ConfigDao{}
}

func (d *ConfigDao) GetByName(ctx context.Context, name string) (cfg *model.Config, err error) {
	err = db.GetRepo().GetDB(ctx).Model(&model.Config{}).Where("name = ?", name).First(&cfg).Error
	if err != nil {
		return nil, err
	}
	return cfg, nil
}
