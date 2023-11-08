package dao

import (
	"context"
	"gorm.io/gorm"
	"xfd-backend/database/db"
	"xfd-backend/database/db/model"
	"xfd-backend/pkg/types"
)

type UserDao struct {
}

func NewUserDao() *UserDao {
	return &UserDao{}
}

func (d *UserDao) ListByOrgID(ctx context.Context, page types.PageRequest, orgID int, username, phone string) (list []*model.User, count int64, err error) {
	sql := db.Get().Model(&model.User{})

	if orgID > 0 {
		sql = sql.Where("organization_id = ? ", orgID)
	}
	if len(username) > 0 {
		sql = sql.Where("username = ? ", username)
	}
	if len(phone) > 0 {
		sql = sql.Where("phone = ? ", phone)
	}
	if err = sql.Offset((page.PageNum - 1) * page.PageSize).Limit(page.PageSize).Find(&list).Error; err != nil {
		return nil, 0, err
	}

	if err = sql.Count(&count).Error; err != nil {
		return nil, 0, err
	}
	return list, count, nil
}

func (d *UserDao) ListByStatus(ctx context.Context, page types.PageRequest, roles []model.UserRole) (list []*model.User, count int64, err error) {
	sql := db.Get().Model(&model.User{}).Where("user_role IN (?)", roles)
	if err = sql.Offset((page.PageNum - 1) * page.PageSize).Limit(page.PageSize).Find(&list).Error; err != nil {
		return nil, 0, err
	}

	if err = sql.Count(&count).Error; err != nil {
		return nil, 0, err
	}
	return list, count, nil
}

func (d *UserDao) ListByUserIDs(ctx context.Context, userIDs []string) (UserList []*model.User, err error) {
	if err = db.Get().Model(&model.User{}).Where("user_id in (?)", userIDs).Find(&UserList).Error; err != nil {
		return nil, err
	}
	return UserList, nil
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
	if err != nil {
		return nil, err
	}
	return User, nil
}
func (d *UserDao) GetByUserIDInTx(tx *gorm.DB, userID string) (User *model.User, err error) {
	err = tx.Model(&model.User{}).Where("user_id = ?", userID).First(&User).Error
	if err != nil {
		return nil, err
	}
	return User, nil
}
func (d *UserDao) GetByPhoneInTx(tx *gorm.DB, phone string) (User *model.User, err error) {
	err = tx.Model(&model.User{}).Where("phone = ?", phone).First(&User).Error
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

func (d *UserDao) SaveInTx(tx *gorm.DB, User *model.User) (err error) {
	err = tx.Model(&model.User{}).Save(User).Error
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

func (d *UserDao) UpdateByUserID(ctx context.Context, userID string, updateValue *model.User) (err error) {
	updateResult := db.Get().Model(&model.User{}).Where("user_id =?", userID).Updates(updateValue)
	if err = updateResult.Error; err != nil {
		return err
	}
	return nil
}

func (d *UserDao) UpdateByUserIDInTx(tx *gorm.DB, userID string, updateValue *model.User) (err error) {
	updateResult := tx.Model(&model.User{}).Where("user_id =?", userID).Updates(updateValue)
	if err = updateResult.Error; err != nil {
		return err
	}
	return nil
}

func (d *UserDao) CountByOrganization(ctx context.Context, orgID int) (count int64, err error) {
	sql := db.Get().Model(&model.User{}).Where("organization_id = ?", orgID)
	if err = sql.Count(&count).Error; err != nil {
		return 0, err
	}
	return count, nil
}

func (d *UserDao) CountByOrganizationAndStatus(ctx context.Context, orgID int, status model.UserPointStatus) (count int64, err error) {
	sql := db.Get().Model(&model.User{}).Where("organization_id = ? AND point_status = ?", orgID, status)
	if err = sql.Count(&count).Error; err != nil {
		return 0, err
	}
	return count, nil
}

func (d *UserDao) DeleteByUserID(ctx context.Context, userID string) (err error) {
	if err = db.Get().Model(&model.User{}).Where("user_id = ?", userID).Delete(&model.User{}).Error; err != nil {
		return err
	}
	return nil
}
