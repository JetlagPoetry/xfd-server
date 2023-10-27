package model

import "gorm.io/gorm"

type OrderQuote struct {
	gorm.Model
	PurchaseOrderID int     `gorm:"column:purchase_order_id;not null" json:"purchase_order_id"`
	QuoteUserID     string  `gorm:"column:quote_user_id;not null" json:"quote_user_id"`
	ItemID          int     `gorm:"column:item_id;not null" json:"item_id"`
	Price           float64 `gorm:"column:price;not null" json:"price"`
	Deleted         int     `gorm:"column:deleted" json:"deleted"`
}

func (u *OrderQuote) TableName() string {
	return "order_quote"
}
