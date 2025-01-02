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

type Category struct {
	ID   int    `json:"id,omitempty"`
	Name string `json:"name,omitempty"`
	Code string `json:"code,omitempty"`
	Type int    `json:"type,omitempty"`
}

type Employee struct {
	FullName string `json:"fullName,omitempty"`
	Email    string `json:"email,omitempty"`
}

type ShippingAddress struct {
	ID       int    `json:"id,omitempty"`
	FirmID   int    `json:"firmid,omitempty"`
	Title    string `json:"title,omitempty"`
	Address  string `json:"address,omitempty"`
	Country  string `json:"country,omitempty"`
	City     string `json:"city,omitempty"`
	District string `json:"district,omitempty"`
	Code     string `json:"code,omitempty"`
	FullName string `json:"fullName,omitempty"`
	Phone    string `json:"phone,omitempty"`
}

type Bank struct {
	Name          string `json:"name,omitempty"`
	Branch        string `json:"branch,omitempty"`
	BranchCode    string `json:"branchCode,omitempty"`
	AccountNumber string `json:"accountNumber,omitempty"`
	Currency      string `json:"currency,omitempty"`
	IBAN          string `json:"iban,omitempty"`
}

type BankAccount struct {
	BankID      int    `json:"bankId,omitempty"`
	AccountID   int    `json:"accountId,omitempty"`
	AccountCode string `json:"accountCode,omitempty"`
	Name        string `json:"name,omitempty"`
	IBAN        string `json:"iban,omitempty"`
	Currency    string `json:"currency,omitempty"`
	BankName    string `json:"bankName,omitempty"`
	BranchName  string `json:"branchName,omitempty"`
}

type EPortalLogin struct {
	UserID   string `json:"userId,omitempty"`
	Password string `json:"password,omitempty"`
}

type FirmRequest struct {
	ID                                   int               `json:"id,omitempty"`
	Code                                 string            `json:"code,omitempty"`
	IsActive                             bool              `json:"isActive,omitempty"`
	IsPersonalCompany                    bool              `json:"isPersonalCompany,omitempty"`
	Name                                 string            `json:"name,omitempty"`
	FirstName                            string            `json:"firstName,omitempty"`
	LastName                             string            `json:"lastName,omitempty"`
	FullName                             string            `json:"fullName,omitempty"`
	DisplayName                          string            `json:"displayName,omitempty"`
	TaxOrPersonalID                      string            `json:"taxOrPersonalId,omitempty"`
	TaxOffice                            string            `json:"taxOffice,omitempty"`
	Country                              string            `json:"country,omitempty"`
	City                                 string            `json:"city,omitempty"`
	ValidateCityAndDistrict              bool              `json:"validateCityAndDistrict,omitempty"`
	District                             string            `json:"district,omitempty"`
	PostalCode                           string            `json:"postalCode,omitempty"`
	Address                              string            `json:"address,omitempty"`
	Phone                                string            `json:"phone,omitempty"`
	WebAddress                           string            `json:"webAddress,omitempty"`
	Tags                                 []string          `json:"tags,omitempty"`
	Category                             Category          `json:"category,omitempty"`
	PhoneNumbers                         []string          `json:"phoneNumbers,omitempty"`
	EmailAddress                         string            `json:"emailAddress,omitempty"`
	Employees                            []Employee        `json:"employees,omitempty"`
	ShippingAddresses                    []ShippingAddress `json:"shippingAddresses,omitempty"`
	Banks                                []Bank            `json:"banks,omitempty"`
	FaxNumber                            string            `json:"faxNumber,omitempty"`
	Icon                                 string            `json:"icon,omitempty"`
	UserID                               string            `json:"UserId,omitempty"`
	EInvoiceResponsible                  bool              `json:"eInvoiceResponsible,omitempty"`
	DefaultReportTemplate                string            `json:"defaultReportTemplate,omitempty"`
	FirmType                             int               `json:"firmType,omitempty"`
	EInvoiceProfile                      int               `json:"eInvoiceProfile,omitempty"`
	EInvoiceSenderLabel                  string            `json:"eInvoiceSenderLabel,omitempty"`
	EInvoicePostLabel                    string            `json:"eInvoicePostLabel,omitempty"`
	ELogoUserName                        string            `json:"eLogoUserName,omitempty"`
	ELogoPassword                        string            `json:"eLogoPassword,omitempty"`
	NaceCode                             string            `json:"naceCode,omitempty"`
	EInvoiceControlType                  int               `json:"eInvoiceControlType,omitempty"`
	EInvoiceCustoms                      bool              `json:"eInvoiceCustoms,omitempty"`
	EInvoiceBrokerComp                   int               `json:"eInvoiceBrokerComp,omitempty"`
	EArchiveResponsible                  bool              `json:"eArchiveResponsible,omitempty"`
	EArchiveWebSite                      string            `json:"eArchiveWebSite,omitempty"`
	AdditionalInvoiceType                int               `json:"additionalInvoiceType,omitempty"`
	SgkResponsibleCode                   string            `json:"sgkResponsibleCode,omitempty"`
	SgkResponsibleName                   string            `json:"sgkResponsibleName,omitempty"`
	EArchiveSendMod                      int               `json:"eArchiveSendMod,omitempty"`
	EGovermentType                       int               `json:"eGovermentType,omitempty"`
	ESmmResponsible                      bool              `json:"eSmmResponsible,omitempty"`
	ESmmSendMod                          int               `json:"eSmmSendMod,omitempty"`
	AcceptEinvPublic                     int               `json:"acceptEinvPublic,omitempty"`
	GenericCustomer                      bool              `json:"genericCustomer,omitempty"`
	NotApplyVat                          bool              `json:"notApplyVat,omitempty"`
	NotApplyWithHolding                  bool              `json:"notApplyWithHolding,omitempty"`
	NotApplyAdditionalTax                bool              `json:"notApplyAdditionalTax,omitempty"`
	MersisNo                             string            `json:"mersisNo,omitempty"`
	TradeRegisterNumber                  string            `json:"tradeRegisterNumber,omitempty"`
	PredefinedDescription                string            `json:"predefinedDescription,omitempty"`
	IsAdmin                              bool              `json:"isAdmin,omitempty"`
	IsCharteredAccountant                bool              `json:"isCharteredAccountant,omitempty"`
	ErrorMessage                         string            `json:"errorMessage,omitempty"`
	BeginningBalance                     float64           `json:"beginningBalance,omitempty"`
	BeginningBalanceDate                 time.Time         `json:"beginningBalanceDate,omitempty"`
	Balance                              float64           `json:"balance,omitempty"`
	CurrencyBalance                      float64           `json:"currencyBalance,omitempty"`
	CostMethodID                         int               `json:"costMethodId,omitempty"`
	Currency                             string            `json:"currency,omitempty"`
	Description                          string            `json:"description,omitempty"`
	EInvoiceBeginDate                    time.Time         `json:"eInvoiceBeginDate,omitempty"`
	EArchiveBeginDate                    time.Time         `json:"eArchiveBeginDate,omitempty"`
	IsSendDispatchInEInvoice             bool              `json:"isSendDispatchInEInvoice,omitempty"`
	SenderIbanBankAccountID              int               `json:"senderIbanBankAccountId,omitempty"`
	PurchaseServicesReceiptReadingMethod int               `json:"purchaseServicesReceiptReadingMethod,omitempty"`
	WasAccessPermissionGranted           bool              `json:"wasAccessPermissionGranted,omitempty"`
	BankAccount                          BankAccount       `json:"bankAccount,omitempty"`
	ParentTenantSetDate                  time.Time         `json:"parentTenantSetDate,omitempty"`
	IsIntegrationFirm                    bool              `json:"isIntegrationFirm,omitempty"`
	HasApiAuthAuthority                  bool              `json:"hasApiAuthAuthority,omitempty"`
	EPortalArchiveResponsible            bool              `json:"ePortalArchiveResponsible,omitempty"`
	EAPortalLoginInformation             EPortalLogin      `json:"eAPortalLoginInformation,omitempty"`
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

type ProductRequest struct {
	ID       int     `json:"id,omitempty"`
	Code     string  `json:"code,omitempty"`
	Name     string  `json:"name,omitempty"`
	Type     int     `json:"type,omitempty"`
	VatRate  float64 `json:"vatRate,omitempty"`
	IsActive bool    `json:"isActive,omitempty"`
}

type FirmResponse struct {
	Code    int         `json:"code,omitempty"`
	Message string      `json:"message,omitempty"`
	IsError bool        `json:"isError,omitempty"`
	Data    FirmRequest `json:"data,omitempty"`
}

type InvoiceResponse struct {
	Code    int            `json:"code,omitempty"`
	Message string         `json:"message,omitempty"`
	IsError bool           `json:"isError,omitempty"`
	Data    InvoiceRequest `json:"data,omitempty"`
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

func (api *API) Login(ctx context.Context, body *LoginRequest) (result LoginResponse, err error) {
	payload, err := json.Marshal(body)
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
		return result, fmt.Errorf("failed to decode login response: %v", err)
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

func (api *API) CreateFirm(ctx context.Context, req *FirmRequest) (result FirmResponse, err error) {
	res, err := api.NewRequest(ctx, "PUT", "/firms", req)
	if err != nil {
		return result, err
	}
	defer res.Body.Close()
	if err := json.NewDecoder(res.Body).Decode(&result); err != nil {
		return result, fmt.Errorf("failed to decode response: %v", err)
	}
	if result.IsError {
		return result, fmt.Errorf("API error: %s", result.Message)
	}
	return result, nil
}

func (api *API) CreateInvoice(ctx context.Context, req *InvoiceRequest) (result InvoiceResponse, err error) {
	res, err := api.NewRequest(ctx, "POST", "/invoices/integrationInvoices", req)
	if err != nil {
		return result, err
	}
	defer res.Body.Close()
	if err := json.NewDecoder(res.Body).Decode(&result); err != nil {
		return result, fmt.Errorf("failed to decode response: %v", err)
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
		return result, fmt.Errorf("failed to decode response: %v", err)
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
		return result, fmt.Errorf("failed to decode response: %v", err)
	}
	if result.IsError {
		return result, fmt.Errorf("API error: %s", result.Message)
	}
	return result, nil
}
