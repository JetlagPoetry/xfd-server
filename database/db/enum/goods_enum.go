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
	ProductAll       ProductVariantType = 0
	ProductWholesale ProductVariantType = 1
	ProductRetail    ProductVariantType = 2
)

func (p ProductVariantType) Code() int {
	switch p {
	case ProductWholesale:
		return 1
	case ProductRetail:
		return 2
	default:
		return 0
	}
}

func (p ProductVariantType) Info() (int, string) {
	switch p {
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
	ProductVariantEnabled  ProductVariantStatus = 1
	ProductVariantDisabled ProductVariantStatus = 2
)

func (p ProductVariantStatus) Code() int {
	switch p {
	case ProductVariantDisabled:
		return 2
	case ProductVariantEnabled:
		return 1
	default:
		return 0
	}
}

func (p ProductVariantStatus) Info() (int, string) {
	switch p {
	case ProductVariantDisabled:
		return 1, "启用"
	case ProductVariantEnabled:
		return 2, "未启用"
	default:
		return 0, "error"
	}
}

type GoodsRetailStatus int

const (
	GoodsRetailNormal  GoodsRetailStatus = 1
	GoodsRetailSoldOut GoodsRetailStatus = 2
)

func (g GoodsRetailStatus) Code() int {
	switch g {
	case GoodsRetailNormal:
		return 1
	case GoodsRetailSoldOut:
		return 2
	default:
		return 0
	}
}

func (g GoodsRetailStatus) Info() (int, string) {
	switch g {
	case GoodsRetailNormal:
		return 1, "正常"
	case GoodsRetailSoldOut:
		return 2, "售磬"
	default:
		return 0, "error"
	}
}

type GoodsListType int

const (
	GoodsListTypeNormal   GoodsListType = 1 // 普通列表，按照创建时间排序
	GoodsListTypeSale     GoodsListType = 2 // 销量列表
	GoodsListTypePrice    GoodsListType = 3 // 价格列表，按照价格排序，一个货品可以有多个sku、多个价格，这里按其sku中的最低价格排序
	GoodsListTypeCategory GoodsListType = 4 // 分类列表，按照分类排序
)

func (g GoodsListType) Code() int {
	switch g {
	case GoodsListTypeNormal:
		return 1
	case GoodsListTypeSale:
		return 2
	case GoodsListTypePrice:
		return 3
	case GoodsListTypeCategory:
		return 4
	default:
		return 0
	}
}

func (g GoodsListType) Info() (int, string) {
	switch g {
	case GoodsListTypeNormal:
		return 1, "普通列表"
	case GoodsListTypeSale:
		return 2, "销量列表"
	case GoodsListTypePrice:
		return 3, "价格列表"
	case GoodsListTypeCategory:
		return 4, "分类列表"
	default:
		return 0, "error"
	}
}
