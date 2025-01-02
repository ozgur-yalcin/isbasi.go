# isbasi.go
Logo İşbaşı API

## Features

- Authentication and session management
- Customer/Supplier (Firm) management
- Invoice operations (e-Invoice, e-Archive)
- Product/Service management
- Stock transaction handling
- Bank and safe operations

## Installation

```bash
go get github.com/ozgur-yalcin/isbasi.go
```

## Getting Started

### Authentication

```go
api := isbasi.Api("your-api-key")

loginReq := &isbasi.LoginRequest{
    Username: "your-username",
    Password: "your-password",
}

ctx := context.Background()
response, err := api.Login(ctx, loginReq)
if err != nil {
    log.Fatal(err)
}
```

### Creating a Customer

```go
customer := &isbasi.FirmRequest{
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
invoice := &isbasi.InvoiceRequest{
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

## API Documentation

### Main Components

- `API`: Main client struct for interacting with the İşbaşı API
- `LoginRequest/LoginResponse`: Authentication structures
- `FirmRequest/FirmResponse`: Customer/Supplier management
- `InvoiceRequest/InvoiceResponse`: Invoice operations
- `ProductRequest/ProductResponse`: Product management

### Common Operations

#### Customer Management
- Create/Update customers
- Get customer details
- List customers
- Delete customers

#### Invoice Operations
- Create regular invoices
- Create e-Invoices
- Create e-Archive invoices
- Get invoice details
- Delete invoices

#### Product Management
- Create/Update products
- Get product details
- List products
- Delete products

## Error Handling

The library uses standard Go error handling patterns. All API methods return an error as their last return value:

```go
if response.IsError {
    // Handle API error
    log.Printf("API Error: %s", response.Message)
}
```

## Environment Setup

The library requires the following environment setup:

- Go 1.16 or later
- API credentials from İşbaşı
- Internet connectivity to access the İşbaşı API endpoints

## Contributing

1. Fork the repository
2. Create your feature branch (`git checkout -b feature/amazing-feature`)
3. Commit your changes (`git commit -m 'Add some amazing feature'`)
4. Push to the branch (`git push origin feature/amazing-feature`)
5. Open a Pull Request

## License

This project is licensed under the MIT License - see the LICENSE file for details.

## Support

For support, please contact yardim@isbasi.com or visit the [İşbaşı documentation](https://isbasi.com/documentation).