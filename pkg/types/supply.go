package types

import "xfd-backend/database/db/model"

type SupplyGetPurchasesReq struct {
	PageRequest
	CategoryA int    `json:"categoryA"`
	CategoryB int    `json:"categoryB"`
	CategoryC int    `json:"categoryC"`
	Like      string `json:"like"`
}

type SupplyGetPurchasesResp struct {
	List     []*PurchaseOrder `json:"list"`
	TotalNum int64            `json:"totalNum"`
}

type SupplyGetQuotesReq struct {
	OrderID int `json:"orderID"`
}

type SupplyGetQuotesResp struct {
	List []*PurchaseQuote `json:"list"`
}

type SupplySubmitQuoteReq struct {
	OrderID int     `json:"orderID"`
	ItemID  int     `json:"itemID"`
	Price   float64 `json:"price"`
}

type SupplySubmitQuoteResp struct {
}

type SupplyGetQuotedPurchasesReq struct {
	PageRequest
	Status model.OrderPurchaseStatus `json:"status"`
}

type SupplyGetQuotedPurchasesResp struct {
	List     []*PurchaseOrder `json:"list"`
	TotalNum int64            `json:"totalNum"`
}

type SupplyGetStatisticsReq struct {
}

type SupplyGetStatisticsResp struct {
	NewPurchase int `json:"newPurchase"`
}
