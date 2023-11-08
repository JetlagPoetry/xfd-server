package types

import (
	"encoding/json"
	"errors"
	"github.com/google/martian/log"
	"xfd-backend/database/db/enum"
	"xfd-backend/database/db/model"
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
	specAName := r.WholesaleProducts[0].SpecAName
	specBName := r.WholesaleProducts[0].SpecBName
	for i, v := range r.WholesaleProducts {
		if err := r.WholesaleProducts[i].CheckParams(); err != nil {
			return err
		}
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
	if len(r.RetailProducts) != 0 {
		specAName = r.RetailProducts[0].SpecAName
		specBName = r.RetailProducts[0].SpecBName
		for i, v := range r.RetailProducts {
			if err := r.RetailProducts[i].CheckParams(); err != nil {
				return err
			}
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
	}
	return nil
}

type ProductWholesale struct {
	SpecAName        string                    `json:"specAName" binding:"required,gte=1,lte=2"`
	SpecBName        string                    `json:"specBName"`
	SpecAValue       string                    `json:"specAValue" binding:"required"`
	SpecBValue       string                    `json:"specBValue"`
	Unit             string                    `json:"unit" binding:"required,gte=1,lte=10"`
	Price            float64                   `json:"price" binding:"required"`
	MinOrderQuantity int                       `json:"minOrderQuantity" binding:"required"`
	Status           enum.ProductVariantStatus `json:"status" binding:"required,gte=1,lte=2"`
}

func (p ProductWholesale) CheckParams() error {
	if (p.SpecBName == "" && p.SpecBValue != "") || (p.SpecBName != "" && p.SpecBValue == "") {
		return errors.New("specBName and specBValue must be either both filled or both empty")
	}
	return nil
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
	ID               int32                `json:"id"`
	Unit             string               `json:"unit"`
	Price            float64              `json:"price"`
	MinOrderQuantity int                  `json:"minOrderQuantity,omitempty"`
	Stock            int                  `json:"stock,omitempty"`
	SKUCode          string               `json:"skuCode"`
	ProductAttr      []*model.ProductAttr `json:"productAttr,omitempty"`
}

type SpecValue struct {
	ID    int32  `json:"id"`
	Value string `json:"value"`
}

type SpecInfo struct {
	SpecID    int32        `json:"specID"`
	SpecName  string       `json:"specName"`
	SpecValue []*SpecValue `json:"specValue"`
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
	RetailShippingFee       float64                 `json:"retailShippingFee"`       // 零售运费 8-0/0-8/8-8
	ChangeRetailShippingFee bool                    `json:"changeRetailShippingFee"` // 是否修改零售运费
	SpecInfo                []*SpecInfo             `json:"specInfo" binding:"required"`
	WholesaleProducts       []*ModifyProduct        `json:"wholesaleProducts"`
	RetailProducts          []*ModifyProduct        `json:"retailProducts"`
}

type ModifyProduct struct {
	ID               int32                     `json:"id"`
	SpecAName        string                    `json:"specAName"`
	SpecANameID      int32                     `json:"specANameID"`
	SpecBName        string                    `json:"specBName" `
	SpecBNameID      int32                     `json:"specBNameID"`
	SpecAValue       string                    `json:"specAValue"`
	SpecAValueID     int32                     `json:"specAValueID"`
	SpecBValue       string                    `json:"specBValue"`
	SpecBValueID     int32                     `json:"specBValueID"`
	Unit             string                    `json:"unit"`
	Price            float64                   `json:"price"`
	MinOrderQuantity int                       `json:"minOrderQuantity"`
	Status           enum.ProductVariantStatus `json:"status"`
	Stock            int                       `json:"stock" binding:"gte=0,lte=9999999"`
}

func (g *GoodsModifyReq) CheckParams() error {
	if g.WholesaleShipping != "" {
		if g.WholesaleAreaCodeA == 0 || g.WholesaleAreaCodeB == 0 || g.WholesaleAreaCodeC == 0 {
			return errors.New("wholesaleAreaCodeA,wholesaleAreaCodeB,wholesaleAreaCodeC must have a value")
		}
	}
	if g.WholesaleProducts != nil {
		for _, v := range g.WholesaleProducts {
			if err := v.CheckParams(); err != nil {
				return err
			}
		}
	}
	if g.RetailProducts != nil {
		for _, v := range g.RetailProducts {
			if err := v.CheckParams(); err != nil {
				return err
			}
		}
	}
	if g.CategoryName != "" {
		if g.CategoryAID == 0 || g.CategoryBID == 0 || g.CategoryCID == 0 {
			return errors.New("categoryAID,categoryBID,categoryCID must have a value")
		}
	}
	return nil
}

func (m *ModifyProduct) CheckParams() error {
	if m.ID == 0 {
		if m.SpecAName == "" || m.SpecBName == "" || m.SpecAValue == "" || m.SpecBValue == "" ||
			m.Unit == "" || m.Price == 0 || m.Status == 0 {
			return errors.New("SpecAName, SpecBName, SpecAValue, SpecBValue, Unit, Price, and Status are required")
		}
		if m.MinOrderQuantity == 0 && m.Stock == 0 {
			return errors.New("either MinOrderQuantity or Stock must have a value")
		}
	}
	if m.SpecAName != "" && m.SpecBName != "" {
		if m.SpecAName == m.SpecBName {
			return errors.New("规格A不能等于规格B")
		}
	}
	if m.SpecAName != "" && m.SpecANameID == 0 {
		return errors.New("specANameID must have a value")
	}
	if m.SpecBName != "" && m.SpecBNameID == 0 {
		return errors.New("specBNameID must have a value")
	}
	if m.SpecAValue != "" && m.SpecBValue != "" {
		if m.SpecAValue == m.SpecAValue {
			return errors.New("规格A的属性值不能等于规格B的属性值")
		}
	}
	return nil
}
