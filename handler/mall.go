package handler

import (
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"xfd-backend/database/db/model"
	"xfd-backend/pkg/response"
	"xfd-backend/pkg/types"
	"xfd-backend/pkg/xerr"
	"xfd-backend/service"
)

type MallHandler struct {
	mallService *service.MallService
	areaService *service.AreaService
}

func NewMallHandler() *MallHandler {
	return &MallHandler{mallService: service.NewMallService(), areaService: service.NewAreaService()}
}

func (h MallHandler) GetCategories(c *gin.Context) {
	var (
		req  types.CategoryListReq
		resp []*model.Category
		xErr xerr.XErr
	)
	err := c.ShouldBindQuery(&req)
	if err != nil {
		c.JSON(http.StatusOK, response.RespError(c, xerr.WithCode(xerr.InvalidParams, err)))
		return
	}
	resp, xErr = h.mallService.GetCategories(c, req)
	if xErr != nil {
		log.Println("[MallHandler] GetCategories failed, err=", xErr)
		c.JSON(http.StatusOK, response.RespError(c, xErr))
		return
	}
	c.JSON(http.StatusOK, response.RespSuccess(c, resp))
}

func (h MallHandler) GetArea(c *gin.Context) {
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
		log.Println("[MallHandler] GetCategories failed, err=", xErr)
		c.JSON(http.StatusOK, response.RespError(c, xErr))
		return
	}
	c.JSON(http.StatusOK, response.RespSuccess(c, resp))
}
