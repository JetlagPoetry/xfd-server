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

type GoodsHandler struct {
	goodsService *service.GoodsService
}

func NewGoodsHandler() *GoodsHandler {
	return &GoodsHandler{goodsService: service.NewGoodsService()}
}

func (h *GoodsHandler) AddGoods(c *gin.Context) {
	var (
		req  types.GoodsAddReq
		xErr xerr.XErr
	)
	err := c.ShouldBindJSON(&req)
	if err != nil {
		c.JSON(http.StatusOK, response.RespError(c, xerr.WithCode(xerr.InvalidParams, err)))
		return
	}
	xErr = h.goodsService.AddGoods(c, req)
	if xErr != nil {
		log.Println("[GoodsHandler] AddGoods failed, err=", xErr)
		c.JSON(http.StatusOK, response.RespError(c, xErr))
		return
	}
	c.JSON(http.StatusOK, response.RespSuccess(c, nil))
}

func (h *GoodsHandler) GetGoodsList(c *gin.Context) {
	var (
		req  types.GoodsListReq
		resp *types.GoodsListResp
		xErr xerr.XErr
	)
	err := c.ShouldBindQuery(&req)
	if err != nil {
		c.JSON(http.StatusOK, response.RespError(c, xerr.WithCode(xerr.InvalidParams, err)))
		return
	}
	resp, xErr = h.goodsService.GetGoodsList(c, req)
	if xErr != nil {
		log.Println("[GoodsHandler] GetGoodsList failed, err=", xErr)
		c.JSON(http.StatusOK, response.RespError(c, xErr))
		return
	}
	c.JSON(http.StatusOK, response.RespSuccess(c, resp))
}

func (h *GoodsHandler) GetGoodsDetail(c *gin.Context) {

}

//后台操作

func (h *GoodsHandler) GetMyGoodsList(c *gin.Context) {

}

func (h *GoodsHandler) GetMyGoodsDetail(c *gin.Context) {

}

func (h *GoodsHandler) ModifyMyGoods(c *gin.Context) {

}

func (h *GoodsHandler) DeleteMyGoods(c *gin.Context) {
	var (
		req  types.GoodsReq
		xErr xerr.XErr
	)
	err := c.ShouldBindJSON(&req)
	if err != nil {
		c.JSON(http.StatusOK, response.RespError(c, xerr.WithCode(xerr.InvalidParams, err)))
		return
	}
	xErr = h.goodsService.DeleteMyGoods(c, req)
	if xErr != nil {
		log.Println("[GoodsHandler] DeleteMyGoods failed, err=", xErr)
		c.JSON(http.StatusOK, response.RespError(c, xErr))
		return
	}
	c.JSON(http.StatusOK, response.RespSuccess(c, nil))

}

func (h *GoodsHandler) ModifyMyGoodsStatus(c *gin.Context) {
	var (
		req  types.GoodsReq
		xErr xerr.XErr
	)
	err := c.ShouldBindJSON(&req)
	if err != nil {
		c.JSON(http.StatusOK, response.RespError(c, xerr.WithCode(xerr.InvalidParams, err)))
		return
	}
	xErr = h.goodsService.ModifyMyGoodsStatus(c, req)
	if xErr != nil {
		log.Println("[GoodsHandler] ModifyMyGoodsStatus failed, err=", xErr)
		c.JSON(http.StatusOK, response.RespError(c, xErr))
		return
	}
	c.JSON(http.StatusOK, response.RespSuccess(c, nil))
}