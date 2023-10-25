package types

import "xfd-backend/database/db/model"

type MessageGetConversationsReq struct {
	BasePage
}

type MessageGetConversationsResp struct {
	List []*Conversation `json:"list"`
}

type Conversation struct {
	ConversationID int64  `json:"conversationID"`
	Username       string `json:"username"`
	AvatarURL      string `json:"avatarURL"`
	LastMessage    string `json:"lastMessage"`
	LastTime       int64  `json:"lastTime"`
	RedDot         bool   `json:"redDot"`
}

type MessageGetMessagesReq struct {
	// todo 使用一个会话内的message编号来控制从第几号开始下发
	ConversationID int `json:"conversationID"`
	BasePage
}

type MessageGetMessagesResp struct {
	List []*Message `json:"list"`
}

type Message struct {
	FromUserID string            `json:"fromUserID"`
	ToUserID   string            `json:"toUserID"`
	Type       model.MessageType `json:"type"`
	Content    string            `json:"content"`
	OrderID    int               `json:"orderID"`
	Time       int64             `json:"time"`
}
