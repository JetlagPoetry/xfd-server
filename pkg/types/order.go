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
	Stock *int `json:"stock"`
}

type ShoppingCartListReq struct {
	PageRequest
	UserID string
}

type CreateOrderReq struct {
	// todo implement
}

type CreateOrderResp struct {
	// todo implement
}
