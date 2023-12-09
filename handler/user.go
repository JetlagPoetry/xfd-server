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
	"xfd-backend/pkg/xerr"
	"xfd-backend/service"
)

type UserHandler struct {
	userService *service.UserService
}

func NewUserHandler() *UserHandler {
	return &UserHandler{userService: service.NewUserService()}
}

func (h *UserHandler) SendCode(c *gin.Context) {
	var (
		req  types.UserSendCodeReq
		resp *types.UserSendCodeResp
		xErr xerr.XErr
	)

	err := c.BindJSON(&req)
	if err != nil {
		c.JSON(http.StatusOK, response.RespError(c, xerr.WithCode(xerr.InvalidParams, err)))
		return
	}

	if !utils.Mobile(req.Phone) {
		c.JSON(http.StatusOK, response.RespError(c, xerr.WithCode(xerr.InvalidParams, errors.New("invalid param"))))
		return
	}

	resp, xErr = h.userService.SendCode(c, req)
	if xErr != nil {
		log.Println("[UserHandler] Login failed, err=", xErr)
		c.JSON(http.StatusOK, response.RespError(c, xErr))
		return
	}
	c.JSON(http.StatusOK, response.RespSuccess(c, resp))
}

func (h *UserHandler) Login(c *gin.Context) {
	var (
		req  types.UserLoginReq
		resp *types.UserLoginResp
		xErr xerr.XErr
	)

	err := c.BindJSON(&req)
	if err != nil {
		c.JSON(http.StatusOK, response.RespError(c, xerr.WithCode(xerr.InvalidParams, err)))
		return
	}

	if !utils.Mobile(req.Phone) || req.Source == types.SourceUnknown {
		c.JSON(http.StatusOK, response.RespError(c, xerr.WithCode(xerr.InvalidParams, errors.New("invalid param"))))
		return
	}

	//if  len(req.Code) == 0{
	//	c.JSON(http.StatusOK, response.RespError(c, xerr.WithCode(xerr.InvalidParams, errors.New("invalid param"))))
	//	return
	//}

	resp, xErr = h.userService.Login(c, req)
	if xErr != nil {
		log.Println("[UserHandler] Login failed, err=", xErr)
		c.JSON(http.StatusOK, response.RespError(c, xErr))
		return
	}
	c.JSON(http.StatusOK, response.RespSuccess(c, resp))
}

func (h *UserHandler) SubmitRole(c *gin.Context) {
	var (
		req  types.UserSubmitRoleReq
		resp *types.UserSubmitRoleResp
		xErr xerr.XErr
	)

	err := c.BindJSON(&req)
	if err != nil {
		c.JSON(http.StatusOK, response.RespError(c, xerr.WithCode(xerr.InvalidParams, err)))
		return
	}

	if req.Role != model.UserRoleBuyer && req.Role != model.UserRoleSupplier && req.Role != model.UserRoleCustomer {
		c.JSON(http.StatusOK, response.RespError(c, xerr.WithCode(xerr.InvalidParams, errors.New("invalid param"))))
		return
	}

	if (req.Role == model.UserRoleSupplier || req.Role == model.UserRoleBuyer) && (len(req.Organization) == 0 || len(req.OrganizationCode) == 0 ||
		len(req.OrganizationURL) == 0 || len(req.IdentityURLA) == 0 || len(req.IdentityURLB) == 0 ||
		len(req.RealName) == 0 || len(req.Position) == 0 || !utils.Mobile(req.Phone)) {
		c.JSON(http.StatusOK, response.RespError(c, xerr.WithCode(xerr.InvalidParams, errors.New("invalid param"))))
		return
	}

	resp, xErr = h.userService.SubmitRole(c, req)
	if xErr != nil {
		log.Println("[UserHandler] SubmitRole failed, err=", xErr)
		c.JSON(http.StatusOK, response.RespError(c, xErr))
		return
	}
	c.JSON(http.StatusOK, response.RespSuccess(c, resp))
}

func (h *UserHandler) RefreshToken(c *gin.Context) {
	var (
		resp *types.UserRefreshTokenResp
		xErr xerr.XErr
	)

	resp, xErr = h.userService.RefreshToken(c)
	if xErr != nil {
		log.Println("[UserHandler] RefreshToken failed, err=", xErr)
		c.JSON(http.StatusOK, response.RespError(c, xErr))
		return
	}
	c.JSON(http.StatusOK, response.RespSuccess(c, resp))
}

func (h *UserHandler) GetVerifyInfo(c *gin.Context) {
	var (
		resp *types.GetVerifyInfoResp
		xErr xerr.XErr
	)

	resp, xErr = h.userService.GetVerifyInfo(c)
	if xErr != nil {
		log.Println("[UserHandler] GetVerifyInfo failed, err=", xErr)
		c.JSON(http.StatusOK, response.RespError(c, xErr))
		return
	}
	c.JSON(http.StatusOK, response.RespSuccess(c, resp))
}

func (h *UserHandler) GetUserInfo(c *gin.Context) {
	var (
		resp *types.GetUserInfoResp
		xErr xerr.XErr
	)

	resp, xErr = h.userService.GetUserInfo(c)
	if xErr != nil {
		log.Println("[UserHandler] GetUserInfo failed, err=", xErr)
		c.JSON(http.StatusOK, response.RespError(c, xErr))
		return
	}
	c.JSON(http.StatusOK, response.RespSuccess(c, resp))
}

func (h *UserHandler) ModifyInfo(c *gin.Context) {
	var (
		req  types.UserModifyInfoReq
		resp *types.UserModifyInfoResp
		xErr xerr.XErr
	)

	err := c.BindJSON(&req)
	if err != nil {
		c.JSON(http.StatusOK, response.RespError(c, xerr.WithCode(xerr.InvalidParams, err)))
		return
	}

	resp, xErr = h.userService.ModifyUserInfo(c, req)
	if xErr != nil {
		log.Println("[UserHandler] ModifyUserInfo failed, err=", xErr)
		c.JSON(http.StatusOK, response.RespError(c, xErr))
		return
	}
	c.JSON(http.StatusOK, response.RespSuccess(c, resp))
}

func (h *UserHandler) AssignAdmin(c *gin.Context) {
	var (
		req  types.UserAssignAdminReq
		resp *types.UserAssignAdminResp
		xErr xerr.XErr
	)

	err := c.BindJSON(&req)
	if err != nil {
		c.JSON(http.StatusOK, response.RespError(c, xerr.WithCode(xerr.InvalidParams, err)))
		return
	}

	if !utils.Mobile(req.Phone) || (req.Role != model.UserRoleAdmin && req.Role != model.UserRoleRoot) {
		c.JSON(http.StatusOK, response.RespError(c, xerr.WithCode(xerr.InvalidParams, errors.New("invalid param"))))
		return
	}

	resp, xErr = h.userService.AssignAdmin(c, req)
	if xErr != nil {
		log.Println("[UserHandler] AssignAdmin failed, err=", xErr)
		c.JSON(http.StatusOK, response.RespError(c, xErr))
		return
	}
	c.JSON(http.StatusOK, response.RespSuccess(c, resp))
}

func (h *UserHandler) GetAdmins(c *gin.Context) {
	var (
		req  types.UserGetAdminsReq
		resp *types.UserGetAdminsResp
		xErr xerr.XErr
	)

	err := c.BindQuery(&req)
	if err != nil {
		c.JSON(http.StatusOK, response.RespError(c, xerr.WithCode(xerr.InvalidParams, err)))
		return
	}

	if err = req.CheckParams(); err != nil {
		c.JSON(http.StatusOK, response.RespError(c, xerr.WithCode(xerr.InvalidParams, err)))
		return
	}

	resp, xErr = h.userService.GetAdmins(c, req)
	if xErr != nil {
		log.Println("[UserHandler] GetAdmins failed, err=", xErr)
		c.JSON(http.StatusOK, response.RespError(c, xErr))
		return
	}
	c.JSON(http.StatusOK, response.RespSuccess(c, resp))
}

func (h *UserHandler) DeleteUser(c *gin.Context) {
	var (
		req  types.UserDeleteUserReq
		resp *types.UserDeleteUserResp

		xErr xerr.XErr
	)

	err := c.BindJSON(&req)
	if err != nil {
		c.JSON(http.StatusOK, response.RespError(c, xerr.WithCode(xerr.InvalidParams, err)))
		return
	}

	if len(req.UserID) == 0 {
		c.JSON(http.StatusOK, response.RespError(c, xerr.WithCode(xerr.InvalidParams, errors.New("invalid param"))))
		return
	}

	resp, xErr = h.userService.DeleteUser(c, req)
	if xErr != nil {
		log.Println("[UserHandler] DeleteUser failed, err=", xErr)
		c.JSON(http.StatusOK, response.RespError(c, xErr))
		return
	}
	c.JSON(http.StatusOK, response.RespSuccess(c, resp))
}

func (h *UserHandler) ImSig(c *gin.Context) {
	var (
		req  types.UserImSigReq
		resp *types.UserImSigResp

		xErr xerr.XErr
	)

	resp, xErr = h.userService.ImSig(c, req)
	if xErr != nil {
		log.Println("[UserHandler] ImSig failed, err=", xErr)
		c.JSON(http.StatusOK, response.RespError(c, xErr))
		return
	}
	c.JSON(http.StatusOK, response.RespSuccess(c, resp))
}

func (h *UserHandler) GetAddressList(c *gin.Context) {
	var (
		req  types.UserGetAddressListReq
		resp *types.UserGetAddressListResp
		xErr xerr.XErr
	)

	err := c.BindQuery(&req)
	if err != nil {
		c.JSON(http.StatusOK, response.RespError(c, xerr.WithCode(xerr.InvalidParams, err)))
		return
	}

	resp, xErr = h.userService.GetAddressList(c, req)
	if xErr != nil {
		log.Println("[UserHandler] GetAddressList failed, err=", xErr)
		c.JSON(http.StatusOK, response.RespError(c, xErr))
		return
	}
	c.JSON(http.StatusOK, response.RespSuccess(c, resp))
}

func (h *UserHandler) GetDefaultAddress(c *gin.Context) {
	var (
		req  types.UserGetDefaultAddressReq
		resp *types.UserGetDefaultAddressResp
		xErr xerr.XErr
	)

	err := c.BindQuery(&req)
	if err != nil {
		c.JSON(http.StatusOK, response.RespError(c, xerr.WithCode(xerr.InvalidParams, err)))
		return
	}

	resp, xErr = h.userService.GetDefaultAddress(c, req)
	if xErr != nil {
		log.Println("[UserHandler] GetDefaultAddress failed, err=", xErr)
		c.JSON(http.StatusOK, response.RespError(c, xErr))
		return
	}
	c.JSON(http.StatusOK, response.RespSuccess(c, resp))
}

func (h *UserHandler) AddAddress(c *gin.Context) {
	var (
		req  types.UserAddAddressReq
		resp *types.UserAddAddressResp
		xErr xerr.XErr
	)

	err := c.BindJSON(&req)
	if err != nil {
		c.JSON(http.StatusOK, response.RespError(c, xerr.WithCode(xerr.InvalidParams, err)))
		return
	}

	if len(req.Name) == 0 || len(req.Province) == 0 || len(req.City) == 0 || len(req.Region) == 0 || len(req.Phone) == 0 || len(req.Address) == 0 {
		c.JSON(http.StatusOK, response.RespError(c, xerr.WithCode(xerr.InvalidParams, errors.New("invalid param"))))
		return
	}

	resp, xErr = h.userService.AddAddress(c, req)
	if xErr != nil {
		log.Println("[UserHandler] AddAddress failed, err=", xErr)
		c.JSON(http.StatusOK, response.RespError(c, xErr))
		return
	}
	c.JSON(http.StatusOK, response.RespSuccess(c, resp))
}

func (h *UserHandler) ModifyAddress(c *gin.Context) {
	var (
		req  types.UserModifyAddressReq
		resp *types.UserModifyAddressResp
		xErr xerr.XErr
	)

	err := c.BindJSON(&req)
	if err != nil {
		c.JSON(http.StatusOK, response.RespError(c, xerr.WithCode(xerr.InvalidParams, err)))
		return
	}

	if req.ID == 0 {
		c.JSON(http.StatusOK, response.RespError(c, xerr.WithCode(xerr.InvalidParams, err)))
		return
	}

	resp, xErr = h.userService.ModifyAddress(c, req)
	if xErr != nil {
		log.Println("[UserHandler] ModifyAddress failed, err=", xErr)
		c.JSON(http.StatusOK, response.RespError(c, xErr))
		return
	}
	c.JSON(http.StatusOK, response.RespSuccess(c, resp))
}

func (h *UserHandler) DeleteAddress(c *gin.Context) {
	var (
		req  types.UserDeleteAddressReq
		resp *types.UserDeleteAddressResp
		xErr xerr.XErr
	)

	err := c.BindJSON(&req)
	if err != nil {
		c.JSON(http.StatusOK, response.RespError(c, xerr.WithCode(xerr.InvalidParams, err)))
		return
	}

	if req.ID == 0 {
		c.JSON(http.StatusOK, response.RespError(c, xerr.WithCode(xerr.InvalidParams, err)))
		return
	}

	resp, xErr = h.userService.DeleteAddress(c, req)
	if xErr != nil {
		log.Println("[UserHandler] DeleteAddress failed, err=", xErr)
		c.JSON(http.StatusOK, response.RespError(c, xErr))
		return
	}
	c.JSON(http.StatusOK, response.RespSuccess(c, resp))
}
