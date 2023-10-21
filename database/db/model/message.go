package model

import "gorm.io/gorm"

type Message struct {
	gorm.Model
	ConversationID int    `gorm:"column:conversation_id;not null" json:"conversation_id"`
	FromUserID     string `gorm:"column:from_user_id;not null" json:"from_user_id"`
	Content        string `gorm:"column:content" json:"content"`
	Status         int    `gorm:"column:status;default:0" json:"status"`
	Deleted        *int   `gorm:"column:deleted" json:"deleted"`
}

func (Message) TableName() string {
	return "message"
}
