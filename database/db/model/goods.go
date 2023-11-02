package model

import "xfd-backend/database/db/enum"

// Category 商品分类表
type Category struct {
	BaseModel
	Name             string                  `gorm:"type:varchar(100);default:'';not null;comment:分类名称" json:"name"`
	ParentCategoryID int32                   `gorm:"type:int;comment:父分类ID;index:parent_status" json:"parentID,omitempty"` // 父级分类ID
	ParentCategory   *Category               `json:"-"`                                                                    // 序列化json数据时忽略该字段
	SubCategory      []*Category             `gorm:"foreignKey:ParentCategoryID;references:ID" json:"children,omitempty"`  // foreignKey：关联"另外一张表的键"；references：另外一张表关联"此表的主键"
	Level            enum.GoodsCategoryLevel `gorm:"type:tinyint(1);not null;default:1;comment:分类级别;index:level_status" json:"level"`
	Image            string                  `gorm:"type:varchar(1000);not null;default:'';comment:分类图片概览" json:"image"`
	Status           int32                   `gorm:"type:tinyint(1);not null;default:1;comment:状态;index:level_status;index:parent_status" json:"-"`
}

func (u *Category) TableName() string {
	return "category"
}

// Goods 商品表
type Goods struct {
	BaseModel
	CategoryID         int32 `gorm:"type:int;default:0;not null"`
	Category           *Category
	BrandsID           int32    `gorm:"type:int;not null"`
	OnSale             bool     `gorm:"default:false;not null;comment:是否已上架"`
	ShipFree           bool     `gorm:"default:false;not null;comment:是否免运费"`
	Name               string   `gorm:"type:varchar(100);not null;comment:商品名称"`
	GoodsSn            string   `gorm:"type:varchar(100);not null;comment:商品编号"`
	GoodsBrief         string   `gorm:"type:varchar(100);not null;comment:商品简介"`
	Images             GormList `gorm:"type:varchar(1000);not null;comment:商品轮播图"`
	DescImages         GormList `gorm:"type:varchar(2000);not null;comment:商品详情图"`
	GoodsFrontImage    string   `gorm:"type:varchar(200);not null;comment:商品封面图"`
	ClickNum           int32    `gorm:"type:int;default:0;not null;comment:商品点击数"`
	IsNew              bool     `gorm:"default:false;not null;comment:是否为新品"`
	IsHot              bool     `gorm:"default:false;not null;comment:是否为热卖商品"`
	SoldNum            int32    `gorm:"type:int;default:0;not null;comment:零售销量"`
	FavNum             int32    `gorm:"type:int;default:0;not null;comment:收藏数量"`
	Origin             string   `gorm:"type:varchar(100);not null;comment:商品产地"`
	WholesaleLogistics string   `gorm:"type:varchar(100);not null;comment:批发物流"`
	WholesaleShipping  string   `gorm:"type:varchar(100);not null;comment:批发发货地"`
	RetailShipping     string   `gorm:"type:varchar(100);not null;comment:零售发货地"`
	RetailSupport      bool     `gorm:"default:false;not null;comment:是否支持零售"`
	//todo:是否根据距离排序，增加发货地的经纬度
	RetailShippingTime enum.RetailDeliveryTime `gorm:"type:int;not null;comment:零售发货时间"`
	//todo:待拓展更复杂信息，自动计算运费
	RetailShippingFee string `gorm:"type:varchar(100);not null;comment:零售运费信息"`
}

// Specification 商品规格表
type Specification struct {
	BaseModel
	Name    string `gorm:"type:varchar(50);not null;comment:规格名称" json:"name"`
	GoodsID int32  `gorm:"type:int;not null;comment:商品ID" json:"goods_id"`
	Goods   *Goods `gorm:"foreignKey:GoodsID;references:ID" json:"goods"`
	Type    int    `gorm:"type:int;not null;comment:类型 1-批发 2-零售" json:"type"`
}

// ProductVariant 商品编号表
type ProductVariant struct {
	BaseModel
	GoodsID          int32          `gorm:"type:int;not null;comment:商品ID" json:"goods_id"`
	Goods            *Goods         `gorm:"foreignKey:GoodsID;references:ID" json:"goods"`
	Specification1ID int32          `gorm:"type:int;not null;comment:规格1ID" json:"specification1_id"`
	Specification1   *Specification `gorm:"foreignKey:Specification1ID;references:ID" json:"specification1"`
	Specification2ID int32          `gorm:"type:int;not null;comment:规格2ID" json:"specification2_id"`
	Specification2   *Specification `gorm:"foreignKey:Specification2ID;references:ID" json:"specification2"`
	Quantity         int            `gorm:"type:int;not null;comment:数量" json:"quantity"`
	Unit             string         `gorm:"type:varchar(10);not null;comment:单位" json:"unit"`
	Price            float64        `gorm:"type:decimal(9,2);not null;comment:价格" json:"price"`
	Stock            int            `gorm:"type:int;comment:库存" json:"stock"`
	MinOrderQuantity int            `gorm:"type:int;comment:起批量" json:"min_order_quantity"`
	Type             int            `gorm:"type:int;not null;comment:类型 1-批发 2-零售" json:"type"`
	Enabled          bool           `gorm:"type:bool;default:true;comment:是否启用" json:"enabled"`
}

// Inventory 商品库存表
type Inventory struct {
	BaseModel
	ProductVariantID int32 `gorm:"type:int;not null;comment:商品编号ID" json:"product_variant_id"`
	Stocks           int32 `gorm:"type:int;comment:商品库存"`
}

// StockSellDetail 库存扣减详情
type StockSellDetail struct {
	// 建立索引，值唯一
	OrderSn string `gorm:"type:varchar(200);index:idx_order_sn,unique;comment:订单编号"`
	// 订单的库存扣减或者归还后都要更新这个字段，
	// 执行库存归还前一定要判断这个状态是否为"已扣减"，
	// 只有"已扣减"状态的订单才可以执行库存归还。
	Status int `gorm:"type:int;comment:库存扣减状态。1：已扣减。2：已归还"`
	// 详细记录这个订单下各个商品扣减之前的库存，至于为什么不把这个字段拆解成Goods和Nums，
	// 是因为拆解后每当对一个订单执行库存扣减时，就需要更新多条记录。
	Detail GoodsDetailList `gorm:"type:varchar(200);comment:该订单下各个商品执行扣减之前的库存"`
}
