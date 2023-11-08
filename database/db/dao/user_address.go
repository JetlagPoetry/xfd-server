package dao

import (
	"context"
	"gorm.io/gorm"
	"xfd-backend/database/db"
	"xfd-backend/database/db/model"
)

type UserAddressDao struct {
}

func NewUserAddressDao() *UserAddressDao {
	return &UserAddressDao{}
}

func (r *UserAddressDao) Lists(ctx context.Context, limit, offset int) (list []*model.UserAddress, count int64, err error) {
	sql := db.Get().Model(&model.UserAddress{}).Where("")

	if err = sql.Limit(limit).Offset(offset).Find(&list).Error; err != nil {
		return nil, 0, err
	}

	if err = sql.Count(&count).Error; err != nil {
		return nil, 0, err
	}
	return list, count, nil
}

func (r *UserAddressDao) ListByUserID(ctx context.Context, userID string) (list []*model.UserAddress, err error) {
	err = db.Get().Model(&model.UserAddress{}).Where("user_id = ?", userID).Find(&list).Error
	if err != nil {
		return nil, err
	}
	return list, nil
}

func (r *UserAddressDao) GetByID(ctx context.Context, id int) (addr *model.UserAddress, err error) {
	err = db.Get().Model(&model.UserAddress{}).Where("id = ?", id).Find(&addr).Error
	if err != nil {
		return nil, err
	}
	return addr, nil
}

func (r *UserAddressDao) Create(ctx context.Context, addr *model.UserAddress) (err error) {
	err = db.Get().Model(&model.UserAddress{}).Create(addr).Error
	if err != nil {
		return err
	}
	return nil
}

func (r *UserAddressDao) CreateInTx(tx *gorm.DB, addr *model.UserAddress) (err error) {
	err = tx.Model(&model.UserAddress{}).Create(addr).Error
	if err != nil {
		return err
	}
	return nil
}

func (r *UserAddressDao) UpdateByID(ctx context.Context, id int, updateValue *model.UserAddress) (err error) {
	updateResult := db.Get().Model(&model.UserAddress{}).Where("id =?", id).Updates(updateValue)
	if err = updateResult.Error; err != nil {
		return err
	}
	return nil
}

func (r *UserAddressDao) UpdateByIDInTx(tx *gorm.DB, id int, updateValue *model.UserAddress) (err error) {
	updateResult := tx.Model(&model.UserAddress{}).Where("id =?", id).Updates(updateValue)
	if err = updateResult.Error; err != nil {
		return err
	}
	return nil
}

func (r *UserAddressDao) UpdateByUserIDInTx(tx *gorm.DB, userID string, updateValue *model.UserAddress) (err error) {
	updateResult := tx.Model(&model.UserAddress{}).Where("user_id =?", userID).Updates(updateValue)
	if err = updateResult.Error; err != nil {
		return err
	}
	return nil
}

func (r *UserAddressDao) Delete(ctx context.Context, id int) (err error) {
	if err = db.Get().Where("id =? ", id).Delete(&model.UserAddress{}).Error; err != nil {
		return err
	}
	return nil
}
