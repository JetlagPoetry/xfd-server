package model

import (
	"gorm.io/gorm"
	"time"
	"xfd-backend/database/db/enum"
)

// ShoppingCart 购物车
type ShoppingCart struct {
	gorm.Model
	User  string `gorm:"type:int;index;comment:用户id"`
	Goods int32  `gorm:"type:int;index;comment:商品id"`
	Nums  int32  `gorm:"type:int;comment:商品的购买个数"`
}

// OrderPayInfo 订单支付信息表
type OrderPayInfo struct {
	gorm.Model
	User       string     `gorm:"type:varchar(100);index;comment:用户id"`
	OrderSn    string     `gorm:"type:varchar(100);index;comment:我们平台自己生成的订单号"`
	PayType    int        `gorm:"type:varchar(20);comment:1:wechat(微信)"`
	Status     string     `gorm:"type:varchar(20);comment:PAYING(待支付), TRADE_SUCCESS(成功)， TRADE_CLOSED(超时关闭), WAIT_BUYER_PAY(交易创建), TRADE_FINISHED(交易结束)"`
	TradeNo    string     `gorm:"type:varchar(100);comment:交易号"`
	OrderMount float32    `gorm:"comment:该订单需支付的总金额"`
	PayTime    *time.Time `gorm:"comment:用户支付该订单的时间"`
}

// OrderGoods 订单商品信息表
type OrderGoods struct {
	BaseModel
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
	BaseModel
	ShipmentID string    `gorm:"column:shipment_id" json:"shipment_id" comment:"快递单号"`
	Times      time.Time `gorm:"column:times" json:"times" comment:"时间"`
	Context    string    `gorm:"column:context" json:"context" comment:"描述信息"`
	CreatedBy  string    `gorm:"column:created_by" json:"created_by" comment:"创建者"`
}

func (Logistics) TableName() string {
	return "logistics"
}
