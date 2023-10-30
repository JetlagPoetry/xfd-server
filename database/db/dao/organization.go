package dao

import (
	"context"
	"gorm.io/gorm"
	"xfd-backend/database/db"
	"xfd-backend/database/db/model"
)

type OrganizationDao struct {
}

func NewOrganizationDao() *OrganizationDao {
	return &OrganizationDao{}
}

func (d *OrganizationDao) Lists(ctx context.Context, limit, offset int) (List []*model.Organization, count int64, err error) {
	if err = db.Get().Model(&model.Organization{}).Find(&List).Limit(limit).Offset(offset).Error; err != nil {
		return nil, 0, err
	}

	if err = db.Get().Model(&model.Organization{}).Count(&count).Error; err != nil {
		return nil, 0, err
	}
	return List, count, nil
}

func (d *OrganizationDao) GetByID(ctx context.Context, id int) (org *model.Organization, err error) {
	err = db.Get().Model(&model.Organization{}).Where("id = ?", id).First(&org).Error
	if err != nil {
		return nil, err
	}
	return org, nil
}

func (d *OrganizationDao) GetByCode(ctx context.Context, code string) (org *model.Organization, err error) {
	err = db.Get().Model(&model.Organization{}).Where("code = ?", code).First(&org).Error
	if err != nil {
		return nil, err
	}
	return org, nil
}
func (d *OrganizationDao) GetByCodeInTx(tx *gorm.DB, code string) (org *model.Organization, err error) {
	err = tx.Model(&model.Organization{}).Where("code = ?", code).First(&org).Error
	if err == gorm.ErrRecordNotFound {
		return nil, nil
	} else if err != nil {
		return nil, err
	}
	return org, nil
}

func (d *OrganizationDao) Create(ctx context.Context, org *model.Organization) (err error) {
	err = db.Get().Model(&model.Organization{}).Create(org).Error
	if err != nil {
		return err
	}
	return nil
}

func (d *OrganizationDao) CreateInTx(tx *gorm.DB, org *model.Organization) (err error) {
	err = tx.Model(&model.Organization{}).Create(org).Error
	if err != nil {
		return err
	}
	return nil
}
