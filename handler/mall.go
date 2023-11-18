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
}

func NewMallHandler() *MallHandler {
	return &MallHandler{mallService: service.NewMallService()}
}

// GetCategories 获取商城分类
func (h *MallHandler) GetCategories(c *gin.Context) {
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

func (h *MallHandler) DeleteCategory(context *gin.Context) {
	var (
		req types.CategoryDeleteReq
		xrr xerr.XErr
	)
	err := context.ShouldBindJSON(&req)
	if err != nil {
		context.JSON(http.StatusOK, response.RespError(context, xerr.WithCode(xerr.InvalidParams, err)))
		return
	}
	xrr = h.mallService.DeleteCategory(context, req)
	if xrr != nil {
		log.Println("[MallHandler] DeleteCategory failed, err=", xrr)
		context.JSON(http.StatusOK, response.RespError(context, xrr))
		return
	}
	context.JSON(http.StatusOK, response.RespSuccess(context, nil))
}

func (h *MallHandler) ModifyCategory(context *gin.Context) {
	var (
		req types.CategoryModifyReq
		xrr xerr.XErr
	)
	err := context.ShouldBindJSON(&req)
	if err != nil {
		context.JSON(http.StatusOK, response.RespError(context, xerr.WithCode(xerr.InvalidParams, err)))
		return
	}
	xrr = h.mallService.ModifyCategory(context, req)
	if xrr != nil {
		log.Println("[MallHandler] ModifyCategory failed, err=", xrr)
		context.JSON(http.StatusOK, response.RespError(context, xrr))
		return
	}
	context.JSON(http.StatusOK, response.RespSuccess(context, nil))
}

func (h *MallHandler) AddCategory(context *gin.Context) {
	var (
		req types.CategoryAddReq
		xrr xerr.XErr
	)
	err := context.ShouldBindJSON(&req)
	if err != nil {
		context.JSON(http.StatusOK, response.RespError(context, xerr.WithCode(xerr.InvalidParams, err)))
		return
	}
	xrr = h.mallService.AddCategory(context, req)
	if xrr != nil {
		log.Println("[MallHandler] AddCategory failed, err=", xrr)
		context.JSON(http.StatusOK, response.RespError(context, xrr))
		return
	}
	context.JSON(http.StatusOK, response.RespSuccess(context, nil))
}
