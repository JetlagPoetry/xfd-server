package service

import (
	"context"
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

func (s *PurchaseService) GetOrders(ctx context.Context, req *types.PurchaseGetOrdersReq) (*types.PurchaseGetOrdersResp, xerr.XErr) {
	purchaseList, _, err := s.purchaseDao.Lists(ctx, req.PageRequest)
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

	list := make([]*types.PurchaseOrder, 0)
	for _, purchase := range purchaseList {
		user := userMap[purchase.UserID]
		list = append(list, &types.PurchaseOrder{
			OrderID:          int(purchase.ID),
			CategoryID:       purchase.CategoryID,
			CategoryName:     purchase.CategoryName,
			Period:           purchase.Period,
			Quantity:         purchase.Quantity,
			Unit:             purchase.Unit,
			Requirement:      purchase.Requirement,
			AreaCodeID:       purchase.AreaCodeID,
			AreaName:         purchase.AreaName,
			UserID:           purchase.UserID,
			UserName:         user.Username,
			UserAvatar:       user.AvatarURL,
			UserOrganization: user.Organization,
			SubmitTime:       purchase.CreatedAt.Unix(),
		})
	}
	return &types.PurchaseGetOrdersResp{List: list}, nil
}

func (s *PurchaseService) SubmitOrder(ctx context.Context, req *types.PurchaseSubmitOrderReq) (*types.PurchaseSubmitOrderResp, xerr.XErr) {
	userID := common.GetUserID(ctx)

	newPurchase := &model.OrderPurchase{
		UserID:       userID,
		CategoryID:   req.CategoryID,
		CategoryName: "", // todo
		Period:       req.Period,
		Quantity:     req.Quantity,
		Unit:         req.Unit,
		Requirement:  req.Requirement,
		AreaCodeID:   req.AreaCodeID,
		AreaName:     "", //todo
		Status:       model.OrderPurchaseStatusSubmitted,
	}
	err := s.purchaseDao.Create(ctx, newPurchase)
	if err != nil {
		return nil, xerr.WithCode(xerr.ErrorDatabase, err)
	}

	return nil, nil
}

func (s *PurchaseService) ModifyOrder(ctx context.Context, req *types.PurchaseModifyOrderReq) (*types.PurchaseModifyOrderResp, xerr.XErr) {
	_, err := s.purchaseDao.GetByID(ctx, req.OrderID)
	if err != nil {
		return nil, xerr.WithCode(xerr.ErrorDatabase, err)
	}

	updateValue := &model.OrderPurchase{
		CategoryID:   req.CategoryID,
		CategoryName: "", // todo 获取category name
		Period:       req.Period,
		Quantity:     req.Quantity,
		Unit:         req.Unit,
		Requirement:  req.Requirement,
		AreaCodeID:   req.AreaCodeID,
	}
	err = s.purchaseDao.UpdateByID(ctx, req.OrderID, updateValue)
	if err != nil {
		return nil, xerr.WithCode(xerr.ErrorDatabase, err)
	}

	return &types.PurchaseModifyOrderResp{}, nil
}

func (s *PurchaseService) ModifyOrderStatus(ctx context.Context, req *types.PurchaseModifyOrderStatusReq) (*types.PurchaseModifyOrderStatusResp, xerr.XErr) {
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

	// 报价单审核通过
	if order.Status == model.OrderPurchaseStatusSubmitted && req.Status == model.OrderPurchaseStatusOngoing {
		err = s.purchaseDao.UpdateByID(ctx, req.OrderID, &model.OrderPurchase{Status: model.OrderPurchaseStatusOngoing, Comment: req.Comment})
		if err != nil {
			return nil, xerr.WithCode(xerr.ErrorDatabase, err)
		}
		return &types.PurchaseModifyOrderStatusResp{}, nil
	}

	// 报价单审核失败
	if order.Status == model.OrderPurchaseStatusSubmitted && req.Status == model.OrderPurchaseStatusRejected {
		err = s.purchaseDao.UpdateByID(ctx, req.OrderID, &model.OrderPurchase{Status: model.OrderPurchaseStatusRejected, Comment: req.Comment})
		if err != nil {
			return nil, xerr.WithCode(xerr.ErrorDatabase, err)
		}
		return &types.PurchaseModifyOrderStatusResp{}, nil
	}

	// 报价单关闭
	if order.Status == model.OrderPurchaseStatusOngoing && req.Status == model.OrderPurchaseStatusClosed {
		err = s.purchaseDao.UpdateByID(ctx, req.OrderID, &model.OrderPurchase{Status: model.OrderPurchaseStatusClosed})
		if err != nil {
			return nil, xerr.WithCode(xerr.ErrorDatabase, err)
		}
		return &types.PurchaseModifyOrderStatusResp{}, nil
	}

	return nil, nil
}

func (s *PurchaseService) GetQuotes(ctx context.Context, req *types.PurchaseGetQuotesReq) (*types.PurchaseGetQuotesResp, xerr.XErr) {
	userID := common.GetUserID(ctx)
	user, err := s.userDao.GetByUserID(ctx, userID)
	if err != nil {
		return nil, xerr.WithCode(xerr.ErrorDatabase, err)
	}
	purchaseOrder, err := s.purchaseDao.GetByID(ctx, req.OrderID)
	if err != nil {
		return nil, xerr.WithCode(xerr.ErrorDatabase, err)
	}
	quoteList, err := s.quoteDao.ListByUserID(ctx, userID)
	if err != nil {
		return nil, xerr.WithCode(xerr.ErrorDatabase, err)
	}

	list := make([]*types.PurchaseQuote, 0)

	for _, quote := range quoteList {
		list = append(list, &types.PurchaseQuote{
			QuoteID:    int(quote.ID),
			OrderID:    int(purchaseOrder.ID),
			ItemID:     quote.ItemID,
			Price:      quote.Price,
			Unit:       purchaseOrder.Unit,
			Time:       quote.CreatedAt.Unix(),
			UserID:     userID,
			UserName:   user.Username,
			UserAvatar: user.AvatarURL,
		})
	}
	return &types.PurchaseGetQuotesResp{List: list}, nil
}

func (s *PurchaseService) SubmitQuote(ctx context.Context, req *types.PurchaseSubmitQuoteReq) (*types.PurchaseSubmitQuoteResp, xerr.XErr) {
	userID := common.GetUserID(ctx)

	newQuote := &model.OrderQuote{
		PurchaseOrderID: req.OrderID,
		QuoteUserID:     userID,
		ItemID:          req.ItemID,
		Price:           req.Price,
	}
	err := s.quoteDao.Create(ctx, newQuote)
	if err != nil {
		return nil, xerr.WithCode(xerr.ErrorDatabase, err)
	}
	return nil, nil
}
