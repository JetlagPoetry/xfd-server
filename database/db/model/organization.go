package model

import (
	"gorm.io/gorm"
)

// todo 所有发放积分并初次审核通过的，会在这里加一条记录，用于追踪其积分总额
type Organization struct {
	gorm.Model
	Name    string `gorm:"column:name;not null" json:"name"`
	Code    string `gorm:"column:code;not null" json:"code"`
	Point   int    `gorm:"column:point;not null" json:"point"`
	Deleted int    `gorm:"column:deleted" json:"deleted"`
}

func (u *Organization) TableName() string {
	return "organization"
}
