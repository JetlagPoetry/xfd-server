package enum

type OrderGoodsStatus int

const (
	OrderGoodsPending   OrderGoodsStatus = iota + 1 //待付款
	OrderGoodsPaid                                  //已付款，待发货//
	OrderGoodsShipped                               //待收货
	OrderGoodsDelivered                             //确认收货
	OrderGoodsCancelled                             //取消订单 //
	OrderGoodsRefunded                              //已退款
	OrderGoodsReturned                              //已退货
	OrderGoodsCompleted                             //订单完成
)

func GetOrderStatusEnumByStatus(status int) (OrderGoodsStatus, string) {
	switch status {
	case 1:
		return OrderGoodsPending, "待付款"
	case 2:
		return OrderGoodsPaid, "已付款，待发货"
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

func (g OrderGoodsStatus) Code() int {
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
