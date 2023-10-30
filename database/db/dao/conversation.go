package dao

import (
	"context"
	"xfd-backend/database/db"
	"xfd-backend/database/db/model"
	"xfd-backend/pkg/types"
)

type ConversationDao struct {
}

func NewConversationDao() *ConversationDao {
	return &ConversationDao{}
}

func (d *ConversationDao) Lists(ctx context.Context, limit, offset int) (list []*model.Conversation, count int64, err error) {
	if err = db.Get().Model(&model.Conversation{}).Find(&list).Limit(limit).Offset(offset).Error; err != nil {
		return nil, 0, err
	}

	if err = db.Get().Model(&model.Conversation{}).Count(&count).Error; err != nil {
		return nil, 0, err
	}
	return list, count, nil
}

func (d *ConversationDao) GetByID(ctx context.Context, id int) (conversation *model.Conversation, err error) {
	err = db.Get().Model(&model.Conversation{}).Where("id = ?", id).First(&conversation).Error
	if err != nil {
		return nil, err
	}
	return conversation, nil
}

func (d *ConversationDao) ListByUserID(ctx context.Context, userID string, page types.PageRequest) (list []*model.Conversation, err error) {
	err = db.Get().Model(&model.Conversation{}).Where("user_a = ? OR user_b = ?", userID, userID).
		Order("updated_at desc").Offset((page.PageNum - 1) * page.PageSize).Limit(page.PageSize).Find(&list).Error
	if err != nil {
		return nil, err
	}
	return list, nil
}
