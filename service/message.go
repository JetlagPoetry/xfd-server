package service

import (
	"context"
	"xfd-backend/database/db/dao"
	"xfd-backend/pkg/common"
	"xfd-backend/pkg/types"
	"xfd-backend/pkg/xerr"
)

type MessageService struct {
	conversationDao *dao.ConversationDao
	messageDao      *dao.MessageDao
	userDao         *dao.UserDao
}

func NewMessageService() *MessageService {
	return &MessageService{
		conversationDao: dao.NewConversationDao(),
		messageDao:      dao.NewMessageDao(),
		userDao:         dao.NewUserDao(),
	}
}

func (s *MessageService) GetConversations(ctx context.Context, req *types.MessageGetConversationsReq) (*types.MessageGetConversationsResp, xerr.XErr) {
	userID := common.GetUserID(ctx)

	conversationList, err := s.conversationDao.ListByUserID(ctx, userID, req.BasePage)
	if err != nil {
		return nil, xerr.WithCode(xerr.ErrorDatabase, err)
	}

	list := make([]*types.Conversation, 0)
	if len(conversationList) == 0 {
		return &types.MessageGetConversationsResp{List: list}, nil
	}

	for _, con := range conversationList {
		// todo lastMessage可以改成redis存储
		message, err := s.messageDao.GetByConversationID(ctx, int(con.ID))
		if err != nil {
			return nil, xerr.WithCode(xerr.ErrorDatabase, err)
		}
		otherUserID := con.UserA
		if con.UserA == userID {
			otherUserID = con.UserB
		}
		// todo userList可以改成批量
		otherUser, err := s.userDao.GetByUserID(ctx, otherUserID)
		if err != nil {
			return nil, xerr.WithCode(xerr.ErrorDatabase, err)
		}
		list = append(list, &types.Conversation{
			ConversationID: int64(con.ID),
			Username:       otherUser.Username,
			AvatarURL:      otherUser.AvatarURL,
			LastMessage:    message.Content,
			LastTime:       con.UpdatedAt.Unix(),
			RedDot:         false, // todo redis保存该用户该会话已读的最大message id
		})
	}

	return &types.MessageGetConversationsResp{List: list}, nil
}

func (s *MessageService) GetMessages(ctx context.Context, req *types.MessageGetMessagesReq) (*types.MessageGetMessagesResp, xerr.XErr) {
	userID := common.GetUserID(ctx)
	list := make([]*types.Message, 0)

	conversation, err := s.conversationDao.GetByID(ctx, req.ConversationID)
	if err != nil {
		return nil, xerr.WithCode(xerr.ErrorDatabase, err)
	}
	if conversation.UserA != userID && conversation.UserB != userID {
		return nil, xerr.WithCode(xerr.ErrorOperationForbidden, err)
	}
	messageList, err := s.messageDao.ListByConversationID(ctx, req.ConversationID)
	if err != nil {
		return nil, xerr.WithCode(xerr.ErrorDatabase, err)
	}
	// todo redis保存该用户该会话->messageList[0].id

	for _, item := range messageList {
		list = append(list, &types.Message{
			FromUserID: item.FromUserID,
			ToUserID:   item.ToUserID,
			Type:       item.Type,
			Content:    item.Content,
			OrderID:    item.OrderID,
			Time:       item.CreatedAt.Unix(),
		})
	}
	return &types.MessageGetMessagesResp{List: list}, nil
}
