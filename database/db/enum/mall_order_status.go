package enum

type OrderInfoStatus int

const (
	OrderInfoPaidWaiting OrderInfoStatus = iota + 1 //等待付款
	OrderInfoPaid                                   //交易成功
	OrderInfoClosed                                 //订单关闭
)

func GetOrderInfoStatusEnumByStatus(status int) (OrderInfoStatus, string) {
	switch status {
	case 1:
		return OrderInfoPaidWaiting, "等待付款"
	case 2:
		return OrderInfoPaid, "交易成功"
	case 3:
		return OrderInfoClosed, "订单关闭"
	default:
		return 0, "error"
	}
}

func (g OrderInfoStatus) Code() int {
	switch g {
	case OrderInfoPaidWaiting:
		return 1
	case OrderInfoPaid:
		return 2
	case OrderInfoClosed:
		return 3
	default:
		return 0
	}
}

type OrderProductVariantDetail int

const (
	OrderProductVariantDetailPending   OrderProductVariantDetail = iota + 1 //待付款
	OrderProductVariantDetailPaid                                           //已付款
	OrderProductVariantDetailShipped                                        //待收货
	OrderProductVariantDetailDelivered                                      //确认收货
	OrderProductVariantDetailClosed                                         //订单关闭
	OrderProductVariantDetailRefunded                                       //已退款
	OrderProductVariantDetailReturned                                       //已退货
	OrderProductVariantDetailCompleted                                      //订单完成
)

func GetOrderProductVariantDetailEnumByStatus(status int) (OrderProductVariantDetail, string) {
	switch status {
	case 1:
		return OrderProductVariantDetailPending, "待付款"
	case 2:
		return OrderProductVariantDetailPaid, "已付款"
	case 3:
		return OrderProductVariantDetailShipped, "待收货"
	case 4:
		return OrderProductVariantDetailDelivered, "确认收货"
	case 5:
		return OrderProductVariantDetailClosed, "订单关闭"
	case 6:
		return OrderProductVariantDetailRefunded, "已退款"
	case 7:
		return OrderProductVariantDetailReturned, "已退货"
	case 8:
		return OrderProductVariantDetailCompleted, "订单完成"
	default:
		return 0, "error"
	}
}

func (g OrderProductVariantDetail) Code() int {
	switch g {
	case OrderProductVariantDetailPending:
		return 1
	case OrderProductVariantDetailPaid:
		return 2
	case OrderProductVariantDetailShipped:
		return 3
	case OrderProductVariantDetailDelivered:
		return 4
	case OrderProductVariantDetailClosed:
		return 5
	case OrderProductVariantDetailRefunded:
		return 6
	case OrderProductVariantDetailReturned:
		return 7
	case OrderProductVariantDetailCompleted:
		return 8
	default:
		return 0
	}
}
