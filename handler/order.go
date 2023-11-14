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

type OrderHandler struct {
	orderService *service.OrderService
}

func NewOrderHandler() *OrderHandler {
	return &OrderHandler{orderService: service.NewOrderService()}
}

func (h *OrderHandler) AddShoppingCart(context *gin.Context) {
	var (
		req  types.ShoppingCartAddReq
		xErr xerr.XErr
	)
	err := context.ShouldBindJSON(&req)
	if err != nil {
		context.JSON(http.StatusOK, response.RespError(context, xerr.WithCode(xerr.InvalidParams, err)))
		return
	}
	xErr = h.orderService.AddShoppingCart(context, req)
	if xErr != nil {
		log.Println("[OrderHandler] AddShoppingCart failed, err=", xErr)
		context.JSON(http.StatusOK, response.RespError(context, xErr))
		return
	}
	context.JSON(http.StatusOK, response.RespSuccess(context, nil))
}

func (h *OrderHandler) GetShoppingCartList(context *gin.Context) {
	var (
		req  types.ShoppingCartListReq
		resp *types.ShoppingCartListResp
		xErr xerr.XErr
	)
	err := context.ShouldBindQuery(&req)
	if err != nil {
		context.JSON(http.StatusOK, response.RespError(context, xerr.WithCode(xerr.InvalidParams, err)))
		return
	}
	resp, xErr = h.orderService.GetShoppingCartList(context, req)
	if xErr != nil {
		log.Println("[OrderHandler] GetShoppingCartList failed, err=", xErr)
		context.JSON(http.StatusOK, response.RespError(context, xErr))
		return
	}
	context.JSON(http.StatusOK, response.RespSuccess(context, resp))

}

func (h *OrderHandler) DeleteShoppingCart(context *gin.Context) {
	var (
		req  types.ShoppingCartDeleteReq
		xErr xerr.XErr
	)
	err := context.ShouldBindJSON(&req)
	if err != nil {
		context.JSON(http.StatusOK, response.RespError(context, xerr.WithCode(xerr.InvalidParams, err)))
		return
	}
	xErr = h.orderService.DeleteShoppingCart(context, req)
	if xErr != nil {
		log.Println("[OrderHandler] DeleteShoppingCart failed, err=", xErr)
		context.JSON(http.StatusOK, response.RespError(context, xErr))
		return
	}
	context.JSON(http.StatusOK, response.RespSuccess(context, nil))
}

func (h *OrderHandler) ModifyShoppingCart(context *gin.Context) {
	var (
		req types.ShoppingCartModifyReq
		xrr xerr.XErr
	)
	err := context.ShouldBindJSON(&req)
	if err != nil {
		context.JSON(http.StatusOK, response.RespError(context, xerr.WithCode(xerr.InvalidParams, err)))
		return
	}
	xrr = h.orderService.ModifyShoppingCart(context, req)
	if xrr != nil {
		log.Println("[OrderHandler] ModifyShoppingCart failed, err=", xrr)
		context.JSON(http.StatusOK, response.RespError(context, xrr))
		return
	}
	context.JSON(http.StatusOK, response.RespSuccess(context, nil))
}
