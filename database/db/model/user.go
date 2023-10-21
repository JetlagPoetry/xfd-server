package model

import (
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	UserID    string   `gorm:"column:user_id;not null" json:"user_id"`
	OpenID    string   `gorm:"column:open_id;not null" json:"open_id"`
	UserRole  UserRole `gorm:"column:user_role;not null" json:"user_role"`
	Username  string   `gorm:"column:username;not null" json:"username"`
	AvatarURL string   `gorm:"column:avatar_url;not null" json:"avatar_url"`
	Deleted   *int     `gorm:"column:deleted" json:"deleted"`
}

type UserRole int

const (
	Unknown  UserRole = 0
	Supplier UserRole = 1
	Buyer    UserRole = 2
	Consumer UserRole = 3
)

func (u *User) TableName() string {
	return "user"
}
