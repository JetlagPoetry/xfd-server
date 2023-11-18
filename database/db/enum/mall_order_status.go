package enum

type OrderInfoStatus int

const (
	OrderInfoPendingCreated OrderInfoStatus = iota + 1 //创建订单
	OrderInfoPaidCreated                               //创建交易
	OrderInfoPaidWaiting                               //等待付款
	OrderInfoPaid                                      //交易成功
	OrderInfoSuccess                                   //订单成功
	OrderInfoPaidCancelled                             //取消付款
	OrderInfoPaidFailed                                //交易失败
	OrderInfoClosed                                    //订单关闭
)

func GetOrderInfoStatusEnumByStatus(status int) (OrderInfoStatus, string) {
	switch status {
	case 1:
		return OrderInfoPendingCreated, "创建订单"
	case 2:
		return OrderInfoPaidCreated, "创建交易"
	case 3:
		return OrderInfoPaidWaiting, "等待付款"
	case 4:
		return OrderInfoPaid, "交易成功"
	case 5:
		return OrderInfoSuccess, "订单成功"
	case 6:
		return OrderInfoPaidCancelled, "取消付款"
	case 7:
		return OrderInfoPaidFailed, "交易失败"
	case 8:
		return OrderInfoClosed, "订单关闭"
	default:
		return 0, ""
	}
}

func (g OrderInfoStatus) Code() int {
	switch g {
	case OrderInfoPendingCreated:
		return 1
	case OrderInfoPaidCreated:
		return 2
	case OrderInfoPaidWaiting:
		return 3
	case OrderInfoPaid:
		return 4
	case OrderInfoSuccess:
		return 5
	case OrderInfoPaidCancelled:
		return 6
	case OrderInfoPaidFailed:
		return 7
	case OrderInfoClosed:
		return 8
	default:
		return 0
	}
}

type OrderProductVariantDetail int

const (
	OrderGoodsPending   OrderProductVariantDetail = iota + 1 //待付款
	OrderGoodsPaid                                           //已付款，待发货
	OrderGoodsShipped                                        //待收货
	OrderGoodsDelivered                                      //确认收货
	OrderGoodsCancelled                                      //取消订单 //
	OrderGoodsRefunded                                       //已退款
	OrderGoodsReturned                                       //已退货
	OrderGoodsCompleted                                      //订单完成
)

func GetOrderProductVariantDetailEnumByStatus(status int) (OrderProductVariantDetail, string) {
	switch status {
	case 1:
		return OrderGoodsPending, "待付款"
	case 2:
		return OrderGoodsPaid, "待发货"
	case 3:
		return OrderGoodsShipped, "待收货"
	case 4:
		return OrderGoodsDelivered, "确认收货"
	case 5:
		return OrderGoodsCancelled, "取消订单"
	case 6:
		return OrderGoodsRefunded, "已退款"
	case 7:
		return OrderGoodsReturned, "已退货"
	case 8:
		return OrderGoodsCompleted, "订单完成"
	default:
		return 0, "未知状态"
	}
}

func (g OrderProductVariantDetail) Code() int {
	switch g {
	case OrderGoodsPending:
		return 1
	case OrderGoodsPaid:
		return 2
	case OrderGoodsShipped:
		return 3
	case OrderGoodsDelivered:
		return 4
	case OrderGoodsCancelled:
		return 5
	case OrderGoodsRefunded:
		return 6
	case OrderGoodsReturned:
		return 7
	case OrderGoodsCompleted:
		return 8
	default:
		return 0
	}
}
