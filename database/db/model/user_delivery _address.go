package model

import "gorm.io/gorm"

// UserDeliveryAddress 用户收货地址管理
type UserDeliveryAddress struct {
	gorm.Model
	UserId       string `gorm:"comment:用户ID;type:varchar(100);index:idx_user_id;default:'';not null"`
	AddressPhone string `gorm:"comment:联系方式;type:varchar(100);default:'';not null"`
	Name         string `gorm:"comment:收货人姓名;type:varchar(100);default:'';not null"`
	Area         string `gorm:"comment:地区;type:varchar(255);default:'';not null"`
	FullAddress  string `gorm:"comment:详细地址;type:varchar(500);default:'';not null"`
	AddressType  int32  `gorm:"comment:地址类型;type:tinyint(1);default:0;not null;comment:地址类型;"`
	DisplayOrder int32  `gorm:"comment:排序;type:int(1);not null;default:0;column:display_order"`
}

func (u UserDeliveryAddress) TableName() string {
	return "user_delivery_address"
}
