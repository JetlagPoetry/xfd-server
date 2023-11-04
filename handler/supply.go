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

type SupplyHandler struct {
	supplyService *service.SupplyService
}

func NewSupplyHandler() *SupplyHandler {
	return &SupplyHandler{supplyService: service.NewSupplyService()}
}

func (h *SupplyHandler) GetPurchases(c *gin.Context) {
	var (
		req  types.SupplyGetPurchasesReq
		resp *types.SupplyGetPurchasesResp
		xErr xerr.XErr
	)

	err := c.BindQuery(&req)
	if err != nil {
		c.JSON(http.StatusOK, response.RespError(c, xerr.WithCode(xerr.InvalidParams, err)))
		return
	}
	resp, xErr = h.supplyService.GetPurchases(c, req)
	if xErr != nil {
		log.Println("[SupplyHandler] GetPurchases failed, err=", xErr)
		c.JSON(http.StatusOK, response.RespError(c, xErr))
		return
	}
	c.JSON(http.StatusOK, response.RespSuccess(c, resp))
}

func (h *SupplyHandler) GetQuotes(c *gin.Context) {
	var (
		req  types.SupplyGetQuotesReq
		resp *types.SupplyGetQuotesResp
		xErr xerr.XErr
	)

	err := c.BindQuery(&req)
	if err != nil {
		c.JSON(http.StatusOK, response.RespError(c, xerr.WithCode(xerr.InvalidParams, err)))
		return
	}
	resp, xErr = h.supplyService.GetQuotes(c, req)
	if xErr != nil {
		log.Println("[SupplyHandler] GetQuotes failed, err=", xErr)
		c.JSON(http.StatusOK, response.RespError(c, xErr))
		return
	}
	c.JSON(http.StatusOK, response.RespSuccess(c, resp))
}

func (h *SupplyHandler) SubmitQuote(c *gin.Context) {
	var (
		req  types.SupplySubmitQuoteReq
		resp *types.SupplySubmitQuoteResp
		xErr xerr.XErr
	)

	err := c.BindJSON(&req)
	if err != nil {
		c.JSON(http.StatusOK, response.RespError(c, xerr.WithCode(xerr.InvalidParams, err)))
		return
	}
	if req.OrderID == 0 || req.ItemID == 0 || req.Price == 0 {
		c.JSON(http.StatusOK, response.RespError(c, xerr.WithCode(xerr.InvalidParams, err)))
		return
	}
	resp, xErr = h.supplyService.SubmitQuote(c, req)
	if xErr != nil {
		log.Println("[SupplyHandler] SubmitQuote failed, err=", xErr)
		c.JSON(http.StatusOK, response.RespError(c, xErr))
		return
	}
	c.JSON(http.StatusOK, response.RespSuccess(c, resp))
}

func (h *SupplyHandler) GetQuotedPurchases(c *gin.Context) {
	var (
		req  types.SupplyGetQuotedPurchasesReq
		resp *types.SupplyGetQuotedPurchasesResp
		xErr xerr.XErr
	)

	err := c.BindQuery(&req)
	if err != nil {
		c.JSON(http.StatusOK, response.RespError(c, xerr.WithCode(xerr.InvalidParams, err)))
		return
	}
	resp, xErr = h.supplyService.GetQuotedPurchases(c, req)
	if xErr != nil {
		log.Println("[SupplyHandler] GetQuotedPurchases failed, err=", xErr)
		c.JSON(http.StatusOK, response.RespError(c, xErr))
		return
	}
	c.JSON(http.StatusOK, response.RespSuccess(c, resp))
}

func (h *SupplyHandler) GetStatistics(c *gin.Context) {
	var (
		req  types.SupplyGetStatisticsReq
		resp *types.SupplyGetStatisticsResp
		xErr xerr.XErr
	)

	err := c.BindQuery(&req)
	if err != nil {
		c.JSON(http.StatusOK, response.RespError(c, xerr.WithCode(xerr.InvalidParams, err)))
		return
	}
	resp, xErr = h.supplyService.GetStatistics(c, req)
	if xErr != nil {
		log.Println("[SupplyHandler] GetStatistics failed, err=", xErr)
		c.JSON(http.StatusOK, response.RespError(c, xErr))
		return
	}
	c.JSON(http.StatusOK, response.RespSuccess(c, resp))
}
