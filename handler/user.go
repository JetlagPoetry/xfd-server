package handler

import (
	"errors"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"xfd-backend/database/db/model"
	"xfd-backend/pkg/response"
	"xfd-backend/pkg/types"
	"xfd-backend/pkg/utils"
	"xfd-backend/service"
)

type UserHandler struct {
	UserService *service.UserService
}

func NewUserHandler() *UserHandler {
	return &UserHandler{UserService: service.NewUserService()}
}

func (h *UserHandler) SendCode(c *gin.Context) {
	var (
		req  *types.UserSendCodeReq
		resp *types.UserSendCodeResp
		err  error
	)

	err = c.BindJSON(&req)
	if err != nil {
		c.JSON(http.StatusOK, response.RespError(c, err, "", response.InvalidParams))
		return
	}

	if !utils.Mobile(req.Phone) {
		c.JSON(http.StatusOK, response.RespError(c, errors.New("invalid param"), "", response.InvalidParams))
		return
	}

	resp, err = h.UserService.SendCode(c, req)
	if err != nil {
		log.Println("[UserHandler] SendCode failed, err=", err)
		c.JSON(http.StatusOK, response.RespSuccess(c, err))
		return
	}
	c.JSON(http.StatusOK, response.RespSuccess(c, resp))
}

func (h *UserHandler) Login(c *gin.Context) {
	var (
		req  *types.UserLoginReq
		resp *types.UserLoginResp
		err  error
	)

	err = c.BindJSON(&req)
	if err != nil {
		c.JSON(http.StatusOK, response.RespError(c, err, "", response.InvalidParams))
		return
	}

	if !utils.Mobile(req.Phone) || len(req.Code) == 0 {
		c.JSON(http.StatusOK, response.RespError(c, errors.New("invalid param"), "", response.InvalidParams))
		return
	}

	resp, err = h.UserService.Login(c, req)
	if err != nil {
		log.Println("[UserHandler] Login failed, err=", err)
		c.JSON(http.StatusOK, response.RespSuccess(c, err))
		return
	}
	c.JSON(http.StatusOK, response.RespSuccess(c, resp))
}

func (h *UserHandler) SubmitRole(c *gin.Context) {
	var (
		req  *types.UserSubmitRoleReq
		resp *types.UserSubmitRoleResp
		err  error
	)

	err = c.BindJSON(&req)
	if err != nil {
		c.JSON(http.StatusOK, response.RespError(c, err, "", response.InvalidParams))
		return
	}

	if req.Role == model.UserRoleUnknown ||
		(req.Role == model.UserRoleSupplier || req.Role == model.UserRoleBuyer) && (len(req.Organization) == 0 || len(req.OrganizationCode) == 0 ||
			len(req.OrganizationURL) == 0 || len(req.CorporationURLA) == 0 || len(req.CorporationURLB) == 0 ||
			len(req.RealName) == 0 || len(req.CertificateNo) == 0 || len(req.Position) == 0 || !utils.Mobile(req.Phone)) {
		c.JSON(http.StatusOK, response.RespError(c, errors.New("invalid param"), "", response.InvalidParams))
		return
	}

	resp, err = h.UserService.SubmitRole(c, req)
	if err != nil {
		log.Println("[UserHandler] SubmitRole failed, err=", err)
		c.JSON(http.StatusOK, response.RespSuccess(c, err))
		return
	}
	c.JSON(http.StatusOK, response.RespSuccess(c, resp))
}

func (h *UserHandler) RefreshToken(c *gin.Context) {
	var (
		resp *types.UserRefreshTokenResp
		err  error
	)

	resp, err = h.UserService.RefreshToken(c)
	if err != nil {
		log.Println("[UserHandler] RefreshToken failed, err=", err)
		c.JSON(http.StatusOK, response.RespSuccess(c, err))
		return
	}
	c.JSON(http.StatusOK, response.RespSuccess(c, resp))
}

func (h *UserHandler) GetUserInfo(c *gin.Context) {
	var (
		resp *types.GetUserInfoResp
		err  error
	)

	resp, err = h.UserService.GetUserInfo(c)
	if err != nil {
		log.Println("[UserHandler] GetUserInfo failed, err=", err)
		c.JSON(http.StatusOK, response.RespSuccess(c, err))
		return
	}
	c.JSON(http.StatusOK, response.RespSuccess(c, resp))
}

func (h *UserHandler) ModifyInfo(c *gin.Context) {
	var (
		req  *types.UserModifyInfoReq
		resp *types.UserModifyInfoResp
		err  error
	)

	err = c.BindJSON(&req)
	if err != nil {
		c.JSON(http.StatusOK, response.RespError(c, err, "", response.InvalidParams))
		return
	}

	resp, err = h.UserService.ModifyUserInfo(c, req)
	if err != nil {
		log.Println("[UserHandler] ModifyUserInfo failed, err=", err)
		c.JSON(http.StatusOK, response.RespSuccess(c, err))
		return
	}
	c.JSON(http.StatusOK, response.RespSuccess(c, resp))
}
