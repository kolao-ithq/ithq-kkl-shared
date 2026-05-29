package kkl_entities

import "gorm.io/gorm"

type PointBalance struct {
	gorm.Model
	CustomerID    uint    `gorm:"uniqueIndex;not null"`
	TotalEarned   float64 `gorm:"not null;default:0"`
	TotalRedeemed float64 `gorm:"not null;default:0"`
	Balance       float64 `gorm:"not null;default:0"`
}

type PointProgram struct {
	gorm.Model
	Name        string `gorm:"not null"`
	Description string `gorm:"type:text"`
	IsActive    bool   `gorm:"not null;default:true"`
}

// TransactionType represents the kind of loyalty point transaction.
type TransactionType string

const (
	TransactionTypeEarn    TransactionType = "EARN"
	TransactionTypeRedeem  TransactionType = "REDEEM"
	TransactionTypeReverse TransactionType = "REVERSE"
)

// TransactionStatus represents the lifecycle status of a transaction.
type TransactionStatus string

const (
	TransactionStatusCompleted TransactionStatus = "COMPLETED"
	TransactionStatusReversed  TransactionStatus = "REVERSED"
)

// PointTransaction records every earn, redeem, and reverse event.
type PointTransaction struct {
	gorm.Model
	CustomerID            uint              `gorm:"index;not null"`
	ProgramID             *uint             `gorm:"index"`
	TransactionType       TransactionType   `gorm:"type:varchar(10);not null"`
	Points                float64           `gorm:"not null"` // positive for EARN, negative for REDEEM/REVERSE
	ReferenceID           string            `gorm:"type:varchar(100);index"`
	Description           string            `gorm:"type:text"`
	Status                TransactionStatus `gorm:"type:varchar(10);not null;default:'COMPLETED'"`
	ReversedTransactionID *uint             `gorm:"index"` // points to the original tx when type=REVERSE
}
