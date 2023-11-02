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

type CommonHandler struct {
	commonService *service.CommonService
}

func NewCommonHandler() *CommonHandler {
	return &CommonHandler{commonService: service.NewCommonService()}
}

func (h *CommonHandler) Upload(c *gin.Context) {
	var (
		req  *types.CommonUploadReq
		resp *types.CommonUploadResp
		xErr xerr.XErr
	)

	err := c.BindJSON(&req)
	if err != nil {
		c.JSON(http.StatusOK, response.RespError(c, xerr.WithCode(xerr.InvalidParams, err)))
		return
	}

	resp, xErr = h.commonService.UploadToOSS(c, req)
	if xErr != nil {
		log.Println("[CommonHandler] UploadToOSS failed, err=", xErr)
		c.JSON(http.StatusOK, response.RespError(c, xErr))
		return
	}
	c.JSON(http.StatusOK, response.RespSuccess(c, resp))
}
