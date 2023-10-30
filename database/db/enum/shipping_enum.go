package enum

type WholesaleLogistics int

const (
	WholesaleLogisticsWholeVehicle WholesaleLogistics = iota + 1
	WholesaleLogisticsSpecialLine
	WholesaleLogisticsExpress
	WholesaleLogisticsAir
	WholesaleLogisticsRailway
	WholesaleLogisticsOther
)

func (w WholesaleLogistics) Info() (int, string) {
	switch w {
	case WholesaleLogisticsWholeVehicle:
		return 1, "整车"
	case WholesaleLogisticsSpecialLine:
		return 2, "物流/专线"
	case WholesaleLogisticsExpress:
		return 3, "快递"
	case WholesaleLogisticsAir:
		return 4, "空运"
	case WholesaleLogisticsRailway:
		return 5, "铁路"
	case WholesaleLogisticsOther:
		return 6, "其他运输方式"
	default:
		return 0, "未知类型"
	}
}

func (w WholesaleLogistics) Code() int {
	switch w {
	case WholesaleLogisticsWholeVehicle:
		return 1
	case WholesaleLogisticsSpecialLine:
		return 2
	case WholesaleLogisticsExpress:
		return 3
	case WholesaleLogisticsAir:
		return 4
	case WholesaleLogisticsRailway:
		return 5
	case WholesaleLogisticsOther:
		return 6
	default:
		return 0
	}
}
