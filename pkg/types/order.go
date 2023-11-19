package types

import (
	"xfd-backend/database/db/enum"
	"xfd-backend/database/db/model"
)

type ShoppingCartAddReq struct {
	ProductVariantID int32 `json:"productVariantID" binding:"required"`
	Quantity         int   `json:"quantity" binding:"required"`
}

type ShoppingCartDeleteReq struct {
	ShoppingCartIDs []int32 `json:"shoppingCartIDs" binding:"required,min=1"`
}

type ShoppingCartModifyReq struct {
	ShoppingCartID     int32                   `json:"shoppingCartID" binding:"required"`
	Quantity           int                     `json:"quantity" binding:"required"`
	ModifyQuantityType enum.ModifyQuantityType `json:"modifyType" binding:"required,oneof=1 2"`
}

type ShoppingCartListResp struct {
	PageResult
	List []*ShoppingCartDetail `json:"list,omitempty"`
}

type ShoppingCartDetail struct {
	*model.ShoppingCart
	Stock       *int   `json:"stock"`
	Price       string `json:"price"`
	Name        string `json:"name"`
	CoverURL    string `json:"coverURL"`
	ProductAttr string `json:"productAttr"`
}

type ShoppingCartListReq struct {
	PageRequest
	UserID string
}

type CreateOrderReq struct {
	UserAddressID   int     `json:"userAddressID" binding:"required"`
	ShoppingCartIDs []int32 `json:"shoppingCartIDs" binding:"required,min=1"`
	Remark          string  `json:"remark"`
}

type CreateOrderResp struct {
	OrderID     int32                `json:"orderID"`
	OrderSn     string               `json:"orderSn"`
	OrderStatus enum.OrderInfoStatus `json:"orderStatus"`
}

type ApplyRefundReq struct {
	OrderID     int  `json:"orderID"`
	RefundPoint bool `json:"refundPoint"`
	// todo RefundType
}

type ApplyRefundResp struct {
	// todo implement
}

type ApplyExchangeReq struct {
	// todo implement
}

type ApplyExchangeResp struct {
	// todo implement
}

type PayOrderReq struct {
	// todo implement
}

type PayOrderResp struct {
}

type CreatePreOrderResp struct {
	PreOrderAddress               *PreOrderAddress                `json:"preOrderAddress"`
	TotalPrice                    string                          `json:"totalPrice"`
	PointPrice                    string                          `json:"pointPrice"`
	WxPrice                       string                          `json:"wxPrice"`
	UserPoint                     string                          `json:"userPoint"`
	PreOrderProductVariantDetails []*PreOrderProductVariantDetail `json:"details"`
}

type PreOrderAddress struct {
	Name    string `json:"name"`
	Phone   string `json:"phone"`
	Address string `json:"address"`
}

type PreOrderProductVariantDetail struct {
	ShoppingCartID int32  `json:"shoppingCartID"`
	SKUCode        string `json:"skuCode"`
	Price          string `json:"price"`
	Quantity       int    `json:"quantity"`
	Name           string `json:"name"`
	CoverURL       string `json:"coverURL"`
	ProductAttr    string `json:"productAttr"`
	PostPrice      string `json:"postPrice"`
}
