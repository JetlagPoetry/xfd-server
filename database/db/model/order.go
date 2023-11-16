package model

import (
	"gorm.io/gorm"
	"time"
	"xfd-backend/database/db/enum"
)

type ShoppingCart struct {
	ID               int32          `gorm:"primary_key;AUTO_INCREMENT;type:int" json:"id"`
	UserID           string         `gorm:"type:varchar(300);not null;default:'';column:user_id;comment:用户ID;index:user_product_sku_code_goods" json:"userID"`
	ProductVariantID int32          `gorm:"type:int;not null;comment:商品编号ID;column:product_variant_id;index:user_product_sku_code_goods" json:"productVariantID"`
	GoodsID          int32          `gorm:"type:int;not null;default:0;comment:商品ID;;column:goods_id;index:user_product_sku_code_goods" json:"goodsID"`
	SKUCode          string         `gorm:"type:varchar(300);not null;default:'';column:sku_code;comment:商品SKU编号;index:user_product_sku_code_goods" json:"skuCode"`
	Name             string         `gorm:"type:varchar(300);not null;default:'';column:name;comment:商品名称" json:"name"`
	CoverURL         string         `gorm:"type:varchar(1000);not null;column:cover_url;default:'';comment:封面URL" json:"coverURL"`
	Quantity         int            `gorm:"type:int;not null;default:0;comment:数量;column:quantity;" json:"quantity"`
	ProductAttr      string         `gorm:"type:varchar(1000);default:'';column:product_attr;comment:商品销售属性:[{\"key\":\"颜色\",\"value\":\"红色\"},{\"key\":\"容量\",\"value\":\"4G\"}]" json:"productAttr"`
	CreatedAt        time.Time      `gorm:"comment:创建时间;not null;column:created_at;default:CURRENT_TIMESTAMP(3);" json:"-"`
	UpdatedAt        time.Time      `gorm:"comment:更新时间;not null;column:updated_at;default:CURRENT_TIMESTAMP(3) ON UPDATE CURRENT_TIMESTAMP(3);" json:"-"`
	DeletedAt        gorm.DeletedAt `gorm:"column:deleted_at;index" json:"-"`
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

// OrderGoods 订单商品信息表
type OrderGoods struct {
	ID               int32                 `gorm:"primary_key;AUTO_INCREMENT;type:int" json:"id"`
	CreatedAt        time.Time             `gorm:"comment:创建时间;not null;column:created_at;default:CURRENT_TIMESTAMP(3);" json:"-"`
	UpdatedAt        time.Time             `gorm:"comment:更新时间;not null;column:updated_at;default:CURRENT_TIMESTAMP(3) ON UPDATE CURRENT_TIMESTAMP(3);" json:"-"`
	DeletedAt        gorm.DeletedAt        `gorm:"column:deleted_at;index" json:"-"`
	UserID           string                `gorm:"type:int;not null;comment:用户ID" json:"user_id"`
	User             *User                 `gorm:"foreignKey:UserID;references:ID" json:"user"`
	TotalPrice       float64               `gorm:"type:decimal(9,2);not null;comment:总价" json:"total_price"`
	UnitPrice        float64               `gorm:"type:decimal(9,2);not null;comment:单价" json:"unit_price"`
	PostPrice        float64               `gorm:"type:decimal(9,2);not null;comment:邮费" json:"post_price"`
	Status           enum.OrderGoodsStatus `gorm:"type:varchar(50);not null;comment:订单状态" json:"status"`
	TradeNo          string                `gorm:"type:varchar(100);comment:交易号"`
	ProductVariantID int32                 `gorm:"type:int;not null;comment:商品编号ID" json:"product_variant_id"`
	ProductVariant   *ProductVariant       `gorm:"foreignKey:ProductVariantID;references:ID" json:"product_variant"`
	Quantity         int                   `gorm:"type:int;not null;comment:购买数量" json:"quantity"`
	GoodsName        string                `gorm:"type:varchar(100);index;comment:商品名称"`
	GoodsImage       string                `gorm:"type:varchar(200);comment:商品封面图"`
	Address          string                `gorm:"type:varchar(100);comment:收件人的地址"`
	SignerName       string                `gorm:"type:varchar(100);comment:收件人的名称"`
	SingerMobile     string                `gorm:"type:varchar(100);comment:收件人的手机号"`
	Post             string                `gorm:"type:varchar(100);comment:留言信息"`
	Context          string                `gorm:"type:varchar(100);comment:描述信息"`
	ShipmentID       string                `gorm:"type:varchar(100);column:shipment_id" json:"shipment_id" comment:"快递单号"`
}

// todo:是否要主动查询快递状态，记录信息
type Logistics struct {
	ShipmentID string    `gorm:"column:shipment_id" json:"shipment_id" comment:"快递单号"`
	Times      time.Time `gorm:"column:times" json:"times" comment:"时间"`
	Context    string    `gorm:"column:context" json:"context" comment:"描述信息"`
	CreatedBy  string    `gorm:"column:created_by" json:"created_by" comment:"创建者"`
}
