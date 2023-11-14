package handler

import (
	"errors"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"time"
	"xfd-backend/database/db/model"
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
		req  types.OrgApplyPointReq
		resp *types.OrgApplyPointResp
		xErr xerr.XErr
	)

	comment := c.PostForm("comment")
	startTime, err := time.Parse("2006-01-02 15:04:05", c.PostForm("startTime"))
	if err != nil {
		c.JSON(http.StatusOK, response.RespError(c, xerr.WithCode(xerr.InvalidParams, err)))
		return
	}
	endTime, err := time.Parse("2006-01-02 15:04:05", c.PostForm("endTime"))
	if err != nil {
		c.JSON(http.StatusOK, response.RespError(c, xerr.WithCode(xerr.InvalidParams, err)))
		return
	}
	if !startTime.Before(endTime) {
		c.JSON(http.StatusOK, response.RespError(c, xerr.WithCode(xerr.InvalidParams, errors.New("生效时间异常"))))
		return
	}
	file, header, err := c.Request.FormFile("file")
	if err != nil {
		c.JSON(http.StatusOK, response.RespError(c, xerr.WithCode(xerr.InvalidParams, err)))
		return
	}
	req = types.OrgApplyPointReq{
		File:       file,
		FileHeader: header,
		Comment:    comment,
		StartTime:  startTime,
		EndTime:    endTime,
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
		req  types.OrgVerifyPointReq
		resp *types.OrgVerifyPointResp
		xErr xerr.XErr
	)

	err := c.BindJSON(&req)
	if err != nil {
		c.JSON(http.StatusOK, response.RespError(c, xerr.WithCode(xerr.InvalidParams, err)))
		return
	}

	if req.ID <= 0 {
		c.JSON(http.StatusOK, response.RespError(c, xerr.WithCode(xerr.InvalidParams, errors.New("invalid param"))))
		return
	}
	if req.VerifyStatus != model.PointApplicationStatusApproved && req.VerifyStatus != model.PointApplicationStatusRejected {
		c.JSON(http.StatusOK, response.RespError(c, xerr.WithCode(xerr.InvalidParams, errors.New("invalid param"))))
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
		req  types.OrgGetApplyToVerifyReq
		resp *types.OrgGetApplyToVerifyResp
		xErr xerr.XErr
	)

	err := c.BindQuery(&req)
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
		req  types.OrgGetApplysReq
		resp *types.OrgGetApplysResp
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

	resp, xErr = h.orgService.GetApplys(c, req)
	if xErr != nil {
		log.Println("[OrgHandler] GetApplys failed, err=", xErr)
		c.JSON(http.StatusOK, response.RespError(c, xErr))
		return
	}
	c.JSON(http.StatusOK, response.RespSuccess(c, resp))
}

func (h *OrgHandler) ClearPoint(c *gin.Context) {
	var (
		req  types.OrgClearPointReq
		resp *types.OrgClearPointResp
		xErr xerr.XErr
	)

	err := c.BindJSON(&req)
	if err != nil {
		c.JSON(http.StatusOK, response.RespError(c, xerr.WithCode(xerr.InvalidParams, err)))
		return
	}

	if req.OrgID <= 0 {
		c.JSON(http.StatusOK, response.RespError(c, xerr.WithCode(xerr.InvalidParams, errors.New("invalid param"))))
		return
	}

	resp, xErr = h.orgService.ClearPoint(c, req)
	if xErr != nil {
		log.Println("[OrgHandler] ClearPoint failed, err=", xErr)
		c.JSON(http.StatusOK, response.RespError(c, xErr))
		return
	}
	c.JSON(http.StatusOK, response.RespSuccess(c, resp))
}

func (h *OrgHandler) VerifyAccount(c *gin.Context) {
	var (
		req  types.VerifyAccountReq
		resp *types.VerifyAccountResp
		xErr xerr.XErr
	)

	err := c.BindJSON(&req)
	if err != nil {
		c.JSON(http.StatusOK, response.RespError(c, xerr.WithCode(xerr.InvalidParams, err)))
		return
	}
	if req.ID == 0 {
		c.JSON(http.StatusOK, response.RespError(c, xerr.WithCode(xerr.InvalidParams, errors.New("invalid param"))))
		return
	}
	if req.Status != model.UserVerifyStatusRejected && req.Status != model.UserVerifyStatusSuccess {
		c.JSON(http.StatusOK, response.RespError(c, xerr.WithCode(xerr.InvalidParams, errors.New("invalid param"))))
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
		req  types.GetAccountToVerifyReq
		resp *types.GetAccountToVerifyResp
		xErr xerr.XErr
	)

	err := c.BindQuery(&req)
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

func (h *OrgHandler) GetAccountVerifyList(c *gin.Context) {
	var (
		req  types.GetAccountVerifyListReq
		resp *types.GetAccountVerifyListResp
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
	resp, xErr = h.orgService.GetAccountVerifyList(c, req)
	if xErr != nil {
		log.Println("[OrgHandler] GetAccountVerifyList failed, err=", xErr)
		c.JSON(http.StatusOK, response.RespError(c, xErr))
		return
	}
	c.JSON(http.StatusOK, response.RespSuccess(c, resp))
}

func (h *OrgHandler) GetOrganizations(c *gin.Context) {
	var (
		req  types.GetOrganizationsReq
		resp *types.GetOrganizationsResp
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
	resp, xErr = h.orgService.GetOrganizations(c, req)
	if xErr != nil {
		log.Println("[OrgHandler] GetOrganizations failed, err=", xErr)
		c.JSON(http.StatusOK, response.RespError(c, xErr))
		return
	}
	c.JSON(http.StatusOK, response.RespSuccess(c, resp))
}

func (h *OrgHandler) GetOrgMembers(c *gin.Context) {
	var (
		req  types.GetOrgMembersReq
		resp *types.GetOrgMembersResp
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
	resp, xErr = h.orgService.GetOrgMembers(c, req)
	if xErr != nil {
		log.Println("[OrgHandler] GetOrgMembers failed, err=", xErr)
		c.JSON(http.StatusOK, response.RespError(c, xErr))
		return
	}
	c.JSON(http.StatusOK, response.RespSuccess(c, resp))
}

func (h *OrgHandler) GetPointRecordsByApply(c *gin.Context) {
	var (
		req  types.GetPointRecordsByApplyReq
		resp *types.GetPointRecordsByApplyResp
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
	if req.ApplyID == 0 {
		c.JSON(http.StatusOK, response.RespError(c, xerr.WithCode(xerr.InvalidParams, errors.New("invalid param"))))
		return
	}
	resp, xErr = h.orgService.GetPointRecordsByApply(c, req)
	if xErr != nil {
		log.Println("[OrgHandler] GetPointRecordsByApply failed, err=", xErr)
		c.JSON(http.StatusOK, response.RespError(c, xErr))
		return
	}
	c.JSON(http.StatusOK, response.RespSuccess(c, resp))
}

func (h *OrgHandler) GetPointRecordsByUser(c *gin.Context) {
	var (
		req  types.GetPointRecordsByUserReq
		resp *types.GetPointRecordsByUserResp
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
	if len(req.UserID) == 0 {
		c.JSON(http.StatusOK, response.RespError(c, xerr.WithCode(xerr.InvalidParams, errors.New("invalid param"))))
		return
	}
	resp, xErr = h.orgService.GetPointRecordsByUser(c, req)
	if xErr != nil {
		log.Println("[OrgHandler] GetPointRecordsByUser failed, err=", xErr)
		c.JSON(http.StatusOK, response.RespError(c, xErr))
		return
	}
	c.JSON(http.StatusOK, response.RespSuccess(c, resp))
}

func (h *OrgHandler) GetPointRecords(c *gin.Context) {
	var (
		req  types.GetPointRecordsReq
		resp *types.GetPointRecordsResp
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
	if req.OrgID <= 0 {
		c.JSON(http.StatusOK, response.RespError(c, xerr.WithCode(xerr.InvalidParams, errors.New("invalid param"))))
		return
	}
	resp, xErr = h.orgService.GetPointRecords(c, req)
	if xErr != nil {
		log.Println("[OrgHandler] GetPointRecords failed, err=", xErr)
		c.JSON(http.StatusOK, response.RespError(c, xErr))
		return
	}
	c.JSON(http.StatusOK, response.RespSuccess(c, resp))
}
