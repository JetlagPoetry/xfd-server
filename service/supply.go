package service

import (
	"context"
	"xfd-backend/database/db/dao"
	"xfd-backend/database/db/model"
	"xfd-backend/pkg/common"
	"xfd-backend/pkg/types"
	"xfd-backend/pkg/xerr"
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

func (s *SupplyService) GetPurchases(ctx context.Context, req *types.SupplyGetPurchasesReq) (*types.SupplyGetPurchasesResp, xerr.XErr) {
	userID := common.GetUserID(ctx)
	purchaseList, count, err := s.purchaseDao.ListByUser(ctx, req.PageRequest, userID, req.Status)
	if err != nil {
		return nil, xerr.WithCode(xerr.ErrorDatabase, err)
	}

	list := make([]*types.PurchaseOrder, 0)
	for _, purchase := range purchaseList {
		// todo 查找品类名

		list = append(list, &types.PurchaseOrder{
			OrderID:       int(purchase.ID),
			CategoryNameA: "", // todo
			CategoryNameB: "", // todo
			CategoryNameC: "", // todo
			Period:        purchase.Period,
			Quantity:      purchase.Quantity,
			Unit:          purchase.Unit,
			Requirement:   purchase.Requirement,
			AreaCodeID:    purchase.AreaCodeID,
			UserID:        purchase.UserID,
			SubmitTime:    purchase.CreatedAt.Unix(),
			NewQuote:      0, // todo
		})
	}
	return &types.SupplyGetPurchasesResp{List: list, TotalNum: count}, nil
}

func (s *SupplyService) GetQuotes(ctx context.Context, req *types.SupplyGetQuotesReq) (*types.SupplyGetQuotesResp, xerr.XErr) {
	userID := common.GetUserID(ctx)
	purchaseOrder, err := s.purchaseDao.GetByID(ctx, req.OrderID)
	if err != nil {
		return nil, xerr.WithCode(xerr.ErrorDatabase, err)
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
		list = append(list, &types.PurchaseQuote{
			QuoteID:          int(quote.ID),
			OrderID:          int(purchaseOrder.ID),
			ItemID:           quote.ItemID,
			Price:            quote.Price,
			Unit:             purchaseOrder.Unit,
			Time:             quote.CreatedAt.Unix(),
			UserID:           userID,
			UserName:         user.Username,
			UserAvatar:       user.AvatarURL,
			UserOrganization: user.Organization,
			IsNew:            false, // todo
		})
	}
	return &types.SupplyGetQuotesResp{List: list, TotalNum: count}, nil
}
