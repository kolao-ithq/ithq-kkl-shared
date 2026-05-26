package kkl_entities

type Bill struct {
	InsertOnly
	TicketID      string `gorm:"unique"`
	TransactionID string `gorm:"index"`
	LotID         uint   `gorm:"index"`
	CreatedBy     uint   `gorm:"index"`
	Amount        uint64
	IsCoreSale    bool
	BillStatus    int8
}

type BillStatus struct {
	InsertOnly
	BillID uint `gorm:"index:idx_bill_status_bill_id"`
	Status uint8
	Remark string
}

type BillDigit struct {
	InsertOnly
	BillID     uint `gorm:"index:idx_bill_digit_bill_id"`
	DigitType  uint
	DigitValue string
	Amount     uint64
	CoreAmount uint64
}
