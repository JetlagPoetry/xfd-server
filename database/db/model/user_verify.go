package model

import "gorm.io/gorm"

type UserVerify struct {
	gorm.Model
	UserID           string           `gorm:"column:user_id;not null" json:"user_id"`
	Organization     string           `gorm:"column:organization;not null" json:"organization"`
	OrganizationCode string           `gorm:"column:organization_code;not null" json:"organization_code"`
	OrganizationURL  string           `gorm:"column:organization_url;not null" json:"organization_url"`
	CorporationURLA  string           `gorm:"column:corporation_url_a;not null" json:"corporation_url_a"`
	CorporationURLB  string           `gorm:"column:corporation_url_b;not null" json:"corporation_url_b"`
	RealName         string           `gorm:"column:real_name;not null" json:"real_name"`
	CertificateNo    string           `gorm:"column:certificate_no;not null" json:"certificate_no"`
	Position         string           `gorm:"column:position;not null" json:"position"`
	Phone            string           `gorm:"column:phone;not null" json:"phone"`
	Status           UserVerifyStatus `gorm:"column:status;not null" json:"status"`
	Comment          string           `gorm:"column:comment;not null" json:"comment"`
	Deleted          *int             `gorm:"column:deleted" json:"deleted"`
}

type UserVerifyStatus int32

const (
	UserVerifyStatusUnknown   UserVerifyStatus = 0
	UserVerifyStatusSubmitted UserVerifyStatus = 1
	UserVerifyStatusRejected  UserVerifyStatus = 2
	UserVerifyStatusSuccess   UserVerifyStatus = 3
)

func (UserVerify) TableName() string {
	return "user_verify"
}