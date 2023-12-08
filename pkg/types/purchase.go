package types

import "xfd-backend/database/db/model"

type PurchaseGetPurchasesReq struct {
	PageRequest
	Status model.OrderPurchaseStatus `form:"status"`
}

type PurchaseGetPurchasesResp struct {
	List     []*PurchaseOrder `json:"list"`
	TotalNum int64            `json:"totalNum"`
}

type PurchaseOrder struct {
	OrderID       int                       `json:"orderId"`
	CategoryNameA string                    `json:"categoryNameA"`
	CategoryNameB string                    `json:"categoryNameB"`
	CategoryNameC string                    `json:"categoryNameC"`
	CategoryA     int                       `json:"categoryA"`
	CategoryB     int                       `json:"categoryB"`
	CategoryC     int                       `json:"categoryC"`
	CategoryName  string                    `json:"categoryName"`
	Period        model.PurchasePeriod      `json:"period"`
	Quantity      int                       `json:"quantity"`
	Unit          string                    `json:"unit"`
	Requirement   string                    `json:"requirement"`
	UserID        string                    `json:"userID"`
	UserName      string                    `json:"userName"`
	UserAvatar    string                    `json:"userAvatar"`
	SubmitTime    int64                     `json:"submitTime"`
	HasQuote      bool                      `json:"hasQuote"`
	TotalQuote    int                       `json:"totalQuote"`
	NewQuote      int                       `json:"newQuote"`
	Status        model.OrderPurchaseStatus `json:"status"`
}

type PurchaseSubmitOrderReq struct {
	CategoryA   int                  `json:"categoryA"`
	CategoryB   int                  `json:"categoryB"`
	CategoryC   int                  `json:"categoryC"`
	Period      model.PurchasePeriod `json:"period"`
	Quantity    int                  `json:"quantity"`
	Unit        string               `json:"unit"`
	Requirement string               `json:"requirement"`
}

type PurchaseSubmitOrderResp struct {
}

//type PurchaseModifyOrderReq struct {
//	OrderID     int                  `json:"orderID"`
//	CategoryA   int                  `json:"categoryA"`
//	CategoryB   int                  `json:"categoryB"`
//	CategoryC   int                  `json:"categoryC"`
//	Period      model.PurchasePeriod `json:"period"`
//	Quantity    int                  `json:"quantity"`
//	Unit        string               `json:"unit"`
//	Requirement string               `json:"requirement"`
//	AreaCodeID  int                  `json:"areaCodeID"`
//}
//
//type PurchaseModifyOrderResp struct {
//}

type PurchaseModifyOrderStatusReq struct {
	OrderID int                       `json:"orderID"`
	Status  model.OrderPurchaseStatus `json:"status"`
	Delete  bool                      `json:"delete"` // 更高优先级
}
type PurchaseModifyOrderStatusResp struct {
}

type PurchaseGetQuotesReq struct {
	PageRequest
	OrderID int `form:"orderID"`
}

type PurchaseGetQuotesResp struct {
	List     []*PurchaseQuote `json:"list"`
	TotalNum int64            `json:"totalNum"`
}

type PurchaseQuote struct {
	QuoteID    int    `json:"quoteID"`
	OrderID    int    `json:"orderID"`
	GoodsID    int    `json:"goodsID"`
	GoodsURL   string `json:"goodsURL"`
	GoodsName  string `json:"goodsName"`
	Price      string `json:"price"`
	Unit       string `json:"unit"`
	Time       int64  `json:"time"`
	UserID     string `json:"userID"`
	UserName   string `json:"userName"`
	UserAvatar string `json:"userAvatar"`
	IsNew      bool   `json:"isNew"`
}

type PurchaseGetStatisticsReq struct {
}

type PurchaseGetStatisticsResp struct {
	NewQuote int `json:"newQuote"`
}

type PurchaseNotifySupplyReq struct {
	SupplyUserID string `json:"supplyUserID"`
	Count        int    `json:"count"`
}

type PurchaseNotifySupplyResp struct {
}

type PurchaseAnswerQuoteReq struct {
	SupplyUserID string `json:"supplyUserID"`
}

type PurchaseAnswerQuoteResp struct {
}
