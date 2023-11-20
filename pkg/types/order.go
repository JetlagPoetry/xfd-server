package types

import (
	"time"
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
	UserAddressID  int    `json:"userAddressID" binding:"required"`
	ShoppingCartID int32  `json:"shoppingCartID" binding:"required"`
	Remark         string `json:"remark"`
	Code           string `json:"code"` // 获取openid
}

type CreateOrderResp struct {
	OrderID     int32                `json:"orderID"`
	OrderSn     string               `json:"orderSn"`
	OrderStatus enum.OrderInfoStatus `json:"orderStatus"`
	PayWx       bool                 `json:"payWx"`
	PayData     *WechatPay           `json:"payData,omitempty"`
}

type WechatPay struct {
	AppID     string `json:"appId"`
	Timestamp int64  `json:"timeStamp"`
	NonceStr  string `json:"nonceStr"`
	Package   string `json:"package"`
	SignType  string `json:"signType"`
	PaySign   string `json:"paySign"`
}

type ApplyRefundReq struct {
	OrderID     int  `json:"orderID"`
	RefundPoint bool `json:"refundPoint"`
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
	PreOrderAddress *PreOrderAddress `json:"preOrderAddress"`
	TotalPrice      string           `json:"totalPrice"`
	PointPrice      string           `json:"pointPrice"`
	WxPrice         string           `json:"wxPrice"`
	UserPoint       string           `json:"userPoint"`
	ShoppingCartID  int32            `json:"shoppingCartID"`
	SKUCode         string           `json:"skuCode"`
	Price           string           `json:"price"`
	Quantity        int              `json:"quantity"`
	Name            string           `json:"name"`
	CoverURL        string           `json:"coverURL"`
	ProductAttr     string           `json:"productAttr"`
	PostPrice       string           `json:"postPrice"`
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

type OrderListReq struct {
	PageRequest
	Status int    `form:"status" binding:"oneof=0 3 4"`
	UserID string `form:"userID"`
}

type OrderListResp struct {
	PageResult
	List []*QueryOrder `json:"list,omitempty"`
}

type QueryOrder struct {
	QueryOrderID                 int32                          `gorm:"column:id" json:"queryOrderID"`
	OrderSn                      string                         `gorm:"column:order_sn" json:"orderSn,omitempty"`
	Status                       enum.OrderProductVariantDetail `gorm:"column:status" json:"status"`
	Name                         string                         `gorm:"column:name" json:"name"`
	Quantity                     int                            `gorm:"column:quantity" json:"quantity"`
	UnitPrice                    float64                        `gorm:"column:unit_price" json:"price"`
	TotalPrice                   float64                        `gorm:"column:total_price" json:"totalPrice"`
	PostPrice                    float64                        `gorm:"column:post_price" json:"postPrice,omitempty"`
	Image                        string                         `gorm:"column:image" json:"image"`
	ProductAttr                  string                         `gorm:"column:product_attr" json:"productAttr"`
	EstimatedDelivery            time.Time                      `gorm:"column:estimated_delivery_time" json:"estimatedDelivery,omitempty"`
	ShipmentCompany              string                         `gorm:"column:shipment_company" json:"shipmentCompany,omitempty"`
	ShipmentSn                   string                         `gorm:"column:shipment_sn" json:"shipmentSn,omitempty"`
	SignerName                   string                         `gorm:"column:signer_name" json:"signerName,omitempty"`
	SignerPhone                  string                         `gorm:"column:singer_mobile" json:"signerPhone,omitempty"`
	SignerAddr                   string                         `gorm:"column:signer_address" json:"signerAddr,omitempty"`
	SupplierUserID               string                         `gorm:"column:goods_supplier_user_id" json:"supplierUserID,omitempty"`
	SupplierOrganizationName     string                         `gorm:"column:organization_name" json:"supplierOrganizationName,omitempty"`
	TotalOrderPrice              string                         `gorm:"column:total_order_price" json:"totalOrderPrice,omitempty"`
	ConsumerUserPhone            string                         `gorm:"column:consumer_user_phone" json:"consumerUserPhone,omitempty"`
	ConsumerUserOrganizationName string                         `gorm:"column:consumer_organization_name" json:"consumerOrganizationName,omitempty"`
	CreatedAt                    time.Time                      `gorm:"column:created_at" json:"createdAt,omitempty"`
	PayedAt                      *time.Time                     `gorm:"column:payed_at" json:"payedAt,omitempty"`
	DeliveryTime                 *time.Time                     `gorm:"column:delivery_time" json:"deliveryTime,omitempty"`
}

type FillShipmentInfoReq struct {
	QueryOrderID    int32  `json:"queryOrderID" binding:"required"`
	ShipmentCompany string `json:"shipmentCompany" binding:"required"`
	ShipmentSn      string `json:"shipmentSn" binding:"required"`
}

type ConfirmReceiptReq struct {
	QueryOrderID int32 `json:"queryOrderID" binding:"required"`
}
