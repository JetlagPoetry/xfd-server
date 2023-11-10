package service

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
	"sort"
	"strings"
	"time"
	"xfd-backend/database/db"
	"xfd-backend/database/db/dao"
	"xfd-backend/database/db/enum"
	"xfd-backend/database/db/model"
	"xfd-backend/database/repo"
	"xfd-backend/pkg/common"
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
	//获取用户ID
	userID := common.GetUserID(ctx)
	//获取用户角色
	user, err := s.userDao.GetByUserID(ctx, userID)
	if err != nil {
		return xerr.WithCode(xerr.ErrorDatabase, err)
	}
	userRole := user.UserRole
	if user.UserRole != model.UserRoleSupplier {
		return xerr.WithCode(xerr.ErrorAuthInsufficientAuthority, fmt.Errorf("user %suserRole:%d is not supplier", userID, userRole))
	}
	if len(req.GoodsDetail.Images) == 0 {
		return xerr.WithCode(xerr.InvalidParams, fmt.Errorf("goods images is empty"))
	}

	spuCode := fmt.Sprintf("SP%s%s", time.Now().Format("20060102150405"), utils.GenSixDigitCode())
	var images, descImages, wholesaleLogistics, goodsFrontImage string
	if len(req.GoodsDetail.Images) != 0 {
		imagesBytes, _ := json.Marshal(req.GoodsDetail.Images)
		images = string(imagesBytes)
		goodsFrontImage = req.GoodsDetail.Images[0]
	}
	if len(req.GoodsDetail.DescImages) == 0 {
		descImages = images
	} else {
		descImagesBytes, _ := json.Marshal(req.GoodsDetail.DescImages)
		descImages = string(descImagesBytes)
	}
	if len(req.GoodsDetail.WholesaleLogistics) != 0 {
		wholesaleLogisticsBytes, _ := json.Marshal(req.GoodsDetail.WholesaleLogistics)
		wholesaleLogistics = string(wholesaleLogisticsBytes)
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
		GoodsFrontImage:    goodsFrontImage,
		RetailStatus:       enum.GoodsRetailNormal,
		WholesaleLogistics: wholesaleLogistics,
		WholesaleShipping:  req.GoodsDetail.WholesaleShipping,
		WholesaleAreaCodeA: req.GoodsDetail.WholesaleAreaCodeA,
		WholesaleAreaCodeB: req.GoodsDetail.WholesaleAreaCodeB,
		WholesaleAreaCodeC: req.GoodsDetail.WholesaleAreaCodeC,
		RetailShippingTime: req.GoodsDetail.RetailShippingTime,
		RetailShippingFee:  &req.GoodsDetail.RetailShippingFee,
	}
	//全国包邮
	if req.GoodsDetail.RetailShippingFee == 0 {
		one := 1
		modelGoods.ShipFree = &one
	}
	totalStock := 0
	for _, v := range req.RetailProducts {
		if v.Status == enum.ProductVariantEnabled {
			one := 1
			modelGoods.IsRetail = &one
			totalStock += v.Stock
		}
	}
	if totalStock == 0 {
		modelGoods.RetailStatus = enum.GoodsRetailSoldOut
	}
	transactionHandler := repo.NewTransactionHandler(db.Get())
	if xErr := transactionHandler.WithTransaction(ctx, func(ctx context.Context) xerr.XErr {
		if _err := s.CreateGoodsWithTransaction(ctx, req, modelGoods); _err != nil {
			return _err
		}
		return nil
	}); xErr != nil {
		return xErr
	}
	return nil
}

func (s *GoodsService) CreateGoodsWithTransaction(ctx context.Context, req types.GoodsAddReq, modelGoods *model.Goods) xerr.XErr {
	//创建商品信息
	goodsID, err := s.goods.CreateGoods(ctx, modelGoods)
	if err != nil {
		return xerr.WithCode(xerr.ErrorDatabase, err)
	}

	xrr := s.processProducts(ctx, req.WholesaleProducts, goodsID, enum.ProductWholesale)
	if xrr != nil {
		return xrr
	}

	xrr = s.processProducts(ctx, req.RetailProducts, goodsID, enum.ProductRetail)
	if xrr != nil {
		return xrr
	}
	return nil
}

func (s *GoodsService) createSpecificationsAndValues(ctx context.Context, products []*types.AddProduct, goodsID int32, productType enum.ProductVariantType) (map[string]int32, map[string]int32, xerr.XErr) {
	valueKeyMap := make(map[string]string)
	keySet := make(map[string]bool)
	valueSet := make(map[string]bool)
	keyIDMap := make(map[string]int32)
	valueIDMap := make(map[string]int32)
	productAttrSet := make(map[string]bool)
	for _, product := range products {
		if len(product.ProductAttr) != 0 {
			var attrs []string
			for _, attr := range product.ProductAttr {
				keySet[attr.Key] = true
				valueSet[attr.Value] = true
				attrs = append(attrs, attr.Key+":"+attr.Value)
				valueKeyMap[attr.Value] = attr.Key
			}
			sort.Strings(attrs) // 对属性进行排序以确保属性的顺序不影响唯一性的判断
			attrString := strings.Join(attrs, ",")
			// 检查这个字符串是否已经存在
			if _, ok := productAttrSet[attrString]; ok {
				return nil, nil, xerr.WithCode(xerr.InvalidParams, fmt.Errorf("duplicate product attributes:%s", attrString))
			}
			productAttrSet[attrString] = true
		}
	}
	for key := range keySet {
		specification := &model.Specification{
			Name:    key,
			GoodsID: goodsID,
			Type:    productType,
		}
		specificationID, err := s.goods.CreateSpecification(ctx, specification)
		if err != nil {
			return nil, nil, xerr.WithCode(xerr.ErrorDatabase, err)
		}
		keyIDMap[key] = specificationID
	}
	for value := range valueSet {
		specificationValue := &model.SpecificationValue{
			Value:           value,
			SpecificationID: keyIDMap[valueKeyMap[value]],
			GoodsID:         goodsID,
			Type:            productType,
		}
		specValueID, err := s.goods.CreateSpecificationValue(ctx, specificationValue)
		if err != nil {
			return nil, nil, xerr.WithCode(xerr.ErrorDatabase, err)
		}
		valueIDMap[value] = specValueID
	}

	return keyIDMap, valueIDMap, nil
}

func (s *GoodsService) processProducts(ctx context.Context, products []*types.AddProduct, goodsID int32, productType enum.ProductVariantType) xerr.XErr {
	keyIDMap, valueIDMap, xrr := s.createSpecificationsAndValues(ctx, products, goodsID, productType)
	if xrr != nil {
		return xrr
	}
	for i := range products {
		for j, attr := range products[i].ProductAttr {
			products[i].ProductAttr[j].KeyID = keyIDMap[attr.Key]
			products[i].ProductAttr[j].ValueID = valueIDMap[attr.Value]
		}
		skuCode := fmt.Sprintf("SK%d%d%s", time.Now().Unix(), goodsID, utils.GenerateRandomNumber(8))
		productAttrBytes, _ := json.Marshal(products[i].ProductAttr)
		productVariant := &model.ProductVariant{
			SKUCode:          skuCode,
			GoodsID:          goodsID,
			Unit:             products[i].Unit,
			Price:            products[i].Price,
			Stock:            &products[i].Stock,
			MinOrderQuantity: products[i].MinOrderQuantity,
			Type:             productType,
			Status:           products[i].Status,
			ProductAttr:      string(productAttrBytes),
		}
		_, err := s.goods.CreateProductVariant(ctx, productVariant)
		if err != nil {
			return xerr.WithCode(xerr.ErrorDatabase, err)
		}
	}
	return nil
}

func (s *GoodsService) GetGoodsList(ctx *gin.Context, req types.GoodsListReq) (*types.GoodsListResp, xerr.XErr) {
	if req.CheckParams() != nil {
		return nil, xerr.WithCode(xerr.InvalidParams, req.CheckParams())
	}
	userID := common.GetUserID(ctx)
	user, err := s.userDao.GetByUserID(ctx, userID)
	if err != nil {
		return nil, xerr.WithCode(xerr.ErrorDatabase, err)
	}
	uerRole := user.UserRole
	//供货商,查看供货信息(批发)
	if uerRole == model.UserRoleSupplier {
		req.ProductVariantType = enum.ProductWholesale
		req.UserID = userID
		switch req.ListType {
		case enum.GoodsListTypeNormal:
			return s.normalGetGoodsList(ctx, req)
		case enum.GoodsListTypeCategory:
			return s.normalGetGoodsList(ctx, req)
		case enum.GoodsListTypeQuery:
			return s.normalGetGoodsList(ctx, req)
		default:
			return nil, xerr.WithCode(xerr.InvalidParams, fmt.Errorf("listType:%d is not supported userRole:%d ", req.ListType, user.UserRole))
		}
	}
	if uerRole == model.UserRoleCustomer {
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
		case enum.GoodsListTypeQuery:
			return s.normalGetGoodsList(ctx, req)
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
		case enum.GoodsListTypeQuery:
			return s.normalGetGoodsList(ctx, req)
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
		productVariants, rr := s.goods.GetProductVariantListByGoodsID(ctx, v.ID, req.ProductVariantType, enum.ProductVariantEnabled)
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
		productVariants, rr := s.goods.GetProductVariantListByGoodsID(ctx, v.ID, req.ProductVariantType, enum.ProductVariantEnabled)
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
	userID := common.GetUserID(ctx)
	user, err := s.userDao.GetByUserID(ctx, userID)
	if err != nil {
		return xerr.WithCode(xerr.ErrorDatabase, err)
	}
	userRole := user.UserRole
	if userRole != model.UserRoleSupplier {
		return xerr.WithCode(xerr.ErrorAuthInsufficientAuthority, fmt.Errorf("user %suserRole:%d is not supplier", userID, userRole))
	}

	_, xrr := s.checkGoodsValid(ctx, req.GoodsID, userID)
	if xrr != nil {
		return xrr
	}

	//删除商品
	err = s.goods.DeleteGoodsByID(ctx, req.GoodsID)
	if err != nil {
		return xerr.WithCode(xerr.ErrorDatabase, err)
	}
	err = s.goods.DeleteProductVariantByGoodsID(ctx, req.GoodsID)
	if err != nil {
		return xerr.WithCode(xerr.ErrorDatabase, err)
	}
	err = s.goods.DeleteSpecificationByGoodsID(ctx, req.GoodsID)
	if err != nil {
		return xerr.WithCode(xerr.ErrorDatabase, err)
	}
	err = s.goods.DeleteSpecificationValueByGoodsID(ctx, req.GoodsID)
	if err != nil {
		return xerr.WithCode(xerr.ErrorDatabase, err)
	}
	return nil

}

func (s *GoodsService) ModifyMyGoodsStatus(c *gin.Context, req types.GoodsReq) xerr.XErr {
	//获取用户ID
	userID := common.GetUserID(c)
	//获取用户角色
	user, err := s.userDao.GetByUserID(c, req.UserID)
	if err != nil {
		return xerr.WithCode(xerr.ErrorDatabase, err)
	}
	userRole := user.UserRole
	if userRole != model.UserRoleSupplier {
		return xerr.WithCode(xerr.ErrorAuthInsufficientAuthority, fmt.Errorf("user %suserRole:%d is not supplier", userID, userRole))
	}
	if req.GoodsStatus != enum.GoodsStatusOnSale && req.GoodsStatus != enum.GoodsStatusOffSale {
		return xerr.WithCode(xerr.InvalidParams, fmt.Errorf("goodsStatus:%d is not supported", req.GoodsStatus))
	}
	//检查用户权限
	_, xrr := s.checkGoodsValid(c, req.GoodsID, userID)
	if xrr != nil {
		return xrr
	}
	//修改商品状态
	rowsAffected, err := s.goods.UpdateGoodsByID(c, req.GoodsID, &model.Goods{Status: req.GoodsStatus})
	if err != nil {
		return xerr.WithCode(xerr.ErrorDatabase, err)
	}
	if rowsAffected == 0 {
		return xerr.WithCode(xerr.ErrorDatabase, fmt.Errorf("update goods status failed,goodsID:%d", req.GoodsID))
	}
	return nil
}

func (s *GoodsService) GetMyGoodsList(c *gin.Context, req types.MyGoodsListReq) (*types.GoodsListResp, xerr.XErr) {
	userID := common.GetUserID(c)
	//获取用户角色
	user, err := s.userDao.GetByUserID(c, req.UserID)
	if err != nil {
		return nil, xerr.WithCode(xerr.ErrorDatabase, err)
	}
	userRole := user.UserRole

	if userRole != model.UserRoleSupplier {
		return nil, xerr.WithCode(xerr.ErrorAuthInsufficientAuthority, fmt.Errorf("user %suserRole:%d is not supplier", userID, userRole))
	}
	req.UserID = userID
	goods, total, err := s.goods.GetMyGoodsList(c, req)
	if err != nil {
		return nil, xerr.WithCode(xerr.ErrorDatabase, err)
	}
	result := types.GoodsListResp{PageResult: types.PageResult{PageNum: req.PageNum, PageSize: req.PageSize, TotalNum: total}}
	for i, v := range goods {
		productVariants, rr := s.goods.GetProductVariantListByGoodsID(c, v.ID, 0, 0)
		if rr != nil {
			return nil, xerr.WithCode(xerr.ErrorDatabase, rr)
		}
		goods[i].GoodsFrontImage = goods[i].GetGoodsFrontImage()
		goods[i].WholesalePriceMax, goods[i].WholesalePriceMin, goods[i].RetailPriceMax, goods[i].RetailPriceMin, goods[i].WholeSaleUnit, goods[i].RetailUnit = s.findPriceBounds(productVariants)
	}
	result.GoodsList = goods
	return &result, nil
}

func (s *GoodsService) GetGoodsDetail(c *gin.Context, req types.GoodsReq) (*types.GoodsDetailResp, xerr.XErr) {
	goods, err := s.goods.GetGoodsByGoodsID(c, req.GoodsID)
	if err != nil {
		return nil, xerr.WithCode(xerr.ErrorDatabase, err)
	}
	if goods == nil {
		return nil, xerr.WithCode(xerr.ErrorNotExistRecord, errors.New("商品不存在"))
	}
	goodsDetail := &types.GoodsDetailResp{
		ID:              goods.ID,
		Name:            goods.Name,
		GoodsFrontImage: goods.GoodsFrontImage,
		Images:          goods.GetImagesList(),
		Description:     goods.Description,
		DescImages:      goods.GetDescImagesList(),
	}

	req.UserID = common.GetUserID(c)
	user, err := s.userDao.GetByUserID(c, req.UserID)
	if err != nil {
		return nil, xerr.WithCode(xerr.ErrorDatabase, err)
	}
	req.UserRole = user.UserRole
	if req.ReqType == 1 {
		return s.handleSupplierGoodsDetailRequest(c, req, goods, goodsDetail)
	} else {
		return s.handleGoodsDetailRequestByRole(c, req, goods, goodsDetail)
	}
}
func (s *GoodsService) handleSupplierGoodsDetailRequest(c *gin.Context, req types.GoodsReq, goods *model.Goods, goodsDetail *types.GoodsDetailResp) (*types.GoodsDetailResp, xerr.XErr) {
	if req.UserRole != model.UserRoleSupplier {
		return nil, xerr.WithCode(xerr.ErrorAuthInsufficientAuthority, fmt.Errorf("user %suserRole:%d is not supplier", req.UserID, req.UserRole))
	}
	if err := s.fillWholesaleDetails(c, req, goods, goodsDetail, 0); err != nil {
		return nil, err
	}

	if *goods.IsRetail == 1 {
		if err := s.fillRetailDetails(c, req, goods, goodsDetail, 0); err != nil {
			return nil, err
		}
	}
	return goodsDetail, nil
}

func (s *GoodsService) fillWholesaleDetails(c *gin.Context, req types.GoodsReq, goods *model.Goods, goodsDetail *types.GoodsDetailResp, status enum.ProductVariantStatus) xerr.XErr {
	goodsDetail.WholesaleShipping = goods.WholesaleShipping
	goodsDetail.WholesaleLogistics = goods.GetWholesaleLogistics()
	specInfo, err := s.getSpecificationByType(c, req, enum.ProductWholesale)
	if err != nil {
		return err
	}
	goodsDetail.SpecInfo = append(goodsDetail.SpecInfo, specInfo...)
	productsInfo, err := s.getProductsInfoByTypeAndStatus(c, req, enum.ProductWholesale, status)
	if err != nil {
		return err
	}
	goodsDetail.WholesaleProducts = productsInfo

	return nil
}

func (s *GoodsService) fillRetailDetails(c *gin.Context, req types.GoodsReq, goods *model.Goods, goodsDetail *types.GoodsDetailResp, status enum.ProductVariantStatus) xerr.XErr {
	goodsDetail.RetailShippingTime = goods.RetailShippingTime
	goodsDetail.RetailShippingFee = *goods.RetailShippingFee
	specInfo, err := s.getSpecificationByType(c, req, enum.ProductRetail)
	if err != nil {
		return err
	}
	goodsDetail.SpecInfo = append(goodsDetail.SpecInfo, specInfo...)
	productsInfo, err := s.getProductsInfoByTypeAndStatus(c, req, enum.ProductRetail, status)
	if err != nil {
		return err
	}
	goodsDetail.RetailProducts = productsInfo

	return nil
}

func (s *GoodsService) handleGoodsDetailRequestByRole(c *gin.Context, req types.GoodsReq, goods *model.Goods, goodsDetail *types.GoodsDetailResp) (*types.GoodsDetailResp, xerr.XErr) {
	if req.UserRole == model.UserRoleCustomer {
		if *goods.IsRetail == 0 || goods.RetailStatus == enum.GoodsRetailSoldOut {
			return nil, xerr.WithCode(xerr.ErrorNotExistRecord, fmt.Errorf("goods:%d is not retail", goods.ID))
		}

		if err := s.fillRetailDetails(c, req, goods, goodsDetail, enum.ProductVariantEnabled); err != nil {
			return nil, err
		}
	}
	if req.UserRole == model.UserRoleSupplier || req.UserRole == model.UserRoleBuyer {
		if err := s.fillWholesaleDetails(c, req, goods, goodsDetail, enum.ProductVariantEnabled); err != nil {
			return nil, err
		}
	}
	return goodsDetail, nil
}

// getSpecificationByType
func (s *GoodsService) getSpecificationByType(c *gin.Context, req types.GoodsReq, productVariantType enum.ProductVariantType) ([]*types.SpecInfo, xerr.XErr) {
	specifications, err := s.goods.GetSpecificationByGoodsID(c, req.GoodsID, productVariantType)
	if err != nil {
		return nil, xerr.WithCode(xerr.ErrorDatabase, err)
	}
	result := make([]*types.SpecInfo, 0)
	for _, v := range specifications {
		specInfo := &types.SpecInfo{
			SpecID:   v.ID,
			SpecName: v.Name,
			Type:     v.Type,
		}
		specificationValues, rr := s.goods.GetSpecificationValueBySpecID(c, v.ID)
		if rr != nil {
			return nil, xerr.WithCode(xerr.ErrorDatabase, rr)
		}
		for _, vv := range specificationValues {
			specInfo.SpecValue = append(specInfo.SpecValue, &types.SpecValue{
				ID:    vv.ID,
				Value: vv.Value,
			})
		}
		result = append(result, specInfo)
	}
	return result, nil
}

func (s *GoodsService) getProductsInfoByTypeAndStatus(c *gin.Context, req types.GoodsReq, productVariantType enum.ProductVariantType, status enum.ProductVariantStatus) ([]*types.ProductVariantInfo, xerr.XErr) {
	products := make([]*types.ProductVariantInfo, 0)
	productVariants, rr1 := s.goods.GetProductVariantListByGoodsID(c, req.GoodsID, productVariantType, status)
	if rr1 != nil {
		return nil, xerr.WithCode(xerr.ErrorDatabase, rr1)
	}
	for _, vv := range productVariants {
		productVariantInfo := &types.ProductVariantInfo{
			ID:               vv.ID,
			Unit:             vv.Unit,
			Price:            vv.Price,
			MinOrderQuantity: vv.MinOrderQuantity,
			Stock:            *vv.Stock,
			SKUCode:          vv.SKUCode,
			ProductAttr:      vv.GetProductAttr(),
			Status:           vv.Status,
		}
		products = append(products, productVariantInfo)
	}
	return products, nil
}

func (s *GoodsService) ModifyMyGoods(c *gin.Context, req types.GoodsModifyReq) xerr.XErr {
	//1.检查基本参数
	if req.CheckParams() != nil {
		return xerr.WithCode(xerr.InvalidParams, req.CheckParams())
	}
	userID := common.GetUserID(c)
	user, err := s.userDao.GetByUserID(c, userID)
	if err != nil {
		return xerr.WithCode(xerr.ErrorDatabase, err)
	}
	if user == nil {
		return xerr.WithCode(xerr.ErrorNotExistUser, fmt.Errorf("user:%s is not exist", userID))
	}
	userRole := user.UserRole
	if userRole != model.UserRoleSupplier {
		return xerr.WithCode(xerr.ErrorAuthInsufficientAuthority, fmt.Errorf("user:%s is not supplier", userID))
	}
	goods, xrr := s.checkGoodsValid(c, req.Id, userID)
	if xrr != nil {
		return xrr
	}
	var wholesaleLogistics, descImages, images, goodsFrontImage string
	if len(req.WholesaleLogistics) != 0 {
		wholesaleLogisticsBytes, _ := json.Marshal(req.WholesaleLogistics)
		wholesaleLogistics = string(wholesaleLogisticsBytes)
	}
	if len(req.DescImages) != 0 {
		descImagesBytes, _ := json.Marshal(req.DescImages)
		descImages = string(descImagesBytes)
	}
	if len(req.Images) != 0 {
		imagesBytes, _ := json.Marshal(req.Images)
		images = string(imagesBytes)
		goodsFrontImage = req.Images[0]
	}
	updateValue := &model.Goods{
		CategoryAID:        req.CategoryAID,
		CategoryBID:        req.CategoryBID,
		CategoryCID:        req.CategoryCID,
		CategoryName:       req.CategoryName,
		Name:               req.Name,
		Description:        req.Description,
		Images:             images,
		DescImages:         descImages,
		GoodsFrontImage:    goodsFrontImage,
		WholesaleLogistics: wholesaleLogistics,
		WholesaleShipping:  req.WholesaleShipping,
		WholesaleAreaCodeA: req.WholesaleAreaCodeA,
		WholesaleAreaCodeB: req.WholesaleAreaCodeB,
		WholesaleAreaCodeC: req.WholesaleAreaCodeC,
		RetailShippingTime: req.RetailShippingTime,
		RetailShippingFee:  req.RetailShippingFee,
	}
	if *goods.ShipFree == 1 && req.RetailShippingFee != nil && *req.RetailShippingFee != 0 {
		zero := 0
		updateValue.ShipFree = &zero
	}
	if *goods.ShipFree == 0 && req.RetailShippingFee != nil && *req.RetailShippingFee == 0 {
		one := 1
		updateValue.ShipFree = &one
	}
	transactionHandler := repo.NewTransactionHandler(db.Get())
	if xErr := transactionHandler.WithTransaction(c, func(ctx context.Context) xerr.XErr {
		if _err := s.ModifyGoodsWithTransaction(ctx, req, updateValue); _err != nil {
			return _err
		}
		return nil
	}); xErr != nil {
		return xErr
	}
	return nil
}

func (s *GoodsService) checkGoodsValid(c context.Context, goodsID int32, userID string) (*model.Goods, xerr.XErr) {
	goods, err := s.goods.GetGoodsByGoodsID(c, goodsID)
	if err != nil {
		return nil, xerr.WithCode(xerr.ErrorDatabase, err)
	}
	if goods == nil {
		return nil, xerr.WithCode(xerr.ErrorNotExistRecord, fmt.Errorf("goods:%d is not exist", goodsID))
	}
	if userID == "" {
		return goods, nil
	} else {
		if goods.UserID != userID {
			return nil, xerr.WithCode(xerr.ErrorAuthInsufficientAuthority, fmt.Errorf("user:%s is not owner of goods:%d", userID, goodsID))
		}
		return goods, nil
	}
}

func (s *GoodsService) ModifyGoodsWithTransaction(c context.Context, req types.GoodsModifyReq, updateValue *model.Goods) xerr.XErr {
	_, _, xrr := s.processModifyProducts(c, req.WholesaleProducts, req.Id, enum.ProductWholesale)
	if xrr != nil {
		return xrr
	}
	one := 1
	totalStock, hasRetail, xrr := s.processModifyProducts(c, req.RetailProducts, req.Id, enum.ProductRetail)
	if xrr != nil {
		return xrr
	}
	if hasRetail {
		updateValue.IsRetail = &one
	}
	if totalStock > 0 {
		updateValue.RetailStatus = enum.GoodsRetailNormal
	}
	if hasRetail && totalStock > 0 {
		rowsAffected, err := s.goods.UpdateGoodsByID(c, req.Id, updateValue)
		if err != nil {
			return xerr.WithCode(xerr.ErrorDatabase, err)
		}
		if rowsAffected == 0 {
			return xerr.WithCode(xerr.ErrorDatabase, fmt.Errorf("update goods status failed,goodsID:%d", req.Id))
		}
	} else {
		stock := 0
		zero := 0
		time.Sleep(50 * time.Millisecond)
		retailProductVariants, rr := s.goods.GetProductVariantListByGoodsID(c, req.Id, enum.ProductRetail, enum.ProductVariantEnabled)
		if rr != nil {
			return xerr.WithCode(xerr.ErrorDatabase, rr)
		}
		if len(retailProductVariants) != 0 {
			updateValue.IsRetail = &one
			for _, v := range retailProductVariants {
				stock += *v.Stock
			}
			if stock > 0 {
				updateValue.RetailStatus = enum.GoodsRetailNormal
			}
			if stock <= 0 {
				updateValue.RetailStatus = enum.GoodsRetailSoldOut
			}
		} else {
			updateValue.IsRetail = &zero
		}
	}

	return nil
}

func (s *GoodsService) processModifyProducts(ctx context.Context, products []*types.ModifyProduct, goodsID int32, productType enum.ProductVariantType) (int, bool, xerr.XErr) {
	if len(products) == 0 {
		return 0, false, nil
	}
	hasValue, hasValid := s.validateProductAttr(products)
	if !hasValid {
		return 0, false, xerr.WithCode(xerr.InvalidParams, fmt.Errorf("productAttr is invalid,params:%+v,must all hasValue or all hasEmpty", products))
	}
	keyIDMap, valueIDMap, xrr := s.updateOrCreateSpecificationsAndValues(ctx, products, goodsID, productType)
	if xrr != nil {
		return 0, false, xrr
	}
	totalStock := 0
	var hasRetail bool
	for i := range products {
		if productType == enum.ProductRetail && products[i].Status == enum.ProductVariantEnabled {
			totalStock += *products[i].Stock
			hasRetail = true
		}
		productVariant := &model.ProductVariant{
			Unit:             products[i].Unit,
			Price:            products[i].Price,
			MinOrderQuantity: products[i].MinOrderQuantity,
			Status:           products[i].Status,
			Stock:            products[i].Stock,
		}
		if products[i].CheckParams(productType) != nil {
			return 0, false, xerr.WithCode(xerr.InvalidParams, products[i].CheckParams(productType))
		}
		if hasValue {
			for j, attr := range products[i].ProductAttr {
				products[i].ProductAttr[j].KeyID = keyIDMap[attr.Key]
				products[i].ProductAttr[j].ValueID = valueIDMap[attr.Value]
			}
			productAttrBytes, _ := json.Marshal(products[i].ProductAttr)
			productVariant.ProductAttr = string(productAttrBytes)
			if products[i].ID == 0 {
				skuCode := fmt.Sprintf("SK%d%d%s", time.Now().Unix(), goodsID, utils.GenerateRandomNumber(8))
				productVariant.SKUCode = skuCode
				productVariant.GoodsID = goodsID
				productVariant.Type = productType
				//创建新的productVariant
				_, err := s.goods.CreateProductVariant(ctx, productVariant)
				if err != nil {
					return 0, false, xerr.WithCode(xerr.ErrorDatabase, err)
				}
			} else {
				//更新productVariant
				rowsAffected, err := s.goods.UpdateProductVariantByID(ctx, products[i].ID, productVariant)
				if err != nil {
					return 0, false, xerr.WithCode(xerr.ErrorDatabase, err)
				}
				if rowsAffected == 0 {
					return 0, false, xerr.WithCode(xerr.ErrorDatabase, fmt.Errorf("update productVariant failed,productVariantID:%d,wait updateValue %v", products[i].ID, productVariant))
				}
			}
		} else {
			if products[i].ID != 0 {
				//更新productVariant
				rowsAffected, err := s.goods.UpdateProductVariantByID(ctx, products[i].ID, productVariant)
				if err != nil {
					return 0, false, xerr.WithCode(xerr.ErrorDatabase, err)
				}
				if rowsAffected == 0 {
					return 0, false, xerr.WithCode(xerr.ErrorDatabase, fmt.Errorf("update productVariant failed,productVariantID:%d,wait updateValue %v", products[i].ID, productVariant))
				}
			}
			return 0, false, xerr.WithCode(xerr.InvalidParams, fmt.Errorf("productAttr is invalid,params:%+v,id is empty,productAttr must all hasValue ", products))
		}
	}
	return totalStock, hasRetail, nil
}

func (s *GoodsService) updateOrCreateSpecificationsAndValues(ctx context.Context, products []*types.ModifyProduct, goodsID int32, productType enum.ProductVariantType) (map[string]int32, map[string]int32, xerr.XErr) {
	valueKeyMap, keySet, valueSet, keyIDMap, valueIDMap, productAttrSet, xrr := s.processProductAttrs(products)
	if xrr != nil {
		return nil, nil, xrr
	}
	fmt.Println(fmt.Sprintf("valueKeyMap:%+v,keySet:%+v,valueSet:%+v,keyIDMap:%+v,valueIDMap:%+v,productAttrSet:%+v", valueKeyMap, keySet, valueSet, keyIDMap, valueIDMap, productAttrSet))
	specifications, err := s.goods.GetSpecificationByGoodsID(ctx, goodsID, productType)
	if err != nil {
		return nil, nil, xerr.WithCode(xerr.ErrorDatabase, err)
	}
	specificationValues, err := s.goods.GetSpecificationValueByGoodID(ctx, goodsID, productType)
	if err != nil {
		return nil, nil, xerr.WithCode(xerr.ErrorDatabase, err)
	}

	oldSpecIDsMap := make(map[int32]*model.Specification)
	for _, specification := range specifications {
		oldSpecIDsMap[specification.ID] = specification
	}
	oldValueIDsMap := make(map[int32]*model.SpecificationValue)
	for _, specificationValue := range specificationValues {
		oldValueIDsMap[specificationValue.ID] = specificationValue
	}
	keyIDMap, xrr = s.handlerSpecification(ctx, oldSpecIDsMap, keySet, keyIDMap, goodsID, productType)
	if xrr != nil {
		return nil, nil, xrr
	}
	valueIDMap, xrr = s.handlerSpecificationValue(ctx, oldValueIDsMap, valueSet, valueIDMap, keyIDMap, valueKeyMap, goodsID, productType)
	if xrr != nil {
		return nil, nil, xrr
	}

	return keyIDMap, valueIDMap, nil
}

func (s *GoodsService) validateProductAttr(products []*types.ModifyProduct) (bool, bool) {
	var hasValue bool
	var hasEmpty bool
	hasValid := true

	for _, product := range products {
		if len(product.ProductAttr) > 0 {
			hasValue = true
		} else {
			hasEmpty = true
		}

		if hasValue && hasEmpty {
			hasValid = false
			break
		}
	}

	return hasValue, hasValid
}

func (s *GoodsService) processProductAttrs(products []*types.ModifyProduct) (map[string]string, map[string]bool, map[string]bool, map[string]int32, map[string]int32, map[string]bool, xerr.XErr) {
	valueKeyMap := make(map[string]string)
	keySet := make(map[string]bool)
	valueSet := make(map[string]bool)
	keyIDMap := make(map[string]int32)
	valueIDMap := make(map[string]int32)
	productAttrSet := make(map[string]bool)
	KeyIDKeyMap := make(map[int32]string)
	ValueIDValueMap := make(map[int32]string)

	for _, product := range products {
		if len(product.ProductAttr) != 0 {
			var attrs []string
			for _, attr := range product.ProductAttr {
				if attr.CheckProductAttr() != nil {
					return nil, nil, nil, nil, nil, nil, xerr.WithCode(xerr.InvalidParams, attr.CheckProductAttr())
				}
				keySet[attr.Key] = true
				valueSet[attr.Value] = true
				valueKeyMap[attr.Value] = attr.Key
				attrs = append(attrs, attr.Key+":"+attr.Value)
				// 更新keyIDMap
				if attr.KeyID != 0 {
					if _, ok := keyIDMap[attr.Key]; !ok {
						keyIDMap[attr.Key] = attr.KeyID
					} else {
						if keyIDMap[attr.Key] != attr.KeyID {
							return nil, nil, nil, nil, nil, nil, xerr.WithCode(xerr.InvalidParams, fmt.Errorf("key %s is already associated with keyID %d, cannot associate with keyID %d", attr.Key, keyIDMap[attr.Key], attr.KeyID))
						}
					}
					if existingKey, ok := KeyIDKeyMap[attr.KeyID]; ok {
						if existingKey != attr.Key {
							return nil, nil, nil, nil, nil, nil, xerr.WithCode(xerr.InvalidParams, fmt.Errorf("keyID %d is already associated with key %s, cannot associate with key %s", attr.KeyID, existingKey, attr.Key))
						}
					} else {
						KeyIDKeyMap[attr.KeyID] = attr.Key
					}

				}
				// 更新valueIDMap
				if attr.ValueID != 0 {
					if _, ok := valueIDMap[attr.Value]; !ok {
						valueIDMap[attr.Value] = attr.ValueID
					} else {
						if valueIDMap[attr.Value] != attr.ValueID {
							return nil, nil, nil, nil, nil, nil, xerr.WithCode(xerr.InvalidParams, fmt.Errorf("value %s is already associated with valueID %d, cannot associate with valueID %d", attr.Value, valueIDMap[attr.Value], attr.ValueID))
						}
					}
					if existingValue, ok := ValueIDValueMap[attr.ValueID]; ok {
						if existingValue != attr.Value {
							return nil, nil, nil, nil, nil, nil, xerr.WithCode(xerr.InvalidParams, fmt.Errorf("valueID %d is already associated with value %s, cannot associate with value %s", attr.ValueID, existingValue, attr.Value))
						}
					} else {
						ValueIDValueMap[attr.ValueID] = attr.Value
					}
				}
			}
			sort.Strings(attrs) // 对属性进行排序以确保属性的顺序不影响唯一性的判断
			attrString := strings.Join(attrs, ",")
			// 检查这个字符串是否已经存在
			if _, ok := productAttrSet[attrString]; ok {
				return nil, nil, nil, nil, nil, nil, xerr.WithCode(xerr.InvalidParams, fmt.Errorf("duplicate product attributes:%s", attrString))
			}
			productAttrSet[attrString] = true
		}
	}

	return valueKeyMap, keySet, valueSet, keyIDMap, valueIDMap, productAttrSet, nil
}

func (s *GoodsService) handlerSpecification(ctx context.Context, oldSpecIDsMap map[int32]*model.Specification, keySet map[string]bool, keyIDMap map[string]int32, goodsID int32, productType enum.ProductVariantType) (map[string]int32, xerr.XErr) {
	for key := range keySet {
		id, ok := keyIDMap[key]
		if ok {
			specification, exists := oldSpecIDsMap[id]
			if !exists {
				return keyIDMap, xerr.WithCode(xerr.InvalidParams, fmt.Errorf("key %s valueId %d is not exist,please check", key, id))
			}
			if specification.Name != key {
				updateValue := &model.Specification{Name: key}
				rowsAffected, err := s.goods.UpdateSpecificationByID(ctx, id, updateValue)
				if err != nil {
					return keyIDMap, xerr.WithCode(xerr.ErrorDatabase, err)
				}
				if rowsAffected == 0 {
					return keyIDMap, xerr.WithCode(xerr.ErrorDatabase, fmt.Errorf("update specification failed,specificationID:%d,wait updateValue %v", id, updateValue))
				}
			}
		} else {
			// 创建新规格，例如：
			newSpecification := &model.Specification{
				Name:    key,
				GoodsID: goodsID,
				Type:    productType,
			}
			specificationID, err := s.goods.CreateSpecification(ctx, newSpecification)
			if err != nil {
				return keyIDMap, xerr.WithCode(xerr.ErrorDatabase, err)
			}
			keyIDMap[key] = specificationID
		}
	}

	keyIDs := make(map[int32]bool)
	for _, id := range keyIDMap {
		keyIDs[id] = true
	}
	missingSpecIDs := make([]int32, 0) // 用于待删去的规格ID
	for _, spec := range oldSpecIDsMap {
		if _, ok := keyIDs[spec.ID]; !ok {
			missingSpecIDs = append(missingSpecIDs, spec.ID)
		}
	}
	if len(missingSpecIDs) > 0 {
		// 删除多余的规格
		rowsAffected, err := s.goods.DeleteSpecificationByIDs(ctx, missingSpecIDs)
		if err != nil {
			return nil, xerr.WithCode(xerr.ErrorDatabase, err)
		}
		if len(missingSpecIDs) != int(rowsAffected) {
			return keyIDMap, xerr.WithCode(xerr.ErrorDatabase, fmt.Errorf("delete specification failed,missingSpecIDs:%v,rowsAffected:%d", missingSpecIDs, rowsAffected))
		}
	}
	return keyIDMap, nil
}

func (s *GoodsService) handlerSpecificationValue(ctx context.Context, oldValueIDsMap map[int32]*model.SpecificationValue, valueSet map[string]bool, valueIDMap map[string]int32, keyIDMap map[string]int32, valueKeyMap map[string]string, goodsID int32, productType enum.ProductVariantType) (map[string]int32, xerr.XErr) {
	missingValueIDs := make([]int32, 0)
	for value := range valueSet {
		id, ok := valueIDMap[value]
		if ok {
			specificationValue, exists := oldValueIDsMap[id]
			if !exists {
				return valueIDMap, xerr.WithCode(xerr.InvalidParams, fmt.Errorf("value %s valueId %d is not exist,please check", value, id))
			}
			if specificationValue.Value != value {
				updateValue := &model.SpecificationValue{Value: value}
				rowsAffected, err := s.goods.UpdateSpecificationValueByID(ctx, id, updateValue)
				if err != nil {
					return valueIDMap, xerr.WithCode(xerr.ErrorDatabase, err)
				}
				if rowsAffected == 0 {
					return valueIDMap, xerr.WithCode(xerr.ErrorDatabase, fmt.Errorf("update specificationValue failed,specificationValueID:%d,wait updateValue %v", id, updateValue))
				}
			}
		} else {
			newSpecificationValue := &model.SpecificationValue{
				Value:           value,
				SpecificationID: keyIDMap[valueKeyMap[value]],
				GoodsID:         goodsID,
				Type:            productType,
			}
			specValueID, err := s.goods.CreateSpecificationValue(ctx, newSpecificationValue)
			if err != nil {
				return nil, xerr.WithCode(xerr.ErrorDatabase, err)
			}
			valueIDMap[value] = specValueID
		}
	}
	valueIDs := make(map[int32]bool)
	for _, id := range valueIDMap {
		valueIDs[id] = true
	}
	for _, specValue := range oldValueIDsMap {
		if _, ok := valueIDs[specValue.ID]; !ok {
			missingValueIDs = append(missingValueIDs, specValue.ID)
		}
	}
	if len(missingValueIDs) > 0 {
		// 删除多余的规格值
		rowsAffected, err := s.goods.DeleteSpecificationValueByIDs(ctx, missingValueIDs)
		if err != nil {
			return valueIDMap, xerr.WithCode(xerr.ErrorDatabase, err)
		}
		if len(missingValueIDs) != int(rowsAffected) {
			return valueIDMap, xerr.WithCode(xerr.ErrorDatabase, fmt.Errorf("delete specificationValue failed,missingValueIDs:%v,rowsAffected:%d", missingValueIDs, rowsAffected))
		}
	}
	return valueIDMap, nil
}
