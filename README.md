# ithq-kkl-shared

Shared Go package containing common entities and models used across ithq-kkl services.

## Installation

```bash
go get github.com/kolao-ithq/ithq-kkl-shared@v0.1.0
```

> Current version: **v0.1.0** — see [VERSION](./VERSION)

## Private Repository Setup

This is a private module. Before you can `go get` it, each developer/CI must configure Go and Git once.

### 1. Tell Go to skip the public proxy for this module

```bash
go env -w GONOSUMCHECK=github.com/kolao-ithq/*
go env -w GONOSUMDB=github.com/kolao-ithq/*
go env -w GOPRIVATE=github.com/kolao-ithq/*
```

Or add to your shell profile (`~/.zshrc` / `~/.bashrc`):
```bash
export GOPRIVATE=github.com/kolao-ithq/*
```

### 2. Configure Git to authenticate with GitHub

**Option A — SSH (recommended)**
```bash
git config --global url."git@github.com:".insteadOf "https://github.com/"
```
Make sure your SSH key is added to your GitHub account.

**Option B — Personal Access Token (HTTPS)**
```bash
git config --global url."https://<YOUR_GITHUB_TOKEN>@github.com/".insteadOf "https://github.com/"
```
Replace `<YOUR_GITHUB_TOKEN>` with a token that has `repo` scope.

### 3. Install the module

```bash
go get github.com/kolao-ithq/ithq-kkl-shared@v0.1.0
```

### CI/CD (GitHub Actions example)

```yaml
- name: Configure private module access
  run: git config --global url."https://${{ secrets.GH_TOKEN }}@github.com/".insteadOf "https://github.com/"
  env:
    GONOSUMDB: "github.com/kolao-ithq/*"
    GOPRIVATE: "github.com/kolao-ithq/*"
```

---

## Packages

| Package | Description |
|---------|-------------|
| `kkl_entities` | GORM database entity structs |
| `kkl_models` | Request/response structs for Core and Lotto APIs |

---

## `kkl_entities` — Database Entities

### InsertOnly (base struct)

A lightweight alternative to `gorm.Model` — has `ID`, `CreatedAt`, and `DeletedAt` but no `UpdatedAt`.

```go
import "github.com/kolao-ithq/ithq-kkl-shared/kkl_entities"

// Bill embeds InsertOnly
bill := kkl_entities.Bill{
    TicketID:      "TKT-0001",
    TransactionID: "TXN-9999",
    LotID:         1,
    CreatedBy:     42,
    Amount:        50000,
    IsCoreSale:    true,
}
```

### Bill, BillStatus, BillDigit

```go
billStatus := kkl_entities.BillStatus{
    BillID: 1,
    Status: 2,
    Remark: "approved by admin",
}

billDigit := kkl_entities.BillDigit{
    BillID:     1,
    DigitType:  3,
    DigitValue: "123",
    Amount:     10000,
    CoreAmount: 9500,
}
```

### Payment & PaymentMethod

```go
payment := kkl_entities.Payment{
    BillID:                1,
    BankReference:         "REF-ABC123",
    PaymentMethodID:       2,
    Amount:                50000,
    Discount:              0,
    ActualAmount:          50000,
    BankCode:              "BCEL",
    AccountType:           "saving",
    BankAccountNumber:     "0001234567",
    CurrentCommissionRate: 1.5,
}
```

### CoreCustomer & SellRatio

```go
customer := kkl_entities.CoreCustomer{
    CustomerID: 101,
}

ratio := kkl_entities.SellRatio{
    DigitType: 2,
    Ratio:     0.85,
}
```

### Auto-migrate example (GORM)

```go
db.AutoMigrate(
    &kkl_entities.Bill{},
    &kkl_entities.BillStatus{},
    &kkl_entities.BillDigit{},
    &kkl_entities.Payment{},
    &kkl_entities.PaymentMethod{},
    &kkl_entities.CoreCustomer{},
    &kkl_entities.SellRatio{},
)
```

---

## `kkl_models` — API Request/Response Models

### Core API

```go
import "github.com/kolao-ithq/ithq-kkl-shared/kkl_models"

// Login
req := kkl_models.CoreLoginRequest{
    Username: "seller01",
    Password: "secret",
}

var res kkl_models.CoreLoginResponse
// json.Unmarshal(body, &res)
fmt.Println(res.Data.Seller.SellerCode)
fmt.Println(res.Data.AccessToken)

// Sell
sellReq := kkl_models.CoreSellRequest{
    FromMember:  "seller01",
    Time:        time.Now(),
    ClientRefNo: "REF-001",
    CustomerID:  "CUST-001",
    Items: []kkl_models.CoreSellRequestItem{
        {Digit: "123", Amount: 10000, LLAmount: 0},
        {Digit: "456", Amount: 5000, LLAmount: 0},
    },
}

// Cancel
cancelReq := kkl_models.CoreCancelRequest{
    SellerCode: "seller01",
    BillNumber: "BILL-0001",
    Reason:     "customer request",
}
```

### Lotto API (V2)

```go
// Sign in
signIn := kkl_models.LottoSignInBodyV2{
    Username: "dealer01",
    Password: "secret",
}

var signInRes kkl_models.LottoSignInResponseV2
// json.Unmarshal(body, &signInRes)
fmt.Println(signInRes.AccessToken)
fmt.Println(signInRes.CurrentDraw.DrawDate)

// Sell ticket
sellReq := kkl_models.SellTicketRequestV2{
    FromMember:  "dealer01",
    Time:        time.Now(),
    ClientRefNo: "REF-002",
    CustomerID:  "CUST-002",
    Items: []kkl_models.SellTicketItemRequestV2{
        {Digit: "789", Amount: 20000},
    },
}

// Cancel ticket
cancelReq := kkl_models.CancelTicketRequestV2{
    FromMember:  "dealer01",
    Time:        time.Now().Format(time.RFC3339),
    BillNumber:  "BILL-0002",
    ClientRefNo: "REF-002",
    Reason:      "mistake",
}
```

### Sell (internal)

```go
sellReq := kkl_models.SellRequest{
    Phone:            "0201234567",
    DaysOfInstallmentID: 1,
    CustomerID:       101,
    LottoDrawsNumber: []kkl_models.LottoDrawsNumber{
        {Digit: "111", Money: 10000},
        {Digit: "222", Money: 5000},
    },
    PaymentRequest: kkl_models.PaymentRequest{
        PaymentMethodID: 2,
        DrawID:          99,
        TotalSale:       15000,
        Phone:           "0201234567",
        Transaction:     "TXN-XYZ",
        BillNumber:      "BILL-0003",
        BillID:          10,
        Account:         "0009876543",
        AccountType:     "saving",
        CustomerID:      101,
        FromApp:         "kkl-app",
        UsePoint:        false,
    },
}
```

---

## Module Path

```
github.com/kolao-ithq/ithq-kkl-shared
```
