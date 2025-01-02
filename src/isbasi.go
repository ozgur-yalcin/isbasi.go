package isbasi

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"time"
)

const (
	baseUrl = "https://isbasimw.isbasi.com/api/v1.0"
)

type API struct {
	BaseUrl   string
	SecretKey string
	TenantId  string
	AuthToken string
	Language  string
}

type LoginRequest struct {
	Username    string `json:"username,omitempty"`
	Password    string `json:"password,omitempty"`
	AccessToken string `json:"accessToken,omitempty"`
	TenantID    string `json:"tenantId,omitempty"`
	BaseUrl     string `json:"baseUrl,omitempty"`
}

type LoginResponse struct {
	Code    int          `json:"code,omitempty"`
	Message string       `json:"message,omitempty"`
	IsError bool         `json:"isError,omitempty"`
	Data    LoginRequest `json:"data,omitempty"`
}

type FirmRequest struct {
	ID              int      `json:"id,omitempty"`
	Code            string   `json:"code,omitempty"`
	IsActive        bool     `json:"isActive,omitempty"`
	Name            string   `json:"name,omitempty"`
	TaxOrPersonalID string   `json:"taxOrPersonalId,omitempty"`
	TaxOffice       string   `json:"taxOffice,omitempty"`
	Country         string   `json:"country,omitempty"`
	City            string   `json:"city,omitempty"`
	District        string   `json:"district,omitempty"`
	Address         string   `json:"address,omitempty"`
	Phone           string   `json:"phone,omitempty"`
	EmailAddress    string   `json:"emailAddress,omitempty"`
	FirmType        int      `json:"firmType,omitempty"`
	Currency        string   `json:"currency,omitempty"`
	Balance         float64  `json:"balance,omitempty"`
	Tags            []string `json:"tags,omitempty"`
}

type FirmResponse struct {
	Code    int         `json:"code,omitempty"`
	Message string      `json:"message,omitempty"`
	IsError bool        `json:"isError,omitempty"`
	Data    FirmRequest `json:"data,omitempty"`
}

type InvoiceRequest struct {
	ID           int       `json:"id,omitempty"`
	Type         string    `json:"type,omitempty"`
	Number       string    `json:"number,omitempty"`
	Date         time.Time `json:"date,omitempty"`
	CustomerID   int       `json:"customerId,omitempty"`
	CustomerName string    `json:"customerName,omitempty"`
	Currency     string    `json:"currency,omitempty"`
	ExchangeRate float64   `json:"exchangeRate,omitempty"`
	Total        float64   `json:"total,omitempty"`
	VatAmount    float64   `json:"vatAmount,omitempty"`
	Description  string    `json:"description,omitempty"`
	IsCancelled  bool      `json:"isCancelled,omitempty"`
}

type InvoiceResponse struct {
	Code    int            `json:"code,omitempty"`
	Message string         `json:"message,omitempty"`
	IsError bool           `json:"isError,omitempty"`
	Data    InvoiceRequest `json:"data,omitempty"`
}

type ProductRequest struct {
	ID       int     `json:"id,omitempty"`
	Code     string  `json:"code,omitempty"`
	Name     string  `json:"name,omitempty"`
	Type     int     `json:"type,omitempty"`
	VatRate  float64 `json:"vatRate,omitempty"`
	IsActive bool    `json:"isActive,omitempty"`
}

type ProductResponse struct {
	Code    int            `json:"code,omitempty"`
	Message string         `json:"message,omitempty"`
	IsError bool           `json:"isError,omitempty"`
	Data    ProductRequest `json:"data,omitempty"`
}

func Api(secretKey string) *API {
	client := new(API)
	client.BaseUrl = baseUrl
	client.SecretKey = secretKey
	client.Language = "tr-TR"
	return client
}

func (api *API) SetBaseUrl(url string) {
	api.BaseUrl = url
}

func (api *API) SetLanguage(lang string) {
	api.Language = lang
}

func (api *API) NewRequest(ctx context.Context, method, path string, body interface{}) (*http.Response, error) {
	payload, err := json.Marshal(body)
	if err != nil {
		return nil, fmt.Errorf("failed to marshal request body: %v", err)
	}
	req, err := http.NewRequestWithContext(ctx, method, api.BaseUrl+path, bytes.NewBuffer(payload))
	if err != nil {
		return nil, fmt.Errorf("failed to create request: %v", err)
	}
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("ApiKey", api.SecretKey)
	req.Header.Set("Authorization", "Bearer "+api.AuthToken)
	req.Header.Set("tenantId", api.TenantId)
	req.Header.Set("lang", api.Language)
	client := new(http.Client)
	res, err := client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("failed to execute request: %v", err)
	}
	defer res.Body.Close()
	return res, nil
}

func (api *API) Login(ctx context.Context, login *LoginRequest) (result LoginResponse, err error) {
	payload, err := json.Marshal(login)
	if err != nil {
		return result, fmt.Errorf("failed to marshal login request: %v", err)
	}
	req, err := http.NewRequestWithContext(ctx, "POST", api.BaseUrl+"/user/integrationLogin", bytes.NewBuffer(payload))
	if err != nil {
		return result, fmt.Errorf("failed to create login request: %v", err)
	}
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("ApiKey", api.SecretKey)
	client := new(http.Client)
	res, err := client.Do(req)
	if err != nil {
		return result, fmt.Errorf("failed to execute login request: %v", err)
	}
	defer res.Body.Close()
	if err := json.NewDecoder(res.Body).Decode(&result); err != nil {
		return result, fmt.Errorf("failed to decode login resonse: %v", err)
	}
	if result.IsError {
		return result, fmt.Errorf("login failed: %s", result.Message)
	}
	api.AuthToken = result.Data.AccessToken
	api.TenantId = result.Data.TenantID
	if result.Data.BaseUrl != "" {
		api.BaseUrl = result.Data.BaseUrl
	}
	return result, nil
}

func (api *API) CreateFirm(ctx context.Context, firm *FirmRequest) (result FirmResponse, err error) {
	res, err := api.NewRequest(ctx, "PUT", "/firms", firm)
	if err != nil {
		return result, err
	}
	defer res.Body.Close()
	if err := json.NewDecoder(res.Body).Decode(&result); err != nil {
		return result, fmt.Errorf("failed to decode resonse: %v", err)
	}
	if result.IsError {
		return result, fmt.Errorf("API error: %s", result.Message)
	}
	return result, nil
}

func (api *API) CreateInvoice(ctx context.Context, invoice *InvoiceRequest) (result InvoiceResponse, err error) {
	res, err := api.NewRequest(ctx, "POST", "/invoices/integrationInvoices", invoice)
	if err != nil {
		return result, err
	}
	defer res.Body.Close()
	if err := json.NewDecoder(res.Body).Decode(&result); err != nil {
		return result, fmt.Errorf("failed to decode resonse: %v", err)
	}
	if result.IsError {
		return result, fmt.Errorf("API error: %s", result.Message)
	}
	return result, nil
}

func (api *API) GetFirm(ctx context.Context, id int) (result FirmResponse, err error) {
	res, err := api.NewRequest(ctx, "GET", fmt.Sprintf("/firms/%d", id), nil)
	if err != nil {
		return result, err
	}
	defer res.Body.Close()
	if err := json.NewDecoder(res.Body).Decode(&result); err != nil {
		return result, fmt.Errorf("failed to decode resonse: %v", err)
	}
	if result.IsError {
		return result, fmt.Errorf("API error: %s", result.Message)
	}
	return result, nil
}

func (api *API) GetProduct(ctx context.Context, id, productType int) (result ProductResponse, err error) {
	res, err := api.NewRequest(ctx, "GET", fmt.Sprintf("/products/%d/%d", id, productType), nil)
	if err != nil {
		return result, err
	}
	defer res.Body.Close()
	if err := json.NewDecoder(res.Body).Decode(&result); err != nil {
		return result, fmt.Errorf("failed to decode resonse: %v", err)
	}
	if result.IsError {
		return result, fmt.Errorf("API error: %s", result.Message)
	}
	return result, nil
}
