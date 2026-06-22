package kkl_entities

import "time"

type Bill struct {
	ID            uint64 `gorm:"primaryKey"`
	CreatedAt     time.Time
	UpdatedAt     time.Time
	DeletedAt     *time.Time `gorm:"index:idx_bill_customer_lot_status,priority:4"`
	TicketID      string     `gorm:"unique"`
	TransactionID string
	LotID         uint `gorm:"index:idx_bill_customer_lot_status,priority:1"`
	CreatedBy     uint `gorm:"index:idx_bill_customer_lot_status,priority:2"`
	Amount        uint64
	IsCoreSale    bool
	BillStatusID  uint `gorm:"index:idx_bill_customer_lot_status,priority:3"`
	CustomerPhone string
	RebuyTicket   *string
}

type BillStatus struct {
	InsertOnly
	BillID uint64 `gorm:"index:idx_bill_status_bill_id"`
	Status uint8
	Remark string
}

type BillDigit struct {
	InsertOnly
	BillID        uint64 `gorm:"index:idx_bill_digit_lookup,priority:1"`
	DigitType     uint   `gorm:"index:idx_bill_digit_lookup,priority:2"`
	DigitValue    string `gorm:"index:idx_bill_digit_lookup,priority:3"`
	Amount        uint64
	CoreAmount    uint64
	LottoAmount   uint64
	MissingAmount uint64
}
