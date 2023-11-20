package model

import (
	"github.com/shopspring/decimal"
	"gorm.io/gorm"
	"time"
	"xfd-backend/database/db/enum"
)

type ShoppingCart struct {
	ID                  int32          `gorm:"primary_key;AUTO_INCREMENT;type:int" json:"id"`
	UserID              string         `gorm:"type:varchar(300);not null;default:'';column:user_id;comment:用户ID;index:user_product_sku_code_goods" json:"userID"`
	ProductVariantID    int32          `gorm:"type:int;not null;comment:商品编号ID;column:product_variant_id;index:user_product_sku_code_goods" json:"productVariantID"`
	GoodsID             int32          `gorm:"type:int;not null;default:0;comment:商品ID;;column:goods_id;index:user_product_sku_code_goods" json:"goodsID"`
	GoodsSupplierUserID string         `gorm:"type:varchar(200);not null;default:'';comment:供应商ID" json:"goodsSupplierUserID"`
	SKUCode             string         `gorm:"type:varchar(300);not null;default:'';column:sku_code;comment:商品SKU编号;index:user_product_sku_code_goods" json:"skuCode"`
	Quantity            int            `gorm:"type:int;not null;default:0;comment:数量;column:quantity;" json:"quantity"`
	CreatedAt           time.Time      `gorm:"comment:创建时间;not null;column:created_at;default:CURRENT_TIMESTAMP(3);" json:"-"`
	UpdatedAt           time.Time      `gorm:"comment:更新时间;not null;column:updated_at;default:CURRENT_TIMESTAMP(3) ON UPDATE CURRENT_TIMESTAMP(3);" json:"-"`
	DeletedAt           gorm.DeletedAt `gorm:"column:deleted_at;index" json:"-"`
}

type OrderPayInfo struct {
	ID         int32          `gorm:"primary_key;AUTO_INCREMENT;type:int" json:"id"`
	CreatedAt  time.Time      `gorm:"comment:创建时间;not null;column:created_at;default:CURRENT_TIMESTAMP(3);" json:"-"`
	UpdatedAt  time.Time      `gorm:"comment:更新时间;not null;column:updated_at;default:CURRENT_TIMESTAMP(3) ON UPDATE CURRENT_TIMESTAMP(3);" json:"-"`
	DeletedAt  gorm.DeletedAt `gorm:"column:deleted_at;index" json:"-"`
	User       string         `gorm:"type:varchar(100);index;comment:用户id"`
	OrderSn    string         `gorm:"type:varchar(100);index;comment:我们平台自己生成的订单号"`
	PayType    int            `gorm:"type:varchar(20);comment:1:wechat(微信)"`
	Status     string         `gorm:"type:varchar(20);comment:PAYING(待支付), TRADE_SUCCESS(成功)， TRADE_CLOSED(超时关闭), WAIT_BUYER_PAY(交易创建), TRADE_FINISHED(交易结束)"`
	TradeNo    string         `gorm:"type:varchar(100);comment:交易号"`
	OrderMount float32        `gorm:"comment:该订单需支付的总金额"`
	PayTime    *time.Time     `gorm:"comment:用户支付该订单的时间"`
}

// OrderInfo 订单信息
type OrderInfo struct {
	ID                int32                  `gorm:"primary_key;AUTO_INCREMENT;type:int" json:"id"`
	CreatedAt         time.Time              `gorm:"comment:创建时间;not null;column:created_at;default:CURRENT_TIMESTAMP(3);" json:"-"`
	UpdatedAt         time.Time              `gorm:"comment:更新时间;not null;column:updated_at;default:CURRENT_TIMESTAMP(3) ON UPDATE CURRENT_TIMESTAMP(3);" json:"-"`
	DeletedAt         gorm.DeletedAt         `gorm:"column:deleted_at;index" json:"-"`
	UserID            string                 `gorm:"column:user_id;type:varchar(200);default:'';not null;comment:下单用户ID;index:order_sn_status_user" json:"userID"`
	TotalPrice        decimal.Decimal        `gorm:"column:total_price;type:decimal(9,2);default:0.0;not null;comment:总订单价" json:"total_price"`
	PostPrice         decimal.Decimal        `gorm:"column:post_price;type:decimal(9,2);default:0.0;not null;comment:总邮费" json:"post_price"`
	PointPrice        decimal.Decimal        `gorm:"column:point_price;type:decimal(9,2);default:0.0;not null;comment:积分实付" json:"point_price"`
	WxPrice           decimal.Decimal        `gorm:"column:wx_price;type:decimal(9,2);default:0.0;not null;comment:微信实付" json:"wx_price"`
	Status            enum.OrderInfoStatus   `gorm:"column:status;type:tinyint(1);default:0;not null;comment:订单状态" json:"status"`
	OrderSn           string                 `gorm:"column:order_sn;unique;type:varchar(500);default:'';not null;index:order_sn_status_user;comment:我们平台自己生成的订单号"`
	TradeNo           string                 `gorm:"column:trade_no;type:varchar(300);default:'';not null;comment:交易号"`
	PayedAt           *time.Time             `gorm:"column:payed_at;comment:用户支付该订单的时间"`
	ExpiredAt         time.Time              `gorm:"column:expired_at;comment:订单过期时间"`
	Address           string                 `gorm:"column:address;type:varchar(300);comment:收件人的地址"`
	SignerName        string                 `gorm:"column:signer_name;type:varchar(300);comment:收件人的名称"`
	SingerMobile      string                 `gorm:"column:singer_mobile;type:varchar(300);comment:收件人的手机号"`
	Message           string                 `gorm:"column:message;type:varchar(500);comment:结束订单留言信息"`
	AfterSaleType     enum.AfterSaleType     `gorm:"column:after_sale_type;type:tinyint(1);default:0;not null;comment:售后方式" json:"after_sale_type"`
	ReturnPointType   enum.ReturnPointType   `gorm:"column:return_point;type:tinyint(1);default:1;not null;comment:是否返还积分" json:"return_point"`
	ReturnPointStatus enum.ReturnPointStatus `gorm:"column:return_point_status;type:tinyint(1);default:0;not null;comment:返还积分状态" json:"return_point_status"`
}

// OrderProductVariantDetail 订单产品详情
type OrderProductVariantDetail struct {
	ID                    int32                          `gorm:"primary_key;AUTO_INCREMENT;type:int" json:"id"`
	CreatedAt             time.Time                      `gorm:"comment:创建时间;not null;column:created_at;default:CURRENT_TIMESTAMP(3);" json:"-"`
	UpdatedAt             time.Time                      `gorm:"comment:更新时间;not null;column:updated_at;default:CURRENT_TIMESTAMP(3) ON UPDATE CURRENT_TIMESTAMP(3);" json:"-"`
	DeletedAt             gorm.DeletedAt                 `gorm:"column:deleted_at;index" json:"-"`
	OrderSn               string                         `gorm:"column:order_sn;type:varchar(500);default:'';not null;comment:平台自己生成的订单号;index:order_sn_status_user" json:"orderSn"`
	OrderProductSn        string                         `gorm:"column:order_product_sn;unique,type:varchar(500);default:'';not null;comment:平台自己生成的SKU订单号;index:order_product_sn_status_user" json:"orderProductSn"`
	Status                enum.OrderProductVariantDetail `gorm:"column:status;type:tinyint(1);default:0;not null;comment:产品状态;index:order_product_sn_status_user;index:supplier_user_id_status" json:"status"`
	UserID                string                         `gorm:"column:user_id;type:varchar(200);default:'';not null;comment:下单用户ID;index:order_product_sn_status_user" json:"userID"`
	ShoppingCartID        int32                          `gorm:"type:int;not null;default:0;comment:购物车ID" json:"shoppingCartID"`
	ProductVariantID      int32                          `gorm:"type:int;not null;default:0;comment:商品编号ID;column:product_variant_id;index:user_product_sku_code_goods" json:"productVariantID"`
	GoodsID               int32                          `gorm:"type:int;not null;default:0;comment:商品ID;;column:goods_id" json:"goodsID"`
	GoodsSupplierUserID   string                         `gorm:"type:varchar(200);not null;default:'';comment:供应商ID;index:supplier_user_id_status" json:"goodsSupplierUserID"`
	SKUCode               string                         `gorm:"type:varchar(300);not null;default:'';column:sku_code;comment:商品SKU编号" json:"skuCode"`
	Quantity              int                            `gorm:"type:int;not null;default:0;comment:数量;column:quantity;" json:"quantity"`
	TotalPrice            decimal.Decimal                `gorm:"type:decimal(9,2);not null;default:0.0;comment:SKU总价" json:"totalPrice"`
	Price                 decimal.Decimal                `gorm:"type:decimal(9,2);not null;default:0.0;comment:SKU单价" json:"price"`
	PostPrice             decimal.Decimal                `gorm:"type:decimal(9,2);not null;default:0.0;comment:邮费" json:"postPrice"`
	Name                  string                         `gorm:"type:varchar(300);not null;default:'';column:name;comment:商品名称"`
	Image                 string                         `gorm:"type:varchar(1000);not null;column:image;comment:商品封面图"`
	ProductAttr           string                         `gorm:"type:varchar(1000);default:'';column:product_attr;comment:商品销售属性:[{\"key\":\"颜色\",\"value\":\"红色\"},{\"key\":\"容量\",\"value\":\"4G\"}]" json:"product_attr"`
	ShipmentCompany       string                         `gorm:"type:varchar(100);default:'';column:shipment_company;comment:快递公司"`
	ShipmentSn            string                         `gorm:"type:varchar(500);default:'';column:shipment_sn;comment:快递单号"`
	EstimatedDeliveryTime time.Time                      `gorm:"column:estimated_delivery_time;comment:预计发货时间" json:"estimated_delivery_time"`
	DeliveryTime          *time.Time                     `gorm:"column:delivery_time;comment:发货时间" json:"delivery_time"`
	ReceiveTime           *time.Time                     `gorm:"column:receive_time;comment:物流收货时间" json:"receive_time"`
	ConfirmTime           *time.Time                     `gorm:"column:confirm_time;comment:确认收货时间" json:"confirm_time"`
	Message               string                         `gorm:"column:message;type:varchar(10000);comment:补充记录信息"`
}
