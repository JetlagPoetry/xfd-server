package service

import (
	"context"
	"sort"
	"xfd-backend/database/db/dao"
	"xfd-backend/database/db/model"
	"xfd-backend/pkg/common"
	"xfd-backend/pkg/types"
	"xfd-backend/pkg/xerr"
)

type PurchaseService struct {
	purchaseDao *dao.OrderPurchaseDao
	quoteDao    *dao.OrderQuoteDao
	userDao     *dao.UserDao
}

func NewPurchaseService() *PurchaseService {
	return &PurchaseService{
		purchaseDao: dao.NewOrderPurchaseDao(),
		quoteDao:    dao.NewOrderQuoteDao(),
		userDao:     dao.NewUserDao(),
	}
}

func (s *PurchaseService) GetPurchases(ctx context.Context, req types.PurchaseGetPurchasesReq) (*types.PurchaseGetPurchasesResp, xerr.XErr) {
	userID := common.GetUserID(ctx)
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

	newPurchase := &model.OrderPurchase{
		UserID:       userID,
		CategoryA:    req.CategoryA,
		CategoryB:    req.CategoryB,
		CategoryC:    req.CategoryC,
		CategoryName: req.CategoryName,
		Period:       req.Period,
		Quantity:     req.Quantity,
		Unit:         req.Unit,
		Requirement:  req.Requirement,
		Status:       model.OrderPurchaseStatusSubmitted,
	}
	err := s.purchaseDao.Create(ctx, newPurchase)
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
	order, err := s.purchaseDao.GetByID(ctx, req.OrderID)
	if err != nil {
		return nil, xerr.WithCode(xerr.ErrorDatabase, err)
	}

	// 删除报价单
	if req.Delete && order.Deleted == 0 {
		err = s.purchaseDao.UpdateByID(ctx, req.OrderID, &model.OrderPurchase{Deleted: 1})
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
	purchaseOrder, err := s.purchaseDao.GetByID(ctx, req.OrderID)
	if err != nil {
		return nil, xerr.WithCode(xerr.ErrorDatabase, err)
	}
	if purchaseOrder.UserID != userID {
		return nil, xerr.WithCode(xerr.ErrorOperationForbidden, err)
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
		// todo 获取商品名、商品url
		list = append(list, &types.PurchaseQuote{
			QuoteID:    int(quote.ID),
			OrderID:    int(purchaseOrder.ID),
			ItemID:     quote.ItemID,
			ItemName:   "",
			ItemURL:    "",
			Price:      quote.Price,
			Unit:       purchaseOrder.Unit,
			Time:       quote.CreatedAt.Unix(),
			UserID:     userID,
			UserName:   user.Username,
			UserAvatar: user.AvatarURL,
			IsNew:      quote.NotifySupply,
		})
	}

	// 修改报价单状态为已读
	go s.quoteDao.UpdateByOrderID(ctx, req.OrderID, &model.OrderQuote{NotifyPurchase: false})

	return &types.PurchaseGetQuotesResp{List: list, TotalNum: count}, nil
}

func (s *PurchaseService) GetStatistics(ctx context.Context, req types.PurchaseGetStatisticsReq) (*types.PurchaseGetStatisticsResp, xerr.XErr) {
	userID := common.GetUserID(ctx)

	count, err := s.quoteDao.CountByPurchaseUserIDAndNotifyPurchase(ctx, userID, true)
	if err != nil {
		return nil, xerr.WithCode(xerr.ErrorDatabase, err)
	}

	return &types.PurchaseGetStatisticsResp{
		NewQuote: int(count),
	}, nil
}

func (s *PurchaseService) AnswerQuote(ctx context.Context, req types.PurchaseAnswerQuoteReq) (*types.PurchaseAnswerQuoteResp, xerr.XErr) {
	userID := common.GetUserID(ctx)

	err := s.quoteDao.UpdateByPurchaseUserAndQuoteUser(ctx, userID, req.SupplyUserID, &model.OrderQuote{NotifySupply: false})
	if err != nil {
		return nil, xerr.WithCode(xerr.ErrorDatabase, err)
	}

	return &types.PurchaseAnswerQuoteResp{}, nil
}
