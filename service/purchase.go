package service

import (
	"context"
	"xfd-backend/database/db/dao"
	"xfd-backend/pkg/types"
	"xfd-backend/pkg/xerr"
)

type PurchaseService struct {
	purchaseDao *dao.OrderPurchaseDao
}

func NewPurchaseService() *PurchaseService {
	return &PurchaseService{
		purchaseDao: dao.NewOrderPurchaseDao(),
	}
}

func (s *PurchaseService) GetOrders(ctx context.Context, req *types.PurchaseGetOrdersReq) (*types.PurchaseGetOrdersResp, xerr.XErr) {
	// todo 获取采购单

	return nil, nil
}

func (s *PurchaseService) SubmitOrder(ctx context.Context, req *types.PurchaseSubmitOrderReq) (*types.PurchaseSubmitOrderResp, xerr.XErr) {
	// todo 提交采购单

	return nil, nil
}

func (s *PurchaseService) ModifyOrder(ctx context.Context, req *types.PurchaseModifyOrderReq) (*types.PurchaseModifyOrderResp, xerr.XErr) {
	// todo 更新采购单

	return nil, nil
}

func (s *PurchaseService) ModifyOrderStatus(ctx context.Context, req *types.PurchaseModifyOrderStatusReq) (*types.PurchaseModifyOrderStatusResp, xerr.XErr) {
	// todo 更新采购单状态

	return nil, nil
}
