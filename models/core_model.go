package models

import "time"

type CoreLoginRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type CoreLoginResponse struct {
	Data struct {
		Seller struct {
			ID         int     `json:"id"`
			SellerCode string  `json:"seller_code"`
			SellerName string  `json:"p_seller_name"`
			Commission float64 `json:"commission"`
			Quota      int     `json:"quata"`
			Status     string  `json:"status"`
			SellerRole string  `json:"seller_role"`
		} `json:"seller"`
		Lot struct {
			DrawDate     string `json:"draw_date"`
			LotCloseDate string `json:"lot_close_date"`
			LotName      string `json:"lot_name"`
			LotID        string `json:"lot_id"`
			LotStatus    string `json:"lot_status"`
			LotDate      string `json:"lot_date"`
		} `json:"lot"`
		AccessToken   string `json:"access_token"`
		VersionApp    string `json:"version_app"`
		VersionLotApp string `json:"version_lot_app"`
	} `json:"data"`
	Status  bool   `json:"status"`
	Message string `json:"message,omitempty"`
}

type CoreLoginRawResponse struct {
	Data struct {
		DigitType      string `json:"DigitType"`
		LOTCLOSETIME   string `json:"LOTCLOSETIME"`
		LOTDATE        string `json:"LOTDATE"`
		LOTID          int    `json:"LOTID"`
		LOT_STATUS     string `json:"LOT_STATUS"`
		STATUSID       int    `json:"STATUSID"`
		SessionExpired string `json:"SessionExpired"`
		TOKEN          string `json:"TOKEN"`
	} `json:"data"`
	Status bool `json:"status"`
}

type CoreErrorResponse struct {
	Status  bool   `json:"status"`
	Message string `json:"message"`
}

type CoreSellRequest struct {
	FromMember  string                `json:"frommember"`
	Time        time.Time             `json:"time"`
	ClientRefNo string                `json:"clientrefno"`
	BillNumber  *string               `json:"billnumber"`
	CustomerID  string                `json:"customerid"`
	Items       []CoreSellRequestItem `json:"items"`
}

type CoreSellRequestItem struct {
	Digit    string `json:"digit"`
	Amount   int    `json:"amount"`
	LLAmount int    `json:"lotlink_amount"`
}

type CoreSellResponse struct {
	Data struct {
		BillNumber  string        `json:"billNumber"`
		SellerID    string        `json:"sellerId"`
		Barcode     string        `json:"barcode"`
		DrawID      int           `json:"drawId"`
		TotalSale   int           `json:"totalSale"`
		OrderDate   string        `json:"orderDate"`
		DrawDate    string        `json:"drawDate"`
		DrawNumber  string        `json:"drawNumber"`
		DealerID    string        `json:"dealerId"`
		ExpireDate  string        `json:"expireDate"`
		Items       []CoreItemRes `json:"items"`
		BranchPhone string        `json:"branch_phone"`
		DealerPhone string        `json:"dealer_phone"`
		Status      string        `json:"status"`
		LotID       string        `json:"lot_id"`
		BranchName  string        `json:"branch_name"`
	} `json:"data"`
	Status bool `json:"status"`
}

type CoreSellRawResponse struct {
	Data struct {
		BillNumber    string        `json:"billNumber"`
		DrawId        int           `json:"drawId"`
		TotalSale     int           `json:"totalSale"`
		OrderDate     string        `json:"orderDate"`
		DrawDate      string        `json:"drawDate"`
		DrawNumber    string        `json:"drawNumber"`
		DealerId      string        `json:"dealerId"`
		ExpireDate    string        `json:"expireDate"`
		Items         []CoreItemRes `json:"items"`
		Status        string        `json:"status"`
		AgentId       string        `json:"agentid"`
		UnitId        string        `json:"unitid"`
		CustomerId    string        `json:"customerid"`
		AgentName     string        `json:"agentname"`
		UnitName      string        `json:"unitname"`
		ContactNumber string        `json:"contactnumber"`
		ClientRefNo   string        `json:"clientrefno"`
		Barcode       string        `json:"barcode"`
		BranchId      string        `json:"branchid"`
		BranchName    string        `json:"branchname"`
		ProductType   int           `json:"producttype"`
	} `json:"data"`
	Status bool `json:"status"`
}

type CoreItemRes struct {
	Digit         string `json:"digit"`
	RequestAmount int    `json:"requestAmount"`
	SaleAmount    int    `json:"saleAmount"`
	Status        int    `json:"status"`
	DigitType     int    `json:"digittype"`
}

type CoreCancelRequest struct {
	SellerCode string `json:"seller_code"`
	BillNumber string `json:"billNumber"`
	Reason     string `json:"reason"`
}

type CoreCancelResponse struct {
	Data struct {
		BillNumber string `json:"billNumber"`
		DrawID     int    `json:"drawId"`
		TotalSale  int    `json:"totalSale"`
		OrderDate  string `json:"orderDate"`
		DrawDate   string `json:"drawDate"`
		DrawNumber string `json:"drawNumber"`
		DealerID   string `json:"dealerId"`
		ExpireDate string `json:"expireDate"`
	} `json:"data"`
	Status bool `json:"status"`
}

type CoreUpdateReq struct {
	Old string `json:"old"`
	New string `json:"new"`
}

type CoreUpdateResponse struct {
	Data   string `json:"data"`
	Status bool   `json:"status"`
}
