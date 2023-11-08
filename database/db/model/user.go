package model

import (
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	UserID           string          `gorm:"column:user_id;not null" json:"user_id"`
	Phone            string          `gorm:"column:phone;not null" json:"phone"`
	UserRole         UserRole        `gorm:"column:user_role;not null" json:"user_role"`
	Username         string          `gorm:"column:username;not null" json:"username"`
	AvatarURL        string          `gorm:"column:avatar_url;not null" json:"avatar_url"`
	OrganizationID   int             `gorm:"column:organization_id;not null" json:"organization_id"`
	OrganizationName string          `gorm:"column:organization_name;not null" json:"organization_name"`
	Point            int             `gorm:"column:point;not null" json:"point"`
	PointStatus      UserPointStatus `gorm:"column:point_status;not null" json:"point_status"`
}

func (u *User) TableName() string {
	return "user"
}

type UserRole int

const (
	UserRoleUnknown  UserRole = 0
	UserRoleCustomer UserRole = 1
	UserRoleSupplier UserRole = 2
	UserRoleBuyer    UserRole = 3
	UserRoleAdmin    UserRole = 4
	UserRoleRoot     UserRole = 5
)

type UserPointStatus int

const (
	UserPointStatusUnknown UserPointStatus = 0
	UserPointStatusOwn     UserPointStatus = 1
)
