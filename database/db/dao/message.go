package dao

import (
	"context"
	"xfd-backend/database/db"
	"xfd-backend/database/db/model"
)

type MessageDao struct {
}

func NewMessageDao() *MessageDao {
	return &MessageDao{}
}

func (d *MessageDao) Lists(ctx context.Context, limit, offset int) (list []*model.Message, count int64, err error) {
	if err = db.Get().Model(&model.Message{}).Find(&list).Limit(limit).Offset(offset).Error; err != nil {
		return nil, 0, err
	}

	if err = db.Get().Model(&model.Message{}).Count(&count).Error; err != nil {
		return nil, 0, err
	}
	return list, count, nil
}

func (d *MessageDao) GetByID(ctx context.Context, id int) (message *model.Message, err error) {
	err = db.Get().Model(&model.Message{}).Where("id = ?", id).First(&message).Error
	if err != nil {
		return nil, err
	}
	return message, nil
}

func (d *MessageDao) GetByMessageID(ctx context.Context, userID string) (message *model.Message, err error) {
	err = db.Get().Model(&model.Message{}).Where("user_id = ?", userID).First(&message).Error
	if err != nil {
		return nil, err
	}
	return message, nil
}

func (d *MessageDao) ListByConversationID(ctx context.Context, conversationID int) (list []*model.Message, err error) {
	err = db.Get().Model(&model.Message{}).Where("conversation_id = ?", conversationID).
		Order("updated_at desc").Find(&list).Error
	if err != nil {
		return nil, err
	}
	return list, nil
}

func (d *MessageDao) GetByConversationID(ctx context.Context, conversationID int) (message *model.Message, err error) {
	err = db.Get().Model(&model.Message{}).Where("conversation_id = ?", conversationID).
		Order("updated_at desc").First(&message).Error
	if err != nil {
		return nil, err
	}
	return message, nil
}
