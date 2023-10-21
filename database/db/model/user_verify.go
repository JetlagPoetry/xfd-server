package model

import "gorm.io/gorm"

type UserVerify struct {
	gorm.Model
	UserID        string `gorm:"column:user_id;not null" json:"user_id"`
	RealName      string `gorm:"column:real_name;not null" json:"real_name"`
	CertificateNo string `gorm:"column:certificate_no;not null" json:"certificate_no"`
	LicenseURL    string `gorm:"column:license_url;not null" json:"license_url"`
	Position      string `gorm:"column:position;not null" json:"position"`
	Phone         string `gorm:"column:phone;not null" json:"phone"`
	Status        int    `gorm:"column:status;not null" json:"status"`
	Deleted       *int   `gorm:"column:deleted" json:"deleted"`
}

func (UserVerify) TableName() string {
	return "user_verify"
}
