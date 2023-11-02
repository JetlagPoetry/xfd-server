package types

import "xfd-backend/database/db/model"

type PurchaseGetOrdersReq struct {
	PageRequest
	CategoryA int `json:"categoryA"`
	CategoryB int `json:"categoryB"`
	CategoryC int `json:"categoryC"`
}

type PurchaseGetOrdersResp struct {
	List     []*PurchaseOrder `json:"list"`
	TotalNum int64            `json:"totalNum"`
}

type PurchaseOrder struct {
	OrderID          int                  `json:"orderId"`
	CategoryNameA    string               `json:"categoryNameA"`
	CategoryNameB    string               `json:"categoryNameB"`
	CategoryNameC    string               `json:"categoryNameC"`
	Period           model.PurchasePeriod `json:"period"`
	Quantity         int                  `json:"quantity"`
	Unit             string               `json:"unit"`
	Requirement      string               `json:"requirement"`
	AreaCodeID       int                  `json:"areaCodeID"`
	AreaName         string               `json:"areaName"`
	UserID           string               `json:"userID"`
	UserName         string               `json:"userName"`
	UserAvatar       string               `json:"userAvatar"`
	UserOrganization string               `json:"userOrganization"`
	SubmitTime       int64                `json:"submitTime"`
	HasQuote         bool                 `json:"hasQuote"`
	NewQuote         int                  `json:"newQuote"`
}

type PurchaseSubmitOrderReq struct {
	CategoryA   int                  `json:"categoryA"`
	CategoryB   int                  `json:"categoryB"`
	CategoryC   int                  `json:"categoryC"`
	Period      model.PurchasePeriod `json:"period"`
	Quantity    int                  `json:"quantity"`
	Unit        string               `json:"unit"`
	Requirement string               `json:"requirement"`
	AreaCodeID  int                  `json:"areaCodeID"`
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
	OrderID int `json:"orderID"`
}

type PurchaseGetQuotesResp struct {
	List []*PurchaseQuote `json:"list"`
}

type PurchaseQuote struct {
	QuoteID int `json:"quoteID"`
	OrderID int `json:"orderID"`
	ItemID  int `json:"itemID"` // 商品id
	// todo 商品标题、图片
	Price            float64 `json:"price"`
	Unit             string  `json:"unit"`
	Time             int64   `json:"time"`
	UserID           string  `json:"userID"`
	UserName         string  `json:"userName"`
	UserAvatar       string  `json:"userAvatar"`
	UserOrganization string  `json:"userOrganization"`
	IsNew            bool    `json:"isNew"` // todo 新增报价
}

type PurchaseSubmitQuoteReq struct {
	OrderID int     `json:"orderID"`
	ItemID  int     `json:"itemID"`
	Price   float64 `json:"price"`
}

type PurchaseSubmitQuoteResp struct {
}

type PurchaseGetStatisticsReq struct {
}

type PurchaseGetStatisticsResp struct {
	RetailNumber int `json:"retailNumber"`
	QuoteNumber  int `json:"quoteNumber"`
}
