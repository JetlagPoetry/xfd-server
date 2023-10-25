package model

import (
	"gorm.io/gorm"
)

// Category 商品分类表
type Category struct {
	BaseModel
	Name             string      `gorm:"type:varchar(50);not null;comment:分类名称" json:"name"`
	ParentCategoryID int32       `json:"parent_category"`                                               // 父级分类ID
	ParentCategory   *Category   `json:"-"`                                                             // 序列化json数据时忽略该字段
	SubCategory      []*Category `gorm:"foreignKey:ParentCategoryID;references:ID" json:"sub_category"` // foreignKey：关联"另外一张表的键"；references：另外一张表关联"此表的主键"
	Level            int32       `gorm:"type:int;not null;default:1;comment:分类级别" json:"level"`
	Image            string      `gorm:"type:varchar(1000);not null;comment:分类图片概览"`
	Status           int32       `gorm:"type:tinyint(1);not null;default:1;comment:状态" json:"status"`
}

// Goods 商品表
type Goods struct {
	gorm.Model
	CategoryID      int32 `gorm:"type:int;default:0;not null"`
	Category        *Category
	BrandsID        int32    `gorm:"type:int;not null"`
	OnSale          bool     `gorm:"default:false;not null;comment:是否已上架"`
	ShipFree        bool     `gorm:"default:false;not null;comment:是否免运费"`
	Name            string   `gorm:"type:varchar(100);not null;comment:商品名称"`
	GoodsSn         string   `gorm:"type:varchar(100);not null;comment:商品编号"`
	GoodsBrief      string   `gorm:"type:varchar(100);not null;comment:商品简介"`
	Images          GormList `gorm:"type:varchar(1000);not null;comment:商品轮播图"`
	DescImages      GormList `gorm:"type:varchar(2000);not null;comment:商品详情图"`
	GoodsFrontImage string   `gorm:"type:varchar(200);not null;comment:商品封面图"`
	ClickNum        int32    `gorm:"type:int;default:0;not null;comment:商品点击数"`
	IsNew           bool     `gorm:"default:false;not null;comment:是否为新品"`
	IsHot           bool     `gorm:"default:false;not null;comment:是否为热卖商品"`
	SoldNum         int32    `gorm:"type:int;default:0;not null;comment:零售销量"`
	FavNum          int32    `gorm:"type:int;default:0;not null;comment:收藏数量"`
	Origin          string   `gorm:"type:varchar(100);not null;comment:商品产地"`
}

// Specification 商品规格表
type Specification struct {
	BaseModel
	Type    int     `gorm:"type:int;not null;comment:类型 1-批发 2-零售" json:"type"`
	Name    string  `gorm:"type:varchar(20);not null;comment:规格名称" json:"name"`
	Unit    string  `gorm:"type:varchar(2);default:斤;comment:单位" json:"unit"`
	Price   float64 `gorm:"type:decimal(9,2);not null;comment:价格" json:"price"`
	Stock   int     `gorm:"type:int;comment:库存" json:"stock"`
	Enabled bool    `gorm:"type:bool;default:true;comment:是否启用" json:"enabled"`
}

//自定义定义品类1 颜色
//
//定义品类2

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
