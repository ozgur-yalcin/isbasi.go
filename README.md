[![license](https://img.shields.io/:license-mit-blue.svg)](https://github.com/ozgur-yalcin/isbasi.go/blob/master/LICENSE.md)
[![documentation](https://pkg.go.dev/badge/github.com/ozgur-yalcin/isbasi.go)](https://pkg.go.dev/github.com/ozgur-yalcin/isbasi.go/src)

# isbasi.go
Logo İşbaşı API

# Installation

```bash
go get github.com/ozgur-yalcin/isbasi.go
```

# Creating a Customer

```go
package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"

	isbasi "github.com/ozgur-yalcin/isbasi.go/src"
)

func main() {
	api := isbasi.Api("your-api-key")
	api.SetLanguage("tr-TR")

	login := &isbasi.Login{
		Username: "your-username",
		Password: "your-password",
	}

	ctx := context.Background()
	_, err := api.Login(ctx, login)
	if err != nil {
		log.Fatal(err)
	}

	customer := &isbasi.Firm{
		Name:       "Test",
		TcknVkn:    "1234567890",
		IsPersonal: true, // Bireysel / Kurumsal
		TaxOffice:  "Vergi dairesi",
		Country:    "Türkiye",
		City:       "İstanbul",
		District:   "Kadıkoy",
		Address:    "Adres",
		FirmType:   1, // 1: Müşteri, 2: Tedarikçi, 3: Her ikisi de
	}

	if res, err := api.CreateFirm(ctx, customer); err == nil {
		pretty, _ := json.MarshalIndent(res, " ", " ")
		fmt.Println(string(pretty))
	} else {
		fmt.Println(err)
	}
}
```

# Creating an Invoice

```go
package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"

	isbasi "github.com/ozgur-yalcin/isbasi.go/src"
)

func main() {
	api := isbasi.Api("your-api-key")
	api.SetLanguage("tr-TR")

	login := &isbasi.Login{
		Username: "your-username",
		Password: "your-password",
	}

	ctx := context.Background()
	_, err := api.Login(ctx, login)
	if err != nil {
		log.Fatal(err)
	}

	invoice := &isbasi.Invoice{
		Customer: &isbasi.Customer{
			Code:      "CUST001",
			Name:      "Test",
			TcknVkn:   "1234567890",
			TaxOffice: "Vergi dairesi",
			Country:   "Türkiye",
			City:      "İstanbul",
			District:  "Kadıkoy",
			Address:   "Adres",
		},
		InvoiceDate:  "2025-01-02",
		Currency:     "TRY",
		Description:  "Fatura açıklaması",
		VatIncluded:  true, // KDV dahil
	}

	salesInvoice := &isbasi.SalesInvoiceDetail{
		Quantity:    1,
		TaxRate:     20,
		Price:       100,
		Description: "Test",
		ProductDetail: &isbasi.ProductDetail{
			ItemCode: "PROD001",
			ItemType: 1,
			Name:     "Test",
			Vat:      20,
		},
	}

	invoice.SalesInvoiceDetails = append(invoice.SalesInvoiceDetails, salesInvoice)

	if res, err := api.CreateInvoice(ctx, invoice); err == nil {
		pretty, _ := json.MarshalIndent(res, " ", " ")
		fmt.Println(string(pretty))
	} else {
		fmt.Println(err)
	}
}
```
