package model

import (
	"gorm.io/gorm"
)

type PointRecord struct {
	gorm.Model
	UserID             string            `gorm:"column:user_id;not null" json:"user_id"`
	OrganizationID     int               `gorm:"column:organization_id;not null" json:"organization_id"`
	ChangePoint        int               `gorm:"column:change_point;not null" json:"change_point"`
	PointApplicationID int               `gorm:"column:point_application_id;not null" json:"point_application_id"`
	PointID            int               `gorm:"column:point;not null" json:"point"`
	PaymentID          int               `gorm:"column:payment_id;not null" json:"payment_id"`
	Type               PointRecordType   `gorm:"column:type;not null" json:"type"`
	Status             PointRecordStatus `gorm:"column:status;not null" json:"status"`
	Comment            string            `gorm:"column:comment;not null" json:"comment"`
	OperateUserID      string            `gorm:"column:operate_user_id;not null" json:"operate_user_id"`
	OperateUsername    string            `gorm:"column:operate_username;not null" json:"operate_username"`
	Deleted            int               `gorm:"column:deleted" json:"deleted"`
}

func (u *PointRecord) TableName() string {
	return "point_record"
}

type PointRecordType int

const (
	PointRecordTypeApplication PointRecordType = 0 // 新增
	PointRecordTypeSpend       PointRecordType = 1 // 消费
	PointRecordTypeExpired     PointRecordType = 2 // 过期
	PointRecordTypeQuit        PointRecordType = 3 // 离职
	PointRecordTypeCancel      PointRecordType = 4 // 清零
)

type PointRecordStatus int

const (
	PointRecordStatusInit      PointRecordStatus = 0
	PointRecordStatusConfirmed PointRecordStatus = 1 //
	PointRecordStatusCancelled PointRecordStatus = 2 //
)
