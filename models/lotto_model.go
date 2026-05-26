package models

import "time"

type LottoSignInBodyV2 struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type LottoSignInResponseV2 struct {
	AccessToken  string                    `json:"accessToken"`
	RefreshToken string                    `json:"refreshToken"`
	ExpiresOn    string                    `json:"expiresOn"`
	CurrentDraw  LottoSignInResponseDrawV2 `json:"currentDraw"`
	DealerInfo   string                    `json:"dealerInfo"`
	ProductName  string                    `json:"productName"`
}

type LottoSignInResponseDrawV2 struct {
	DrawID      int    `json:"drawid"`
	DrawDate    string `json:"draw_date"`
	CloseTime   string `json:"close_time"`
	Result1t    string `json:"result_1t"`
	Result2t    string `json:"result_2t"`
	Result3t    string `json:"result_3t"`
	Result4t    string `json:"result_4t"`
	Result5t    string `json:"result_5t"`
	Result6t    string `json:"result_6t"`
	DrawStatus  string `json:"draw_status"`
	DrawNumber  string `json:"drawnumber"`
	ProductType int    `json:"producttype"`
}

type LottoSignInDataResponse struct {
	DigitType      string `json:"DigitType"`
	LOTCLOSETIME   string `json:"LOTCLOSETIME"`
	LOTDATE        string `json:"LOTDATE"`
	LOTID          int64  `json:"LOTID"`
	LOTSTATUS      string `json:"LOT_STATUS"`
	STATUSID       int64  `json:"STATUSID"`
	SessionExpired string `json:"SessionExpired"`
	TOKEN          string `json:"TOKEN"`
}

type SellTicketItemRequestV2 struct {
	Digit  string `json:"digit"`
	Amount int    `json:"amount"`
}

type SellTicketRequestV2 struct {
	FromMember  string                    `json:"frommember"`
	Time        time.Time                 `json:"time"`
	ClientRefNo string                    `json:"clientrefno"`
	CustomerID  string                    `json:"customerid"`
	Items       []SellTicketItemRequestV2 `json:"items"`
}

type SellTicketItemResponseV2 struct {
	Digit         string `json:"digit"`
	RequestAmount int    `json:"requestAmount"`
	SaleAmount    int    `json:"saleAmount"`
	Status        int    `json:"status"`
	DigitType     int    `json:"digittype"`
}

type SellTicketResponseV2 struct {
	BillNumber    *string                    `json:"billNumber"`
	DrawId        *int                       `json:"drawId"`
	TotalSale     *int                       `json:"totalSale"`
	OrderDate     *string                    `json:"orderDate"`
	DrawDate      *string                    `json:"drawDate"`
	DrawNumber    *string                    `json:"drawNumber"`
	DealerId      *string                    `json:"dealerId"`
	ExpireDate    *string                    `json:"expireDate"`
	Items         []SellTicketItemResponseV2 `json:"items"`
	Status        *string                    `json:"status"`
	AgentId       *string                    `json:"agentid"`
	UnitId        *string                    `json:"unitid"`
	CustomerId    *string                    `json:"customerid"`
	AgentName     *string                    `json:"agentname"`
	UnitName      *string                    `json:"unitname"`
	ContactNumber *string                    `json:"contactnumber"`
	ClientRefNo   *string                    `json:"clientrefno"`
	Barcode       *string                    `json:"barcode"`
	BranchId      *string                    `json:"branchid"`
	BranchName    *string                    `json:"branchname"`
	ProductType   *int                       `json:"producttype"`
}

type CancelTicketRequestV2 struct {
	FromMember  string `json:"frommember"`
	Time        string `json:"time"`
	BillNumber  string `json:"billNumber"`
	ClientRefNo string `json:"clientrefno"`
	Reason      string `json:"reason"`
}

type CancelTicketItemResponseV2 struct {
	Digit         string      `json:"digit"`
	RequestAmount int         `json:"requestAmount"`
	SaleAmount    int         `json:"saleAmount"`
	Status        interface{} `json:"status"`
}

type CancelTicketResponseV2 struct {
	BillNumber    string                       `json:"billNumber"`
	DrawId        int                          `json:"drawId"`
	TotalSale     int                          `json:"totalSale"`
	OrderDate     string                       `json:"orderDate"`
	DrawDate      string                       `json:"drawDate"`
	DrawNumber    string                       `json:"drawNumber"`
	DealerId      string                       `json:"dealerId"`
	ExpireDate    string                       `json:"expireDate"`
	Items         []CancelTicketItemResponseV2 `json:"items"`
	Status        string                       `json:"status"`
	AgentID       string                       `json:"agentid"`
	UnitID        string                       `json:"unitid"`
	CustomerID    *string                      `json:"customerid"`
	AgentName     *string                      `json:"agentname"`
	UnitName      *string                      `json:"unitname"`
	ContactNumber *string                      `json:"contactnumber"`
	ClientRefNo   string                       `json:"clientrefno"`
	Barcode       string                       `json:"barcode"`
	BranchID      *string                      `json:"branchid"`
}
