package model

import "gorm.io/gorm"

type Config struct {
	gorm.Model
	Key     string `gorm:"column:key" json:"key"`
	Value   string `gorm:"column:value" json:"value"`
	Deleted *int   `gorm:"column:deleted" json:"deleted"`
}

func (Config) TableName() string {
	return "config"
}
