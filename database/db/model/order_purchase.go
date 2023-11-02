package model

import "gorm.io/gorm"

type OrderPurchase struct {
	gorm.Model
	UserID      string              `gorm:"column:user_id;not null" json:"user_id"`
	CategoryA   int                 `gorm:"column:category_a;not null" json:"category_a"`
	CategoryB   int                 `gorm:"column:category_b;not null" json:"category_b"`
	CategoryC   int                 `gorm:"column:category_c;not null" json:"category_c"`
	Period      PurchasePeriod      `gorm:"column:period;not null" json:"period"`
	Quantity    int                 `gorm:"column:quantity;not null" json:"quantity"`
	Unit        string              `gorm:"column:unit;not null" json:"unit"`
	Requirement string              `gorm:"column:requirement" json:"requirement"`
	AreaCodeID  int                 `gorm:"column:area_code_id" json:"area_code_id"` // todo
	Status      OrderPurchaseStatus `gorm:"column:status" json:"status"`
	Deleted     int                 `gorm:"column:deleted" json:"deleted"`
}

type PurchasePeriod int

const (
	PurchasePeriodUnknown PurchasePeriod = 0
	PurchasePeriodBatch   PurchasePeriod = 1
	PurchasePeriodDay     PurchasePeriod = 2
	PurchasePeriodWeek    PurchasePeriod = 3
	PurchasePeriodMonth   PurchasePeriod = 4
)

type OrderPurchaseStatus int

const (
	OrderPurchaseStatusUnknown   OrderPurchaseStatus = 0
	OrderPurchaseStatusSubmitted OrderPurchaseStatus = 1 // 已发布
	OrderPurchaseStatusClosed    OrderPurchaseStatus = 2 // 已结束
	//OrderPurchaseStatusOngoing   OrderPurchaseStatus = 2 // 审核通过
	//OrderPurchaseStatusRejected  OrderPurchaseStatus = 4 // 审核失败
)

func (u *OrderPurchase) TableName() string {
	return "order_purchase"
}
