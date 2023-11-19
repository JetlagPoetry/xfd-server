package enum

type OrderInfoStatus int

const (
	OrderInfoCreated     OrderInfoStatus = iota + 1 //订单创建
	OrderInfoPaidWaiting                            //交易创建 todo 待补充
	OderInfoPaidSuccess                             //交易成功
	OrderRefundAndReturn                            //退货退款
	OrderRefund                                     //退款
	OrderReturn                                     //退货
	OrderInfoClosed                                 //订单关闭
)

func GetOrderInfoStatusEnumByStatus(status int) (OrderInfoStatus, string) {
	switch status {
	case 1:
		return OrderInfoCreated, "订单创建"
	case 2:
		return OrderInfoPaidWaiting, "交易创建"
	case 3:
		return OderInfoPaidSuccess, "交易成功"
	case 4:
		return OrderRefundAndReturn, "退货退款"
	case 5:
		return OrderRefund, "退款"
	case 6:
		return OrderReturn, "退货"
	case 7:
		return OrderInfoClosed, "订单关闭"
	default:
		return OrderInfoCreated, "订单创建"
	}
}

func (g OrderInfoStatus) Code() int {
	switch g {
	case OrderInfoCreated:
		return 1
	case OrderInfoPaidWaiting:
		return 2
	case OderInfoPaidSuccess:
		return 3
	case OrderRefundAndReturn:
		return 4
	case OrderRefund:
		return 5
	case OrderReturn:
		return 6
	case OrderInfoClosed:
		return 7
	default:
		return 0
	}
}

type OrderProductVariantDetail int

const (
	OrderProductVariantDetailCreated         OrderProductVariantDetail = iota + 1 //订单创建
	OrderProductVariantDetailPending                                              //待付款
	OrderProductVariantDetailPaid                                                 //已付款
	OrderProductVariantDetailShipped                                              //待收货
	OrderProductVariantDetailDelivered                                            //确认收货
	OrderProductVariantDetailClosed                                               //订单关闭
	OrderProductVariantDetailRefund                                               //退款
	OrderProductVariantDetailReturn                                               //退货
	OrderProductVariantDetailRefundAndReturn                                      //退货退款
	OrderProductVariantDetailCompleted                                            //订单完成
)

func GetOrderProductVariantDetailEnumByStatus(status int) (OrderProductVariantDetail, string) {
	switch status {
	case 1:
		return OrderProductVariantDetailCreated, "订单创建"
	case 2:
		return OrderProductVariantDetailPending, "待付款"
	case 3:
		return OrderProductVariantDetailPaid, "已付款"
	case 4:
		return OrderProductVariantDetailShipped, "待收货"
	case 5:
		return OrderProductVariantDetailDelivered, "确认收货"
	case 6:
		return OrderProductVariantDetailClosed, "订单关闭"
	case 7:
		return OrderProductVariantDetailRefund, "退款"
	case 8:
		return OrderProductVariantDetailReturn, "退货"
	case 9:
		return OrderProductVariantDetailRefundAndReturn, "退货退款"
	case 10:
		return OrderProductVariantDetailCompleted, "订单完成"
	default:
		return OrderProductVariantDetailCreated, "订单创建"
	}
}

func (g OrderProductVariantDetail) Code() int {
	switch g {
	case OrderProductVariantDetailCreated:
		return 1
	case OrderProductVariantDetailPending:
		return 2
	case OrderProductVariantDetailPaid:
		return 3
	case OrderProductVariantDetailShipped:
		return 4
	case OrderProductVariantDetailDelivered:
		return 5
	case OrderProductVariantDetailClosed:
		return 6
	case OrderProductVariantDetailRefund:
		return 7
	case OrderProductVariantDetailReturn:
		return 8
	case OrderProductVariantDetailRefundAndReturn:
		return 9
	case OrderProductVariantDetailCompleted:
		return 10
	default:
		return 0
	}
}
