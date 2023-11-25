package enum

type OrderInfoStatus int

const (
	OrderInfoCreated                 OrderInfoStatus = iota + 1 //订单创建
	OrderInfoPaidWaiting                                        //交易创建，待微信支付
	OderInfoPaidSuccess                                         //交易成功，待发货
	OderInfoShipped                                             //已发货，待收货
	OderInfoReceived                                            //已收货，已签收，确认收货
	OderInfoAfterSale                                           //售后/结束
	OderInfoClosed                                              //交易关闭
	OrderInfoPaidPointConfirmWaiting                            // 交易成功，待实际抵扣积分
)

func GetOrderInfoStatusEnumByStatus(status int) (OrderInfoStatus, string) {
	switch status {
	case 1:
		return OrderInfoCreated, "订单创建"
	case 2:
		return OrderInfoPaidWaiting, "交易创建"
	case 3:
		return OderInfoPaidSuccess, "交易成功，待发货"
	case 4:
		return OderInfoShipped, "已发货"
	case 5:
		return OderInfoReceived, "已收货，已签收，确认收货"
	case 6:
		return OderInfoAfterSale, "售后/结束,换"
	case 7:
		return OderInfoClosed, "交易关闭"
	case 8:
		return OrderInfoPaidPointConfirmWaiting, "交易成功，待发货"
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
	case OderInfoShipped:
		return 4
	case OderInfoReceived:
		return 5
	case OderInfoAfterSale:
		return 6
	case OderInfoClosed:
		return 7
	case OrderInfoPaidPointConfirmWaiting:
		return 8
	default:
		return 0
	}
}

type OrderProductVariantDetail int

const (
	OrderProductVariantDetailCreated   OrderProductVariantDetail = iota + 1 //订单创建
	OrderProductVariantDetailPending                                        //待付款
	OrderProductVariantDetailPaid                                           //已付款
	OrderProductVariantDetailShipped                                        //待收货
	OrderProductVariantDetailDelivered                                      //已签收
	OrderProductVariantDetailAfterSale                                      //售后/结束
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
		return OrderProductVariantDetailAfterSale, "售后/结束"
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
	case OrderProductVariantDetailAfterSale:
		return 6
	default:
		return 0
	}
}

type QueryOrderStatus int

const (
	QueryOrderStatusWaitingShipped QueryOrderStatus = iota + 1 //待发货
	QueryOrderStatusShipped                                    //待收货
	QueryOrderStatusDelivered                                  //已签收
	QueryOrderStatusAfterSale                                  //售后/结束
)

func GetQueryOrderStatusEnumByStatus(status int) (QueryOrderStatus, string) {
	switch status {
	case 1:
		return QueryOrderStatusWaitingShipped, "待发货"
	case 2:
		return QueryOrderStatusShipped, "待收货"
	case 3:
		return QueryOrderStatusDelivered, "已签收"
	case 4:
		return QueryOrderStatusAfterSale, "售后/结束"
	default:
		return QueryOrderStatusWaitingShipped, "待发货"
	}
}

func (g QueryOrderStatus) Code() int {
	switch g {
	case QueryOrderStatusWaitingShipped:
		return 1
	case QueryOrderStatusShipped:
		return 2
	case QueryOrderStatusDelivered:
		return 3
	case QueryOrderStatusAfterSale:
		return 4
	default:
		return 0
	}
}

type AfterSaleType int

const (
	AfterSaleTypeExchange AfterSaleType = iota + 1
	AfterSaleTypeReturnAndRefund
)

func GetAfterSaleType(status int) (AfterSaleType, string) {
	switch status {
	case 1:
		return AfterSaleTypeExchange, "换货"
	case 2:
		return AfterSaleTypeReturnAndRefund, "退货退款"
	default:
		return AfterSaleTypeExchange, "换货"
	}
}

func (g AfterSaleType) Code() int {
	switch g {
	case AfterSaleTypeExchange:
		return 1
	case AfterSaleTypeReturnAndRefund:
		return 2
	default:
		return 0
	}
}

type ReturnPointType int

const (
	ReturnPointNo ReturnPointType = iota + 1
	ReturnPointYes
)

type ReturnPointStatus int

// 不需要返回积分/待返回积分/已返回积分
const (
	ReturnPointStatusNoReturn ReturnPointStatus = iota
	ReturnPointStatusWaitReturn
	ReturnPointStatusReturned
)

type ManuallyClosed int

const (
	ManuallyClosedNo ManuallyClosed = iota + 1
	ManuallyClosedYes
)
