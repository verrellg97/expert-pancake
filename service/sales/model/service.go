package model

import "net/http"

type SalesService interface {
	UpsertPOS(w http.ResponseWriter, r *http.Request) error
	DeletePOS(w http.ResponseWriter, r *http.Request) error
}

type POS struct {
	Id                 string `json:"id" validate:"required"`
	CompanyId          string `json:"company_id" validate:"required"`
	BranchId           string `json:"branch_id" validate:"required"`
	WarehouseId        string `json:"warehouse_id" validate:"required"`
	WarehouseName      string `json:"warehouse_name" validate:"required"`
	FormNumber         string `json:"form_number" validate:"required"`
	TransactionDate    string `json:"transaction_date" validate:"required"`
	ContactBookId      string `json:"contact_book_id" validate:"required"`
	ContactBookName    string `json:"contact_book_name" validate:"required"`
	SecondaryCompanyId string `json:"secondary_company_id" validate:"required"`
	KonekinId          string `json:"konekin_id" validate:"required"`
	CurrencyCode       string `json:"currency_code" validate:"required"`
	ChartOfAccountId   string `json:"chart_of_account_id" validate:"required"`
	ChartOfAccountName string `json:"chart_of_account_name" validate:"required"`
	TotalItems         string `json:"total_items" validate:"required"`
	Total              string `json:"total" validate:"required"`
}

type POSItem struct {
	DetailId          string `json:"detail_id" validate:"required"`
	POSId             string `json:"point_of_sale_id" validate:"required"`
	WarehouseRackId   string `json:"warehouse_rack_id" validate:"required"`
	WarehouseRackName string `json:"warehouse_rack_name" validate:"required"`
	ItemVariantId     string `json:"item_variant_id" validate:"required"`
	ItemVariantName   string `json:"item_variant_name" validate:"required"`
	ItemUnitId        string `json:"item_unit_id" validate:"required"`
	ItemUnitValue     string `json:"item_unit_value" validate:"required"`
	ItemCode          string `json:"item_code" validate:"required"`
	ItemName          string `json:"item_name" validate:"required"`
	Batch             string `json:"batch" validate:"required"`
	ExpiredDate       string `json:"expired_date" validate:"required"`
	ItemBarcodeId     string `json:"item_barcode_id" validate:"required"`
	Amount            string `json:"amount" validate:"required"`
	Price             string `json:"price" validate:"required"`
}

type POSRequest struct {
	Id                 string `json:"id" validate:"required"`
	CompanyId          string `json:"company_id" validate:"required"`
	BranchId           string `json:"branch_id" validate:"required"`
	WarehouseId        string `json:"warehouse_id" validate:"required"`
	FormNumber         string `json:"form_number" validate:"required"`
	TransactionDate    string `json:"transaction_date" validate:"required"`
	ContactBookId      string `json:"contact_book_id" validate:"required"`
	SecondaryCompanyId string `json:"secondary_company_id" validate:"required"`
	KonekinId          string `json:"konekin_id" validate:"required"`
	CurrencyCode       string `json:"currency_code" validate:"required"`
	ChartOfAccountId   string `json:"chart_of_account_id" validate:"required"`
	TotalItems         string `json:"total_items" validate:"required"`
	Total              string `json:"total" validate:"required"`
}

type POSItemRequest struct {
	POSId           string `json:"point_of_sale_id" validate:"required"`
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
	POS      POSRequest       `json:"pos" validate:"required"`
	POSItems []POSItemRequest `json:"pos_items" validate:"required"`
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

