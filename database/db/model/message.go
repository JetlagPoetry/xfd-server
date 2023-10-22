package model

import "gorm.io/gorm"

type Message struct {
	gorm.Model
	ConversationID int         `gorm:"column:conversation_id;not null" json:"conversation_id"`
	FromUserID     string      `gorm:"column:from_user_id;not null" json:"from_user_id"`
	ToUserID       string      `gorm:"column:to_user_id;not null" json:"to_user_id"`
	Type           MessageType `gorm:"column:type;default:0" json:"type"`
	Content        string      `gorm:"column:content" json:"content"`
	OrderID        int         `gorm:"column:order_id;not null" json:"order_id"`
	Status         int         `gorm:"column:status;default:0" json:"status"`
	Deleted        *int        `gorm:"column:deleted" json:"deleted"`
}

func (Message) TableName() string {
	return "message"
}

type MessageType int32

const (
	MessageTypeText  MessageType = 0
	MessageTypeImage MessageType = 1
	MessageTypeOrder MessageType = 2
)
