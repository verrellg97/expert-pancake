package model

import (
	"net/http"
)

type SalesService interface {
	UpsertPOS(w http.ResponseWriter, r *http.Request) error
	DeletePOS(w http.ResponseWriter, r *http.Request) error
	GetPOS(w http.ResponseWriter, r *http.Request) error
	GetPOSItems(w http.ResponseWriter, r *http.Request) error

	GetPOSUserSetting(w http.ResponseWriter, r *http.Request) error
	UpdatePOSUserSetting(w http.ResponseWriter, r *http.Request) error

	UpdatePOSCOASetting(w http.ResponseWriter, r *http.Request) error
	GetPOSCOASetting(w http.ResponseWriter, r *http.Request) error

	UpdatePOSCustomerSetting(w http.ResponseWriter, r *http.Request) error
	GetPOSCustomerSetting(w http.ResponseWriter, r *http.Request) error

	UpsertPOSPaymentMethod(w http.ResponseWriter, r *http.Request) error
	DeletePOSPaymentMethod(w http.ResponseWriter, r *http.Request) error
	GetPOSPaymentMethod(w http.ResponseWriter, r *http.Request) error
}

type POS struct {
	Id                   string `json:"id" validate:"required"`
	CompanyId            string `json:"company_id" validate:"required"`
	BranchId             string `json:"branch_id" validate:"required"`
	WarehouseId          string `json:"warehouse_id" validate:"required"`
	WarehouseName        string `json:"warehouse_name" validate:"required"`
	FormNumber           string `json:"form_number" validate:"required"`
	TransactionDate      string `json:"transaction_date" validate:"required"`
	ContactBookId        string `json:"contact_book_id" validate:"required"`
	ContactBookName      string `json:"contact_book_name" validate:"required"`
	SecondaryCompanyId   string `json:"secondary_company_id" validate:"required"`
	KonekinId            string `json:"konekin_id" validate:"required"`
	CurrencyCode         string `json:"currency_code" validate:"required"`
	POSPaymentMethodId   string `json:"pos_payment_method_id" validate:"required"`
	POSPaymentMethodName string `json:"pos_payment_method_name" validate:"required"`
	ChartOfAccountId   string `json:"chart_of_account_id" validate:"required"`
	ChartOfAccountName string `json:"chart_of_account_name" validate:"required"`
	TotalItems           string `json:"total_items" validate:"required"`
	Total                string `json:"total" validate:"required"`
}

type POSItem struct {
	DetailId          string  `json:"detail_id" validate:"required"`
	POSId             string  `json:"point_of_sale_id" validate:"required"`
	WarehouseRackId   string  `json:"warehouse_rack_id" validate:"required"`
	WarehouseRackName string  `json:"warehouse_rack_name" validate:"required"`
	ItemVariantId     string  `json:"item_variant_id" validate:"required"`
	ItemVariantName   string  `json:"item_variant_name" validate:"required"`
	ItemUnitId        string  `json:"item_unit_id" validate:"required"`
	ItemUnitName      string  `json:"item_unit_name" validate:"required"`
	ItemUnitValue     string  `json:"item_unit_value" validate:"required"`
	ItemCode          string  `json:"item_code" validate:"required"`
	ItemName          string  `json:"item_name" validate:"required"`
	Batch             *string `json:"batch" validate:"required"`
	ExpiredDate       *string `json:"expired_date" validate:"required"`
	ItemBarcodeId     string  `json:"item_barcode_id" validate:"required"`
	Amount            string  `json:"amount" validate:"required"`
	Price             string  `json:"price" validate:"required"`
}

type POSItemRequest struct {
	WarehouseRackId string `json:"warehouse_rack_id" validate:"required"`
	ItemVariantId   string `json:"item_variant_id" validate:"required"`
	ItemUnitId      string `json:"item_unit_id" validate:"required"`
	ItemUnitValue   string `json:"item_unit_value" validate:"required"`
	Batch           string `json:"batch" validate:"required"`
	ExpiredDate     string `json:"expired_date" validate:"required"`
	ItemBarcodeId   string `json:"item_barcode_id" validate:"required"`
	Amount          string `json:"amount" validate:"required"`
	Price           string `json:"price" validate:"required"`
}

type UpsertPOSRequest struct {
	Id                 string           `json:"id"`
	CompanyId          string           `json:"company_id" validate:"required"`
	BranchId           string           `json:"branch_id" validate:"required"`
	WarehouseId        string           `json:"warehouse_id" validate:"required"`
	FormNumber         string           `json:"form_number"`
	TransactionDate    string           `json:"transaction_date" validate:"required"`
	ContactBookId      string           `json:"contact_book_id" validate:"required"`
	SecondaryCompanyId string           `json:"secondary_company_id" validate:"required"`
	KonekinId          string           `json:"konekin_id" validate:"required"`
	CurrencyCode       string           `json:"currency_code" validate:"required"`
	POSPaymentMethodId string           `json:"pos_payment_method_id" validate:"required"`
	TotalItems         string           `json:"total_items" validate:"required"`
	Total              string           `json:"total" validate:"required"`
	POSItems           []POSItemRequest `json:"pos_items" validate:"required"`
}

type UpsertPOSResponse struct {
	Message string `json:"message"`
}

type DeletePOSRequest struct {
	Id string `json:"id" validate:"required"`
}

type DeletePOSResponse struct {
	Message string `json:"message"`
}

type GetPOSRequest struct {
	CompanyId string `json:"company_id" validate:"required"`
	BranchId  string `json:"branch_id" validate:"required"`
	StartDate string `json:"start_date" validate:"required"`
	EndDate   string `json:"end_date" validate:"required"`
}

type GetPOSResponse struct {
	POS []POS `json:"pos"`
}

type GetPOSItemsRequest struct {
	POSId string `json:"pos_id" validate:"required"`
}

type GetPOSItemsResponse struct {
	POSItems []POSItem `json:"pos_items"`
}

type POSUserSetting struct {
	UserId            string `json:"user_id" validate:"required"`
	BranchId          string `json:"branch_id" validate:"required"`
	WarehouseId       string `json:"warehouse_id" validate:"required"`
	WarehouseName     string `json:"warehouse_name" validate:"required"`
	WarehouseRackId   string `json:"warehouse_rack_id" validate:"required"`
	WarehouseRackName string `json:"warehouse_rack_name" validate:"required"`
}

type GetPOSUserSettingRequest struct {
	UserId   string `json:"user_id" validate:"required"`
	BranchId string `json:"branch_id" validate:"required"`
}

type GetPOSUserSettingResponse struct {
	POSUserSetting
}

type UpdatePOSUserSettingRequest struct {
	UserId          string `json:"user_id" validate:"required"`
	BranchId        string `json:"branch_id" validate:"required"`
	WarehouseId     string `json:"warehouse_id" validate:"required"`
	WarehouseRackId string `json:"warehouse_rack_id" validate:"required"`
}

type UpdatePOSUserSettingResponse struct {
	POSUserSetting
}

type UpdatePOSCOASettingRequest struct {
	BranchId        string   `json:"branch_id" validate:"required"`
	ChartOfAccounts []string `json:"chart_of_accounts" validate:"required"`
}

type UpdatePOSCOASettingResponse struct {
	Message string `json:"message"`
}

type POSCOA struct {
	ChartOfAccountId   string `json:"chart_of_account_id" validate:"required"`
	ChartOfAccountName string `json:"chart_of_account_name" validate:"required"`
}

type GetPOSCOASettingRequest struct {
	BranchId string `json:"branch_id" validate:"required"`
}

type GetPOSCOASettingResponse struct {
	POSCOAs []POSCOA `json:"pos_coas"`
}

type UpdatePOSCustomerSettingRequest struct {
	BranchId  string   `json:"branch_id" validate:"required"`
	Customers []string `json:"customers" validate:"required"`
}

type UpdatePOSCustomerSettingResponse struct {
	Message string `json:"message"`
}

type POSCustomer struct {
	ContactBookId      string `json:"contact_book_id" validate:"required"`
	ContactBookName    string `json:"contact_book_name" validate:"required"`
	SecondaryCompanyId string `json:"secondary_company_id" validate:"required"`
	KonekinId          string `json:"konekin_id" validate:"required"`
}

type GetPOSCustomerSettingRequest struct {
	BranchId string `json:"branch_id" validate:"required"`
}

type GetPOSCustomerSettingResponse struct {
	POSCustomers []POSCustomer `json:"pos_customers"`
}

type POSPaymentMethod struct {
	Id                 string `json:"id" validate:"required"`
	CompanyId          string `json:"company_id" validate:"required"`
	ChartOfAccountId   string `json:"chart_of_account_id" validate:"required"`
	ChartOfAccountName string `json:"chart_of_account_name" validate:"required"`
	Name               string `json:"name" validate:"required"`
}

type UpsertPOSPaymentMethodRequest struct {
	Id               string `json:"id"`
	CompanyId        string `json:"company_id" validate:"required"`
	ChartOfAccountId string `json:"chart_of_account_id" validate:"required"`
	Name             string `json:"name" validate:"required"`
}

type UpsertPOSPaymentMethodResponse struct {
	Message string `json:"message"`
}

type DeletePOSPaymentMethodRequest struct {
	Id string `json:"id" validate:"required"`
}

type DeletePOSPaymentMethodResponse struct {
	Message string `json:"message"`
}

type GetPOSPaymentMethodRequest struct {
	CompanyId string `json:"company_id" validate:"required"`
	Keyword   string `json:"keyword"`
}

type GetPOSPaymentMethodResponse struct {
	POSPaymentMethods []POSPaymentMethod `json:"pos_payment_methods"`
}
