package model

import "gorm.io/gorm"

type OrderQuote struct {
	gorm.Model
	PurchaseOrderID int     `gorm:"column:purchase_order_id;not null" json:"purchase_order_id"`
	PurchaseUserID  string  `gorm:"column:purchase_user_id;not null" json:"purchase_user_id"`
	QuoteUserID     string  `gorm:"column:quote_user_id;not null" json:"quote_user_id"`
	GoodsID         int     `gorm:"column:goods_id;not null" json:"goods_id"`
	Price           float64 `gorm:"column:price;not null" json:"price"`
	NotifySupply    bool    `gorm:"column:notify_supply;not null" json:"notify_supply"`
	NotifyPurchase  bool    `gorm:"column:notify_purchase;not null" json:"notify_purchase"`
}

func (u *OrderQuote) TableName() string {
	return "order_quote"
}
