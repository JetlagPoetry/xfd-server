package types

import (
	"xfd-backend/database/db/model"
)

type SupplyGetPurchasesReq struct {
	PageRequest
	CategoryA int    `form:"categoryA"`
	CategoryB int    `form:"categoryB"`
	CategoryC int    `form:"categoryC"`
	Like      string `form:"like"`
}

type SupplyGetPurchasesResp struct {
	List     []*PurchaseOrder `json:"list"`
	TotalNum int64            `json:"totalNum"`
}

type SupplyGetQuotesReq struct {
	OrderID int `form:"orderID"`
}

type SupplyGetQuotesResp struct {
	List []*PurchaseQuote `json:"list"`
}

type SupplySubmitQuoteReq struct {
	OrderID int    `json:"orderID"`
	GoodsID int    `json:"goodsID"`
	Price   string `json:"price"`
}

type SupplySubmitQuoteResp struct {
}

type SupplyGetQuotedPurchasesReq struct {
	PageRequest
	Status model.OrderPurchaseStatus `form:"status"`
}

type SupplyGetQuotedPurchasesResp struct {
	List     []*PurchaseOrder `json:"list"`
	TotalNum int64            `json:"totalNum"`
}

type SupplyGetStatisticsReq struct {
}

type SupplyGetStatisticsResp struct {
	NewPurchase           int `json:"newPurchase"`
	NewWaitingForDelivery int `json:"newWaitingForDelivery"`
}

type SupplyAnswerQuoteReq struct {
	PurchaseUserID string `json:"purchaseUserID"`
}

type SupplyAnswerQuoteResp struct {
}
