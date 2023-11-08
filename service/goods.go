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
	////获取用户ID
	//userID := common.GetUserID(ctx)
	////获取用户角色
	//userRole := common.GetUserRole(ctx)
	//if userRole != model.UserRoleSupplier {
	//	return xerr.WithCode(xerr.ErrorAuthInsufficientAuthority, errors.New("用户权限不是供应商"))
	//}
	if len(req.GoodsDetail.Images) == 0 {
		return xerr.WithCode(xerr.InvalidParams, errors.New("商品图片不能为空"))
	}
	userID := "123456"
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
		if _err := s.CreateWithTransaction(ctx, req, modelGoods); _err != nil {
			return _err
		}
		return nil
	}); xErr != nil {
		return xErr
	}
	return nil
}

func (s *GoodsService) CreateWithTransaction(ctx context.Context, req types.GoodsAddReq, modelGoods *model.Goods) xerr.XErr {
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
				valueKeyMap[attr.Value] = attr.Key
				attrs = append(attrs, attr.Key+":"+attr.Value)
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
		skuCode := fmt.Sprintf("SK%d%d%s", time.Now().Unix(), goodsID, utils.GenerateRandomNumber(3))
		productAttrBytes, _ := json.Marshal(products[i].ProductAttr)
		productVariant := &model.ProductVariant{
			SKUCode:          skuCode,
			GoodsID:          goodsID,
			Unit:             products[i].Unit,
			Price:            products[i].Price,
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
	userID := "123456"

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
	//获取用户ID
	//userID := common.GetUserID(ctx)
	////获取用户角色
	//userRole := common.GetUserRole(ctx)
	userID := "123456"
	user, err := s.userDao.GetByUserID(ctx, userID)
	if err != nil {
		return xerr.WithCode(xerr.ErrorDatabase, err)
	}
	userRole := user.UserRole
	if userRole != model.UserRoleSupplier {
		return xerr.WithCode(xerr.ErrorAuthInsufficientAuthority, errors.New("用户权限不是供应商"))
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
	return nil

}

func (s *GoodsService) ModifyMyGoodsStatus(c *gin.Context, req types.GoodsReq) xerr.XErr {
	//获取用户ID
	//userID := common.GetUserID(c)
	////获取用户角色
	//userRole := common.GetUserRole(c)
	userID := "123456"
	user, err := s.userDao.GetByUserID(c, userID)
	if err != nil {
		return xerr.WithCode(xerr.ErrorDatabase, err)
	}

	userRole := user.UserRole
	if userRole != model.UserRoleSupplier {
		return xerr.WithCode(xerr.ErrorAuthInsufficientAuthority, errors.New("用户不是供应商"))
	}
	if req.GoodsStatus != enum.GoodsStatusOnSale && req.GoodsStatus != enum.GoodsStatusOffSale {
		return xerr.WithCode(xerr.InvalidParams, errors.New("商品状态不合法"))
	}
	//检查用户权限
	_, xrr := s.checkGoodsValid(c, req.GoodsID, userID)
	if xrr != nil {
		return xrr
	}
	//修改商品状态
	err = s.goods.UpdateGoodsByID(c, req.GoodsID, &model.Goods{Status: req.GoodsStatus})
	if err != nil {
		return xerr.WithCode(xerr.ErrorDatabase, err)
	}
	return nil
}

func (s *GoodsService) GetMyGoodsList(c *gin.Context, req types.MyGoodsListReq) (*types.GoodsListResp, xerr.XErr) {
	//获取用户ID
	//userID := common.GetUserID(c)
	////获取用户角色
	//userRole := common.GetUserRole(c)
	userID := "123456"
	user, err := s.userDao.GetByUserID(c, userID)
	if err != nil {
		return nil, xerr.WithCode(xerr.ErrorDatabase, err)
	}
	if user == nil {
		return nil, xerr.WithCode(xerr.ErrorNotExistUser, errors.New("用户不存在"))
	}
	req.UserID = userID
	userRole := user.UserRole
	if userRole != model.UserRoleSupplier {
		return nil, xerr.WithCode(xerr.ErrorAuthInsufficientAuthority, errors.New("用户不是供应商"))
	}
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
	userID := "123456"
	user, err := s.userDao.GetByUserID(c, userID)
	if err != nil {
		return nil, xerr.WithCode(xerr.ErrorDatabase, err)
	}
	if user == nil {
		return nil, xerr.WithCode(xerr.ErrorNotExistUser, errors.New("用户不存在"))
	}
	if req.ReqType == 1 {
		return s.handleSupplierGoodsDetailRequest(c, req, user, goods, goodsDetail)
	} else {
		return s.handleGoodsDetailRequestByRole(c, req, user, goods, goodsDetail)
	}
}
func (s *GoodsService) handleSupplierGoodsDetailRequest(c *gin.Context, req types.GoodsReq, user *model.User, goods *model.Goods, goodsDetail *types.GoodsDetailResp) (*types.GoodsDetailResp, xerr.XErr) {
	if user.UserRole != model.UserRoleSupplier {
		return nil, xerr.WithCode(xerr.ErrorAuthInsufficientAuthority, errors.New("用户不是供应商"))
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

func (s *GoodsService) handleGoodsDetailRequestByRole(c *gin.Context, req types.GoodsReq, user *model.User, goods *model.Goods, goodsDetail *types.GoodsDetailResp) (*types.GoodsDetailResp, xerr.XErr) {
	if user.UserRole == model.UserRoleCustomer {
		if *goods.IsRetail == 0 || goods.RetailStatus == enum.GoodsRetailSoldOut {
			return nil, xerr.WithCode(xerr.ErrorNotExistRecord, errors.New("商品不存在"))
		}

		if err := s.fillRetailDetails(c, req, goods, goodsDetail, enum.ProductVariantEnabled); err != nil {
			return nil, err
		}
	}
	if user.UserRole == model.UserRoleSupplier || user.UserRole == model.UserRoleBuyer {
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
	userID := "123456"

	user, err := s.userDao.GetByUserID(c, userID)
	if err != nil {
		return xerr.WithCode(xerr.ErrorDatabase, err)
	}
	if user == nil {
		return xerr.WithCode(xerr.ErrorNotExistUser, errors.New("用户不存在"))
	}
	if user.UserRole != model.UserRoleSupplier {
		return xerr.WithCode(xerr.ErrorAuthInsufficientAuthority, errors.New("用户不是供应商"))
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

	//
	////2.如果不修改规格信息,只修改SKU信息,根据ID直接更新
	////todo:更新是商品否售罄，是否下架，是否零售状态
	//if len(req.SpecInfo) == 0 {
	//	if len(req.WholesaleProducts) != 0 {
	//		for i := range req.WholesaleProducts {
	//			if req.WholesaleProducts[i].ID == 0 {
	//				return xerr.WithCode(xerr.InvalidParams, fmt.Errorf("wholesaleProducts[%d].ID is invalid", i))
	//			}
	//			if req.WholesaleProducts[i].ID != 0 {
	//				updateProductVariant := &model.ProductVariant{
	//					ID:               req.WholesaleProducts[i].ID,
	//					Unit:             req.WholesaleProducts[i].Unit,
	//					Price:            req.WholesaleProducts[i].Price,
	//					MinOrderQuantity: req.WholesaleProducts[i].MinOrderQuantity,
	//					Status:           req.WholesaleProducts[i].Status,
	//				}
	//				rowsAffected, err := s.goods.UpdateProductVariantByID(c, req.WholesaleProducts[i].ID, updateProductVariant)
	//				if err != nil {
	//					return xerr.WithCode(xerr.ErrorDatabase, err)
	//				}
	//				if rowsAffected == 0 {
	//					return xerr.WithCode(xerr.InvalidParams, fmt.Errorf("wholesaleProducts[%d].ID[%d] is invalid", i, req.WholesaleProducts[i].ID))
	//				}
	//			}
	//		}
	//	}
	//	if len(req.RetailProducts) != 0 {
	//		for i := range req.RetailProducts {
	//			if req.RetailProducts[i].ID == 0 {
	//				return xerr.WithCode(xerr.InvalidParams, fmt.Errorf("retailProducts[%d].ID is invalid", i))
	//			}
	//			if req.RetailProducts[i].ID != 0 {
	//				updateProductVariant := &model.ProductVariant{
	//					ID:     req.RetailProducts[i].ID,
	//					Unit:   req.RetailProducts[i].Unit,
	//					Price:  req.RetailProducts[i].Price,
	//					Stock:  req.RetailProducts[i].Stock,
	//					Status: req.RetailProducts[i].Status,
	//				}
	//				rowsAffected, err := s.goods.UpdateProductVariantByID(c, req.RetailProducts[i].ID, updateProductVariant)
	//				if err != nil {
	//					return xerr.WithCode(xerr.ErrorDatabase, err)
	//				}
	//				if rowsAffected == 0 {
	//					return xerr.WithCode(xerr.InvalidParams, fmt.Errorf("retailProducts[%d].ID[%d] is invalid", i, req.RetailProducts[i].ID))
	//				}
	//			}
	//		}
	//	}
	//}
	////3.如果修改规格信息
	//if len(req.SpecInfo) != 0 {
	//	specAllMap := make(map[int32]*model.Specification)
	//	specWholesaleMap := make(map[int32]*model.Specification)
	//	specRetailTypeIDMap := make(map[int32]*model.Specification)
	//	specDeletedMap := make(map[int32]*model.Specification)
	//	specifications, err := s.goods.GetSpecificationByGoodsID(c, req.Id, 0)
	//	if err != nil {
	//		return xerr.WithCode(xerr.ErrorDatabase, err)
	//	}
	//	for _, spec := range specifications {
	//		specAllMap[spec.ID] = spec
	//		if spec.Type == enum.ProductWholesale {
	//			specWholesaleMap[spec.ID] = spec
	//		}
	//		if spec.Type == enum.ProductRetail {
	//			specRetailTypeIDMap[spec.ID] = spec
	//		}
	//	}
	//	for i, spec := range req.SpecInfo {
	//		if spec.DeletedSpecID == 0 {
	//			if spec.SpecID != 0 {
	//				//检查规格ID是否存在
	//				if _, ok := specAllMap[spec.SpecID]; !ok {
	//					return xerr.WithCode(xerr.InvalidParams, fmt.Errorf("SpecInfo[%d].SpecID:%d is invalid,not exist", i, spec.SpecID))
	//				}
	//				if spec.SpecName == "" && len(spec.SpecValue) == 0 {
	//					continue
	//				}
	//			}
	//		}
	//		if spec.DeletedSpecID != 0 {
	//			//检查规格ID是否存在
	//			if _, ok := specAllMap[spec.DeletedSpecID]; !ok {
	//				return xerr.WithCode(xerr.InvalidParams, fmt.Errorf("SpecInfo[%d].SpecID:%d is invalid,not exist", i, spec.SpecID))
	//			}
	//			if _, ok := specDeletedMap[spec.DeletedSpecID]; !ok {
	//				//删除规格信息
	//				err = s.goods.DeleteSpecificationByID(c, spec.DeletedSpecID)
	//				if err != nil {
	//					return xerr.WithCode(xerr.ErrorDatabase, err)
	//				}
	//				specDeletedMap[spec.DeletedSpecID] = specAllMap[spec.DeletedSpecID]
	//			}
	//		}
	//	}
	//}
	//var updateWholesaleSpec, updateRetailSpec bool
	//for i := range req.SpecInfo {
	//	if req.SpecInfo[i].Type == enum.ProductWholesale {
	//		updateWholesaleSpec = true
	//	}
	//	if req.SpecInfo[i].Type == enum.ProductRetail {
	//		updateRetailSpec = true
	//	}
	//}
	//if updateWholesaleSpec {
	//	specNameToIDW := make(map[string]int32)
	//	specValueToModelW := make(map[string]*model.SpecificationValue)
	//	for j := range req.SpecInfo {
	//		if req.SpecInfo[j].Type == enum.ProductWholesale {
	//			if req.SpecInfo[j].DeletedSpecID == 0 {
	//				//新增规格
	//				if req.SpecInfo[j].SpecID == 0 && req.SpecInfo[j].SpecName != "" {
	//					if _, ok := specNameToIDW[req.SpecInfo[j].SpecName]; !ok {
	//						specification := &model.Specification{
	//							Name:    req.SpecInfo[j].SpecName,
	//							Type:    enum.ProductWholesale,
	//							GoodsID: req.Id,
	//						}
	//						specificationID, err := s.goods.CreateSpecification(c, specification)
	//						if err != nil {
	//							return xerr.WithCode(xerr.ErrorDatabase, err)
	//						}
	//						specNameToIDW[req.SpecInfo[j].SpecName] = specificationID
	//					}
	//					//新增规格值
	//					if len(req.SpecInfo[j].SpecValue) == 0 {
	//						return xerr.WithCode(xerr.InvalidParams, errors.New("新增规格，规格值不能为空"))
	//					}
	//					for k := range req.SpecInfo[j].SpecValue {
	//						if req.SpecInfo[j].SpecValue[k].Value == "" {
	//							return xerr.WithCode(xerr.InvalidParams, errors.New("新增规格，规格值不能为空"))
	//						}
	//						if _, ok := specValueToModelW[req.SpecInfo[j].SpecValue[k].Value]; !ok {
	//							specificationValue := &model.SpecificationValue{
	//								SpecificationID: specNameToIDW[req.SpecInfo[j].SpecName],
	//								Value:           req.SpecInfo[j].SpecValue[k].Value,
	//								Type:            enum.ProductWholesale,
	//								GoodsID:         req.Id,
	//							}
	//							specificationValueID, err := s.goods.CreateSpecificationValue(c, specificationValue)
	//							if err != nil {
	//								return xerr.WithCode(xerr.ErrorDatabase, err)
	//							}
	//							specificationValue.ID = specificationValueID
	//							specificationValue.SpecificationName = req.SpecInfo[j].SpecName
	//							specValueToModelW[req.SpecInfo[j].SpecValue[k].Value] = specificationValue
	//						}
	//					}
	//				}
	//				//修改规格信息
	//				if req.SpecInfo[j].SpecID != 0 {
	//					//修改规格名称
	//					if req.SpecInfo[j].SpecName != "" {
	//						err = s.goods.UpdateSpecificationByID(c, req.SpecInfo[j].SpecID, &model.Specification{Name: req.SpecInfo[j].SpecName})
	//						if err != nil {
	//							return xerr.WithCode(xerr.ErrorDatabase, err)
	//						}
	//						specNameToIDW[req.SpecInfo[j].SpecName] = req.SpecInfo[j].SpecID
	//					} else {
	//						specNameToIDW[req.SpecInfo[j].SpecName] = req.SpecInfo[j].SpecID
	//					}
	//
	//					if len(req.SpecInfo[j].SpecValue) != 0 {
	//						for k := range req.SpecInfo[j].SpecValue {
	//							//1.删除规格值信息
	//							if req.SpecInfo[j].SpecValue[k].DeletedSpecVID != 0 {
	//								err = s.goods.DeleteSpecificationValueByID(c, req.SpecInfo[j].SpecValue[k].DeletedSpecVID)
	//								if err != nil {
	//									return xerr.WithCode(xerr.ErrorDatabase, err)
	//								}
	//							}
	//							//2.修改规格值信息
	//							if req.SpecInfo[j].SpecValue[k].DeletedSpecVID != 0 {
	//								if req.SpecInfo[j].SpecValue[k].Value != "" && req.SpecInfo[j].SpecValue[k].ID != 0 {
	//									if _, ok := specValueToModelW[req.SpecInfo[j].SpecValue[k].Value]; !ok {
	//										err = s.goods.UpdateSpecificationValueByID(c, req.SpecInfo[j].SpecValue[k].ID, &model.SpecificationValue{Value: req.SpecInfo[j].SpecValue[k].Value})
	//										if err != nil {
	//											return xerr.WithCode(xerr.ErrorDatabase, err)
	//										}
	//										specValueToModelW[req.SpecInfo[j].SpecValue[k].Value] = &model.SpecificationValue{
	//											ID:                req.SpecInfo[j].SpecValue[k].ID,
	//											SpecificationID:   req.SpecInfo[j].SpecID,
	//											SpecificationName: req.SpecInfo[j].SpecName,
	//											Value:             req.SpecInfo[j].SpecValue[k].Value,
	//											GoodsID:           req.Id,
	//											Type:              enum.ProductWholesale,
	//										}
	//									}
	//								}
	//								if req.SpecInfo[j].SpecValue[k].Value != "" && req.SpecInfo[j].SpecValue[k].ID == 0 {
	//									if _, ok := specValueToModelW[req.SpecInfo[j].SpecValue[k].Value]; !ok {
	//										specificationValue := &model.SpecificationValue{
	//											SpecificationID:   req.SpecInfo[j].SpecID,
	//											SpecificationName: req.SpecInfo[j].SpecName,
	//											Value:             req.SpecInfo[j].SpecValue[k].Value,
	//											GoodsID:           req.Id,
	//											Type:              enum.ProductWholesale,
	//										}
	//										specificationValueID, err := s.goods.CreateSpecificationValue(c, specificationValue)
	//										if err != nil {
	//											return xerr.WithCode(xerr.ErrorDatabase, err)
	//										}
	//										specificationValue.ID = specificationValueID
	//										specValueToModelW[req.SpecInfo[j].SpecValue[k].Value] = specificationValue
	//									}
	//								}
	//							}
	//
	//						}
	//					}
	//
	//				}
	//
	//			}
	//			if req.SpecInfo[j].DeletedSpecID != 0 {
	//				//1.删除规格信息
	//				err = s.goods.DeleteSpecificationByID(c, req.SpecInfo[j].DeletedSpecID)
	//				if err != nil {
	//					return xerr.WithCode(xerr.ErrorDatabase, err)
	//				}
	//				//2.删除规格值信息
	//				err = s.goods.DeleteSpecificationValueBySpecID(c, req.SpecInfo[j].DeletedSpecID)
	//				if err != nil {
	//					return xerr.WithCode(xerr.ErrorDatabase, err)
	//				}
	//			}
	//		}
	//	}
	//	if len(req.WholesaleProducts) == 0 {
	//		return xerr.WithCode(xerr.InvalidParams, errors.New("批发产品不能为空"))
	//	}
	//	for _, wholesaleProduct := range req.WholesaleProducts {
	//		if wholesaleProduct.ID == 0 {
	//			if wholesaleProduct.MinOrderQuantity == 0 {
	//				return xerr.WithCode(xerr.InvalidParams, errors.New("最小起订量不能为空"))
	//			}
	//			if wholesaleProduct.CheckParams() != nil {
	//				return xerr.WithCode(xerr.InvalidParams, wholesaleProduct.CheckParams())
	//			}
	//			if len(wholesaleProduct.ProductAttr) == 0 {
	//				return xerr.WithCode(xerr.InvalidParams, errors.New("批发产品规格不能为空"))
	//			}
	//			for i, v := range wholesaleProduct.ProductAttr {
	//				if v.CheckProductAttr() != nil {
	//					return xerr.WithCode(xerr.InvalidParams, v.CheckProductAttr())
	//				}
	//				if v.KeyID == 0 {
	//					wholesaleProduct.ProductAttr[i].KeyID = specNameToIDW[v.Key]
	//				}
	//				if v.ValueID == 0 {
	//					wholesaleProduct.ProductAttr[i].ValueID = specValueToModelW[v.Value].ID
	//				}
	//			}
	//			var specAValueID, specBValueID int32
	//			if len(wholesaleProduct.ProductAttr) != 0 {
	//				specAValueID = wholesaleProduct.ProductAttr[0].ValueID
	//			}
	//			if len(wholesaleProduct.ProductAttr) == 2 {
	//				specBValueID = wholesaleProduct.ProductAttr[1].ValueID
	//			}
	//			productAttrBytes, _ := json.Marshal(wholesaleProduct.ProductAttr)
	//			skuCode := fmt.Sprintf("SK%d%d%d", time.Now().UnixNano(), specAValueID, specBValueID)
	//			productVariant := &model.ProductVariant{
	//				SKUCode:          skuCode,
	//				GoodsID:          req.Id,
	//				Unit:             wholesaleProduct.Unit,
	//				Price:            wholesaleProduct.Price,
	//				MinOrderQuantity: wholesaleProduct.MinOrderQuantity,
	//				Type:             enum.ProductWholesale,
	//				Status:           wholesaleProduct.Status,
	//				ProductAttr:      string(productAttrBytes),
	//			}
	//			_, err = s.goods.CreateProductVariant(c, productVariant)
	//			if err != nil {
	//				return xerr.WithCode(xerr.ErrorDatabase, err)
	//			}
	//		} else {
	//			if len(wholesaleProduct.ProductAttr) == 0 {
	//				return xerr.WithCode(xerr.InvalidParams, errors.New("批发产品规格不能为空"))
	//			}
	//			for i, v := range wholesaleProduct.ProductAttr {
	//				if v.CheckProductAttr() != nil {
	//					return xerr.WithCode(xerr.InvalidParams, v.CheckProductAttr())
	//				}
	//				if v.KeyID == 0 {
	//					wholesaleProduct.ProductAttr[i].KeyID = specNameToIDW[v.Key]
	//				}
	//				if v.ValueID == 0 {
	//					wholesaleProduct.ProductAttr[i].ValueID = specValueToModelW[v.Value].ID
	//				}
	//			}
	//			productAttrBytes, _ := json.Marshal(wholesaleProduct.ProductAttr)
	//			productVariantUpdate := &model.ProductVariant{
	//				Unit:             wholesaleProduct.Unit,
	//				Price:            wholesaleProduct.Price,
	//				MinOrderQuantity: wholesaleProduct.MinOrderQuantity,
	//				Type:             enum.ProductWholesale,
	//				Status:           wholesaleProduct.Status,
	//				ProductAttr:      string(productAttrBytes),
	//			}
	//			rowsAffected, err := s.goods.UpdateProductVariantByID(c, wholesaleProduct.ID, productVariantUpdate)
	//			if err != nil {
	//				return xerr.WithCode(xerr.ErrorDatabase, err)
	//			}
	//			if rowsAffected == 0 {
	//				return xerr.WithCode(xerr.InvalidParams, fmt.Errorf("wholesaleProducts.ID[%d] is invalid", wholesaleProduct.ID))
	//			}
	//		}
	//	}
	//} else {
	//	if len(req.RetailProducts) != 0 {
	//		for i := range req.RetailProducts {
	//			if req.RetailProducts[i].ID == 0 {
	//				return xerr.WithCode(xerr.InvalidParams, errors.New("批发产品ID不能为空"))
	//			}
	//			if req.RetailProducts[i].ID != 0 {
	//				updateProductVariant := &model.ProductVariant{
	//					ID:     req.RetailProducts[i].ID,
	//					Unit:   req.RetailProducts[i].Unit,
	//					Price:  req.RetailProducts[i].Price,
	//					Stock:  req.RetailProducts[i].Stock,
	//					Status: req.RetailProducts[i].Status,
	//				}
	//				rowsAffected, err := s.goods.UpdateProductVariantByID(c, req.RetailProducts[i].ID, updateProductVariant)
	//				if err != nil {
	//					return xerr.WithCode(xerr.ErrorDatabase, err)
	//				}
	//				if rowsAffected == 0 {
	//					return xerr.WithCode(xerr.InvalidParams, fmt.Errorf("wholesaleProducts.ID[%d] is invalid", req.RetailProducts[i].ID))
	//				}
	//			}
	//		}
	//	}
	//}
	//if updateRetailSpec {
	//	specNameToIDR := make(map[string]int32)
	//	specValueToModelR := make(map[string]*model.SpecificationValue)
	//	for j := range req.SpecInfo {
	//		if req.SpecInfo[j].Type == enum.ProductRetail {
	//			if req.SpecInfo[j].DeletedSpecID == 0 {
	//				//新增规格
	//				if req.SpecInfo[j].SpecID == 0 && req.SpecInfo[j].SpecName != "" {
	//					if _, ok := specNameToIDR[req.SpecInfo[j].SpecName]; !ok {
	//						specification := &model.Specification{
	//							Name:    req.SpecInfo[j].SpecName,
	//							Type:    enum.ProductRetail,
	//							GoodsID: req.Id,
	//						}
	//						specificationID, err := s.goods.CreateSpecification(c, specification)
	//						if err != nil {
	//							return xerr.WithCode(xerr.ErrorDatabase, err)
	//						}
	//						specNameToIDR[req.SpecInfo[j].SpecName] = specificationID
	//					}
	//					//新增规格值
	//					if len(req.SpecInfo[j].SpecValue) == 0 {
	//						return xerr.WithCode(xerr.InvalidParams, errors.New("新增规格，规格值不能为空"))
	//					}
	//					for k := range req.SpecInfo[j].SpecValue {
	//						if req.SpecInfo[j].SpecValue[k].Value == "" {
	//							return xerr.WithCode(xerr.InvalidParams, errors.New("新增规格，规格值不能为空"))
	//						}
	//						if _, ok := specNameToIDR[req.SpecInfo[j].SpecValue[k].Value]; !ok {
	//							specificationValue := &model.SpecificationValue{
	//								SpecificationID: specNameToIDR[req.SpecInfo[j].SpecName],
	//								Value:           req.SpecInfo[j].SpecValue[k].Value,
	//								Type:            enum.ProductRetail,
	//								GoodsID:         req.Id,
	//							}
	//							specificationValueID, err := s.goods.CreateSpecificationValue(c, specificationValue)
	//							if err != nil {
	//								return xerr.WithCode(xerr.ErrorDatabase, err)
	//							}
	//							specificationValue.ID = specificationValueID
	//							specificationValue.SpecificationName = req.SpecInfo[j].SpecName
	//							specValueToModelR[req.SpecInfo[j].SpecValue[k].Value] = specificationValue
	//						}
	//					}
	//				}
	//				//修改规格信息
	//				if req.SpecInfo[j].SpecID != 0 {
	//					//修改规格名称
	//					if req.SpecInfo[j].SpecName != "" {
	//						err = s.goods.UpdateSpecificationByID(c, req.SpecInfo[j].SpecID, &model.Specification{Name: req.SpecInfo[j].SpecName})
	//						if err != nil {
	//							return xerr.WithCode(xerr.ErrorDatabase, err)
	//						}
	//						specNameToIDR[req.SpecInfo[j].SpecName] = req.SpecInfo[j].SpecID
	//					} else {
	//						specNameToIDR[req.SpecInfo[j].SpecName] = req.SpecInfo[j].SpecID
	//					}
	//
	//					if len(req.SpecInfo[j].SpecValue) != 0 {
	//						for k := range req.SpecInfo[j].SpecValue {
	//							//1.删除规格值信息
	//							if req.SpecInfo[j].SpecValue[k].DeletedSpecVID != 0 {
	//								err = s.goods.DeleteSpecificationValueByID(c, req.SpecInfo[j].SpecValue[k].DeletedSpecVID)
	//								if err != nil {
	//									return xerr.WithCode(xerr.ErrorDatabase, err)
	//								}
	//							}
	//							//2.修改规格值信息
	//							if req.SpecInfo[j].SpecValue[k].DeletedSpecVID != 0 {
	//								if req.SpecInfo[j].SpecValue[k].Value != "" && req.SpecInfo[j].SpecValue[k].ID != 0 {
	//									if _, ok := specValueToModelR[req.SpecInfo[j].SpecValue[k].Value]; !ok {
	//										err = s.goods.UpdateSpecificationValueByID(c, req.SpecInfo[j].SpecValue[k].ID, &model.SpecificationValue{Value: req.SpecInfo[j].SpecValue[k].Value})
	//										if err != nil {
	//											return xerr.WithCode(xerr.ErrorDatabase, err)
	//										}
	//										specValueToModelR[req.SpecInfo[j].SpecValue[k].Value] = &model.SpecificationValue{
	//											ID:                req.SpecInfo[j].SpecValue[k].ID,
	//											SpecificationID:   req.SpecInfo[j].SpecID,
	//											SpecificationName: req.SpecInfo[j].SpecName,
	//											Value:             req.SpecInfo[j].SpecValue[k].Value,
	//											GoodsID:           req.Id,
	//											Type:              enum.ProductRetail,
	//										}
	//									}
	//								}
	//								if req.SpecInfo[j].SpecValue[k].Value != "" && req.SpecInfo[j].SpecValue[k].ID == 0 {
	//									if _, ok := specValueToModelR[req.SpecInfo[j].SpecValue[k].Value]; !ok {
	//										specificationValue := &model.SpecificationValue{
	//											SpecificationID:   req.SpecInfo[j].SpecID,
	//											SpecificationName: req.SpecInfo[j].SpecName,
	//											Value:             req.SpecInfo[j].SpecValue[k].Value,
	//											GoodsID:           req.Id,
	//											Type:              enum.ProductRetail,
	//										}
	//										specificationValueID, err := s.goods.CreateSpecificationValue(c, specificationValue)
	//										if err != nil {
	//											return xerr.WithCode(xerr.ErrorDatabase, err)
	//										}
	//										specificationValue.ID = specificationValueID
	//										specValueToModelR[req.SpecInfo[j].SpecValue[k].Value] = specificationValue
	//									}
	//								}
	//							}
	//
	//						}
	//					}
	//
	//				}
	//
	//			}
	//			if req.SpecInfo[j].DeletedSpecID != 0 {
	//				//1.删除规格信息
	//				err = s.goods.DeleteSpecificationByID(c, req.SpecInfo[j].DeletedSpecID)
	//				if err != nil {
	//					return xerr.WithCode(xerr.ErrorDatabase, err)
	//				}
	//				//2.删除规格值信息
	//				err = s.goods.DeleteSpecificationValueBySpecID(c, req.SpecInfo[j].DeletedSpecID)
	//				if err != nil {
	//					return xerr.WithCode(xerr.ErrorDatabase, err)
	//				}
	//			}
	//		}
	//	}
	//	if len(req.RetailProducts) == 0 {
	//		return xerr.WithCode(xerr.InvalidParams, errors.New("零售产品不能为空"))
	//	}
	//	for _, retailProduct := range req.RetailProducts {
	//		if retailProduct.ID == 0 {
	//			if *retailProduct.Stock == 0 {
	//				return xerr.WithCode(xerr.InvalidParams, errors.New("最小起订量不能为空"))
	//			}
	//			if retailProduct.CheckParams() != nil {
	//				return xerr.WithCode(xerr.InvalidParams, retailProduct.CheckParams())
	//			}
	//			if len(retailProduct.ProductAttr) == 0 {
	//				return xerr.WithCode(xerr.InvalidParams, errors.New("批发产品规格不能为空"))
	//			}
	//			for i, v := range retailProduct.ProductAttr {
	//				if v.CheckProductAttr() != nil {
	//					return xerr.WithCode(xerr.InvalidParams, v.CheckProductAttr())
	//				}
	//				if v.KeyID == 0 {
	//					retailProduct.ProductAttr[i].KeyID = specNameToIDR[v.Key]
	//				}
	//				if v.ValueID == 0 {
	//					retailProduct.ProductAttr[i].ValueID = specValueToModelR[v.Value].ID
	//				}
	//			}
	//			var specAValueID, specBValueID int32
	//			if len(retailProduct.ProductAttr) != 0 {
	//				specAValueID = retailProduct.ProductAttr[0].ValueID
	//			}
	//			if len(retailProduct.ProductAttr) == 2 {
	//				specBValueID = retailProduct.ProductAttr[1].ValueID
	//			}
	//			productAttrBytes, _ := json.Marshal(retailProduct.ProductAttr)
	//			skuCode := fmt.Sprintf("SK%d%d%d", time.Now().UnixNano(), specAValueID, specBValueID)
	//			productVariant := &model.ProductVariant{
	//				SKUCode:     skuCode,
	//				GoodsID:     req.Id,
	//				Unit:        retailProduct.Unit,
	//				Price:       retailProduct.Price,
	//				Stock:       retailProduct.Stock,
	//				Type:        enum.ProductRetail,
	//				Status:      retailProduct.Status,
	//				ProductAttr: string(productAttrBytes),
	//			}
	//			_, err = s.goods.CreateProductVariant(c, productVariant)
	//			if err != nil {
	//				return xerr.WithCode(xerr.ErrorDatabase, err)
	//			}
	//		} else {
	//			if len(retailProduct.ProductAttr) == 0 {
	//				return xerr.WithCode(xerr.InvalidParams, errors.New("批发产品规格不能为空"))
	//			}
	//			for i, v := range retailProduct.ProductAttr {
	//				if v.CheckProductAttr() != nil {
	//					return xerr.WithCode(xerr.InvalidParams, v.CheckProductAttr())
	//				}
	//				if v.KeyID == 0 {
	//					retailProduct.ProductAttr[i].KeyID = specNameToIDR[v.Key]
	//				}
	//				if v.ValueID == 0 {
	//					retailProduct.ProductAttr[i].ValueID = specValueToModelR[v.Value].ID
	//				}
	//			}
	//			productAttrBytes, _ := json.Marshal(retailProduct.ProductAttr)
	//			productVariantUpdate := &model.ProductVariant{
	//				Unit:             retailProduct.Unit,
	//				Price:            retailProduct.Price,
	//				MinOrderQuantity: retailProduct.MinOrderQuantity,
	//				Status:           retailProduct.Status,
	//				ProductAttr:      string(productAttrBytes),
	//			}
	//
	//			rowsAffected, err := s.goods.UpdateProductVariantByID(c, retailProduct.ID, productVariantUpdate)
	//			if err != nil {
	//				return xerr.WithCode(xerr.ErrorDatabase, err)
	//			}
	//			if rowsAffected == 0 {
	//				return xerr.WithCode(xerr.InvalidParams, fmt.Errorf("retailProducts.ID[%d] is invalid", retailProduct.ID))
	//			}
	//		}
	//	}
	//} else {
	//	if len(req.WholesaleProducts) != 0 {
	//		for i := range req.WholesaleProducts {
	//			if req.WholesaleProducts[i].ID == 0 {
	//				return xerr.WithCode(xerr.InvalidParams, errors.New("批发产品ID不能为空"))
	//			}
	//
	//			if req.WholesaleProducts[i].ID != 0 {
	//				updateProductVariant := &model.ProductVariant{
	//					ID:               req.WholesaleProducts[i].ID,
	//					Unit:             req.WholesaleProducts[i].Unit,
	//					Price:            req.WholesaleProducts[i].Price,
	//					MinOrderQuantity: req.WholesaleProducts[i].MinOrderQuantity,
	//					Status:           req.WholesaleProducts[i].Status,
	//				}
	//
	//				rowsAffected, err := s.goods.UpdateProductVariantByID(c, req.WholesaleProducts[i].ID, updateProductVariant)
	//				if err != nil {
	//					return xerr.WithCode(xerr.ErrorDatabase, err)
	//				}
	//				if rowsAffected == 0 {
	//					return xerr.WithCode(xerr.InvalidParams, fmt.Errorf("wholesaleProducts[%d].ID[%d] is invalid", i, req.WholesaleProducts[i].ID))
	//				}
	//			}
	//		}
	//	}
	//}
	s.handlerModifyProducts(c, req.WholesaleProducts, enum.ProductWholesale)
	if len(req.WholesaleProducts) != 0 {

	}
	if len(req.RetailProducts) != 0 {

	}
	err = s.goods.UpdateGoodsByID(c, req.Id, updateValue)
	if err != nil {
		return xerr.WithCode(xerr.ErrorDatabase, err)
	}
	return nil
}

func (s *GoodsService) checkGoodsValid(c context.Context, goodsID int32, userID string) (*model.Goods, xerr.XErr) {
	goods, err := s.goods.GetGoodsByGoodsID(c, goodsID)
	if err != nil {
		return nil, xerr.WithCode(xerr.ErrorDatabase, err)
	}
	if goods == nil {
		return nil, xerr.WithCode(xerr.ErrorNotExistRecord, errors.New("商品不存在"))
	}
	if userID == "" {
		return goods, nil
	} else {
		if goods.UserID != userID {
			return nil, xerr.WithCode(xerr.ErrorAuthInsufficientAuthority, errors.New("用户没有权限修改该商品信息"))
		}
		return goods, nil
	}
}

func (s *GoodsService) handlerModifyProducts(c *gin.Context, products []*types.ModifyProduct, status enum.ProductVariantType) xerr.XErr {
	keyValueMap := make(map[string][]string)
	keyIDMap := make(map[string]int32)
	valueIDMap := make(map[string]int32)
	keySet := make(map[string]bool)
	valueSet := make(map[string]bool)
	for _, product := range products {
		if len(product.ProductAttr) != 0 {
			for _, attr := range product.ProductAttr {
				keySet[attr.Key] = true
				valueSet[attr.Value] = true

				keyIDMap[attr.Key] = attr.KeyID
				valueIDMap[attr.Value] = attr.ValueID

				if _, ok := keyValueMap[attr.Key]; !ok {
					keyValueMap[attr.Key] = []string{}
				}
				keyValueMap[attr.Key] = append(keyValueMap[attr.Key], attr.Value)
			}
		}

	}
	return nil
}
