package types

import "xfd-backend/database/db/model"

type PurchaseGetOrdersReq struct {
	BasePage
	Status     model.OrderPurchaseStatus `json:"status"`
	CategoryID int                       `json:"categoryID"`
	SortBy     string                    `json:"sortBy"` // todo
}

type PurchaseGetOrdersResp struct {
	List []*PurchaseOrder `json:"list"`
}

type PurchaseOrder struct {
	OrderID          int                  `json:"orderId"`
	CategoryID       int                  `json:"categoryID"`
	CategoryName     string               `json:"categoryName"`
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
}

type PurchaseSubmitOrderReq struct {
	CategoryID  int                  `json:"categoryID"`
	Period      model.PurchasePeriod `json:"period"`
	Quantity    int                  `json:"quantity"`
	Unit        string               `json:"unit"`
	Requirement string               `json:"requirement"`
	AreaCodeID  int                  `json:"areaCodeID"`
}

type PurchaseSubmitOrderResp struct {
}

type PurchaseModifyOrderReq struct {
	OrderID     int                  `json:"orderID"`
	CategoryID  int                  `json:"categoryID"`
	Period      model.PurchasePeriod `json:"period"`
	Quantity    int                  `json:"quantity"`
	Unit        string               `json:"unit"`
	Requirement string               `json:"requirement"`
	AreaCodeID  int                  `json:"areaCodeID"`
}

type PurchaseModifyOrderResp struct {
}

type PurchaseModifyOrderStatusReq struct {
	OrderID int                       `json:"orderID"`
	Status  model.OrderPurchaseStatus `json:"status"`
	Comment string                    `json:"comment"` // 审核意见
	Delete  bool                      `json:"delete"`  // 更高优先级
}
type PurchaseModifyOrderStatusResp struct {
}

type PurchaseGetQuotesReq struct {
	BasePage
	OrderID int `json:"orderID"`
}

type PurchaseGetQuotesResp struct {
	List []*PurchaseQuote `json:"list"`
}

type PurchaseQuote struct {
	QuoteID    int     `json:"quoteID"`
	OrderID    int     `json:"orderID"`
	ItemID     int     `json:"itemID"` // 商品id
	Price      float64 `json:"price"`
	Unit       string  `json:"unit"`
	Time       int64   `json:"time"`
	UserID     string  `json:"userID"`
	UserName   string  `json:"userName"`
	UserAvatar string  `json:"userAvatar"`
}

type PurchaseSubmitQuoteReq struct {
	OrderID int     `json:"orderID"`
	ItemID  int     `json:"itemID"`
	Price   float64 `json:"price"`
}

type PurchaseSubmitQuoteResp struct {
}
