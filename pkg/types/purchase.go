package types

import "xfd-backend/database/db/model"

type PurchaseGetOrdersReq struct {
	BasePage
	Status model.OrderPurchaseStatus `json:"status"`
}

type PurchaseGetOrdersResp struct {
	List []*PurchaseOrder `json:"list"`
}

type PurchaseOrder struct {
	OrderID     int                  `json:"orderId"`
	CategoryID  int                  `json:"categoryID"` // todo
	Period      model.PurchasePeriod `json:"period"`
	Quantity    int                  `json:"quantity"`
	Unit        string               `json:"unit"`
	Requirement string               `json:"requirement"`
	Location    string               `json:"location"` // todo
}

type PurchaseSubmitOrderReq struct {
	CategoryID  int                  `json:"categoryID"` // todo
	Period      model.PurchasePeriod `json:"period"`
	Quantity    int                  `json:"quantity"`
	Unit        string               `json:"unit"`
	Requirement string               `json:"requirement"`
	Location    string               `json:"location"` // todo
}

type PurchaseSubmitOrderResp struct {
}

type PurchaseModifyOrderReq struct {
	OrderID model.OrderPurchaseStatus `json:"orderID"`

	CategoryID  int                  `json:"categoryId"` // todo
	Period      model.PurchasePeriod `json:"period"`
	Quantity    int                  `json:"quantity"`
	Unit        string               `json:"unit"`
	Requirement string               `json:"requirement"`
	Location    string               `json:"location"` // todo
}

type PurchaseModifyOrderResp struct {
}

type PurchaseModifyOrderStatusReq struct {
	OrderID model.OrderPurchaseStatus `json:"orderID"`
	Status  model.OrderPurchaseStatus `json:"status"`
	Delete  bool                      `json:"delete"` // 更高优先级
}

type PurchaseModifyOrderStatusResp struct {
}
