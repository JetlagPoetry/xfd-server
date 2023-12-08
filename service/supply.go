package service

import (
	"context"
	"errors"
	"fmt"
	goredis "github.com/go-redis/redis"
	"strconv"
	"xfd-backend/database/db/dao"
	"xfd-backend/database/db/model"
	"xfd-backend/database/redis"
	"xfd-backend/pkg/common"
	"xfd-backend/pkg/types"
	"xfd-backend/pkg/utils"
	"xfd-backend/pkg/xerr"
	"xfd-backend/service/cache"
)

type SupplyService struct {
	purchaseDao *dao.OrderPurchaseDao
	quoteDao    *dao.OrderQuoteDao
	userDao     *dao.UserDao
	orderDao    *dao.OrderDao
	goodsDao    *dao.GoodsDao
}

func NewSupplyService() *SupplyService {
	return &SupplyService{
		purchaseDao: dao.NewOrderPurchaseDao(),
		quoteDao:    dao.NewOrderQuoteDao(),
		userDao:     dao.NewUserDao(),
		orderDao:    dao.NewOrderDao(),
		goodsDao:    dao.NewGoodsDao(),
	}
}

func (s *SupplyService) GetPurchases(ctx context.Context, req types.SupplyGetPurchasesReq) (*types.SupplyGetPurchasesResp, xerr.XErr) {
	userID := common.GetUserID(ctx)
	user, err := s.userDao.GetByUserID(ctx, userID)
	if err != nil {
		return nil, xerr.WithCode(xerr.ErrorUserNotFound, err)
	}
	if user.UserRole != model.UserRoleSupplier {
		return nil, xerr.WithCode(xerr.ErrorOperationForbidden, errors.New("user is not supplier"))
	}

	purchaseList, count, err := s.purchaseDao.List(ctx, req.PageRequest, req.CategoryA, req.CategoryB, req.CategoryC, req.Like)
	if err != nil {
		return nil, xerr.WithCode(xerr.ErrorDatabase, err)
	}

	userMap := make(map[string]*model.User)
	userIDList := make([]string, 0)
	for _, purchase := range purchaseList {
		userMap[purchase.UserID] = &model.User{}
	}
	for id := range userMap {
		userIDList = append(userIDList, id)
	}
	userList, err := s.userDao.ListByUserIDs(ctx, userIDList)
	if err != nil {
		return nil, xerr.WithCode(xerr.ErrorDatabase, err)
	}

	for _, user := range userList {
		userMap[user.UserID] = user
	}

	purchaseIDList := make([]int, 0)
	for _, purchase := range purchaseList {
		purchaseIDList = append(purchaseIDList, int(purchase.ID))
	}

	quoteList, err := s.quoteDao.ListByQuoteUserIDAndOrderIDs(ctx, userID, purchaseIDList)
	if err != nil {
		return nil, xerr.WithCode(xerr.ErrorDatabase, err)
	}

	quoteMap := make(map[int]*model.OrderQuote)
	for _, quote := range quoteList {
		quoteMap[int(quote.ID)] = quote
	}

	category := cache.GetCategory()

	list := make([]*types.PurchaseOrder, 0)
	for _, purchase := range purchaseList {
		user := userMap[purchase.UserID]
		list = append(list, &types.PurchaseOrder{
			OrderID:       int(purchase.ID),
			CategoryNameA: category[int32(purchase.CategoryA)].Name,
			CategoryNameB: category[int32(purchase.CategoryB)].Name,
			CategoryNameC: category[int32(purchase.CategoryC)].Name,
			CategoryA:     purchase.CategoryA,
			CategoryB:     purchase.CategoryB,
			CategoryC:     purchase.CategoryC,
			CategoryName:  purchase.CategoryName,
			Period:        purchase.Period,
			Quantity:      purchase.Quantity,
			Unit:          purchase.Unit,
			Requirement:   purchase.Requirement,
			UserID:        purchase.UserID,
			UserName:      user.Username,
			UserAvatar:    user.AvatarURL,
			SubmitTime:    purchase.CreatedAt.Unix(),
			HasQuote:      quoteMap[int(purchase.ID)] != nil,
		})
	}
	return &types.SupplyGetPurchasesResp{List: list, TotalNum: count}, nil
}

func (s *SupplyService) GetQuotes(ctx context.Context, req types.SupplyGetQuotesReq) (*types.SupplyGetQuotesResp, xerr.XErr) {
	userID := common.GetUserID(ctx)
	user, err := s.userDao.GetByUserID(ctx, userID)
	if err != nil {
		return nil, xerr.WithCode(xerr.ErrorUserNotFound, err)
	}
	if user.UserRole != model.UserRoleSupplier {
		return nil, xerr.WithCode(xerr.ErrorOperationForbidden, errors.New("user is not supplier"))
	}

	purchaseOrder, err := s.purchaseDao.GetByID(ctx, req.OrderID)
	if err != nil {
		return nil, xerr.WithCode(xerr.ErrorDatabase, err)
	}
	quoteList, err := s.quoteDao.ListByUserIDAndOrderID(ctx, userID, req.OrderID)
	if err != nil {
		return nil, xerr.WithCode(xerr.ErrorDatabase, err)
	}

	list := make([]*types.PurchaseQuote, 0)
	for _, quote := range quoteList {

		newQuote := &types.PurchaseQuote{
			QuoteID: int(quote.ID),
			OrderID: int(purchaseOrder.ID),
			GoodsID: quote.GoodsID,

			Price:      quote.Price.Round(2).String(),
			Unit:       purchaseOrder.Unit,
			Time:       quote.CreatedAt.Unix(),
			UserID:     userID,
			UserName:   user.Username,
			UserAvatar: user.AvatarURL,
		}
		goods, err := s.goodsDao.GetGoodsByGoodsID(ctx, int32(quote.GoodsID))
		if err == nil && goods != nil {
			newQuote.GoodsURL = goods.GoodsFrontImage
			newQuote.GoodsName = goods.Name
		}
		list = append(list, newQuote)
	}
	return &types.SupplyGetQuotesResp{List: list}, nil
}

func (s *SupplyService) SubmitQuote(ctx context.Context, req types.SupplySubmitQuoteReq) (*types.SupplySubmitQuoteResp, xerr.XErr) {
	userID := common.GetUserID(ctx)
	user, err := s.userDao.GetByUserID(ctx, userID)
	if err != nil {
		return nil, xerr.WithCode(xerr.ErrorUserNotFound, err)
	}
	if user.UserRole != model.UserRoleSupplier {
		return nil, xerr.WithCode(xerr.ErrorOperationForbidden, errors.New("user is not supplier"))
	}

	purchaseOrder, err := s.purchaseDao.GetByID(ctx, req.OrderID)
	if err != nil {
		return nil, xerr.WithCode(xerr.ErrorDatabase, err)
	}

	newQuote := &model.OrderQuote{
		PurchaseOrderID: req.OrderID,
		PurchaseUserID:  purchaseOrder.UserID,
		QuoteUserID:     userID,
		GoodsID:         req.GoodsID,
		Price:           utils.StringToDecimal(req.Price),
		NotifySupply:    utils.BoolPtr(true),
		NotifyPurchase:  utils.BoolPtr(true),
	}
	err = s.quoteDao.Create(ctx, newQuote)
	if err != nil {
		return nil, xerr.WithCode(xerr.ErrorDatabase, err)
	}
	return &types.SupplySubmitQuoteResp{}, nil
}

func (s *SupplyService) GetQuotedPurchases(ctx context.Context, req types.SupplyGetQuotedPurchasesReq) (*types.SupplyGetQuotedPurchasesResp, xerr.XErr) {
	userID := common.GetUserID(ctx)
	user, err := s.userDao.GetByUserID(ctx, userID)
	if err != nil {
		return nil, xerr.WithCode(xerr.ErrorUserNotFound, err)
	}
	if user.UserRole != model.UserRoleSupplier {
		return nil, xerr.WithCode(xerr.ErrorOperationForbidden, errors.New("user is not supplier"))
	}

	quoteList, err := s.quoteDao.ListByQuoteUserID(ctx, userID)
	if err != nil {
		return nil, xerr.WithCode(xerr.ErrorDatabase, err)
	}
	orderIDs := make([]int, 0)
	for _, quote := range quoteList {
		orderIDs = append(orderIDs, quote.PurchaseOrderID)
	}

	purchaseList, count, err := s.purchaseDao.ListByOrderIDs(ctx, req.PageRequest, orderIDs, req.Status)
	if err != nil {
		return nil, xerr.WithCode(xerr.ErrorDatabase, err)
	}

	list := make([]*types.PurchaseOrder, 0)
	for _, purchase := range purchaseList {
		list = append(list, &types.PurchaseOrder{
			OrderID:      int(purchase.ID),
			CategoryName: purchase.CategoryName,
			Period:       purchase.Period,
			Quantity:     purchase.Quantity,
			Unit:         purchase.Unit,
			Requirement:  purchase.Requirement,
			UserID:       purchase.UserID,
			SubmitTime:   purchase.CreatedAt.Unix(),
		})
	}
	return &types.SupplyGetQuotedPurchasesResp{List: list, TotalNum: count}, nil
}

func (s *SupplyService) GetStatistics(ctx context.Context, req types.SupplyGetStatisticsReq) (*types.SupplyGetStatisticsResp, xerr.XErr) {
	userID := common.GetUserID(ctx)
	user, err := s.userDao.GetByUserID(ctx, userID)
	if err != nil {
		return nil, xerr.WithCode(xerr.ErrorUserNotFound, err)
	}
	if user.UserRole != model.UserRoleSupplier {
		return nil, xerr.WithCode(xerr.ErrorOperationForbidden, errors.New("user is not supplier"))
	}

	//count, err := s.quoteDao.CountBySupplyUserIDAndNotifySupply(ctx, userID, true)
	//if err != nil {
	//	return nil, xerr.WithCode(xerr.ErrorDatabase, err)
	//}
	result, err := redis.RedisClient.Get(fmt.Sprintf(redis.SUPPLY_NOTIFY_NUMBER, user.UserID)).Result()
	if err != goredis.Nil && err != nil {
		return nil, xerr.WithCode(xerr.ErrorRedis, err)
	}
	newPurchase, _ := strconv.Atoi(result)

	//查看零售待发货的订单数量
	count, err := s.orderDao.CountWaitingForDeliveryOrderBySupplyUserID(ctx, userID)
	if err != nil {
		return nil, xerr.WithCode(xerr.ErrorDatabase, err)
	}

	return &types.SupplyGetStatisticsResp{NewPurchase: newPurchase, NewWaitingForDelivery: int(count)}, nil
}

func (s *SupplyService) AnswerQuote(ctx context.Context, req types.SupplyAnswerQuoteReq) (*types.SupplyAnswerQuoteResp, xerr.XErr) {
	userID := common.GetUserID(ctx)
	user, err := s.userDao.GetByUserID(ctx, userID)
	if err != nil {
		return nil, xerr.WithCode(xerr.ErrorDatabase, err)
	}
	if user.UserRole != model.UserRoleSupplier {
		return nil, xerr.WithCode(xerr.ErrorOperationForbidden, errors.New("user is not supplier"))
	}

	purchaser, err := s.userDao.GetByUserID(ctx, req.PurchaseUserID)
	if err != nil {
		return nil, xerr.WithCode(xerr.ErrorDatabase, err)
	}
	if purchaser.UserRole != model.UserRoleBuyer {
		return nil, xerr.WithCode(xerr.ErrorOperationForbidden, errors.New("user is not buyer"))
	}

	value := 0
	result, err := redis.RedisClient.HGet(fmt.Sprintf(redis.SUPPLY_NOTIFY_WITH_PURCHASE, user.UserID), purchaser.UserID).Result()
	if err != goredis.Nil && err != nil {
		return nil, xerr.WithCode(xerr.ErrorRedis, err)
	}
	value, _ = strconv.Atoi(result)
	if err := redis.RedisClient.HSet(fmt.Sprintf(redis.SUPPLY_NOTIFY_WITH_PURCHASE, user.UserID), purchaser.UserID, 0).Err(); err != nil {
		return nil, xerr.WithCode(xerr.ErrorRedis, err)
	}
	// 设置redis hash，减少供应商的提示数字
	if err := redis.RedisClient.DecrBy(fmt.Sprintf(redis.SUPPLY_NOTIFY_NUMBER, user.UserID), int64(value)).Err(); err != nil {
		return nil, xerr.WithCode(xerr.ErrorRedis, err)
	}

	// 触发该supplier的数字校准
	go s.notifyCalibration(user.UserID)

	return &types.SupplyAnswerQuoteResp{}, nil
}

func (s *SupplyService) notifyCalibration(supplyUserID string) error {
	result, err := redis.RedisClient.HGetAll(fmt.Sprintf(redis.SUPPLY_NOTIFY_WITH_PURCHASE, supplyUserID)).Result()
	if err != nil {
		return err
	}
	total := 0
	for _, v := range result {
		num, _ := strconv.Atoi(v)
		total += num
	}
	_, err = redis.RedisClient.Set(fmt.Sprintf(redis.SUPPLY_NOTIFY_NUMBER, supplyUserID), total, 0).Result()
	if err != nil {
		return err
	}
	return nil
}
