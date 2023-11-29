package service

import (
	"context"
	"errors"
	"fmt"
	"sort"
	"xfd-backend/database/db/dao"
	"xfd-backend/database/db/model"
	"xfd-backend/database/redis"
	"xfd-backend/pkg/common"
	"xfd-backend/pkg/types"
	"xfd-backend/pkg/utils"
	"xfd-backend/pkg/xerr"
	"xfd-backend/service/cache"
)

type PurchaseService struct {
	purchaseDao *dao.OrderPurchaseDao
	quoteDao    *dao.OrderQuoteDao
	userDao     *dao.UserDao
	goodsDao    *dao.GoodsDao
}

func NewPurchaseService() *PurchaseService {
	return &PurchaseService{
		purchaseDao: dao.NewOrderPurchaseDao(),
		quoteDao:    dao.NewOrderQuoteDao(),
		userDao:     dao.NewUserDao(),
		goodsDao:    dao.NewGoodsDao(),
	}
}

func (s *PurchaseService) GetPurchases(ctx context.Context, req types.PurchaseGetPurchasesReq) (*types.PurchaseGetPurchasesResp, xerr.XErr) {
	userID := common.GetUserID(ctx)

	user, err := s.userDao.GetByUserID(ctx, userID)
	if err != nil {
		return nil, xerr.WithCode(xerr.ErrorDatabase, err)
	}
	if user.UserRole != model.UserRoleBuyer {
		return nil, xerr.WithCode(xerr.ErrorOperationForbidden, errors.New("user is not buyer"))
	}

	purchaseList, count, err := s.purchaseDao.ListByUser(ctx, req.PageRequest, userID, req.Status)
	if err != nil {
		return nil, xerr.WithCode(xerr.ErrorDatabase, err)
	}

	list := make([]*types.PurchaseOrder, 0)
	for _, purchase := range purchaseList {
		// todo 通过redis优化
		totalQuote, err := s.quoteDao.CountByOrderID(ctx, int(purchase.ID))
		if err != nil {
			return nil, xerr.WithCode(xerr.ErrorDatabase, err)
		}
		newQuote, err := s.quoteDao.CountByOrderIDAndNotifyPurchase(ctx, int(purchase.ID), true)
		if err != nil {
			return nil, xerr.WithCode(xerr.ErrorDatabase, err)
		}
		list = append(list, &types.PurchaseOrder{
			OrderID:      int(purchase.ID),
			CategoryName: purchase.CategoryName,
			Period:       purchase.Period,
			Quantity:     purchase.Quantity,
			Unit:         purchase.Unit,
			Requirement:  purchase.Requirement,
			UserID:       purchase.UserID,
			SubmitTime:   purchase.CreatedAt.Unix(),
			NewQuote:     int(newQuote),
			TotalQuote:   int(totalQuote),
			Status:       purchase.Status,
		})
	}
	// 有新增在无新增前，时间由近到远
	sort.Slice(list, func(i, j int) bool {
		if list[i].NewQuote == 0 && list[j].NewQuote > 0 {
			return false
		} else if list[i].NewQuote > 0 && list[j].NewQuote == 0 {
			return true
		}
		return list[i].SubmitTime > list[j].SubmitTime
	})
	return &types.PurchaseGetPurchasesResp{List: list, TotalNum: count}, nil
}

func (s *PurchaseService) SubmitOrder(ctx context.Context, req types.PurchaseSubmitOrderReq) (*types.PurchaseSubmitOrderResp, xerr.XErr) {
	userID := common.GetUserID(ctx)

	user, err := s.userDao.GetByUserID(ctx, userID)
	if err != nil {
		return nil, xerr.WithCode(xerr.ErrorDatabase, err)
	}
	if user.UserRole != model.UserRoleBuyer {
		return nil, xerr.WithCode(xerr.ErrorOperationForbidden, errors.New("user is not buyer"))
	}

	category := cache.GetCategory()
	if category[int32(req.CategoryB)] == nil {
		return nil, xerr.WithCode(xerr.ErrorNotExistRecord, errors.New("category not found"))
	}
	orderName := category[int32(req.CategoryB)].Name
	if req.CategoryC > 0 {
		if category[int32(req.CategoryC)] == nil {
			return nil, xerr.WithCode(xerr.ErrorNotExistRecord, errors.New("category not found"))
		}
		orderName += " - " + category[int32(req.CategoryC)].Name
	}

	newPurchase := &model.OrderPurchase{
		UserID:       userID,
		CategoryA:    req.CategoryA,
		CategoryB:    req.CategoryB,
		CategoryC:    req.CategoryC,
		CategoryName: orderName,
		Period:       req.Period,
		Quantity:     req.Quantity,
		Unit:         req.Unit,
		Requirement:  req.Requirement,
		Status:       model.OrderPurchaseStatusSubmitted,
	}
	err = s.purchaseDao.Create(ctx, newPurchase)
	if err != nil {
		return nil, xerr.WithCode(xerr.ErrorDatabase, err)
	}

	return nil, nil
}

//func (s *PurchaseService) ModifyOrder(ctx context.Context, req types.PurchaseModifyOrderReq) (*types.PurchaseModifyOrderResp, xerr.XErr) {
//	_, err := s.purchaseDao.GetByID(ctx, req.OrderID)
//	if err != nil {
//		return nil, xerr.WithCode(xerr.ErrorDatabase, err)
//	}
//
//	updateValue := &model.OrderPurchase{
//		CategoryID:   req.CategoryID,
//		CategoryName: "", //
//		Period:       req.Period,
//		Quantity:     req.Quantity,
//		Unit:         req.Unit,
//		Requirement:  req.Requirement,
//		AreaCodeID:   req.AreaCodeID,
//	}
//	err = s.purchaseDao.UpdateByID(ctx, req.OrderID, updateValue)
//	if err != nil {
//		return nil, xerr.WithCode(xerr.ErrorDatabase, err)
//	}
//
//	return &types.PurchaseModifyOrderResp{}, nil
//}

func (s *PurchaseService) ModifyOrderStatus(ctx context.Context, req types.PurchaseModifyOrderStatusReq) (*types.PurchaseModifyOrderStatusResp, xerr.XErr) {
	userID := common.GetUserID(ctx)
	user, err := s.userDao.GetByUserID(ctx, userID)
	if err != nil {
		return nil, xerr.WithCode(xerr.ErrorDatabase, err)
	}
	if user.UserRole != model.UserRoleBuyer {
		return nil, xerr.WithCode(xerr.ErrorOperationForbidden, errors.New("user is not buyer"))
	}

	order, err := s.purchaseDao.GetByID(ctx, req.OrderID)
	if err != nil {
		return nil, xerr.WithCode(xerr.ErrorDatabase, err)
	}

	if order.UserID != userID {
		return nil, xerr.WithCode(xerr.ErrorOperationForbidden, errors.New("not belong to user"))
	}

	// 删除报价单
	if req.Delete {
		err = s.purchaseDao.Delete(ctx, req.OrderID)
		if err != nil {
			return nil, xerr.WithCode(xerr.ErrorDatabase, err)
		}
		return &types.PurchaseModifyOrderStatusResp{}, nil
	}

	//// 报价单审核通过
	//if order.Status == model.OrderPurchaseStatusSubmitted && req.Status == model.OrderPurchaseStatusOngoing {
	//	err = s.purchaseDao.UpdateByID(ctx, req.OrderID, &model.OrderPurchase{Status: model.OrderPurchaseStatusOngoing, Comment: req.Comment})
	//	if err != nil {
	//		return nil, xerr.WithCode(xerr.ErrorDatabase, err)
	//	}
	//	return &types.PurchaseModifyOrderStatusResp{}, nil
	//}
	//
	//// 报价单审核失败
	//if order.Status == model.OrderPurchaseStatusSubmitted && req.Status == model.OrderPurchaseStatusRejected {
	//	err = s.purchaseDao.UpdateByID(ctx, req.OrderID, &model.OrderPurchase{Status: model.OrderPurchaseStatusRejected, Comment: req.Comment})
	//	if err != nil {
	//		return nil, xerr.WithCode(xerr.ErrorDatabase, err)
	//	}
	//	return &types.PurchaseModifyOrderStatusResp{}, nil
	//}

	// 报价单关闭
	if order.Status == model.OrderPurchaseStatusSubmitted && req.Status == model.OrderPurchaseStatusClosed {
		err = s.purchaseDao.UpdateByID(ctx, req.OrderID, &model.OrderPurchase{Status: model.OrderPurchaseStatusClosed})
		if err != nil {
			return nil, xerr.WithCode(xerr.ErrorDatabase, err)
		}
		return &types.PurchaseModifyOrderStatusResp{}, nil
	}

	return nil, nil
}

func (s *PurchaseService) GetQuotes(ctx context.Context, req types.PurchaseGetQuotesReq) (*types.PurchaseGetQuotesResp, xerr.XErr) {
	userID := common.GetUserID(ctx)
	user, err := s.userDao.GetByUserID(ctx, userID)
	if err != nil {
		return nil, xerr.WithCode(xerr.ErrorDatabase, err)
	}
	if user.UserRole != model.UserRoleBuyer {
		return nil, xerr.WithCode(xerr.ErrorOperationForbidden, errors.New("user is not buyer"))
	}

	purchaseOrder, err := s.purchaseDao.GetByID(ctx, req.OrderID)
	if err != nil {
		return nil, xerr.WithCode(xerr.ErrorDatabase, err)
	}
	if purchaseOrder.UserID != userID {
		return nil, xerr.WithCode(xerr.ErrorOperationForbidden, errors.New("not belong to user"))
	}
	quoteList, count, err := s.quoteDao.ListByOrderID(ctx, req.OrderID, req.PageRequest)
	if err != nil {
		return nil, xerr.WithCode(xerr.ErrorDatabase, err)
	}
	quoteUserIDs := make([]string, 0)
	for _, quote := range quoteList {
		quoteUserIDs = append(quoteUserIDs, quote.QuoteUserID)
	}
	userList, err := s.userDao.ListByUserIDs(ctx, quoteUserIDs)
	if err != nil {
		return nil, xerr.WithCode(xerr.ErrorDatabase, err)
	}
	userMap := make(map[string]*model.User)
	for _, user := range userList {
		userMap[user.UserID] = user
	}

	list := make([]*types.PurchaseQuote, 0)
	for _, quote := range quoteList {
		user := userMap[quote.QuoteUserID]

		goods, err := s.goodsDao.GetGoodsByGoodsID(ctx, int32(quote.GoodsID))
		if err != nil || goods == nil {
			continue
		}
		list = append(list, &types.PurchaseQuote{
			QuoteID:    int(quote.ID),
			OrderID:    int(purchaseOrder.ID),
			GoodsID:    quote.GoodsID,
			GoodsName:  goods.Name,
			GoodsURL:   goods.GoodsFrontImage,
			Price:      quote.Price.Round(2).String(),
			Unit:       purchaseOrder.Unit,
			Time:       quote.CreatedAt.Unix(),
			UserID:     userID,
			UserName:   user.Username,
			UserAvatar: user.AvatarURL,
			IsNew:      *quote.NotifyPurchase,
		})
	}

	// 修改报价单状态为已读
	go s.quoteDao.UpdateByOrderID(ctx, req.OrderID, &model.OrderQuote{NotifyPurchase: utils.BoolPtr(false)})

	return &types.PurchaseGetQuotesResp{List: list, TotalNum: count}, nil
}

func (s *PurchaseService) GetStatistics(ctx context.Context, req types.PurchaseGetStatisticsReq) (*types.PurchaseGetStatisticsResp, xerr.XErr) {
	userID := common.GetUserID(ctx)
	user, err := s.userDao.GetByUserID(ctx, userID)
	if err != nil {
		return nil, xerr.WithCode(xerr.ErrorDatabase, err)
	}
	if user.UserRole != model.UserRoleBuyer {
		return nil, xerr.WithCode(xerr.ErrorOperationForbidden, errors.New("user is not buyer"))
	}

	count, err := s.quoteDao.CountByPurchaseUserIDAndNotifyPurchase(ctx, userID, true)
	if err != nil {
		return nil, xerr.WithCode(xerr.ErrorDatabase, err)
	}

	return &types.PurchaseGetStatisticsResp{
		NewQuote: int(count),
	}, nil
}

func (s *PurchaseService) NotifySupply(ctx context.Context, req types.PurchaseNotifySupplyReq) (*types.PurchaseNotifySupplyResp, xerr.XErr) {
	userID := common.GetUserID(ctx)
	user, err := s.userDao.GetByUserID(ctx, userID)
	if err != nil {
		return nil, xerr.WithCode(xerr.ErrorDatabase, err)
	}
	if user.UserRole != model.UserRoleBuyer {
		return nil, xerr.WithCode(xerr.ErrorOperationForbidden, errors.New("user is not buyer"))
	}

	supplier, err := s.userDao.GetByUserID(ctx, req.SupplyUserID)
	if err != nil {
		return nil, xerr.WithCode(xerr.ErrorDatabase, err)
	}
	if supplier.UserRole != model.UserRoleSupplier {
		return nil, xerr.WithCode(xerr.ErrorOperationForbidden, errors.New("user is not supplier"))
	}

	// 设置redis hash，保存当前用户的提示数
	if err := redis.RedisClient.HIncrBy(fmt.Sprintf(redis.SUPPLY_NOTIFY_WITH_PURCHASE, supplier.UserID), user.UserID, int64(req.Count)).Err(); err != nil {
		return nil, xerr.WithCode(xerr.ErrorRedis, err)
	}

	// 设置redis hash，增加供应商的提示数字
	if err := redis.RedisClient.IncrBy(fmt.Sprintf(redis.SUPPLY_NOTIFY_NUMBER, supplier.UserID), int64(req.Count)).Err(); err != nil {
		return nil, xerr.WithCode(xerr.ErrorRedis, err)
	}

	return &types.PurchaseNotifySupplyResp{}, nil
}

func (s *PurchaseService) AnswerQuote(ctx context.Context, req types.PurchaseAnswerQuoteReq) (*types.PurchaseAnswerQuoteResp, xerr.XErr) {
	userID := common.GetUserID(ctx)
	user, err := s.userDao.GetByUserID(ctx, userID)
	if err != nil {
		return nil, xerr.WithCode(xerr.ErrorDatabase, err)
	}
	if user.UserRole != model.UserRoleBuyer {
		return nil, xerr.WithCode(xerr.ErrorOperationForbidden, errors.New("user is not buyer"))
	}

	err = s.quoteDao.UpdateByPurchaseUserAndQuoteUser(ctx, userID, req.SupplyUserID, &model.OrderQuote{NotifySupply: utils.BoolPtr(false)})
	if err != nil {
		return nil, xerr.WithCode(xerr.ErrorDatabase, err)
	}

	return &types.PurchaseAnswerQuoteResp{}, nil
}
