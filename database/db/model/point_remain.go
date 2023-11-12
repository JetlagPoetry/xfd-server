package model

import "gorm.io/gorm"

type PointRemain struct {
	gorm.Model
	UserID             string   `gorm:"column:user_id;not null" json:"user_id"`
	OrganizationID     int      `gorm:"column:organization_id;not null" json:"organization_id"`
	PointApplicationID int      `gorm:"column:point_application_id;not null" json:"point_application_id"`
	Point              *float64 `gorm:"column:point;not null" json:"point"`
	PointRemain        *float64 `gorm:"column:point_remain;not null" json:"point_remain"`
}
