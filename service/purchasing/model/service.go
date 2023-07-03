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

	GetCheckPurchaseOrders(w http.ResponseWriter, r *http.Request) error

	UpdatePurchaseOrderStatus(w http.ResponseWriter, r *http.Request) error

	UpsertReceiptOrder(w http.ResponseWriter, r *http.Request) error
	GetReceiptOrders(w http.ResponseWriter, r *http.Request) error
	UpdateReceiptOrderItems(w http.ResponseWriter, r *http.Request) error
	GetReceiptOrderItems(w http.ResponseWriter, r *http.Request) error
	DeleteReceiptOrder(w http.ResponseWriter, r *http.Request) error

	UpsertPurchaseInvoice(w http.ResponseWriter, r *http.Request) error
	GetPurchaseInvoices(w http.ResponseWriter, r *http.Request) error
	GetPurchaseInvoiceItems(w http.ResponseWriter, r *http.Request) error
}

type PurchaseOrder struct {
	TransactionId             string `json:"transaction_id" validate:"required"`
	CompanyId                 string `json:"company_id" validate:"required"`
	BranchId                  string `json:"branch_id" validate:"required"`
	FormNumber                string `json:"form_number" validate:"required"`
	TransactionDate           string `json:"transaction_date" validate:"required"`
	ContactBookId             string `json:"contact_book_id" validate:"required"`
	SecondaryCompanyId        string `json:"secondary_company_id" validate:"required"`
	KonekinId                 string `json:"konekin_id" validate:"required"`
	SupplierName              string `json:"supplier_name" validate:"required"`
	CurrencyCode              string `json:"currency_code" validate:"required"`
	ShippingDate              string `json:"shipping_date" validate:"required"`
	ReceivingWarehouseId      string `json:"receiving_warehouse_id" validate:"required"`
	ReceivingWarehouseName    string `json:"receiving_warehouse_name" validate:"required"`
	ReceivingWarehouseAddress string `json:"receiving_warehouse_address" validate:"required"`
	TotalItems                string `json:"total_items" validate:"required"`
	Total                     string `json:"total" validate:"required"`
	Status                    string `json:"status" validate:"required"`
}

type UpsertPurchaseOrderRequest struct {
	Id                   string `json:"id"`
	CompanyId            string `json:"company_id" validate:"required"`
	BranchId             string `json:"branch_id" validate:"required"`
	TransactionDate      string `json:"transaction_date" validate:"required"`
	ContactBookId        string `json:"contact_book_id" validate:"required"`
	SecondaryCompanyId   string `json:"secondary_company_id"`
	KonekinId            string `json:"konekin_id"`
	CurrencyCode         string `json:"currency_code" validate:"required"`
	ShippingDate         string `json:"shipping_date" validate:"required"`
	ReceivingWarehouseId string `json:"receiving_warehouse_id" validate:"required"`
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

type GetCheckPurchaseOrdersRequest struct {
	CompanyId string `json:"company_id" validate:"required"`
}

type GetCheckPurchaseOrdersResponse struct {
	Status bool `json:"status" validate:"required"`
}

type UpdatePurchaseOrderStatusRequest struct {
	PurchaseOrderId string `json:"purchase_order_id" validate:"required"`
	Status          string `json:"status" validate:"required"`
}

type UpdatePurchaseOrderStatusResponse struct {
	Message string `json:"message" validate:"required"`
}

type ReceiptOrder struct {
	Id                 string `json:"id" validate:"required"`
	DeliveryOrderId    string `json:"delivery_order_id" validate:"required"`
	WarehouseId        string `json:"warehouse_id" validate:"required"`
	WarehouseName      string `json:"warehouse_name" validate:"required"`
	CompanyId          string `json:"company_id" validate:"required"`
	BranchId           string `json:"branch_id" validate:"required"`
	FormNumber         string `json:"form_number" validate:"required"`
	TransactionDate    string `json:"transaction_date" validate:"required"`
	ContactBookId      string `json:"contact_book_id" validate:"required"`
	SupplierName       string `json:"supplier_name" validate:"required"`
	SecondaryCompanyId string `json:"secondary_company_id" validate:"required"`
	KonekinId          string `json:"konekin_id" validate:"required"`
	TotalItems         string `json:"total_items" validate:"required"`
	Status             string `json:"status" validate:"required"`
}

type ReceiptOrderItem struct {
	DetailId               string  `json:"detail_id" validate:"required"`
	PurchaseOrderItemId    string  `json:"purchase_order_item_id" validate:"required"`
	SalesOrderItemId       string  `json:"sales_order_item_id" validate:"required"`
	DeliveryOrderItemId    string  `json:"delivery_order_item_id" validate:"required"`
	ReceiptOrderId         string  `json:"receipt_order_id" validate:"required"`
	PrimaryItemVariantId   string  `json:"primary_item_variant_id" validate:"required"`
	ItemCode               string  `json:"item_code" validate:"required"`
	ItemName               string  `json:"item_name" validate:"required"`
	ItemVariantName        string  `json:"item_variant_name" validate:"required"`
	WarehouseRackId        string  `json:"warehouse_rack_id" validate:"required"`
	WarehouseRackName      string  `json:"warehouse_rack_name" validate:"required"`
	Batch                  *string `json:"batch" validate:"required"`
	ExpiredDate            *string `json:"expired_date" validate:"required"`
	ItemBarcodeId          string  `json:"item_barcode_id" validate:"required"`
	SecondaryItemVariantId string  `json:"secondary_item_variant_id" validate:"required"`
	PrimaryItemUnitId      string  `json:"primary_item_unit_id" validate:"required"`
	ItemUnitName           string  `json:"item_unit_name" validate:"required"`
	SecondaryItemUnitId    string  `json:"secondary_item_unit_id" validate:"required"`
	PrimaryItemUnitValue   string  `json:"primary_item_unit_value" validate:"required"`
	SecondaryItemUnitValue string  `json:"secondary_item_unit_value" validate:"required"`
	Amount                 string  `json:"amount" validate:"required"`
}

type UpsertReceiptOrderItemsRequest struct {
	PurchaseOrderItemId    string `json:"purchase_order_item_id"`
	SalesOrderItemId       string `json:"sales_order_item_id"`
	DeliveryOrderItemId    string `json:"delivery_order_item_id"`
	PrimaryItemVariantId   string `json:"primary_item_variant_id" validate:"required"`
	WarehouseRackId        string `json:"warehouse_rack_id" validate:"required"`
	Batch                  string `json:"batch"`
	ExpiredDate            string `json:"expired_date"`
	ItemBarcodeId          string `json:"item_barcode_id"`
	SecondaryItemVariantId string `json:"secondary_item_variant_id"`
	PrimaryItemUnitId      string `json:"primary_item_unit_id" validate:"required"`
	SecondaryItemUnitId    string `json:"secondary_item_unit_id"`
	PrimaryItemUnitValue   string `json:"primary_item_unit_value" validate:"required"`
	SecondaryItemUnitValue string `json:"secondary_item_unit_value"`
	Amount                 string `json:"amount" validate:"required"`
}
type UpsertReceiptOrderRequest struct {
	Id                             string                           `json:"id"`
	DeliveryOrderId                string                           `json:"delivery_order_id"`
	WarehouseId                    string                           `json:"warehouse_id" validate:"required"`
	CompanyId                      string                           `json:"company_id" validate:"required"`
	BranchId                       string                           `json:"branch_id" validate:"required"`
	TransactionDate                string                           `json:"transaction_date" validate:"required"`
	ContactBookId                  string                           `json:"contact_book_id" validate:"required"`
	SecondaryCompanyId             string                           `json:"secondary_company_id"`
	KonekinId                      string                           `json:"konekin_id"`
	TotalItems                     string                           `json:"total_items"`
	UpsertReceiptOrderItemsRequest []UpsertReceiptOrderItemsRequest `json:"receipt_order_items" validate:"required"`
}

type UpsertReceiptOrderResponse struct {
	Message string `json:"message" validate:"required"`
}

type GetReceiptOrdersRequest struct {
	CompanyId string `json:"company_id" validate:"required"`
	BranchId  string `json:"branch_id" validate:"required"`
	StartDate string `json:"start_date" validate:"required"`
	EndDate   string `json:"end_date" validate:"required"`
}

type GetReceiptOrdersResponse struct {
	ReceiptOrders []ReceiptOrder
}

type ReceiptOrderItemsRequest struct {
	PurchaseOrderItemId    string `json:"purchase_order_item_id" validate:"required"`
	SalesOrderItemId       string `json:"sales_order_item_id"`
	DeliveryOrderItemId    string `json:"delivery_order_item_id"`
	PrimaryItemVariantId   string `json:"primary_item_variant_id" validate:"required"`
	WarehouseRackId        string `json:"warehouse_rack_id" validate:"required"`
	Batch                  string `json:"batch" validate:"required"`
	ExpiredDate            string `json:"expired_date" validate:"required"`
	ItemBarcodeId          string `json:"item_barcode_id" validate:"required"`
	SecondaryItemVariantId string `json:"secondary_item_variant_id"`
	PrimaryItemUnitId      string `json:"primary_item_unit_id" validate:"required"`
	SecondaryItemUnitId    string `json:"secondary_item_unit_id"`
	PrimaryItemUnitValue   string `json:"primary_item_unit_value" validate:"required"`
	SecondaryItemUnitValue string `json:"secondary_item_unit_value"`
	Amount                 string `json:"amount" validate:"required"`
}

type UpdateReceiptOrderItemsRequest struct {
	ReceiptOrderId    string                     `json:"receipt_order_id" validate:"required"`
	ReceiptOrderItems []ReceiptOrderItemsRequest `json:"receipt_order_items" validate:"required"`
}

type UpdateReceiptOrderItemsResponse struct {
	Message string `json:"message" validate:"required"`
}

type GetReceiptOrderItemsRequest struct {
	ReceiptOrderId string `json:"receipt_order_id" validate:"required"`
}

type GetReceiptOrderItemsResponse struct {
	ReceiptOrderItems []ReceiptOrderItem
}

type DeleteReceiptOrderRequest struct {
	ReceiptOrderId string `json:"receipt_order_id" validate:"required"`
}

type DeleteReceiptOrderResponse struct {
	Message string `json:"message" validate:"required"`
}
