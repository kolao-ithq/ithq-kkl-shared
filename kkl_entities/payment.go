package kkl_entities

import "gorm.io/gorm"

type Payment struct {
	InsertOnly
	BillID                uint64 `gorm:"uniqueIndex"`
	PaymentMethodID       uint   `gorm:"index:idx_payment_method_id"`
	PaymentMethod         PaymentMethod
	Amount                uint64
	Discount              uint64
	ActualAmount          uint64
	BankCode              string
	AccountType           string
	BankAccountNumber     string
	CurrentCommissionRate float64
	PointTransactionID    uint
	LotID         string    `gorm:"type:varchar(25)"`
}

type PaymentMethod struct {
	gorm.Model
	Name           string
	CommissionRate float64
	IsActive       bool
}
