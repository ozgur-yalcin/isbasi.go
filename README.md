[![license](https://img.shields.io/:license-mit-blue.svg)](https://github.com/ozgur-yalcin/isbasi.go/blob/master/LICENSE.md)
[![documentation](https://pkg.go.dev/badge/github.com/ozgur-yalcin/isbasi.go)](https://pkg.go.dev/github.com/ozgur-yalcin/isbasi.go/src)

# isbasi.go
Logo İşbaşı API

## Installation

```bash
go get github.com/ozgur-yalcin/isbasi.go
```

## Getting Started

### Authentication

```go
api := isbasi.Api("your-api-key")

login := &isbasi.Login{
    Username: "your-username",
    Password: "your-password",
}

ctx := context.Background()
response, err := api.Login(ctx, login)
if err != nil {
    log.Fatal(err)
}
```

### Creating a Customer

```go
customer := &isbasi.Firm{
    Name: "Test Company",
    TaxOrPersonalID: "1234567890",
    TaxOffice: "Test Tax Office",
    Country: "Turkey",
    City: "Istanbul",
    District: "Kadikoy",
    Address: "Test Address",
    FirmType: 1, // 1: Customer, 2: Supplier, 3: Both
}

response, err := api.CreateFirm(ctx, customer)
if err != nil {
    log.Fatal(err)
}
```

### Creating an Invoice

```go
invoice := &isbasi.Invoice{
    Customer: &isbasi.Customer{
        Code: "CUST001",
        Name: "Test Customer",
        TcknVkn: "1234567890",
        TaxOffice: "Test Tax Office",
        Country: "Turkey",
        City: "Istanbul",
        District: "Kadikoy",
        Address: "Test Address",
    },
    InvoiceDate: "2025-01-02",
    Currency: "TRY",
    ExchangeRate: 1,
    Description: "Test Invoice",
    SalesInvoiceDetails: []*isbasi.SalesInvoiceDetail{
        {
            Quantity: 1,
            TaxRate: 18,
            Price: 100,
            Description: "Test Product",
            ProductDetail: &isbasi.ProductDetail{
                ItemCode: "PROD001",
                ItemType: 1,
                Name: "Test Product",
                Vat: 18,
            },
        },
    },
}

response, err := api.CreateInvoice(ctx, invoice)
if err != nil {
    log.Fatal(err)
}
```
