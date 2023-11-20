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

func (h *OrderHandler) CreateOrder(context *gin.Context) {
	var (
		req  types.CreateOrderReq
		resp *types.CreateOrderResp
		xrr  xerr.XErr
	)
	err := context.ShouldBindJSON(&req)
	if err != nil {
		context.JSON(http.StatusOK, response.RespError(context, xerr.WithCode(xerr.InvalidParams, err)))
		return
	}

	resp, xrr = h.orderService.CreateOrder(context, req)
	if xrr != nil {
		log.Println("[OrderHandler] CreateOrder failed, err=", xrr)
		context.JSON(http.StatusOK, response.RespError(context, xrr))
		return
	}
	context.JSON(http.StatusOK, response.RespSuccess(context, resp))
}

func (h *OrderHandler) ApplyRefund(context *gin.Context) {
	var (
		req  types.ApplyRefundReq
		resp *types.ApplyRefundResp
		xrr  xerr.XErr
	)
	err := context.ShouldBindJSON(&req)
	if err != nil {
		context.JSON(http.StatusOK, response.RespError(context, xerr.WithCode(xerr.InvalidParams, err)))
		return
	}
	resp, xrr = h.orderService.ApplyRefund(context, req)
	if xrr != nil {
		log.Println("[OrderHandler] ApplyRefund failed, err=", xrr)
		context.JSON(http.StatusOK, response.RespError(context, xrr))
		return
	}
	context.JSON(http.StatusOK, response.RespSuccess(context, resp))
}

func (h *OrderHandler) ApplyExchange(ctx *gin.Context) {
	var (
		req  types.ApplyExchangeReq
		resp *types.ApplyExchangeResp
		xrr  xerr.XErr
	)
	err := ctx.ShouldBindJSON(&req)
	if err != nil {
		ctx.JSON(http.StatusOK, response.RespError(ctx, xerr.WithCode(xerr.InvalidParams, err)))
		return
	}
	resp, xrr = h.orderService.ApplyExchange(ctx, req)
	if xrr != nil {
		log.Println("[OrderHandler] ApplyExchange failed, err=", xrr)
		ctx.JSON(http.StatusOK, response.RespError(ctx, xrr))
		return
	}
	ctx.JSON(http.StatusOK, response.RespSuccess(ctx, resp))

}

func (h *OrderHandler) PayOrder(context *gin.Context) {
	var (
		req  types.PayOrderReq
		resp *types.PayOrderResp
		xrr  xerr.XErr
	)
	err := context.ShouldBindJSON(&req)
	if err != nil {
		context.JSON(http.StatusOK, response.RespError(context, xerr.WithCode(xerr.InvalidParams, err)))
		return
	}
	resp, xrr = h.orderService.PayOrder(context, req)
	if xrr != nil {
		log.Println("[OrderHandler] PayOrder failed, err=", xrr)
		context.JSON(http.StatusOK, response.RespError(context, xrr))
		return
	}
	context.JSON(http.StatusOK, response.RespSuccess(context, resp))
}

func (h *OrderHandler) CreatePreOrder(context *gin.Context) {
	var (
		req  types.CreateOrderReq
		resp *types.CreatePreOrderResp
		xrr  xerr.XErr
	)
	err := context.ShouldBindJSON(&req)
	if err != nil {
		context.JSON(http.StatusOK, response.RespError(context, xerr.WithCode(xerr.InvalidParams, err)))
		return
	}
	resp, xrr = h.orderService.CreatePreOrder(context, req)
	if xrr != nil {
		log.Println("[OrderHandler] CreatePreOrder failed, err=", xrr)
		context.JSON(http.StatusOK, response.RespError(context, xrr))
		return
	}
	context.JSON(http.StatusOK, response.RespSuccess(context, resp))
}

// GetOrderList 获取订单列表
func (h *OrderHandler) GetOrderList(context *gin.Context) {
	//消费者 待发货，已发货，全部
	//供应商 待发货，已发货，全部
	//官方 待发货，待收货，已签收，售后/结束  全部
	var (
		req  types.OrderListReq
		resp *types.OrderListResp
		xrr  xerr.XErr
	)
	err := context.ShouldBindQuery(&req)
	if err != nil {
		context.JSON(http.StatusOK, response.RespError(context, xerr.WithCode(xerr.InvalidParams, err)))
		return
	}
	resp, xrr = h.orderService.GetOrderList(context, req)
	if xrr != nil {
		log.Println("[OrderHandler] GetOrderList failed, err=", xrr)
		context.JSON(http.StatusOK, response.RespError(context, xrr))
		return
	}
	context.JSON(http.StatusOK, response.RespSuccess(context, resp))
}

func (h *OrderHandler) FillShipmentInfo(context *gin.Context) {
	var (
		req types.FillShipmentInfoReq
		xrr xerr.XErr
	)
	err := context.ShouldBindJSON(&req)
	if err != nil {
		context.JSON(http.StatusOK, response.RespError(context, xerr.WithCode(xerr.InvalidParams, err)))
		return
	}
	xrr = h.orderService.FillShipmentInfo(context, req)
	if xrr != nil {
		log.Println("[OrderHandler] FillShipmentInfo failed, err=", xrr)
		context.JSON(http.StatusOK, response.RespError(context, xrr))
		return
	}
	context.JSON(http.StatusOK, response.RespSuccess(context, nil))
}

func (h *OrderHandler) ConfirmReceipt(ctx *gin.Context) {
	var (
		req types.ConfirmReceiptReq
		xrr xerr.XErr
	)
	err := ctx.ShouldBindJSON(&req)
	if err != nil {
		ctx.JSON(http.StatusOK, response.RespError(ctx, xerr.WithCode(xerr.InvalidParams, err)))
		return
	}
	xrr = h.orderService.ConfirmReceipt(ctx, req)
	if xrr != nil {
		log.Println("[OrderHandler] ConfirmReceipt failed, err=", xrr)
		ctx.JSON(http.StatusOK, response.RespError(ctx, xrr))
		return
	}
	ctx.JSON(http.StatusOK, response.RespSuccess(ctx, nil))
}
