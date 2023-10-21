package handler

import (
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"xfd-backend/pkg/response"
	"xfd-backend/pkg/types"
	"xfd-backend/service"
)

type UserHandler struct {
	UserService *service.UserService
}

func NewUserHandler() *UserHandler {
	return &UserHandler{UserService: service.NewUserService()}
}

func (h *UserHandler) Login(c *gin.Context) {
	var (
		req  *types.UserLoginReq
		resp *types.UserLoginResp
		err  error
	)

	err = c.BindQuery(&req)
	if err != nil {
		c.JSON(http.StatusOK, response.RespError(c, err, "", response.InvalidParams))
		return
	}

	resp, err = h.UserService.Login(c, req)
	if err != nil {
		log.Println("[UserHandler] GetUserInfo failed, err=", err)
		c.JSON(http.StatusOK, response.RespSuccess(c, err))
		return
	}
	c.JSON(http.StatusOK, response.RespSuccess(c, resp))
}

func (h *UserHandler) RefreshToken(c *gin.Context) {
	var (
		req  *types.UserLoginReq
		resp *types.UserLoginResp
		err  error
	)

	err = c.BindQuery(&req)
	if err != nil {
		c.JSON(http.StatusOK, response.RespError(c, err, "", response.InvalidParams))
		return
	}

	resp, err = h.UserService.Login(c, req)
	if err != nil {
		log.Println("[UserHandler] GetUserInfo failed, err=", err)
		c.JSON(http.StatusOK, response.RespSuccess(c, err))
		return
	}
	c.JSON(http.StatusOK, response.RespSuccess(c, resp))
}
