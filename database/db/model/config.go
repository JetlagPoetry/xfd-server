package model

import "gorm.io/gorm"

type Config struct {
	gorm.Model
	Name  string `gorm:"column:name" json:"name"`
	Value string `gorm:"column:value" json:"value"`
}

func (Config) TableName() string {
	return "config"
}
