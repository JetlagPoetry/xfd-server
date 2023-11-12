package dao

import (
	"context"
	"gorm.io/gorm"
	"xfd-backend/database/db"
	"xfd-backend/database/db/model"
	"xfd-backend/pkg/types"
)

type UserVerifyDao struct {
}

func NewUserVerifyDao() *UserVerifyDao {
	return &UserVerifyDao{}
}

func (d *UserVerifyDao) ListUserVerifyByUserID(ctx context.Context, userID string) (userVerifyList []*model.UserVerify, err error) {
	if err = db.Get().Model(&model.UserVerify{}).Where("user_id = ?", userID).
		Order("id desc").Find(&userVerifyList).Error; err != nil {
		return nil, err
	}
	return userVerifyList, nil
}

func (d *UserVerifyDao) List(ctx context.Context, page types.PageRequest) (list []*model.UserVerify, count int64, err error) {
	sql := db.Get().Model(&model.UserVerify{}).Order("updated_at desc")
	err = sql.Offset((page.PageNum - 1) * page.PageSize).Limit(page.PageSize).Find(&list).Error
	if err != nil {
		return nil, 0, err
	}

	if err = sql.Count(&count).Error; err != nil {
		return nil, 0, err
	}
	return list, count, nil
}

func (d *UserVerifyDao) GetByID(ctx context.Context, id int) (verify *model.UserVerify, err error) {
	err = db.Get().Model(&model.UserVerify{}).Where("id = ?", id).First(&verify).Error
	if err != nil {
		return nil, err
	}
	return verify, nil
}

func (d *UserVerifyDao) GetByUserID(ctx context.Context, userID string) (verify *model.UserVerify, err error) {
	err = db.Get().Model(&model.UserVerify{}).Where("user_id = ?", userID).First(&verify).Error
	if err != nil {
		return nil, err
	}
	return verify, nil
}

func (d *UserVerifyDao) GetByStatus(ctx context.Context, status model.UserVerifyStatus) (verify *model.UserVerify, err error) {
	err = db.Get().Model(&model.UserVerify{}).Where("status = ?", status).
		Order("id asc").First(&verify).Error
	if err == gorm.ErrRecordNotFound {
		return nil, nil
	} else if err != nil {
		return nil, err
	}
	return verify, nil
}

func (d *UserVerifyDao) CountByStatus(ctx context.Context, status model.UserVerifyStatus) (count int64, err error) {
	err = db.Get().Model(&model.UserVerify{}).Where("status = ?", status).Count(&count).Error
	if err != nil {
		return 0, err
	}
	return count, nil
}

func (d *UserVerifyDao) Create(ctx context.Context, User *model.UserVerify) (err error) {
	err = db.Get().Model(&model.UserVerify{}).Create(User).Error
	if err != nil {
		return err
	}
	return nil
}

func (d *UserVerifyDao) CreateInTx(tx *gorm.DB, User *model.UserVerify) (err error) {
	err = tx.Model(&model.UserVerify{}).Create(User).Error
	if err != nil {
		return err
	}
	return nil
}

func (d *UserVerifyDao) UpdateByID(ctx context.Context, id int, updateValue *model.UserVerify) (err error) {
	updateResult := db.Get().Model(&model.UserVerify{}).Where("id =?", id).Updates(updateValue)
	if err = updateResult.Error; err != nil {
		return err
	}
	return nil
}

func (d *UserVerifyDao) UpdateByIDInTx(tx *gorm.DB, id int, updateValue *model.UserVerify) (err error) {
	updateResult := tx.Model(&model.UserVerify{}).Where("id =?", id).Updates(updateValue)
	if err = updateResult.Error; err != nil {
		return err
	}
	return nil
}
