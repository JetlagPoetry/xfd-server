package handler

import (
	"errors"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"path/filepath"
	"xfd-backend/pkg/response"
	"xfd-backend/pkg/types"
	"xfd-backend/pkg/utils"
	"xfd-backend/pkg/xerr"
	"xfd-backend/service"
)

type CommonHandler struct {
	commonService *service.CommonService
	areaService   *service.AreaService
}

func NewCommonHandler() *CommonHandler {
	return &CommonHandler{commonService: service.NewCommonService(), areaService: service.NewAreaService()}
}

const (
	BucketNameT       = "xfd-t"
	DefaultFolderName = "temp"
)

func (h *CommonHandler) GetArea(c *gin.Context) {
	var (
		req  types.AreaReq
		resp []*types.AreaList
		xErr xerr.XErr
	)
	err := c.ShouldBindQuery(&req)
	if err != nil {
		c.JSON(http.StatusOK, response.RespError(c, xerr.WithCode(xerr.InvalidParams, err)))
		return
	}
	resp, xErr = h.areaService.GetAreaInfo(c, req)
	if xErr != nil {
		log.Println("[CommonHandler] GetArea failed, err=", xErr)
		c.JSON(http.StatusOK, response.RespError(c, xErr))
		return
	}
	c.JSON(http.StatusOK, response.RespSuccess(c, resp))
}

func (h *CommonHandler) UploadFile(c *gin.Context) {
	var resp *types.CommonUploadResp
	file, header, err := c.Request.FormFile("file")
	if err != nil || file == nil {
		c.JSON(http.StatusOK, response.RespError(c, xerr.WithCode(xerr.InvalidParams, err)))
		return
	}
	fileSize := header.Size
	maxSize := int64(50 * 1024 * 1024) // 50MB
	if fileSize > maxSize {
		c.JSON(http.StatusOK, response.RespError(c, xerr.WithCode(xerr.InvalidParams, errors.New("文件大小超过限制"))))
	}
	folderName := c.Request.FormValue("folderName")
	if folderName == "" {
		folderName = DefaultFolderName
	}
	link, err := utils.Upload(c, BucketNameT, folderName+"/"+utils.GenerateFileName()+filepath.Ext(header.Filename), file)
	if err != nil {
		c.JSON(http.StatusOK, response.RespError(c, xerr.WithCode(xerr.ErrorUploadFile, err)))
		return
	}
	resp = &types.CommonUploadResp{Link: link}
	c.JSON(http.StatusOK, response.RespSuccess(c, resp))
}

func (h *CommonHandler) DeleteUploadFile(c *gin.Context) {
	var req types.CommonDeleteUploadReq
	err := c.BindQuery(&req)
	if err != nil {
		c.JSON(http.StatusOK, response.RespError(c, xerr.WithCode(xerr.InvalidParams, err)))
		return
	}
	rr := utils.Delete(req.Link)
	if rr != nil {
		c.JSON(http.StatusOK, response.RespError(c, xerr.WithCode(xerr.ErrorDeleteFile, rr)))
		return
	}
	c.JSON(http.StatusOK, response.RespSuccess(c, nil))
}
