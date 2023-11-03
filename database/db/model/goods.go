package model

import (
	"time"
	"xfd-backend/database/db/enum"
)

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
	ID                 int32                   `gorm:"primary_key;AUTO_INCREMENT;column:id;type:int"`
	UserID             string                  `gorm:"type:varchar(100);default:'';not null;column:user_id;not null;comment:供货商ID;index:user_id" `
	CategoryAID        int32                   `gorm:"type:int;not null;default:0;column:category_a_id;comment:一级分类ID;index:level_status_deleted" `
	CategoryBID        int32                   `gorm:"type:int;not null;default:0;column:category_b_id;comment:二级分类ID;index:level_status_deleted" `
	CategoryCID        int32                   `gorm:"type:int;not null;default:0;column:category_c_id;comment:三级分类ID;index:level_status_deleted" `
	CategoryName       string                  `gorm:"type:varchar(100);not null;default:'';column:category_name;comment:分类名称"`
	SPUCode            string                  `gorm:"type:varchar(300);not null;default:'';column:spu_code;comment:商品SPU编号"`
	Name               string                  `gorm:"type:varchar(300);not null;default:'';column:name;comment:商品名称"`
	Status             enum.GoodsStatus        `gorm:"type:tinyint(1);not null;default:0;column:status;comment:状态:1-在售2-下架;index:level_status_deleted"`
	ShipFree           int                     `gorm:"type:tinyint(1);not null;default:0;column:ship_free;comment:是否包邮"`
	Description        string                  `gorm:"default:'';type:varchar(1000);column:description;not null;comment:商品详情"`
	Images             string                  `gorm:"default:'';type:varchar(5000);column:images;not null;comment:商品轮播图"`
	DescImages         string                  `gorm:"type:varchar(5000);not null;column:desc_images;comment:商品详情图"`
	GoodsFrontImage    string                  `gorm:"type:varchar(1000);not null;column:goods_front_image;comment:商品封面图"`
	IsRetail           int                     `gorm:"type:tinyint(1);not null;default:0;column:is_retail;comment:是否支持零售;index:level_status_deleted"`
	RetailStatus       enum.GoodsRetailStatus  `gorm:"type:tinyint(1);not null;default:0;column:status;comment:零售状态:1-正常2-售磬;index:level_status_deleted"`
	IsNew              int                     `gorm:"type:tinyint(1);not null;default:0;column:is_new;comment:是否为新品"`
	IsHot              int                     `gorm:"default:0;not null;column:is_hot;comment:是否为热卖商品"`
	SoldNum            int                     `gorm:"type:bigint;default:0;not null;column:sold_num;comment:零售销量"`
	WholesaleLogistics enum.WholesaleLogistics `gorm:"type:int;default:0;not null;column:wholesale_logistics;comment:批发物流"`
	WholesaleShipping  string                  `gorm:"type:varchar(500);default:'';not null;comment:批发发货地"`
	WholesaleAreaCodeA int                     `gorm:"type:bigint;not null;default:0;not null;column:wholesale_area_code_a;comment:筛选code省"`
	WholesaleAreaCodeB int                     `gorm:"type:bigint;not null;default:0;not null;column:wholesale_area_code_b;comment:筛选code区"`
	WholesaleAreaCodeC int                     `gorm:"type:bigint;not null;default:0;not null;column:wholesale_area_code_c;comment:筛选code县/市"`
	RetailShippingTime enum.RetailDeliveryTime `gorm:"type:int;not null;default:0;column:retail_shipping_time;comment:零售发货时间"`
	RetailShippingFee  float64                 `gorm:"type:decimal(9,2);not null;default:0.0;column:retail_shipping_fee;comment:零售运费"`
	CreatedAt          time.Time               `gorm:"comment:创建时间;not null;column:created_at;default:CURRENT_TIMESTAMP(3);index:created_at;index:level_status_deleted"`
	UpdatedAt          time.Time               `gorm:"comment:更新时间;not null;column:updated_at;default:CURRENT_TIMESTAMP(3) ON UPDATE CURRENT_TIMESTAMP(3);index:updated_at;index:level_status_deleted"`
	Deleted            int                     `gorm:"type:tinyint(1);not null;default:0;column:deleted;index:deleted;index:level_status_deleted"`
}

// Specification 商品规格表
type Specification struct {
	ID        int32                   `gorm:"primary_key;AUTO_INCREMENT;type:int" json:"id"`
	CreatedAt time.Time               `gorm:"comment:创建时间;not null;column:created_at;default:CURRENT_TIMESTAMP(3);" json:"-"`
	UpdatedAt time.Time               `gorm:"comment:更新时间;not null;column:updated_at;default:CURRENT_TIMESTAMP(3) ON UPDATE CURRENT_TIMESTAMP(3);" json:"-"`
	Deleted   int                     `gorm:"type:tinyint(1);not null;default:0;column:deleted;index:deleted;index:type_deleted" json:"-"`
	Name      string                  `gorm:"type:varchar(100);not null;column:name;comment:规格名称" json:"name"`
	GoodsID   int32                   `gorm:"type:int;not null;default:0;column:goods_id;comment:商品ID" json:"goods_id"`
	Type      enum.ProductVariantType `gorm:"type:tinyint(1);default:0;not null;column:type;comment:类型 1-批发 2-零售;index:type_deleted" json:"type"`
}

type SpecificationValue struct {
	ID              int32                   `gorm:"primary_key;AUTO_INCREMENT;type:int" json:"id"`
	CreatedAt       time.Time               `gorm:"comment:创建时间;not null;column:created_at;default:CURRENT_TIMESTAMP(3);" json:"-"`
	UpdatedAt       time.Time               `gorm:"comment:更新时间;not null;column:updated_at;default:CURRENT_TIMESTAMP(3) ON UPDATE CURRENT_TIMESTAMP(3);" json:"-"`
	Deleted         int                     `gorm:"type:tinyint(1);not null;default:0;column:deleted;index:deleted" json:"-"`
	Value           string                  `gorm:"type:varchar(100);not null;comment:规格属性值;index:specification_type_value_deleted" json:"value"`
	SpecificationID int32                   `gorm:"type:int;not null;default:0;column:specification_id;comment:规格ID;index:specification_type_value_deleted"`
	GoodsID         int32                   `gorm:"type:int;not null;default:0;column:goods_id;comment:商品ID"`
	Type            enum.ProductVariantType `gorm:"type:tinyint(1);default:0;not null;column:type;comment:类型 1-批发 2-零售;index:specification_type_value_deleted" json:"type"`
}

// ProductVariant 产品SKU表
type ProductVariant struct {
	BaseModel
	SKUCode          string                    `gorm:"type:varchar(300);not null;default:'';column:sku_code;comment:产品SKU编号;index:sku_code_deleted" json:"sku_code"`
	GoodsID          int32                     `gorm:"type:int;not null;default:0;comment:商品ID;index:goods_id_deleted_type_status" json:"goods_id"`
	Goods            *Goods                    `gorm:"foreignKey:GoodsID;references:ID" json:"goods"`
	SpecValueAID     int32                     `gorm:"type:int;comment:规格属性值A的ID" json:"specValueAID;index:spec_value_a_id_deleted"`
	SpecValueA       *SpecificationValue       `gorm:"foreignKey:SpecValueAID;references:ID" json:"specValueA"`
	SpecValueBID     int32                     `gorm:"type:int;comment:规格属性值B的ID" json:"specValueBID;index:spec_value_b_id_deleted"`
	SpecValueB       *SpecificationValue       `gorm:"foreignKey:SpecValueBID;references:ID" json:"specValueB"`
	Unit             string                    `gorm:"type:varchar(100);default:'';not null;column:unit;comment:单位" json:"unit"`
	Price            float64                   `gorm:"type:decimal(9,2);default:0.0;not null;column:price;comment:价格" json:"price"`
	Stock            int                       `gorm:"type:int;default:0;column:stock;not null;comment:库存" json:"stock"`
	MinOrderQuantity int                       `gorm:"type:int;default:0;column:min_order_quantity;comment:起批量" json:"min_order_quantity"`
	Type             enum.ProductVariantType   `gorm:"type:tinyint(1);default:0;not null;column:type;comment:类型 1-批发 2-零售;index:goods_id_deleted_type_status" json:"type"`
	Status           enum.ProductVariantStatus `gorm:"type:tinyint(1);not null;default:1;column:status;comment:状态 1-未启用 2-启用" json:"status"`
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

// GoodsApplication 商品申请表
type GoodsApplication struct {
	BaseModel
	VerifierID string `gorm:"type:varchar(100);not null;default:'';comment:审核人ID"`
	GoodsID    int32  `gorm:"type:bigint;not null;default:0;comment:商品ID"`
	Status     int    `gorm:"type:tinyint(1);not null;default:0;column:status;comment:状态" json:"status"`
	Comment    string `gorm:"column:comment;not null" json:"comment"`
}
