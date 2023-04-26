package model

import "net/http"

type PurchasingService interface {
	UpsertPurchaseOrder(w http.ResponseWriter, r *http.Request) error
	UpsertPurchaseOrderItem(w http.ResponseWriter, r *http.Request) error
	UpdatePurchaseOrderItems(w http.ResponseWriter, r *http.Request) error
	GetPurchaseOrders(w http.ResponseWriter, r *http.Request) error
	GetPurchaseOrderItems(w http.ResponseWriter, r *http.Request) error
	UpdatePurchaseSetting(w http.ResponseWriter, r *http.Request) error
	GetPurchaseSetting(w http.ResponseWriter, r *http.Request) error
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

type PurchaseOrderItem struct {
	DetailId               string `json:"detail_id" validate:"required"`
	PurchaseOrderId        string `json:"purchase_order_id" validate:"required"`
	PrimaryItemVariantId   string `json:"primary_item_variant_id" validate:"required"`
	ItemCode               string `json:"item_code" validate:"required"`
	ItemName               string `json:"item_name" validate:"required"`
	ItemVariantName        string `json:"item_variant_name" validate:"required"`
	SecondaryItemVariantId string `json:"secondary_item_variant_id" validate:"required"`
	PrimaryItemUnitId      string `json:"primary_item_unit_id" validate:"required"`
	ItemUnitName           string `json:"item_unit_name" validate:"required"`
	SecondaryItemUnitId    string `json:"secondary_item_unit_id" validate:"required"`
	PrimaryItemUnitValue   string `json:"primary_item_unit_value" validate:"required"`
	SecondaryItemUnitValue string `json:"secondary_item_unit_value" validate:"required"`
	Amount                 string `json:"amount" validate:"required"`
	Price                  string `json:"price" validate:"required"`
}

type UpsertPurchaseOrderItemRequest struct {
	Id                     string `json:"id"`
	PurchaseOrderId        string `json:"purchase_order_id" validate:"required"`
	PrimaryItemVariantId   string `json:"primary_item_variant_id" validate:"required"`
	SecondaryItemVariantId string `json:"secondary_item_variant_id" validate:"required"`
	PrimaryItemUnitId      string `json:"primary_item_unit_id" validate:"required"`
	SecondaryItemUnitId    string `json:"secondary_item_unit_id" validate:"required"`
	PrimaryItemUnitValue   string `json:"primary_item_unit_value" validate:"required"`
	SecondaryItemUnitValue string `json:"secondary_item_unit_value" validate:"required"`
	Amount                 string `json:"amount" validate:"required"`
	Price                  string `json:"price" validate:"required"`
}

type UpsertPurchaseOrderItemResponse struct {
	PurchaseOrderItem
}

type PurchaseOrderItemsRequest struct {
	PrimaryItemVariantId   string `json:"primary_item_variant_id" validate:"required"`
	SecondaryItemVariantId string `json:"secondary_item_variant_id" validate:"required"`
	PrimaryItemUnitId      string `json:"primary_item_unit_id" validate:"required"`
	SecondaryItemUnitId    string `json:"secondary_item_unit_id" validate:"required"`
	PrimaryItemUnitValue   string `json:"primary_item_unit_value" validate:"required"`
	SecondaryItemUnitValue string `json:"secondary_item_unit_value" validate:"required"`
	Amount                 string `json:"amount" validate:"required"`
	Price                  string `json:"price" validate:"required"`
}

type UpdatePurchaseOrderItemsRequest struct {
	PurchaseOrderId    string                      `json:"purchase_order_id" validate:"required"`
	PurchaseOrderItems []PurchaseOrderItemsRequest `json:"purchase_order_items" validate:"required"`
}

type UpdatePurchaseOrderItemsResponse struct {
	PurchaseOrderItems []PurchaseOrderItem `json:"purchase_order_items"`
}

type GetPurchaseOrdersRequest struct {
	CompanyId string `json:"company_id" validate:"required"`
	BranchId  string `json:"branch_id" validate:"required"`
	StartDate string `json:"start_date" validate:"required"`
	EndDate   string `json:"end_date" validate:"required"`
}

type GetPurchaseOrdersResponse struct {
	PurchaseOrders []PurchaseOrder `json:"purchase_orders"`
}

type GetPurchaseOrderItemsRequest struct {
	PurchaseOrderId string `json:"purchase_order_id" validate:"required"`
}

type GetPurchaseOrderItemsResponse struct {
	PurchaseOrderItems []PurchaseOrderItem `json:"purchase_order_items"`
}

type PurchaseSetting struct {
	CompanyId          string `json:"company_id" validate:"required"`
	IsAutoApproveOrder bool   `json:"is_auto_approve_order" validate:"required"`
}

type UpdatePurchaseSettingRequest struct {
	CompanyId          string `json:"company_id" validate:"required"`
	IsAutoApproveOrder bool   `json:"is_auto_approve_order"`
}

type UpdatePurchaseSettingResponse struct {
	PurchaseSetting PurchaseSetting `json:"purchase_setting"`
}

type GetPurchaseSettingRequest struct {
	CompanyId string `json:"company_id" validate:"required"`
}

type GetPurchaseSettingResponse struct {
	PurchaseSetting PurchaseSetting `json:"purchase_setting"`
}
