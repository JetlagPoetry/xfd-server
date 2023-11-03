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

type GoodsStatus int

const (
	GoodsStatusOnSale  GoodsStatus = 1
	GoodsStatusOffSale GoodsStatus = 2
)

func (g GoodsStatus) Code() int {
	switch g {
	case GoodsStatusOnSale:
		return 1
	case GoodsStatusOffSale:
		return 2
	default:
		return 0
	}
}

func (g GoodsStatus) Info() (int, string) {
	switch g {
	case GoodsStatusOnSale:
		return 1, "在售"
	case GoodsStatusOffSale:
		return 2, "下架"
	default:
		return 0, "error"
	}
}

type ProductVariantType int

const (
	ProductWholesale ProductVariantType = 1
	ProductRetail    ProductVariantType = 2
)

func (g ProductVariantType) Code() int {
	switch g {
	case ProductWholesale:
		return 1
	case ProductRetail:
		return 2
	default:
		return 0
	}
}

func (g ProductVariantType) Info() (int, string) {
	switch g {
	case ProductWholesale:
		return 1, "批发"
	case ProductRetail:
		return 2, "零售"
	default:
		return 0, "error"
	}
}

type ProductVariantStatus int

const (
	ProductVariantDisabled ProductVariantStatus = 1
	ProductVariantEnabled  ProductVariantStatus = 2
)

func (g ProductVariantStatus) Code() int {
	switch g {
	case ProductVariantDisabled:
		return 1
	case ProductVariantEnabled:
		return 2
	default:
		return 0
	}
}

func (g ProductVariantStatus) Info() (int, string) {
	switch g {
	case ProductVariantDisabled:
		return 1, "未启用"
	case ProductVariantEnabled:
		return 2, "启用"
	default:
		return 0, "error"
	}
}

type GoodsRetailStatus int

const (
	GoodsRetailNormal  GoodsRetailStatus = 1
	GoodsRetailSoldOut GoodsRetailStatus = 2
)

func (r GoodsRetailStatus) Code() int {
	switch r {
	case GoodsRetailNormal:
		return 1
	case GoodsRetailSoldOut:
		return 2
	default:
		return 0
	}
}

func (r GoodsRetailStatus) Info() (int, string) {
	switch r {
	case GoodsRetailNormal:
		return 1, "正常"
	case GoodsRetailSoldOut:
		return 2, "售磬"
	default:
		return 0, "error"
	}
}
