package types

import (
	"encoding/json"
	"errors"
	"github.com/google/martian/log"
	"xfd-backend/database/db/enum"
	"xfd-backend/database/db/model"
)

type GoodsAddReq struct {
	GoodsDetail       GoodsDetail   `json:"goodsDetail"`
	WholesaleProducts []*AddProduct `json:"wholesaleProducts"`
	RetailProducts    []*AddProduct `json:"retailProducts"`
}

type ProductWholesale struct {
	ProductAttr      []*model.ProductAttr      `json:"productAttr" binding:"required"`
	Unit             string                    `json:"unit" binding:"required,gte=1,lte=10"`
	Price            float64                   `json:"price" binding:"required"`
	Status           enum.ProductVariantStatus `json:"status" binding:"required,gte=1,lte=2"`
	MinOrderQuantity int                       `json:"minOrderQuantity" binding:"required"`
	Stock            int                       `json:"stock" binding:"required,gte=1,lte=9999999"`
}
type AddProduct struct {
	ProductAttr      []*model.ProductAttr      `json:"productAttr" binding:"required"`
	Unit             string                    `json:"unit" binding:"required,gte=1,lte=10"`
	Price            float64                   `json:"price" binding:"required"`
	Status           enum.ProductVariantStatus `json:"status" binding:"required,gte=1,lte=2"`
	MinOrderQuantity int                       `json:"minOrderQuantity" binding:"required"`
	Stock            int                       `json:"stock" binding:"required,gte=1,lte=9999999"`
}

type ProductRetail struct {
	SpecAName  string                    `json:"specAName" binding:"required,gte=1,lte=2"`
	SpecBName  string                    `json:"specBName"`
	SpecAValue string                    `json:"specAValue" binding:"required"`
	SpecBValue string                    `json:"specBValue"`
	Unit       string                    `json:"unit" binding:"required,gte=1,lte=10"`
	Price      float64                   `json:"price" binding:"required"`
	Stock      int                       `json:"stock" binding:"required,gte=1,lte=9999999"`
	Status     enum.ProductVariantStatus `json:"status" binding:"required,gte=1,lte=2"`
}

func (p ProductRetail) CheckParams() error {
	if (p.SpecBName == "" && p.SpecBValue != "") || (p.SpecBName != "" && p.SpecBValue == "") {
		return errors.New("specBName and specBValue must be either both filled or both empty")
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
	WholesaleLogistics []int                   `json:"wholesaleLogistics" binding:"required"`             // 批发物流
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
	GoodsID     int32            `json:"goodsID" form:"goodsID" binding:"required,numeric,gte=1"`
	GoodsStatus enum.GoodsStatus `json:"goodsStatus" binding:"numeric,gte=0,lte=2"`
	ReqType     int
}

type MyGoodsListReq struct {
	PageRequest
	QueryGoodsListStatus enum.QueryGoodsListStatus `form:"queryGoodsListStatus" binding:"numeric,gte=0,lte=3"`
	UserID               string
}

// GoodsDetailResp 商品详情信息
type GoodsDetailResp struct {
	ID                 int32                   `json:"id"`
	Name               string                  `json:"name"`
	GoodsFrontImage    string                  `json:"goodsFrontImage"`              // 商品首图
	Images             []string                `json:"images"`                       // 商品轮播图
	Description        string                  `json:"description"`                  // 商品详情
	DescImages         []string                `json:"descImages"`                   // 商品详情图
	WholesaleLogistics []int                   `json:"wholesaleLogistics,omitempty"` // 批发物流
	WholesaleShipping  string                  `json:"wholesaleShipping,omitempty"`  // 批发发货地
	RetailShippingTime enum.RetailDeliveryTime `json:"retailShippingTime,omitempty"`
	RetailShippingFee  float64                 `json:"retailShippingFee,omitempty"`
	SpecInfo           []*SpecInfo             `json:"specInfo"`
	WholesaleProducts  []*ProductVariantInfo   `json:"wholesaleProducts,omitempty"`
	RetailProducts     []*ProductVariantInfo   `json:"retailProduct,omitempty"`
}

type ProductVariantInfo struct {
	ID               int32                     `json:"id"`
	Unit             string                    `json:"unit"`
	Price            float64                   `json:"price"`
	MinOrderQuantity int                       `json:"minOrderQuantity,omitempty"`
	Stock            int                       `json:"stock,omitempty"`
	SKUCode          string                    `json:"skuCode"`
	Status           enum.ProductVariantStatus `json:"status"`
	ProductAttr      []*model.ProductAttr      `json:"productAttr,omitempty"`
}

type SpecValue struct {
	ID    int32  `json:"id"`
	Value string `json:"value"`
}

type SpecInfo struct {
	SpecID    int32                   `json:"specID"`
	SpecName  string                  `json:"specName"`
	Type      enum.ProductVariantType `json:"type"`
	SpecValue []*SpecValue            `json:"specValue"`
}

type GoodsModifyReq struct {
	Id                      int32                   `json:"id" binding:"required,gte=1,numeric"`
	CategoryAID             int32                   `json:"categoryAID"`
	CategoryBID             int32                   `json:"categoryBID"`
	CategoryCID             int32                   `json:"categoryCID"`
	CategoryName            string                  `json:"categoryName"`
	Name                    string                  `json:"name"`
	GoodsFrontImage         string                  `json:"goodsFrontImage"`
	Images                  []string                `json:"images"`
	Description             string                  `json:"description"`
	DescImages              []string                `json:"descImages"`
	WholesaleLogistics      []int                   `json:"wholesaleLogistics"`
	WholesaleShipping       string                  `json:"wholesaleShipping"`
	WholesaleAreaCodeA      int                     `json:"wholesaleAreaCodeA" `     // 筛选code省
	WholesaleAreaCodeB      int                     `json:"wholesaleAreaCodeB" `     // 筛选code区
	WholesaleAreaCodeC      int                     `json:"wholesaleAreaCodeC"`      // 筛选code县/市
	RetailShippingTime      enum.RetailDeliveryTime `json:"retailShippingTime"`      // 零售发货时间
	RetailShippingFee       *float64                `json:"retailShippingFee"`       // 零售运费 8-0/0-8/8-8
	ChangeRetailShippingFee bool                    `json:"changeRetailShippingFee"` // 是否修改零售运费
	WholesaleProducts       []*ModifyProduct        `json:"wholesaleProducts"`
	RetailProducts          []*ModifyProduct        `json:"retailProducts"`
}

type ModifyProduct struct {
	ID               int32                     `json:"id"`
	Unit             string                    `json:"unit"`
	Price            float64                   `json:"price"`
	MinOrderQuantity int                       `json:"minOrderQuantity"`
	Status           enum.ProductVariantStatus `json:"status"`
	Stock            *int                      `json:"stock" binding:"gte=1,lte=9999999"`
	ProductAttr      []*model.ProductAttr      `json:"productAttr"`
}

func (m *ModifyProduct) HandleProductAttr() (map[string]string, map[string]bool, map[string]bool) {
	valueKeyMap := make(map[string]string)
	keySet := make(map[string]bool)
	valueSet := make(map[string]bool)
	if len(m.ProductAttr) != 0 {
		for _, attr := range m.ProductAttr {
			keySet[attr.Key] = true
			valueSet[attr.Value] = true
			valueKeyMap[attr.Value] = attr.Key
		}
	}
	return valueKeyMap, keySet, valueSet
}

func (m *ModifyProduct) CheckParams() error {
	if m.Stock != nil {
		if *m.Stock < 1 || *m.Stock > 9999999 {
			return errors.New("库存数量必须在1-9999999之间")
		}
	}
	if m.ID == 0 {
		if m.Unit == "" || m.Price == 0 || m.Status == 0 || m.ProductAttr == nil {
			return errors.New("商品规格信息不完整")
		}
	}
	return nil
}

func (g *GoodsModifyReq) CheckParams() error {
	if g.WholesaleShipping != "" {
		if g.WholesaleAreaCodeA == 0 || g.WholesaleAreaCodeB == 0 || g.WholesaleAreaCodeC == 0 {
			return errors.New("修改了批发发货地，批发发货地code不能为空")
		}
	}
	if g.CategoryName != "" {
		if g.CategoryAID == 0 || g.CategoryBID == 0 || g.CategoryCID == 0 {
			return errors.New("修改了分类名称，分类ID不能为空")
		}
	}
	return nil
}
