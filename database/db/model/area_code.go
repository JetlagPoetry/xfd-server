package model

import "time"

type AreaCode struct {
	Code      int       `gorm:"comment:区划代码;type:bigint(12);not null;column:code;default:0;index:code;index:code_status"`
	Name      string    `gorm:"comment:名称;type:varchar(128);default:'';not null;column:name;index:name;"`
	Level     int       `gorm:"comment:级别1-5,省市县镇村;type:tinyint(1);default:0;not null;column:level;index:level;"`
	Pcode     int       `gorm:"comment:父级区划代码;not null;type:bigint(12);default:0;column:pcode;index:pcode;"`
	Status    int       `gorm:"comment:地址是否有效1:有效,0:无效;type:tinyint(1);not null;default:1;column:status;index:code_status"`
	ID        int       `gorm:"primary_key;AUTO_INCREMENT;not null;column:id"`
	CreatedAt time.Time `gorm:"column:created_at;default:CURRENT_TIMESTAMP(3);comment:创建时间"`
	UpdatedAt time.Time `gorm:"column:updated_at;default:CURRENT_TIMESTAMP(3) ON UPDATE CURRENT_TIMESTAMP(3);comment:更新时间"`
}

func (u *AreaCode) TableName() string {
	return "area_code"
}
