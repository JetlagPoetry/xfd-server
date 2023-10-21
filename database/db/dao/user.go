package dao

import (
	"context"
	"gorm.io/gorm"
	"xfd-backend/database/db"
	"xfd-backend/database/db/model"
)

type UserDao struct {
}

func NewUserDB() *UserDao {
	return &UserDao{}
}

func (d *UserDao) GetUserInfo(ctx context.Context, userID string) (user *model.User, err error) {
	if err = db.Get().Model(&model.User{}).Where("id = ?", userID).Find(&user).Error; err != nil {
		return nil, err
	}
	return user, nil
}

func (d *UserDao) Lists(ctx context.Context, limit, offset int) (UserList []*model.User, count int64, err error) {
	if err = db.Get().Model(&model.User{}).Find(&UserList).Limit(limit).Offset(offset).Error; err != nil {
		return nil, 0, err
	}

	if err = db.Get().Model(&model.User{}).Count(&count).Error; err != nil {
		return nil, 0, err
	}
	return UserList, count, nil
}

func (d *UserDao) GetByID(ctx context.Context, id int) (User *model.User, err error) {
	err = db.Get().Model(&model.User{}).Where("id = ?", id).First(&User).Error
	if err != nil {
		return nil, err
	}
	return User, nil
}

func (d *UserDao) GetByUserID(ctx context.Context, userID string) (User *model.User, err error) {
	err = db.Get().Model(&model.User{}).Where("user_id = ?", userID).First(&User).Error
	if err == gorm.ErrRecordNotFound {
		return nil, nil
	} else if err != nil {

		return nil, err
	}
	return User, nil
}

func (d *UserDao) GetByOpenIDAndRoleInTx(tx *gorm.DB, openID string, role model.UserRole) (User *model.User, err error) {
	err = tx.Model(&model.User{}).Where("open_id = ? AND user_role = ?", openID, role).First(&User).Error
	if err == gorm.ErrRecordNotFound {
		return nil, nil
	} else if err != nil {
		return nil, err
	}
	return User, nil
}

func (d *UserDao) GetByOpenIDAndRole(ctx context.Context, openID string, role model.UserRole) (User *model.User, err error) {
	err = db.Get().Model(&model.User{}).Where("open_id = ? AND user_role = ?", openID, role).First(&User).Error
	if err == gorm.ErrRecordNotFound {
		return nil, nil
	} else if err != nil {

		return nil, err
	}
	return User, nil
}

func (d *UserDao) Create(ctx context.Context, User *model.User) (err error) {
	err = db.Get().Model(&model.User{}).Create(User).Error
	if err != nil {
		return err
	}
	return nil
}
func (d *UserDao) CreateInTx(tx *gorm.DB, User *model.User) (err error) {
	err = tx.Model(&model.User{}).Create(User).Error
	if err != nil {
		return err
	}
	return nil
}

func (d *UserDao) Upsert(ctx context.Context, User *model.User) (err error) {
	err = db.Get().Save(User).Error
	if err != nil {
		return err
	}
	return nil
}

func (d *UserDao) UpdateByID(ctx context.Context, id int, updateValue *model.User) (err error) {
	updateResult := db.Get().Model(&model.User{}).Where("id =?", id).Updates(updateValue)
	if err = updateResult.Error; err != nil {
		return err
	}
	return nil
}
