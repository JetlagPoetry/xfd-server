package types

import "xfd-backend/database/db/model"

type SupplyGetPurchasesReq struct {
	PageRequest
	Status model.OrderPurchaseStatus `json:"status"`
}

type SupplyGetPurchasesResp struct {
	List     []*PurchaseOrder `json:"list"`
	TotalNum int64            `json:"totalNum"`
}

type SupplyGetQuotesReq struct {
	PageRequest
	OrderID int `json:"orderID"`
}

type SupplyGetQuotesResp struct {
	List     []*PurchaseQuote `json:"list"`
	TotalNum int64            `json:"totalNum"`
}
