package kkl_entities

import "gorm.io/gorm"

type CoreCustomer struct {
	gorm.Model
	CustomerID int64  `gorm:"column:customer_id;type:bigint(20)"`
	Customer   string `gorm:"column:customer;type:varchar(50)"`
}

func (CoreCustomer) TableName() string { return "core_customers" }