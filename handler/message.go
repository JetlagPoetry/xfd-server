package handler

import (
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"xfd-backend/pkg/response"
	"xfd-backend/pkg/types"
	"xfd-backend/pkg/xerr"
	"xfd-backend/service"
)

type MessageHandler struct {
	MessageService *service.MessageService
}

func NewMessageHandler() *MessageHandler {
	return &MessageHandler{MessageService: service.NewMessageService()}
}

func (h *MessageHandler) GetConversations(c *gin.Context) {
	var (
		req  *types.MessageGetConversationsReq
		resp *types.MessageGetConversationsResp
		xErr xerr.XErr
	)

	err := c.BindQuery(&req)
	if err != nil {
		c.JSON(http.StatusOK, response.RespError(c, xerr.WithCode(xerr.InvalidParams, err)))
		return
	}

	resp, xErr = h.MessageService.GetConversations(c, req)
	if xErr != nil {
		log.Println("[MessageHandler] GetConversations failed, err=", xErr)
		c.JSON(http.StatusOK, response.RespError(c, xErr))
		return
	}
	c.JSON(http.StatusOK, response.RespSuccess(c, resp))
}

func (h *MessageHandler) GetMessages(c *gin.Context) {
	var (
		req  *types.MessageGetMessagesReq
		resp *types.MessageGetMessagesResp
		xErr xerr.XErr
	)

	err := c.BindQuery(&req)
	if err != nil {
		c.JSON(http.StatusOK, response.RespError(c, xerr.WithCode(xerr.InvalidParams, err)))
		return
	}

	resp, xErr = h.MessageService.GetMessages(c, req)
	if xErr != nil {
		log.Println("[MessageHandler] GetConversations failed, err=", xErr)
		c.JSON(http.StatusOK, response.RespError(c, xErr))
		return
	}
	c.JSON(http.StatusOK, response.RespSuccess(c, resp))
}
