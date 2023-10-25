package model

import "gorm.io/gorm"

type OrderPurchase struct {
	gorm.Model
	UserID      string              `gorm:"column:user_id;not null" json:"user_id"`
	CategoryID  int                 `gorm:"column:category_id;not null" json:"category_id"` // todo 可以是二级id或三级id
	Period      PurchasePeriod      `gorm:"column:period;not null" json:"period"`
	Quantity    int                 `gorm:"column:quantity;not null" json:"quantity"`
	Unit        string              `gorm:"column:unit;not null" json:"unit"`
	Requirement string              `gorm:"column:requirement" json:"requirement"`
	Location    string              `gorm:"column:location" json:"location"` // todo
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
	OrderPurchaseStatusUnknown   PurchasePeriod = 0
	OrderPurchaseStatusSubmitted PurchasePeriod = 1 // 已提交
	OrderPurchaseStatusOngoing   PurchasePeriod = 2 // 已通过
	OrderPurchaseStatusClosed    PurchasePeriod = 3 // 已结束
)

func (u *OrderPurchase) TableName() string {
	return "order_purchase"
}
