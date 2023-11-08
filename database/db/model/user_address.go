package model

import "gorm.io/gorm"

type UserAddress struct {
	gorm.Model
	UserID    string `gorm:"column:user_id" json:"user_id"`
	Name      string `gorm:"column:name" json:"name"`
	Phone     string `gorm:"column:phone" json:"phone"`
	Province  string `gorm:"column:province" json:"province"`
	City      string `gorm:"column:city" json:"city"`
	Region    string `gorm:"column:region" json:"region"`
	Address   string `gorm:"column:address" json:"address"`
	IsDefault *int   `gorm:"column:is_default" json:"is_default"` // 0非默认，1默认
}

func (UserAddress) TableName() string {
	return "user_address"
}
