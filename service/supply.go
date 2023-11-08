package service

import (
	"context"
	"errors"
	"xfd-backend/database/db/dao"
	"xfd-backend/database/db/model"
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
}

func NewSupplyService() *SupplyService {
	return &SupplyService{
		purchaseDao: dao.NewOrderPurchaseDao(),
		quoteDao:    dao.NewOrderQuoteDao(),
		userDao:     dao.NewUserDao(),
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
		list = append(list, &types.PurchaseQuote{
			QuoteID:    int(quote.ID),
			OrderID:    int(purchaseOrder.ID),
			GoodsID:    quote.GoodsID,
			GoodsURL:   "", // todo
			GoodsName:  "",
			Price:      quote.Price,
			Unit:       purchaseOrder.Unit,
			Time:       quote.CreatedAt.Unix(),
			UserID:     userID,
			UserName:   user.Username,
			UserAvatar: user.AvatarURL,
		})
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
		Price:           req.Price,
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

	count, err := s.quoteDao.CountBySupplyUserIDAndNotifySupply(ctx, userID, true)
	if err != nil {
		return nil, xerr.WithCode(xerr.ErrorDatabase, err)
	}

	return &types.SupplyGetStatisticsResp{NewPurchase: int(count)}, nil
}
