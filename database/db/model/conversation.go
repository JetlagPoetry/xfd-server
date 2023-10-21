package model

import "gorm.io/gorm"

type Conversation struct {
	gorm.Model
	UserA   string `gorm:"column:user_a;not null" json:"user_a"`
	UserB   string `gorm:"column:user_b;not null" json:"user_b"`
	Deleted *int   `gorm:"column:deleted" json:"deleted"`
}

func (Conversation) TableName() string {
	return "conversation"
}
