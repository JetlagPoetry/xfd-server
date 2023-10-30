package model

import (
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	UserID         string   `gorm:"column:user_id;not null" json:"user_id"`
	Phone          string   `gorm:"column:phone;not null" json:"phone"`
	UserRole       UserRole `gorm:"column:user_role;not null" json:"user_role"`
	Username       string   `gorm:"column:username;not null" json:"username"`
	AvatarURL      string   `gorm:"column:avatar_url;not null" json:"avatar_url"`
	OrganizationID int      `gorm:"column:organization_id;not null" json:"organization_id"`
	Organization   string   `gorm:"column:organization;not null" json:"organization"`
	Point          int      `gorm:"column:point;not null" json:"point"`
	Deleted        int      `gorm:"column:deleted" json:"deleted"`
}

type UserRole int

const (
	UserRoleUnknown  UserRole = 0
	UserRoleSupplier UserRole = 1
	UserRoleBuyer    UserRole = 2
	UserRoleCustomer UserRole = 3
	UserRoleAdmin    UserRole = 4
)

func (u *User) TableName() string {
	return "user"
}
