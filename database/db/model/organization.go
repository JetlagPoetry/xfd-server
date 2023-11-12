package model

import (
	"gorm.io/gorm"
)

type Organization struct {
	gorm.Model
	Name  string   `gorm:"column:name;not null" json:"name"`
	Code  string   `gorm:"column:code;not null" json:"code"` // 社会信用代码
	Point *float64 `gorm:"column:point;not null" json:"point"`
}

func (u *Organization) TableName() string {
	return "organization"
}
