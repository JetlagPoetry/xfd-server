package model

import "gorm.io/gorm"

type OrgPointApplication struct {
	gorm.Model
	OrganizationID string                    `gorm:"column:organization_id;not null" json:"organization_id"`
	FileURL        string                    `gorm:"column:file_url;not null" json:"file_url"`
	Status         OrgPointApplicationStatus `gorm:"column:status;not null" json:"status"`
	Comment        string                    `gorm:"column:comment;not null" json:"comment"`
	Deleted        int                       `gorm:"column:deleted" json:"deleted"`
}

func (u *OrgPointApplication) TableName() string {
	return "org_point_application"
}

type OrgPointApplicationStatus int

const (
	OrgPointApplicationStatusUnknown  OrderPurchaseStatus = 0
	OrgPointApplicationStatusApproved OrderPurchaseStatus = 1 // 审核通过
	OrgPointApplicationStatusRejected OrderPurchaseStatus = 2 // 审核失败
)
