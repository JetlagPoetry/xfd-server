package enum

type GoodsCategoryLevel int32

const (
	Default    GoodsCategoryLevel = 0
	LevelOne   GoodsCategoryLevel = 1
	LevelTwo   GoodsCategoryLevel = 2
	LevelThree GoodsCategoryLevel = 3
)

func (g GoodsCategoryLevel) Info() (int, string) {
	switch g {
	case LevelOne:
		return 1, "一级分类"
	case LevelTwo:
		return 2, "二级分类"
	case LevelThree:
		return 3, "三级分类"
	default:
		return 0, "error"
	}
}

func (g GoodsCategoryLevel) Code() int {
	switch g {
	case LevelOne:
		return 1
	case LevelTwo:
		return 2
	case LevelThree:
		return 3
	default:
		return 0
	}
}

type RetailDeliveryTime int

const (
	RetailDeliveryTimeWithin24Hours RetailDeliveryTime = iota + 1
	RetailDeliveryTimeWithin48Hours
	RetailDeliveryTimeWithin7Days
)

func (r RetailDeliveryTime) Info() (int, string) {
	switch r {
	case RetailDeliveryTimeWithin24Hours:
		return 1, "24小时内"
	case RetailDeliveryTimeWithin48Hours:
		return 2, "48小时内"
	case RetailDeliveryTimeWithin7Days:
		return 3, "7天内"
	default:
		return 0, "未知时间"
	}
}

func (r RetailDeliveryTime) Code() int {
	switch r {
	case RetailDeliveryTimeWithin24Hours:
		return 1
	case RetailDeliveryTimeWithin48Hours:
		return 2
	case RetailDeliveryTimeWithin7Days:
		return 3
	default:
		return 0
	}
}
