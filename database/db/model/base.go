package model

import (
	"database/sql/driver"
	"encoding/json"
	"time"
)

type GormList []string

func (g *GormList) Value() (driver.Value, error) {
	return json.Marshal(g)
}

// Scan 实现 sql.Scanner 接口，Scan 将 value 扫描至 Jsonb
func (g *GormList) Scan(value interface{}) error {
	return json.Unmarshal(value.([]byte), &g)
}

type BaseModel struct {
	ID        int32     `gorm:"primary_key;AUTO_INCREMENT;type:int" json:"id"`
	CreatedAt time.Time `gorm:"comment:创建时间;not null;column:created_at;default:CURRENT_TIMESTAMP(3);" json:"-"`
	UpdatedAt time.Time `gorm:"comment:更新时间;not null;column:updated_at;default:CURRENT_TIMESTAMP(3) ON UPDATE CURRENT_TIMESTAMP(3);" json:"-"`
	Deleted   *int      `gorm:"type:tinyint(1);not null;default:0;column:deleted;index:deleted" json:"-"`
}

type GoodsDetail struct {
	Goods            int32
	ProductVariantID int32
	Num              int32
}

// GoodsDetailList 自定义gorm类型
type GoodsDetailList []GoodsDetail

func (g *GoodsDetailList) Value() (driver.Value, error) {
	return json.Marshal(g)
}

func (g *GoodsDetailList) Scan(value interface{}) error {
	return json.Unmarshal(value.([]byte), &g)
}
