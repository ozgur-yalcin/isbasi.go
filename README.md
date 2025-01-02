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
		Name:       "Test",       // Müşteri adı
		TcknVkn:    "1234567890", // TCKN / VKN
		IsPersonal: true,         // Bireysel / Kurumsal
		TaxOffice:  "Maslak",     // Vergi dairesi
		Country:    "Türkiye",    // Ülke
		City:       "İstanbul",   // Şehir
		District:   "Kadıkoy",    // İlçe
		Address:    "No:1",       // Adres
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
		InvoiceDate: "2025-01-02", // Fatura tarihi
		Description: "",           // Fatura açıklaması
		Currency:    "TRY",        // Para birimi
		VatIncluded: true,         // KDV dahil
		Customer: &isbasi.Customer{
			Code:       "CUST001",    // Müşteri kodu
			Name:       "Test",       // Müşteri adı
			TcknVkn:    "1234567890", // TCKN / VKN
			IsPersonal: true,         // Bireysel / Kurumsal
			TaxOffice:  "Maslak",     // Vergi Dairesi
			Country:    "Türkiye",    // Ülke
			City:       "İstanbul",   // Şehir
			District:   "Kadıkoy",    // İlçe
			Address:    "No:1",       // Adres
		},
	}

	salesInvoice := &isbasi.SalesInvoiceDetail{
		Quantity:    1,      // Miktar
		TaxRate:     20,     // KDV Oranı
		Price:       1.00,   // Fiyat
		Name:        "Test", // Ürün adı
		Description: "Test", // Ürün açıklaması
		ProductDetail: &isbasi.ProductDetail{
			ItemCode: "PROD001", // Ürün kodu
			ItemType: 1,         // Ürün tipi (1: Mal, 2: Hizmet)
			Name:     "Test",    // Ürün adı
			Vat:      20,        // KDV Oranı
			Unit:     "Adet",    // Birim
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
