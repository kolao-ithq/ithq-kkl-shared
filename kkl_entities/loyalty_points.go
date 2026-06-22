package kkl_entities

import "gorm.io/gorm"

type PointBalance struct {
	gorm.Model
	CustomerID    uint  `gorm:"column:customer_id;uniqueIndex;not null"` // cross-service ref, no FK constraint
	TotalEarned   int64 `gorm:"not null;default:0"`
	TotalRedeemed int64 `gorm:"not null;default:0"`
	Balance       int64 `gorm:"not null;default:0"`
}

type PointProgram struct {
	gorm.Model
	Name        string `gorm:"type:varchar(255);not null"`
	Description string `gorm:"type:text"`
	IsActive    bool   `gorm:"not null;default:true"`
}

type TransactionType string

const (
	TransactionTypeEarn    TransactionType = "EARN"
	TransactionTypeRedeem  TransactionType = "REDEEM"
	TransactionTypeReverse TransactionType = "REVERSE"
)

type TransactionStatus string

const (
	TransactionStatusCompleted TransactionStatus = "COMPLETED"
	TransactionStatusReversed  TransactionStatus = "REVERSED"
)

type PointTransaction struct {
	gorm.Model
	CustomerID            uint              `gorm:"column:customer_id;index;not null"`
	ProgramID             *uint             `gorm:"column:program_id;uniqueIndex:idx_point_tx_ref_unique"`
	TransactionType       TransactionType   `gorm:"type:varchar(20);not null;uniqueIndex:idx_point_tx_ref_unique"`
	Points                int64             `gorm:"not null"`
	ReferenceID           string            `gorm:"type:varchar(100);uniqueIndex:idx_point_tx_ref_unique"`
	Description           string            `gorm:"type:text"`
	Status                TransactionStatus `gorm:"type:varchar(20);not null;default:'COMPLETED'"`
	ReversedTransactionID *uint             `gorm:"column:reversed_transaction_id;index"`
}