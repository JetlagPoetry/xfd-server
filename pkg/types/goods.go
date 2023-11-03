package types

import (
	"errors"
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
	Status           enum.ProductVariantStatus `json:"status" binding:"required"`
}

type ProductRetail struct {
	SpecAName  string                    `json:"specAName" binding:"required,gte=1,lte=2"`
	SpecBName  string                    `json:"specBName" binding:"required,gte=1,lte=2"`
	SpecAValue string                    `json:"specAValue" binding:"required"`
	SpecBValue string                    `json:"specBValue" binding:"required"`
	Unit       string                    `json:"unit" binding:"required,gte=1,lte=10"`
	Price      float64                   `json:"price" binding:"required"`
	Stock      int                       `json:"stock" binding:"required,gte=1,lte=9999999"`
	Status     enum.ProductVariantStatus `json:"status" binding:"required"`
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
	CategoryAID   int32                  `form:"categoryAID" binding:"numeric"`
	CategoryBID   int32                  `form:"categoryBID" binding:"numeric"`
	CategoryCID   int32                  `form:"categoryCID" binding:"numeric"`
	RequestSource enum.RequestSourceType `form:"requestSource" binding:"required,gte=1,lte=2"`
}
