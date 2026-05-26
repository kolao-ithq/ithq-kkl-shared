package kkl_entities

import "gorm.io/gorm"

type CoreCustomer struct {
	gorm.Model
	CustomerID uint `gorm:"unique"`
}
