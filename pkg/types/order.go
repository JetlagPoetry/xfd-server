package types

import (
	"fmt"
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
	QueryOrderID    int32                `json:"queryOrderID" binding:"required"`
	ReturnPointType enum.ReturnPointType `json:"returnPointType" binding:"required,oneof=1 2"`
	Reason          string               `json:"reason" binding:"required"`
}

type ApplyExchangeReq struct {
	QueryOrderID int32  `json:"queryOrderID" binding:"required"`
	Reason       string `json:"reason" binding:"required"`
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
	Status                   enum.OrderInfoStatus `form:"status" binding:"oneof=0 3 4 5 6"`
	UserID                   string               `form:"userID"`
	OrderSn                  string               `form:"orderSn"`
	GoodName                 string               `form:"goodName"`
	SupplierOrganizationName string               `form:"supplierOrganizationName"`
	UserPhone                string               `form:"userPhone"`
	UserOrganizationName     string               `form:"userOrganizationName"`
}

func (o *OrderListReq) CheckMiniAppParams() error {
	if o.Status == enum.OderInfoShipped || o.Status == enum.OderInfoReceived || o.Status == enum.OderInfoAfterSale {
		return fmt.Errorf("invalid query status, status=%d,please check", o.Status)
	}
	if o.OrderSn != "" || o.GoodName != "" || o.SupplierOrganizationName != "" || o.UserPhone != "" || o.UserOrganizationName != "" {
		return fmt.Errorf("only support query by status")
	}
	return nil
}

type OrderListResp struct {
	PageResult
	List []*QueryOrder `json:"list,omitempty"`
}

type QueryOrder struct {
	QueryOrderID                 int32                `gorm:"column:id" json:"queryOrderID"`
	OrderSn                      string               `gorm:"column:order_sn" json:"orderSn,omitempty"`
	Status                       enum.OrderInfoStatus `gorm:"column:status" json:"status"`
	Name                         string               `gorm:"column:name" json:"name"`
	Quantity                     int                  `gorm:"column:quantity" json:"quantity,omitempty"`
	UnitPrice                    float64              `gorm:"column:unit_price" json:"price,omitempty"`
	TotalPrice                   float64              `gorm:"column:total_price" json:"totalPrice"`
	PostPrice                    float64              `gorm:"column:post_price" json:"postPrice,omitempty"`
	Image                        string               `gorm:"column:image" json:"image,omitempty"`
	ProductAttr                  string               `gorm:"column:product_attr" json:"productAttr,omitempty"`
	EstimatedDelivery            *time.Time           `gorm:"column:estimated_delivery_time" json:"estimatedDelivery,omitempty"`
	ShipmentCompany              string               `gorm:"column:shipment_company" json:"shipmentCompany,omitempty"`
	ShipmentSn                   string               `gorm:"column:shipment_sn" json:"shipmentSn,omitempty"`
	SignerName                   string               `gorm:"column:signer_name" json:"signerName,omitempty"`
	SignerPhone                  string               `gorm:"column:singer_mobile" json:"signerPhone,omitempty"`
	SignerAddr                   string               `gorm:"column:signer_address" json:"signerAddr,omitempty"`
	SupplierUserID               string               `gorm:"column:goods_supplier_user_id" json:"supplierUserID,omitempty"`
	SupplierOrganizationName     string               `gorm:"column:goods_supplier_organization_name" json:"supplierOrganizationName,omitempty"`
	TotalOrderPrice              string               `gorm:"column:total_order_price" json:"totalOrderPrice,omitempty"`
	ConsumerUserPhone            string               `gorm:"column:user_phone" json:"consumerUserPhone,omitempty"`
	ConsumerUserOrganizationName string               `gorm:"column:user_organization_name" json:"consumerOrganizationName,omitempty"`
	CreatedAt                    *time.Time           `gorm:"column:created_at" json:"createdAt,omitempty"`
	PayedAt                      *time.Time           `gorm:"column:payed_at" json:"payedAt,omitempty"`
	DeliveryTime                 *time.Time           `gorm:"column:delivery_time" json:"deliveryTime,omitempty"`
}

type FillShipmentInfoReq struct {
	QueryOrderID    int32  `json:"queryOrderID" binding:"required"`
	ShipmentCompany string `json:"shipmentCompany" binding:"required"`
	ShipmentSn      string `json:"shipmentSn" binding:"required"`
}

type ConfirmReceiptReq struct {
	QueryOrderID int32 `json:"queryOrderID" form:"queryOrderID" binding:"required"`
}

type OrderDetailResp struct {
	IsOrderClosed bool `json:"isOrderClosed"`
	OrderInfo     `json:"order"`
	GoodsInfo     `json:"goods"`
	BuyerInfo     `json:"buyer"`
	SellerInfo    `json:"seller"`
	OrderRecord   `json:"orderRecord"`
}

type OrderInfo struct {
	OrderID    int32                `json:"orderID"`
	OrderSn    string               `json:"orderSn"`
	Status     enum.OrderInfoStatus `gorm:"column:status" json:"status"`
	CreatedAt  *time.Time           `json:"createdAt,omitempty"`
	PayedAt    *time.Time           `json:"payedAt,omitempty"`
	TotalPrice string               `json:"totalPrice"`
	WxPrice    string               `json:"wxPrice"`
	PointPrice string               `json:"pointPrice"`
}

type GoodsInfo struct {
	Name        string `json:"name"`
	GoodsID     int32  `json:"goodsID"`
	SKUCode     string `json:"skuCode"`
	Image       string `json:"image"`
	ProductAttr string `json:"productAttr"`
	Quantity    int    `json:"quantity"`
	UintPrice   string `json:"unitPrice"`
	PostPrice   string `json:"postPrice"`
}

type BuyerInfo struct {
	UserName             string `json:"userName"`
	UserPhone            string `json:"userPhone"`
	UserOrganizationName string `json:"userOrganizationName"`
	SingerName           string `json:"singerName"`
	SingerPhone          string `json:"singerPhone"`
	SingerAddr           string `json:"singerAddr"`
}

type SellerInfo struct {
	Name             string `json:"name"`
	Phone            string `json:"phone"`
	OrganizationName string `json:"organizationName"`
	Position         string `json:"position"`
}

type OrderRecord struct {
	CreatedAt          time.Time           `json:"createdAt,omitempty"`
	PayedAt            *time.Time          `json:"payedAt,omitempty"`
	DeliveryTime       *time.Time          `json:"deliveryTime,omitempty"`
	ConfirmTime        *time.Time          `json:"confirmTime,omitempty"`
	ManuallyCloseOrder *ManuallyCloseOrder `json:"manuallyCloseOrder,omitempty"`
	AfterSaleRecords   []*AfterSaleRecord  `json:"afterSaleRecords,omitempty"`
}

type AfterSaleRecord struct {
	CreatedAt       *time.Time           `json:"createdAt,omitempty"`
	Reason          string               `json:"reason,omitempty"`
	AfterSaleType   enum.AfterSaleType   `json:"afterSaleType,omitempty"`
	ReturnPointType enum.ReturnPointType `json:"returnPointType,omitempty"`
}

type ManuallyCloseOrder struct {
	CreatedAt       *time.Time           `json:"createdAt,omitempty"`
	Reason          string               `json:"reason,omitempty"`
	ReturnPointType enum.ReturnPointType `json:"returnPointType,omitempty"`
}

type CloseOrderReq struct {
	QueryOrderID    int32                `json:"queryOrderID" binding:"required"`
	ReturnPointType enum.ReturnPointType `json:"returnPointType" binding:"required,oneof=1 2"`
	Reason          string               `json:"reason" binding:"required"`
}
type GetCustomerServiceReq struct {
	SupplierUserID string `form:"supplierUserID" binding:"required"`
}

type GetCustomerServiceResp struct {
	Phone string `json:"phone"`
}
