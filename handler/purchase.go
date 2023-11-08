package handler

import (
	"errors"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"xfd-backend/pkg/response"
	"xfd-backend/pkg/types"
	"xfd-backend/pkg/xerr"
	"xfd-backend/service"
)

type PurchaseHandler struct {
	purchaseService *service.PurchaseService
}

func NewPurchaseHandler() *PurchaseHandler {
	return &PurchaseHandler{purchaseService: service.NewPurchaseService()}
}

func (h *PurchaseHandler) GetPurchases(c *gin.Context) {
	var (
		req  types.PurchaseGetPurchasesReq
		resp *types.PurchaseGetPurchasesResp
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

	resp, xErr = h.purchaseService.GetPurchases(c, req)
	if xErr != nil {
		log.Println("[PurchaseHandler] GetPurchases failed, err=", xErr)
		c.JSON(http.StatusOK, response.RespError(c, xErr))
		return
	}
	c.JSON(http.StatusOK, response.RespSuccess(c, resp))
}

func (h *PurchaseHandler) SubmitOrder(c *gin.Context) {
	var (
		req  types.PurchaseSubmitOrderReq
		resp *types.PurchaseSubmitOrderResp
		xErr xerr.XErr
	)

	err := c.BindJSON(&req)
	if err != nil {
		c.JSON(http.StatusOK, response.RespError(c, xerr.WithCode(xerr.InvalidParams, err)))
		return
	}

	if req.Period == 0 || req.Quantity == 0 || req.Unit == "" || len(req.CategoryName) == 0 {
		c.JSON(http.StatusOK, response.RespError(c, xerr.WithCode(xerr.InvalidParams, errors.New("invalid param"))))
		return
	}
	if req.CategoryA == 0 || req.CategoryB == 0 {
		c.JSON(http.StatusOK, response.RespError(c, xerr.WithCode(xerr.InvalidParams, errors.New("invalid param"))))
		return
	}

	resp, xErr = h.purchaseService.SubmitOrder(c, req)
	if xErr != nil {
		log.Println("[PurchaseHandler] SubmitOrder failed, err=", xErr)
		c.JSON(http.StatusOK, response.RespError(c, xErr))
		return
	}
	c.JSON(http.StatusOK, response.RespSuccess(c, resp))
}

//func (h *PurchaseHandler) ModifyOrder(c *gin.Context) {
//	var (
//		req types.PurchaseModifyOrderReq
//		resp *types.PurchaseModifyOrderResp
//		xErr xerr.XErr
//	)
//
//	err := c.BindJSON(&req)
//	if err != nil {
//		c.JSON(http.StatusOK, response.RespError(c, xerr.WithCode(xerr.InvalidParams, err)))
//		return
//	}
//	if req.Period == 0 || req.Quantity == 0 || req.Unit == "" || req.AreaCodeID == 0 {
//		c.JSON(http.StatusOK, response.RespError(c, xerr.WithCode(xerr.InvalidParams, err)))
//		return
//	}
//	if req.CategoryA == 0 || req.CategoryB == 0 {
//		c.JSON(http.StatusOK, response.RespError(c, xerr.WithCode(xerr.InvalidParams, err)))
//		return
//	}
//	resp, xErr = h.purchaseService.ModifyOrder(c, req)
//	if xErr != nil {
//		log.Println("[PurchaseHandler] ModifyOrder failed, err=", xErr)
//		c.JSON(http.StatusOK, response.RespError(c, xErr))
//		return
//	}
//	c.JSON(http.StatusOK, response.RespSuccess(c, resp))
//}

func (h *PurchaseHandler) ModifyOrderStatus(c *gin.Context) {
	var (
		req  types.PurchaseModifyOrderStatusReq
		resp *types.PurchaseModifyOrderStatusResp
		xErr xerr.XErr
	)

	err := c.BindJSON(&req)
	if err != nil {
		c.JSON(http.StatusOK, response.RespError(c, xerr.WithCode(xerr.InvalidParams, err)))
		return
	}
	if req.OrderID == 0 {
		c.JSON(http.StatusOK, response.RespError(c, xerr.WithCode(xerr.InvalidParams, errors.New("invalid param"))))
		return
	}
	resp, xErr = h.purchaseService.ModifyOrderStatus(c, req)
	if xErr != nil {
		log.Println("[PurchaseHandler] ModifyOrderStatus failed, err=", xErr)
		c.JSON(http.StatusOK, response.RespError(c, xErr))
		return
	}
	c.JSON(http.StatusOK, response.RespSuccess(c, resp))
}

func (h *PurchaseHandler) GetQuotes(c *gin.Context) {
	var (
		req  types.PurchaseGetQuotesReq
		resp *types.PurchaseGetQuotesResp
		xErr xerr.XErr
	)

	err := c.BindQuery(&req)
	if err != nil {
		c.JSON(http.StatusOK, response.RespError(c, xerr.WithCode(xerr.InvalidParams, err)))
		return
	}
	if req.OrderID == 0 {
		c.JSON(http.StatusOK, response.RespError(c, xerr.WithCode(xerr.InvalidParams, errors.New("invalid param"))))
		return
	}
	if err = req.CheckParams(); err != nil {
		c.JSON(http.StatusOK, response.RespError(c, xerr.WithCode(xerr.InvalidParams, err)))
		return
	}
	resp, xErr = h.purchaseService.GetQuotes(c, req)
	if xErr != nil {
		log.Println("[PurchaseHandler] GetQuotes failed, err=", xErr)
		c.JSON(http.StatusOK, response.RespError(c, xErr))
		return
	}
	c.JSON(http.StatusOK, response.RespSuccess(c, resp))
}

func (h *PurchaseHandler) GetStatistics(c *gin.Context) {
	var (
		req  types.PurchaseGetStatisticsReq
		resp *types.PurchaseGetStatisticsResp
		xErr xerr.XErr
	)

	err := c.BindJSON(&req)
	if err != nil {
		c.JSON(http.StatusOK, response.RespError(c, xerr.WithCode(xerr.InvalidParams, err)))
		return
	}
	resp, xErr = h.purchaseService.GetStatistics(c, req)
	if xErr != nil {
		log.Println("[PurchaseHandler] GetStatistics failed, err=", xErr)
		c.JSON(http.StatusOK, response.RespError(c, xErr))
		return
	}
	c.JSON(http.StatusOK, response.RespSuccess(c, resp))
}

func (h *PurchaseHandler) AnswerQuote(c *gin.Context) {
	var (
		req  types.PurchaseAnswerQuoteReq
		resp *types.PurchaseAnswerQuoteResp
		xErr xerr.XErr
	)

	err := c.BindJSON(&req)
	if err != nil {
		c.JSON(http.StatusOK, response.RespError(c, xerr.WithCode(xerr.InvalidParams, err)))
		return
	}
	resp, xErr = h.purchaseService.AnswerQuote(c, req)
	if xErr != nil {
		log.Println("[PurchaseHandler] AnswerQuote failed, err=", xErr)
		c.JSON(http.StatusOK, response.RespError(c, xErr))
		return
	}
	c.JSON(http.StatusOK, response.RespSuccess(c, resp))
}
