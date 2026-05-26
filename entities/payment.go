package entities

import "gorm.io/gorm"

type Payment struct {
	InsertOnly
	BillID                uint `gorm:"uniqueIndex"`
	BankReference         string
	PaymentMethodID       uint `gorm:"index:idx_payment_method_id"`
	PaymentMethod         PaymentMethod
	Amount                uint64
	Discount              uint64
	ActualAmount          uint64
	BankCode              string
	AccountType           string
	BankAccountNumber     string
	CurrentCommissionRate float64
}

type PaymentMethod struct {
	gorm.Model
	Name           string
	CommissionRate float64
	IsActive       bool
}
