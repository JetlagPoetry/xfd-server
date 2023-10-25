package model

import (
	"gorm.io/gorm"
	"time"
)

// 购物车表结构
type ShoppingCart struct {
	gorm.Model
	User    int32 `gorm:"type:int;index;comment:用户id"` //在购物车列表中我们需要查询当前用户的购物车记录
	Goods   int32 `gorm:"type:int;index;comment:商品id"` //如果表的字段没有被多次查询，就不要加索引。原因如下： 1. 会影响插入性能 2. 会占用磁盘
	Nums    int32 `gorm:"type:int;comment:商品的购买个数"`
	Checked bool  `gorm:"comment:是否选中"`
}

//func (ShoppingCart) TableName() string {
//	return "shoppingcart"
//}

// 订单信息表结构
type OrderInfo struct {
	gorm.Model
	User         int32      `gorm:"type:int;index;comment:用户id"`
	OrderSn      string     `gorm:"type:varchar(30);index;comment:我们平台自己生成的订单号"`
	PayType      string     `gorm:"type:varchar(20);comment:alipay(支付宝)，wechat(微信)"`
	Status       string     `gorm:"type:varchar(20);comment:PAYING(待支付), TRADE_SUCCESS(成功)， TRADE_CLOSED(超时关闭), WAIT_BUYER_PAY(交易创建), TRADE_FINISHED(交易结束)"`
	TradeNo      string     `gorm:"type:varchar(100);comment:交易号，可理解为支付宝的订单号，用于查账"`
	OrderMount   float32    `gorm:"comment:该订单需支付的总金额"`
	PayTime      *time.Time `gorm:"comment:用户支付该订单的时间"`
	Address      string     `gorm:"type:varchar(100);comment:收件人的地址"`
	SignerName   string     `gorm:"type:varchar(20);comment:收件人的名称"`
	SingerMobile string     `gorm:"type:varchar(11);comment:收件人的手机号"`
	Post         string     `gorm:"type:varchar(20);comment:留言信息"`
}

//func (OrderInfo) TableName() string {
//	return "orderinfo"
//}

// 订单-商品对应关系表，一条记录对应"一个订单id和一个商品id"
type OrderGoods struct {
	gorm.Model

	Order int32 `gorm:"type:int;index;comment:订单id"`
	Goods int32 `gorm:"type:int;index;comment:商品id"`

	// 为什么把商品的信息保存下来了？这会造成字段冗余，因为在商品表中已存在相同的字段。
	// 高并发系统中我们一般都不会遵循表设计的三范式。
	// 这样做的好处有两点：
	// 1. 避免频繁跨服务获取商品信息。
	// 2. 做镜像、记录下订单时商品的价格，因为商品的价格会被调整，而我们要查的是下单时商品的价格。
	GoodsName  string  `gorm:"type:varchar(100);index;comment:商品名称"`
	GoodsImage string  `gorm:"type:varchar(200);comment:商品封面图"`
	GoodsPrice float32 `gorm:"comment:商品价格"`
	Nums       int32   `gorm:"type:int;comment:商品的购买个数"`
}

//func (OrderGoods) TableName() string {
//	return "ordergoods"
//}
