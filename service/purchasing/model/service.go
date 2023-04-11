package model

import "net/http"

type PurchasingService interface {
	UpsertPurchaseOrder(w http.ResponseWriter, r *http.Request) error
}

type PurchaseOrder struct {
	TransactionId      string `json:"transaction_id" validate:"required"`
	CompanyId          string `json:"company_id" validate:"required"`
	BranchId           string `json:"branch_id" validate:"required"`
	FormNumber         string `json:"form_number" validate:"required"`
	TransactionDate    string `json:"transaction_date" validate:"required"`
	ContactBookId      string `json:"contact_book_id" validate:"required"`
	SecondaryCompanyId string `json:"secondary_company_id" validate:"required"`
	KonekinId          string `json:"konekin_id" validate:"required"`
	SupplierName       string `json:"supplier_name" validate:"required"`
	CurrencyCode       string `json:"currency_code" validate:"required"`
	TotalItems         string `json:"total_items" validate:"required"`
	Total              string `json:"total" validate:"required"`
	Status             string `json:"status" validate:"required"`
}

type UpsertPurchaseOrderRequest struct {
	Id                 string `json:"id"`
	CompanyId          string `json:"company_id" validate:"required"`
	BranchId           string `json:"branch_id" validate:"required"`
	TransactionDate    string `json:"transaction_date" validate:"required"`
	ContactBookId      string `json:"contact_book_id" validate:"required"`
	SecondaryCompanyId string `json:"secondary_company_id"`
	KonekinId          string `json:"konekin_id"`
	CurrencyCode       string `json:"currency_code" validate:"required"`
}

type UpsertPurchaseOrderResponse struct {
	PurchaseOrder
}
