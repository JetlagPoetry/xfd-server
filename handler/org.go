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

type OrgHandler struct {
	orgService *service.OrgService
}

func NewOrgHandler() *OrgHandler {
	return &OrgHandler{orgService: service.NewOrgService()}
}

func (h *OrgHandler) ApplyPoint(c *gin.Context) {
	var (
		req  *types.OrgApplyPointReq
		resp *types.OrgApplyPointResp
		xErr xerr.XErr
	)

	err := c.BindJSON(&req)
	if err != nil {
		c.JSON(http.StatusOK, response.RespError(c, xerr.WithCode(xerr.InvalidParams, err)))
		return
	}
	comment := c.PostForm("comment")
	file, header, err := c.Request.FormFile("file")
	req = &types.OrgApplyPointReq{
		File:       file,
		FileHeader: header,
		Comment:    comment,
	}
	resp, xErr = h.orgService.ApplyPoint(c, req)
	if xErr != nil {
		log.Println("[OrgHandler] ApplyPoint failed, err=", xErr)
		c.JSON(http.StatusOK, response.RespError(c, xErr))
		return
	}
	c.JSON(http.StatusOK, response.RespSuccess(c, resp))
}

func (h *OrgHandler) VerifyPoint(c *gin.Context) {
	var (
		req  *types.OrgVerifyPointReq
		resp *types.OrgVerifyPointResp
		xErr xerr.XErr
	)

	err := c.BindJSON(&req)
	if err != nil {
		c.JSON(http.StatusOK, response.RespError(c, xerr.WithCode(xerr.InvalidParams, err)))
		return
	}
	resp, xErr = h.orgService.VerifyPoint(c, req)
	if xErr != nil {
		log.Println("[OrgHandler] VerifyPoint failed, err=", xErr)
		c.JSON(http.StatusOK, response.RespError(c, xErr))
		return
	}
	c.JSON(http.StatusOK, response.RespSuccess(c, resp))
}

func (h *OrgHandler) GetApplyToVerify(c *gin.Context) {
	var (
		req  *types.OrgGetApplyToVerifyReq
		resp *types.OrgGetApplyToVerifyResp
		xErr xerr.XErr
	)

	err := c.BindJSON(&req)
	if err != nil {
		c.JSON(http.StatusOK, response.RespError(c, xerr.WithCode(xerr.InvalidParams, err)))
		return
	}
	resp, xErr = h.orgService.GetApplyToVerify(c, req)
	if xErr != nil {
		log.Println("[OrgHandler] GetApplyToVerify failed, err=", xErr)
		c.JSON(http.StatusOK, response.RespError(c, xErr))
		return
	}
	c.JSON(http.StatusOK, response.RespSuccess(c, resp))
}

func (h *OrgHandler) GetApplys(c *gin.Context) {
	var (
		req  *types.OrgGetApplysReq
		resp *types.OrgGetApplysResp
		xErr xerr.XErr
	)

	err := c.BindJSON(&req)
	if err != nil {
		c.JSON(http.StatusOK, response.RespError(c, xerr.WithCode(xerr.InvalidParams, err)))
		return
	}
	resp, xErr = h.orgService.GetApplys(c, req)
	if xErr != nil {
		log.Println("[OrgHandler] GetApplys failed, err=", xErr)
		c.JSON(http.StatusOK, response.RespError(c, xErr))
		return
	}
	c.JSON(http.StatusOK, response.RespSuccess(c, resp))
}

func (h *OrgHandler) VerifyAccount(c *gin.Context) {
	var (
		req  *types.VerifyAccountReq
		resp *types.VerifyAccountResp
		xErr xerr.XErr
	)

	err := c.BindJSON(&req)
	if err != nil {
		c.JSON(http.StatusOK, response.RespError(c, xerr.WithCode(xerr.InvalidParams, err)))
		return
	}
	resp, xErr = h.orgService.VerifyAccount(c, req)
	if xErr != nil {
		log.Println("[OrgHandler] VerifyAccount failed, err=", xErr)
		c.JSON(http.StatusOK, response.RespError(c, xErr))
		return
	}
	c.JSON(http.StatusOK, response.RespSuccess(c, resp))
}

func (h *OrgHandler) GetAccountToVerify(c *gin.Context) {
	var (
		req  *types.GetAccountToVerifyReq
		resp *types.GetAccountToVerifyResp
		xErr xerr.XErr
	)

	err := c.BindJSON(&req)
	if err != nil {
		c.JSON(http.StatusOK, response.RespError(c, xerr.WithCode(xerr.InvalidParams, err)))
		return
	}
	resp, xErr = h.orgService.GetAccountToVerify(c, req)
	if xErr != nil {
		log.Println("[OrgHandler] GetAccountToVerify failed, err=", xErr)
		c.JSON(http.StatusOK, response.RespError(c, xErr))
		return
	}
	c.JSON(http.StatusOK, response.RespSuccess(c, resp))
}

func (h *OrgHandler) GetAccounts(c *gin.Context) {
	var (
		req  *types.GetAccountsReq
		resp *types.GetAccountsResp
		xErr xerr.XErr
	)

	err := c.BindJSON(&req)
	if err != nil {
		c.JSON(http.StatusOK, response.RespError(c, xerr.WithCode(xerr.InvalidParams, err)))
		return
	}
	resp, xErr = h.orgService.GetAccounts(c, req)
	if xErr != nil {
		log.Println("[OrgHandler] GetAccounts failed, err=", xErr)
		c.JSON(http.StatusOK, response.RespError(c, xErr))
		return
	}
	c.JSON(http.StatusOK, response.RespSuccess(c, resp))
}
