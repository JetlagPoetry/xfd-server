package model

import "gorm.io/gorm"

type OrgPointApply struct {
	gorm.Model
	OrganizationID string `gorm:"column:organization_id;not null" json:"organization_id"`
	Point          int    `gorm:"column:point;not null" json:"point"`
	Deleted        int    `gorm:"column:deleted" json:"deleted"`
}

func (u *OrgPointApply) TableName() string {
	return "org_point_apply"
}
