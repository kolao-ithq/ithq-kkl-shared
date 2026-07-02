package kkl_entities

import (
	"time"

	"gorm.io/gorm"
)

type CustomerOtp struct {
	ID           uint `gorm:"primarykey"`
	CreatedAt    time.Time
	UpdatedAt    time.Time
	DeletedAt    gorm.DeletedAt `gorm:"index"`
	OTPNumber    string         `gorm:"column:otp_number;type:varchar(10)"`
	ActiveStatus bool           `gorm:"column:active_status"`
	ExpiredTime  time.Time      `gorm:"column:expired_time"`
	CustomerID   uint           `gorm:"column:customer_id"`
	RefCode      string         `gorm:"column:ref_code;type:longtext"`
}

func (CustomerOtp) TableName() string { return "customer_otps" }
