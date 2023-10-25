package model

import (
	"gorm.io/gorm"
)

type Category struct {
	gorm.Model
	Name             string      `gorm:"type:varchar(20);not null" json:"name"`
	ParentCategoryID int32       `json:"parent"`
	ParentCategory   *Category   `json:"-"`
	SubCategory      []*Category `gorm:"foreignKey:ParentCategoryID;references:ID" json:"sub_category"`
	Level            int32       `gorm:"type:int;not null;default:1" json:"level"`
	IsTab            bool        `gorm:"default:false;not null" json:"is_tab"`
}

// 商品分类表
type Category2 struct {
	gorm.Model
	Name             string      `gorm:"type:varchar(50);not null;comment:分类名称" json:"name"`
	ParentCategoryID int32       `json:"parent_category"`                                               // 父级分类ID
	ParentCategory   *Category   `json:"-"`                                                             // 序列化json数据时忽略该字段
	SubCategory      []*Category `gorm:"foreignKey:ParentCategoryID;references:ID" json:"sub_category"` // foreignKey：关联"另外一张表的键"；references：另外一张表关联"此表的主键"
	Level            int32       `gorm:"type:int;not null;default:1;comment:分类级别" json:"level"`
	IsTab            bool        `gorm:"default:false;not null;comment:是否在tab栏展示" json:"is_tab"`
}

// 商品品牌表
type Brands struct {
	gorm.Model
	Name string `gorm:"type:varchar(20);not null;comment:品牌名称"`
	Logo string `gorm:"type:varchar(200);default:'';not null;comment:品牌商标的url"`
}

// 商品分类-商品品牌 关系表
// CategoryID 与 BrandsID 组成了联合主键
type GoodsCategoryBrand struct {
	gorm.Model
	CategoryID int32 `gorm:"type:int;index:idx_category_brand,unique"`
	Category   *Category
	BrandsID   int32 `gorm:"type:int;index:idx_category_brand,unique"`
	Brands     *Brands
}

// 表名重载
//func (GoodsCategoryBrand) TableName() string {
//	return "goodscategorybrand"
//}

// 轮播图的表结构
type Banner struct {
	gorm.Model
	Image string `gorm:"type:varchar(200);not null;comment:品牌图片的url"`
	Url   string `gorm:"type:varchar(200);not null;comment:品牌详情页的url"`
	Index int32  `gorm:"type:int;default:1;not null;comment:轮播序号"`
}

// 商品表
type Goods struct {
	gorm.Model
	CategoryID      int32 `gorm:"type:int;not null"`
	Category        *Category
	BrandsID        int32 `gorm:"type:int;not null"`
	Brands          *Brands
	OnSale          bool     `gorm:"default:false;not null;comment:是否已上架"`
	ShipFree        bool     `gorm:"default:false;not null;comment:是否免运费"`
	IsNew           bool     `gorm:"default:false;not null;comment:是否为新品"`
	IsHot           bool     `gorm:"default:false;not null;comment:是否为热卖商品"`
	Name            string   `gorm:"type:varchar(100);not null;comment:商品名称"`
	GoodsSn         string   `gorm:"type:varchar(100);not null;comment:商品编号"`
	ClickNum        int32    `gorm:"type:int;default:0;not null;comment:商品点击数"`
	SoldNum         int32    `gorm:"type:int;default:0;not null;comment:销量"`
	FavNum          int32    `gorm:"type:int;default:0;not null;comment:收藏数量"`
	MarketPrice     float32  `gorm:"not null;comment:市场价(折扣前)"`
	ShopPrice       float32  `gorm:"not null;comment:本店价格(折扣后)"`
	GoodsBrief      string   `gorm:"type:varchar(100);not null;comment:商品简介"`
	Images          GormList `gorm:"type:varchar(1000);not null;comment:商品图概览"`
	DescImages      GormList `gorm:"type:varchar(2000);not null;comment:商品描述图"`
	GoodsFrontImage string   `gorm:"type:varchar(200);not null;comment:商品封面图"`
}

// gorm的钩子函数

//// 创建完商品后，设置商品库存，并将商品信息同步到es
//func (g *Goods) AfterCreate(tx *gorm.DB) (err error) {
//	// 调用库存服务，设置商品库存
//	// 设置商品库存必须在同步es之前，否则若先同步es，es同步成功后，
//	// 库存服务一旦调用失败，mysql中的商品数据会回滚，但是es中的商品数据不会回滚。
//	var stocks int32
//	if val, ok := tx.Get("good_stock"); ok {
//		stocks = val.(int32)
//	}
//
//	_, err = global.InventorySrvClient.SetInv(context.Background(), &proto.GoodsInvInfo{
//		GoodsId: g.ID,
//		Num:     stocks,
//	})
//	if err != nil {
//		zap.S().Error("调用库存服务出错:", err)
//		tx.Rollback()
//		return
//	}
//
//	// 将商品信息同步到es
//	good := &EsGoods{
//		ID:          g.ID,
//		CategoryID:  g.CategoryID,
//		BrandsID:    g.BrandsID,
//		OnSale:      g.OnSale,
//		ShipFree:    g.ShipFree,
//		IsNew:       g.IsNew,
//		IsHot:       g.IsHot,
//		Name:        g.Name,
//		ClickNum:    g.ClickNum,
//		SoldNum:     g.SoldNum,
//		FavNum:      g.FavNum,
//		MarketPrice: g.MarketPrice,
//		GoodsBrief:  g.GoodsBrief,
//		ShopPrice:   g.ShopPrice,
//	}
//	if _, err = global.EsClient.Index().
//		Index(good.GetIndexName()).
//		Id(strconv.FormatInt(int64(good.ID), 10)).
//		BodyJson(good).
//		Do(context.Background()); err != nil {
//		return
//	}
//	return nil
//}
//
//func (g *Goods) AfterUpdate(tx *gorm.DB) (err error) {
//	// 将商品信息同步到es
//	good := &EsGoods{
//		ID:          g.ID,
//		CategoryID:  g.CategoryID,
//		BrandsID:    g.BrandsID,
//		OnSale:      g.OnSale,
//		ShipFree:    g.ShipFree,
//		IsNew:       g.IsNew,
//		IsHot:       g.IsHot,
//		Name:        g.Name,
//		ClickNum:    g.ClickNum,
//		SoldNum:     g.SoldNum,
//		FavNum:      g.FavNum,
//		MarketPrice: g.MarketPrice,
//		GoodsBrief:  g.GoodsBrief,
//		ShopPrice:   g.ShopPrice,
//	}
//
//	_, err = global.EsClient.Update().
//		// 不执行立即刷新、保证性能
//		//Refresh("false").
//		Index(good.GetIndexName()).
//		Id(strconv.FormatInt(int64(g.ID), 10)).
//		// data为结构体或map, 需注意的是如果使用结构体零值也会去更新原记录
//		Doc(good).
//		// true 无则插入, 有则更新, 设置为false时记录不存在将报错
//		DocAsUpsert(false).
//		Do(context.Background())
//	return err
//}
//
//func (g *Goods) AfterDelete(tx *gorm.DB) (err error) {
//	_, err = global.EsClient.Delete().
//		Index(EsGoods{}.GetIndexName()).
//		Id(strconv.FormatInt(int64(g.ID), 10)).
//		Do(context.Background())
//	return err
//}
