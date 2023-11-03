package service

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"strconv"
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
	goods *dao.GoodsDao
}

func NewGoodsService() *GoodsService {
	return &GoodsService{
		goods: dao.NewGoodsDao(),
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
	if len(req.GoodsDetail.DescImages) != 0 {
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
		RetailStatus:       enum.GoodsRetailSoldOut,
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
	if len(req.RetailProducts) != 0 {
		modelGoods.IsRetail = 1
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
			}
			_, err = s.goods.CreateProductVariant(ctx, productVariant)
			if err != nil {
				return xerr.WithCode(xerr.ErrorDatabase, err)
			}
		}
	}
	return nil
}

func (s *GoodsService) GetGoodsList(ctx *gin.Context, req types.GoodsListReq) xerr.XErr {
	//获取用户ID
	//userID := common.GetUserID(ctx)
	//获取用户角色
	userRole := common.GetUserRole(ctx)
	if req.RequestSource == enum.MiniProgramSource {
		//小程序后台
		if userRole == model.UserRoleSupplier {
			//        1. 弹窗内按发布时间由新到旧展示满足报价条件的所有商品，可上下划动
			//供货商只能看自己的产品,(发布时间由新到旧展示满足报价条件的所有批发商品)
		}
		if userRole == model.UserRoleCustomer {
			//消费者只能看到上架的的零售商品,默认按照发布时间由新到旧展示,筛选满足条件的商品
			//    2. 综合排序：按货品发布时间由新到旧排列
			//    3. 销量优先：按货品近30天累计支付成功的订单总量降序排列
			//    4. 低价优先：按货品最低价格升序排列（一个货品可以有多个sku、多个价格，这里按其sku中的最低价格排序）
		}
		if userRole == model.UserRoleBuyer {
			//采购商只能看到上架的批发商品,默认按照发布时间由新到旧展示,筛选满足条件的商品,由新到旧排列
		}
		// 2. 综合排序、最新发布：都按需求单提交时间由新到旧排列（本期不做区分）
		// 3. 低价优先：按货品最低价格升序排列（一个货品可以有多个sku、多个价格，这里按其sku中的最低价格排序）
	}
	if req.RequestSource == enum.BackendSource {
		if userRole == model.UserRoleCustomer || userRole == model.UserRoleBuyer {
			//消费者采购商都没有后台
		}
		if userRole == model.UserRoleSupplier {
			//供货商只能看自己的产品
		}

	}
	return nil
}
