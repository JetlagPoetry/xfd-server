package handler

import (
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"xfd-backend/pkg/response"
	"xfd-backend/service"
	"xfd-backend/types"
)

type UserHandler struct {
	UserService *service.UserService
}

func NewUserHandler() *UserHandler {
	return &UserHandler{UserService: service.NewUserService()}
}

func (h *UserHandler) GetUserInfo(c *gin.Context) {
	var (
		//req  *types.GetUserInfoReq
		resp *types.GetUserInfoResp
		err  error
	)

	userID := c.Query("userID")
	resp, err = h.UserService.GetUserInfo(c, userID)
	if err != nil {
		log.Println("[UserHandler] GetUserInfo failed, err=", err)
		c.JSON(http.StatusOK, response.RespSuccess(c, err))
		return
	}
	c.JSON(http.StatusOK, response.RespSuccess(c, resp))
}
