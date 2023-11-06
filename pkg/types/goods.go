package types

import (
	"encoding/json"
	"errors"
	"github.com/google/martian/log"
	"xfd-backend/database/db/enum"
)

type GoodsAddReq struct {
	GoodsDetail       GoodsDetail        `json:"goodsDetail"`
	WholesaleProducts []ProductWholesale `json:"wholesaleProducts"`
	RetailProducts    []ProductRetail    `json:"retailProducts"`
}

func (r GoodsAddReq) CheckParams() error {
	if err := r.GoodsDetail.CheckParams(); err != nil {
		return err
	}
	if len(r.WholesaleProducts) == 0 {
		return errors.New("批发商品不能为空")
	}

	// 检查第一个 ProductVariant 的 SpecAName和 SpecBName
	specAName := r.WholesaleProducts[0].SpecAName
	specBName := r.WholesaleProducts[0].SpecBName
	for _, v := range r.WholesaleProducts {
		// 检查 SpecAName 和 SpecBName 是否不相等
		if v.SpecAName == v.SpecBName {
			return errors.New("SpecAName 和 SpecBName 不能相等")
		}
		// 检查 SpecAName 是否一致
		if v.SpecAName != specAName {
			return errors.New("所有的 SpecAName 必须相同")
		}
		// 检查 SpecBName 是否一致
		if v.SpecBName != specBName {
			return errors.New("所有的 SpecBName 必须相同")
		}
	}
	specAName = r.RetailProducts[0].SpecAName
	specBName = r.RetailProducts[0].SpecBName
	for _, v := range r.RetailProducts {
		// 检查 SpecAName 和 SpecBName 是否不相等
		if v.SpecAName == v.SpecBName {
			return errors.New("SpecAName 和 SpecBName 不能相等")
		}
		// 检查 SpecAName 是否一致
		if v.SpecAName != specAName {
			return errors.New("所有的 SpecAName 必须相同")
		}
		// 检查 SpecBName 是否一致
		if v.SpecBName != specBName {
			return errors.New("所有的 SpecBName 必须相同")
		}
	}
	return nil
}

type ProductWholesale struct {
	SpecAName        string                    `json:"specAName" binding:"required,gte=1,lte=2"`
	SpecBName        string                    `json:"specBName" binding:"required,gte=1,lte=2"`
	SpecAValue       string                    `json:"specAValue" binding:"required"`
	SpecBValue       string                    `json:"specBValue" binding:"required"`
	Unit             string                    `json:"unit" binding:"required,gte=1,lte=10"`
	Price            float64                   `json:"price" binding:"required"`
	MinOrderQuantity int                       `json:"minOrderQuantity" binding:"required"`
	Status           enum.ProductVariantStatus `json:"status" binding:"required,gte=1,lte=2"`
}

type ProductRetail struct {
	SpecAName  string                    `json:"specAName" binding:"required,gte=1,lte=2"`
	SpecBName  string                    `json:"specBName" binding:"required,gte=1,lte=2"`
	SpecAValue string                    `json:"specAValue" binding:"required"`
	SpecBValue string                    `json:"specBValue" binding:"required"`
	Unit       string                    `json:"unit" binding:"required,gte=1,lte=10"`
	Price      float64                   `json:"price" binding:"required"`
	Stock      int                       `json:"stock" binding:"required,gte=1,lte=9999999"`
	Status     enum.ProductVariantStatus `json:"status" binding:"required,gte=1,lte=2"`
}

func (r ProductRetail) CheckParams() error {
	if r.Price <= 0 {
		return errors.New("商品价格不能小于等于0")
	}
	if r.SpecAName == r.SpecBName {
		return errors.New("规格A不能等于规格B")
	}
	if r.SpecAName == r.SpecBName {
		return errors.New("规格A不能等于规格B")
	}
	if r.SpecAValue == r.SpecAValue {
		return errors.New("规格A的属性值不能等于规格B的属性值")
	}
	return nil
}

type GoodsDetail struct {
	CategoryAID        int32                   `json:"categoryAID" binding:"required"`
	CategoryBID        int32                   `json:"categoryBID" binding:"required"`
	CategoryCID        int32                   `json:"categoryCID"`
	CategoryName       string                  `json:"categoryName" binding:"required"`
	Name               string                  `json:"name" binding:"required,gte=1,lte=100"`             // 商品名称
	Description        string                  `json:"description"`                                       // 商品详情
	Images             []string                `json:"images" binding:"required"`                         // 商品轮播图
	DescImages         []string                `json:"descImages"`                                        // 商品详情图
	WholesaleLogistics enum.WholesaleLogistics `json:"wholesaleLogistics" binding:"required,gte=1,lte=6"` // 批发物流
	WholesaleShipping  string                  `json:"wholesaleShipping" binding:"required"`              // 批发发货地
	WholesaleAreaCodeA int                     `json:"wholesaleAreaCodeA" binding:"required"`             // 筛选code省
	WholesaleAreaCodeB int                     `json:"wholesaleAreaCodeB" binding:"required"`             // 筛选code区
	WholesaleAreaCodeC int                     `json:"wholesaleAreaCodeC" binding:"required"`             // 筛选code县/市
	RetailShippingTime enum.RetailDeliveryTime `json:"retailShippingTime" binding:"required,gte=1,lte=3"` // 零售发货时间
	RetailShippingFee  float64                 `json:"retailShippingFee"`                                 // 零售运费
}

func (r GoodsDetail) CheckParams() error {
	if len(r.Images) == 0 {
		return errors.New("商品轮播图不能为空")
	}
	return nil
}

type GoodsListReq struct {
	PageRequest
	CategoryAID        int32              `form:"categoryAID" binding:"numeric"`
	CategoryBID        int32              `form:"categoryBID" binding:"numeric"`
	CategoryCID        int32              `form:"categoryCID" binding:"numeric"`
	ListType           enum.GoodsListType `form:"listType" binding:"required,gte=1,lte=4"`
	ProductVariantType enum.ProductVariantType
	IsRetail           int
	UserID             string
	GoodsIDs           []int32
}

func (r GoodsListReq) CheckParams() error {
	if r.ListType == enum.GoodsListTypeCategory {
		if r.CategoryAID == 0 {
			return errors.New("一级分类ID不能为空")
		}
	}
	return nil
}

type GoodsListResp struct {
	PageResult
	GoodsList []*GoodsList `json:"goodsList"`
}

type GoodsList struct {
	ID                int32            `json:"id" gorm:"column:id"`
	Name              string           `json:"name" gorm:"column:name"`
	GoodsFrontImage   string           `json:"goodsFrontImage" gorm:"column:goods_front_image"`
	Images            string           `json:"-" gorm:"column:images"`
	RetailPriceMax    float64          `json:"retailPriceMax,omitempty"`
	RetailPriceMin    float64          `json:"retailPriceMin,omitempty"`
	WholesalePriceMax float64          `json:"wholesalePriceMax,omitempty"`
	WholesalePriceMin float64          `json:"wholesalePriceMin,omitempty"`
	WholeSaleUnit     string           `json:"wholeSaleUnit,omitempty"`
	RetailUnit        string           `json:"retailUnit,omitempty"`
	CreatedAt         string           `json:"createdAt,omitempty" gorm:"column:created_at"`
	UpdatedAt         string           `json:"updatedAt,omitempty" gorm:"column:updated_at"`
	Status            enum.GoodsStatus `json:"status,omitempty" gorm:"column:status"`
	SoldNum           int              `json:"soldNum,omitempty" gorm:"column:sold_num"`
	SPUCode           string           `json:"spuCode,omitempty" gorm:"column:spu_code"`
}

func (g *GoodsList) GetGoodsFrontImage() string {
	if g.GoodsFrontImage != "" {
		return g.GoodsFrontImage
	}
	if g.Images != "" {
		var images []string
		err := json.Unmarshal([]byte(g.Images), &images)
		if err != nil {
			log.Errorf("GetGoodsFrontImage json.Unmarshal error: %v, data: %s", err, g.Images)
			return ""
		}
		if len(images) > 0 {
		}
		return images[0]
	}
	return ""
}

type MinPriceResult struct {
	GoodsID         int32   `gorm:"column:goods_id"`
	MinPrice        float64 `gorm:"column:min_price"`
	Name            string  `gorm:"column:name"`
	GoodsFrontImage string  `gorm:"column:goods_front_image"`
	Images          string  `gorm:"column:images"`
}

type GoodsReq struct {
	GoodsID     int32            `json:"goodsID" binding:"required,numeric,gte=1"`
	GoodsStatus enum.GoodsStatus `json:"goodsStatus" binding:"numeric,gte=0,lte=2"`
}

type MyGoodsListReq struct {
	PageRequest
	QueryGoodsListStatus enum.QueryGoodsListStatus `form:"queryGoodsListStatus" binding:"numeric,gte=0,lte=3"`
	UserID               string
}
