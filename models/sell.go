package models

import "time"

type SellRequest struct {
	Phone               string             `json:"phone"`
	DaysOfInstallmentID uint               `json:"doi_id"`
	LottoDrawsNumber    []LottoDrawsNumber `json:"lotto_draw_number"`
	CustomerID          int                `json:"customer_id"`
	MuteLotto           []struct {
		ID uint `json:"id"`
	} `json:"mute_lotto"`
	PaymentRequest PaymentRequest `json:"payment_request"`
}

type LottoDrawsNumber struct {
	Digit string `json:"digit"`
	Money int    `json:"money"`
}

type PaymentRequest struct {
	PaymentMethodID uint   `json:"payment_method_id"`
	DrawID          int    `json:"draw_id"`
	PaymentLinkID   uint   `json:"payment_link_id"`
	TotalSale       int    `json:"total_sale"`
	Phone           string `json:"phone"`
	Transaction     string `json:"transaction"`
	BillNumber      string `json:"bill_number"`
	BillID          uint   `json:"bill_id"`
	Account         string `json:"customer_account"`
	AccountType     string `json:"account_type"`
	CustomerID      uint   `json:"customer_id"`
	FromApp         string `json:"from_app"`
	UsePoint        bool   `json:"use_point"`
}

type LottoBillHistoryResponse struct {
	ID             uint      `json:"id"`
	ExpiredTime    time.Time `json:"expire_time"`
	BillNumber     string    `json:"bill_number"`
	RequestID      string    `json:"request_id"`
	SecurityCode   string    `json:"security_code"`
	BuyTypeID      uint      `json:"buy_type_id"`
	TotalAmount    int       `json:"total_amount"`
	BuyTime        string    `json:"buy_time"`
	BillStatusName string    `json:"bill_status_name"`

	TblDaysOfInstallment TblDaysOfInstallmentBill `json:"tbl_days_of_installment"`
	LottoDrawsNumber     []LottoDrawsNumberBill   `json:"lotto_draw_number"`
	PaymentHistory       PaymentHistoryResponse   `json:"payment_history"`
	RatioV1              interface{}              `json:"ratio_v1"`
	RatioV2              interface{}              `json:"ratio_v2"`
}

type TblDaysOfInstallmentBill struct {
	ID            uint   `json:"id"`
	IntName       string `json:"int_name"`
	StartDatetime string `json:"start_datetime"`
	EndDatetime   string `json:"end_datetime"`
	DrawNumber    *int   `json:"draw_number"`
}

type LottoDrawsNumberBill struct {
	ID              uint   `json:"id"`
	QuotaID         uint   `json:"quota_id"`
	Digit           string `json:"digit"`
	Money           int    `json:"money"`
	CreatedBy       uint   `json:"created_by"`
	UpdateBy        uint   `json:"update_by"`
	Status          bool   `json:"status"`
	SideStatus      bool   `json:"side_status"`
	AlternateStatus bool   `json:"alternate_status"`
}

type PaymentHistoryResponse struct {
	ID                  uint                      `json:"id"`
	Amount              int                       `json:"amount"`
	Description         string                    `json:"description"`
	Currency            string                    `json:"currency"`
	CurrencyCode        string                    `json:"currency_code"`
	PayStatusID         string                    `json:"pay_status_id"`
	PayStatusName       string                    `json:"pay_status_name"`
	TransactionTime     string                    `json:"transaction_time"`
	CustomerName        string                    `json:"customer_name"`
	CustomerPhone       string                    `json:"customer_phone"`
	CustomerAccount     string                    `json:"customer_account"`
	PaymentMethod       PaymentMethodReponse      `json:"payment_method"`
	CustomerAccountType string                    `json:"customer_account_type"`
	Created_by          uint                      `json:"created_by"`
	Updated_by          uint                      `json:"updated_by"`
	QrCode              *string                   `json:"qr_code"`
	DeepLink            *string                   `json:"deep_link"`
	OtpRef              *PyamentConfirmOTPRequest `json:"otp_ref"`
}

type PaymentMethodReponse struct {
	ID                  uint   `json:"id"`
	PaymentMethodName   string `json:"payment_method_name"`
	PaymentMethodLogo   string `json:"payment_method_logo"`
	PaymentMethodStatus bool   `json:"payment_method_status"`
}

type PyamentConfirmOTPRequest struct {
	BillID          uint   `json:"bill_id"`
	TransID         string `json:"trans_id"`
	TransCashOutID  string `json:"transCashOutID"`
	PaymentMethodID uint   `json:"payment_method_id"`
	Amount          int64  `json:"amount"`
	Otp             string `json:"otp"`
	OtpRef          string `json:"otp_ref"`
	OtpRefNo        string `json:"otp_ref_no"`
}
