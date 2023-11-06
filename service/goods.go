package service

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
	"strconv"
	"time"
	"xfd-backend/database/db"
	"xfd-backend/database/db/dao"
	"xfd-backend/database/db/enum"
	"xfd-backend/database/db/model"
	"xfd-backend/database/repo"
	"xfd-backend/pkg/types"
	"xfd-backend/pkg/utils"
	"xfd-backend/pkg/xerr"
)

type GoodsService struct {
	goods   *dao.GoodsDao
	userDao *dao.UserDao
}

func NewGoodsService() *GoodsService {
	return &GoodsService{
		goods:   dao.NewGoodsDao(),
		userDao: dao.NewUserDao(),
	}
}

func (s *GoodsService) AddGoods(ctx *gin.Context, req types.GoodsAddReq) xerr.XErr {
	//检查参数
	if err := req.CheckParams(); err != nil {
		return xerr.WithCode(xerr.InvalidParams, err)
	}
	////获取用户ID
	//userID := common.GetUserID(ctx)
	////获取用户角色
	//userRole := common.GetUserRole(ctx)
	//if userRole != model.UserRoleSupplier {
	//	return xerr.WithCode(xerr.ErrorAuthInsufficientAuthority, errors.New("用户权限不是供应商"))
	//}
	userID := "w2wwww"
	spuCode := fmt.Sprintf("SP%s%s%s", utils.TimeFormatUs(), utils.GenSixDigitCode(), utils.GenerateRandCode("", 5))
	var images, descImages string
	if len(req.GoodsDetail.Images) != 0 {
		imagesBytes, _ := json.Marshal(req.GoodsDetail.Images)
		images = string(imagesBytes)
	}
	if len(req.GoodsDetail.DescImages) == 0 {
		descImages = images
	} else {
		descImagesBytes, _ := json.Marshal(req.GoodsDetail.DescImages)
		descImages = string(descImagesBytes)
	}

	modelGoods := &model.Goods{
		UserID:             userID,
		CategoryAID:        req.GoodsDetail.CategoryAID,
		CategoryBID:        req.GoodsDetail.CategoryBID,
		CategoryCID:        req.GoodsDetail.CategoryCID,
		CategoryName:       req.GoodsDetail.CategoryName,
		SPUCode:            spuCode,
		Name:               req.GoodsDetail.Name,
		Status:             enum.GoodsStatusOnSale,
		Description:        req.GoodsDetail.Description,
		Images:             images,
		DescImages:         descImages,
		GoodsFrontImage:    req.GoodsDetail.Images[0],
		RetailStatus:       enum.GoodsRetailNormal,
		IsRetail:           0,
		WholesaleLogistics: req.GoodsDetail.WholesaleLogistics,
		WholesaleShipping:  req.GoodsDetail.WholesaleShipping,
		WholesaleAreaCodeA: req.GoodsDetail.WholesaleAreaCodeA,
		WholesaleAreaCodeB: req.GoodsDetail.WholesaleAreaCodeB,
		WholesaleAreaCodeC: req.GoodsDetail.WholesaleAreaCodeC,
		RetailShippingTime: req.GoodsDetail.RetailShippingTime,
		RetailShippingFee:  req.GoodsDetail.RetailShippingFee,
	}
	//全国包邮
	if req.GoodsDetail.RetailShippingFee == 0 {
		modelGoods.ShipFree = 1
	}
	totalStock := 0
	for _, v := range req.RetailProducts {
		if v.Status == enum.ProductVariantEnabled {
			modelGoods.IsRetail = 1
			totalStock += v.Stock
		}
	}
	if totalStock == 0 {
		modelGoods.RetailStatus = enum.GoodsRetailSoldOut
	}
	transactionHandler := repo.NewTransactionHandler(db.Get())
	if xErr := transactionHandler.WithTransaction(ctx, func(ctx context.Context) xerr.XErr {
		if _err := s.CreateWithTransaction(ctx, req, modelGoods); _err != nil {
			return xerr.WithCode(xerr.ErrorDatabase, _err)
		}
		return nil
	}); xErr != nil {
		return xErr
	}
	return nil
}

func (s *GoodsService) CreateWithTransaction(ctx context.Context, req types.GoodsAddReq, modelGoods *model.Goods) error {
	//创建商品信息
	goodsID, err := s.goods.CreateGoods(ctx, modelGoods)
	if err != nil {
		return xerr.WithCode(xerr.ErrorDatabase, err)
	}
	specificationWA := &model.Specification{
		Name:    req.WholesaleProducts[0].SpecAName,
		GoodsID: goodsID,
		Type:    enum.ProductWholesale,
	}
	specificationWB := &model.Specification{
		Name:    req.WholesaleProducts[0].SpecBName,
		GoodsID: goodsID,
		Type:    enum.ProductWholesale,
	}

	//创建批发商品规格
	specificationWAID, err := s.goods.CreateSpecification(ctx, specificationWA)
	if err != nil {
		return xerr.WithCode(xerr.ErrorDatabase, err)
	}
	specificationWBID, err := s.goods.CreateSpecification(ctx, specificationWB)
	if err != nil {
		return xerr.WithCode(xerr.ErrorDatabase, err)
	}

	//创建批发商品规格属性
	specificationAValueMap := make(map[string]int32)
	specificationBValueMap := make(map[string]int32)
	for _, v := range req.WholesaleProducts {
		specificationAValue := &model.SpecificationValue{
			Value:           v.SpecAValue,
			SpecificationID: specificationWAID,
			GoodsID:         goodsID,
			Type:            enum.ProductWholesale,
		}
		specificationBValue := &model.SpecificationValue{
			Value:           v.SpecBValue,
			SpecificationID: specificationWBID,
			GoodsID:         goodsID,
			Type:            enum.ProductWholesale,
		}
		if _, ok := specificationAValueMap[v.SpecAValue]; !ok {
			specValueAID, err := s.goods.CreateSpecificationValue(ctx, specificationAValue)
			if err != nil {
				return xerr.WithCode(xerr.ErrorDatabase, err)
			}
			specificationAValueMap[v.SpecAValue] = specValueAID
		}
		if _, ok := specificationBValueMap[v.SpecBValue]; !ok {
			specValueBID, err := s.goods.CreateSpecificationValue(ctx, specificationBValue)
			if err != nil {
				return xerr.WithCode(xerr.ErrorDatabase, err)
			}
			specificationBValueMap[v.SpecBValue] = specValueBID
		}
		skuCode := fmt.Sprintf("SK%d%s%d%d", time.Now().UnixNano(), utils.GenerateRandCode("", 5), specificationBValueMap[v.SpecAValue], specificationBValueMap[v.SpecBValue])
		productVariant := &model.ProductVariant{
			SKUCode:          skuCode,
			GoodsID:          goodsID,
			SpecValueAID:     specificationAValueMap[v.SpecAValue],
			SpecValueBID:     specificationBValueMap[v.SpecBValue],
			Unit:             v.Unit,
			Price:            v.Price,
			MinOrderQuantity: v.MinOrderQuantity,
			Type:             enum.ProductWholesale,
			Status:           v.Status,
		}

		_, err = s.goods.CreateProductVariant(ctx, productVariant)
		if err != nil {
			return xerr.WithCode(xerr.ErrorDatabase, err)
		}
	}
	//创建零售产品
	if len(req.RetailProducts) != 0 {
		specificationRA := &model.Specification{
			Name:    req.RetailProducts[0].SpecAName,
			GoodsID: goodsID,
			Type:    enum.ProductRetail,
		}
		specificationRB := &model.Specification{
			Name:    req.RetailProducts[0].SpecBName,
			GoodsID: goodsID,
			Type:    enum.ProductRetail,
		}
		//创建零售规格
		specificationRAID, err := s.goods.CreateSpecification(ctx, specificationRA)
		if err != nil {
			return xerr.WithCode(xerr.ErrorDatabase, err)
		}
		specificationRBID, err := s.goods.CreateSpecification(ctx, specificationRB)
		if err != nil {
			return xerr.WithCode(xerr.ErrorDatabase, err)
		}
		//创建零售规格属性
		specificationAValueRMap := make(map[string]int32)
		specificationBValueRMap := make(map[string]int32)
		for _, v := range req.RetailProducts {
			specificationAValue := &model.SpecificationValue{
				Value:           v.SpecAValue,
				SpecificationID: specificationRAID,
				GoodsID:         goodsID,
				Type:            enum.ProductRetail,
			}
			specificationBValue := &model.SpecificationValue{
				Value:           v.SpecBValue,
				SpecificationID: specificationRBID,
				GoodsID:         goodsID,
				Type:            enum.ProductRetail,
			}
			if _, ok := specificationAValueRMap[v.SpecAValue]; !ok {
				specValueAID, err := s.goods.CreateSpecificationValue(ctx, specificationAValue)
				if err != nil {
					return xerr.WithCode(xerr.ErrorDatabase, err)
				}
				specificationAValueRMap[v.SpecAValue] = specValueAID
			}
			if _, ok := specificationBValueRMap[v.SpecBValue]; !ok {
				specValueBID, err := s.goods.CreateSpecificationValue(ctx, specificationBValue)
				if err != nil {
					return xerr.WithCode(xerr.ErrorDatabase, err)
				}
				specificationBValueRMap[v.SpecBValue] = specValueBID
			}
			skuCode := fmt.Sprintf("SK%s%s%d%d", strconv.FormatInt(time.Now().UnixNano(), 10), utils.GenerateRandCode("", 5), specificationAValueRMap[v.SpecAValue], specificationBValueRMap[v.SpecBValue])
			productVariant := &model.ProductVariant{
				SKUCode:      skuCode,
				GoodsID:      goodsID,
				SpecValueAID: specificationAValueRMap[v.SpecAValue],
				SpecValueBID: specificationBValueRMap[v.SpecBValue],
				Unit:         v.Unit,
				Price:        v.Price,
				Type:         enum.ProductRetail,
				Status:       v.Status,
				Stock:        v.Stock,
			}
			_, err = s.goods.CreateProductVariant(ctx, productVariant)
			if err != nil {
				return xerr.WithCode(xerr.ErrorDatabase, err)
			}
		}
	}
	return nil
}

func (s *GoodsService) GetGoodsList(ctx *gin.Context, req types.GoodsListReq) (*types.GoodsListResp, xerr.XErr) {
	if req.CheckParams() != nil {
		return nil, xerr.WithCode(xerr.InvalidParams, req.CheckParams())
	}
	userID := "w2wwww"

	user, err := s.userDao.GetByUserID(ctx, userID)
	if err != nil {
		return nil, xerr.WithCode(xerr.ErrorDatabase, err)
	}
	//供货商,查看供货信息(批发)
	if user.UserRole == model.UserRoleSupplier {
		req.ProductVariantType = enum.ProductWholesale
		req.UserID = userID
		switch req.ListType {
		case enum.GoodsListTypeNormal:
			return s.normalGetGoodsList(ctx, req)
		case enum.GoodsListTypeCategory:
			return s.normalGetGoodsList(ctx, req)
		default:
			return nil, xerr.WithCode(xerr.InvalidParams, fmt.Errorf("listType:%d is not supported userRole:%d ", req.ListType, user.UserRole))
		}
	}
	if user.UserRole == model.UserRoleCustomer {
		req.IsRetail = 1
		req.ProductVariantType = enum.ProductRetail
		//消费者只能看到上架的的零售商品,默认按照发布时间由新到旧展示,筛选满足条件的商品
		// 2. 综合排序：按货品发布时间由新到旧排列
		// 3. 销量优先：按货品近30天累计支付成功的订单总量降序排列
		// 4. 低价优先：按货品最低价格升序排列（一个货品可以有多个sku、多个价格，这里按其sku中的最低价格排序）
		switch req.ListType {
		case enum.GoodsListTypeNormal:
			return s.normalGetGoodsList(ctx, req)
		case enum.GoodsListTypeCategory:
			return s.normalGetGoodsList(ctx, req)
		case enum.GoodsListTypeSale:
			return s.normalGetGoodsList(ctx, req)
		case enum.GoodsListTypePrice:
			return s.getGoodsListByPrice(ctx, req)
		default:
			return nil, xerr.WithCode(xerr.InvalidParams, fmt.Errorf("listType:%d is not supported userRole:%d ", req.ListType, user.UserRole))
		}
	}

	//采购商查看批发信息
	// 2. 综合排序、最新发布：都按需求单提交时间由新到旧排列（本期不做区分）
	// 3. 低价优先：按货品最低价格升序排列（一个货品可以有多个sku、多个价格，这里按其sku中的最低价格排序）
	if user.UserRole == model.UserRoleBuyer {
		req.ProductVariantType = enum.ProductWholesale
		switch req.ListType {
		case enum.GoodsListTypeNormal:
			return s.normalGetGoodsList(ctx, req)
		case enum.GoodsListTypePrice:
			return s.getGoodsListByPrice(ctx, req)
		default:
			return nil, xerr.WithCode(xerr.InvalidParams, fmt.Errorf("listType:%d is not supported userRole:%d ", req.ListType, user.UserRole))
		}

	}
	return nil, xerr.WithCode(xerr.InvalidParams, fmt.Errorf("listType:%d is not supported userRole:%d", req.ListType, user.UserRole))
}

func (s *GoodsService) normalGetGoodsList(ctx *gin.Context, req types.GoodsListReq) (*types.GoodsListResp, xerr.XErr) {
	goods, total, err := s.goods.GetGoodsList(ctx, req)
	if err != nil {
		return nil, xerr.WithCode(xerr.ErrorDatabase, err)
	}

	result := types.GoodsListResp{PageResult: types.PageResult{PageNum: req.PageNum, PageSize: req.PageSize, TotalNum: total}}
	for i, v := range goods {
		productVariants, rr := s.goods.GetProductVariantListByGoodsID(ctx, v.ID, req.ProductVariantType)
		if rr != nil {
			return nil, xerr.WithCode(xerr.ErrorDatabase, rr)
		}
		goods[i].GoodsFrontImage = goods[i].GetGoodsFrontImage()
		goods[i].WholesalePriceMax, goods[i].WholesalePriceMin, goods[i].RetailPriceMax, goods[i].RetailPriceMin, goods[i].WholeSaleUnit, goods[i].RetailUnit = s.findPriceBounds(productVariants)
	}
	result.GoodsList = goods
	return &result, nil
}

func (s *GoodsService) findPriceBounds(productVariants []*model.ProductVariant) (float64, float64, float64, float64, string, string) {
	var wholesalePriceMax, wholesalePriceMin, retailPriceMax, retailPriceMin float64
	var wholesaleUnit, retailUnit string
	initW := false
	initR := false
	for k := range productVariants {
		if productVariants[k].Type == enum.ProductWholesale {
			wholesaleUnit = productVariants[k].Unit
			if !initW || productVariants[k].Price > wholesalePriceMax {
				wholesalePriceMax = productVariants[k].Price
			}
			if !initW || productVariants[k].Price < wholesalePriceMin {
				wholesalePriceMin = productVariants[k].Price
			}
			initW = true
		} else if productVariants[k].Type == enum.ProductRetail {
			retailUnit = productVariants[k].Unit
			if !initR || productVariants[k].Price > retailPriceMax {
				retailPriceMax = productVariants[k].Price
			}
			if !initR || productVariants[k].Price < retailPriceMin {
				retailPriceMin = productVariants[k].Price
			}
			initR = true
		}
	}
	return wholesalePriceMax, wholesalePriceMin, retailPriceMax, retailPriceMin, wholesaleUnit, retailUnit
}

func (s *GoodsService) getGoodsListByPrice(ctx *gin.Context, req types.GoodsListReq) (*types.GoodsListResp, xerr.XErr) {
	minPriceResult, total, err := s.goods.GetMinPriceList(ctx, req)
	if err != nil {
		return nil, xerr.WithCode(xerr.ErrorDatabase, err)
	}
	goods := s.ConvertMinPriceResultToGoodsList(minPriceResult)
	result := types.GoodsListResp{PageResult: types.PageResult{PageNum: req.PageNum, PageSize: req.PageSize, TotalNum: total}}
	for i, v := range goods {
		productVariants, rr := s.goods.GetProductVariantListByGoodsID(ctx, v.ID, req.ProductVariantType)
		if rr != nil {
			return nil, xerr.WithCode(xerr.ErrorDatabase, rr)
		}
		goods[i].GoodsFrontImage = goods[i].GetGoodsFrontImage()
		goods[i].WholesalePriceMax, goods[i].WholesalePriceMin, goods[i].RetailPriceMax, goods[i].RetailPriceMin, goods[i].WholeSaleUnit, goods[i].RetailUnit = s.findPriceBounds(productVariants)
	}
	result.GoodsList = goods
	return &result, nil
}

func (s *GoodsService) ConvertMinPriceResultToGoodsList(minPriceResults []*types.MinPriceResult) []*types.GoodsList {
	goodsList := make([]*types.GoodsList, len(minPriceResults))
	for i, result := range minPriceResults {
		goodsList[i] = &types.GoodsList{
			ID:              result.GoodsID,
			Name:            result.Name,
			GoodsFrontImage: result.GoodsFrontImage,
			Images:          result.Images,
		}
	}
	return goodsList
}

func (s *GoodsService) DeleteMyGoods(ctx *gin.Context, req types.GoodsReq) xerr.XErr {
	//获取用户ID
	//userID := common.GetUserID(ctx)
	////获取用户角色
	//userRole := common.GetUserRole(ctx)
	userID := "w2wwww"
	user, err := s.userDao.GetByUserID(ctx, userID)
	if err != nil {
		return xerr.WithCode(xerr.ErrorDatabase, err)
	}
	userRole := user.UserRole
	if userRole != model.UserRoleSupplier {
		return xerr.WithCode(xerr.ErrorAuthInsufficientAuthority, errors.New("用户权限不是供应商"))
	}
	//检查用户权限
	goods, err := s.goods.GetGoodsByGoodsID(ctx, req.GoodsID)
	if err != nil {
		return xerr.WithCode(xerr.ErrorDatabase, err)
	}
	if goods == nil {
		return xerr.WithCode(xerr.ErrorNotExistRecord, errors.New("商品不存在"))
	}
	if goods.UserID != userID {
		return xerr.WithCode(xerr.ErrorAuthInsufficientAuthority, errors.New("用户没有权限删除该商品"))
	}

	//删除商品
	err = s.goods.UpdateGoodsByID(ctx, req.GoodsID, &model.Goods{Deleted: 1})
	if err != nil {
		return xerr.WithCode(xerr.ErrorDatabase, err)
	}
	return nil

}

func (s *GoodsService) ModifyMyGoodsStatus(c *gin.Context, req types.GoodsReq) xerr.XErr {
	//获取用户ID
	//userID := common.GetUserID(c)
	////获取用户角色
	//userRole := common.GetUserRole(c)
	userID := "w2wwww"
	user, err := s.userDao.GetByUserID(c, userID)
	if err != nil {
		return xerr.WithCode(xerr.ErrorDatabase, err)
	}
	userRole := user.UserRole
	if userRole != model.UserRoleSupplier {
		return xerr.WithCode(xerr.ErrorAuthInsufficientAuthority, errors.New("用户权限不是供应商"))
	}
	if req.GoodsStatus != enum.GoodsStatusOnSale && req.GoodsStatus != enum.GoodsStatusOffSale {
		return xerr.WithCode(xerr.InvalidParams, errors.New("商品状态不合法"))
	}
	//检查用户权限
	goods, err := s.goods.GetGoodsByGoodsID(c, req.GoodsID)
	if err != nil {
		return xerr.WithCode(xerr.ErrorDatabase, err)
	}
	if goods == nil {
		return xerr.WithCode(xerr.ErrorNotExistRecord, errors.New("商品不存在"))
	}
	if goods.UserID != userID {
		return xerr.WithCode(xerr.ErrorAuthInsufficientAuthority, errors.New("用户没有权限删除该商品"))
	}
	if goods.Status == req.GoodsStatus {
		return xerr.WithCode(xerr.InvalidParams, errors.New("商品状态不正确"))
	}
	//修改商品状态
	err = s.goods.UpdateGoodsByID(c, req.GoodsID, &model.Goods{Status: req.GoodsStatus})
	if err != nil {
		return xerr.WithCode(xerr.ErrorDatabase, err)
	}
	return nil
}
