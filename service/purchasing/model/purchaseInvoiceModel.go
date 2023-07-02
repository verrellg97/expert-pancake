package model

const (
	UpsertPurchaseInvoiceError   = 800401
	GetPurchaseInvoicesError     = 800402
	GetPurchaseInvoiceItemsError = 800403
)

type UpsertPurchaseInvoiceItemRequest struct {
	PurchaseOrderItemId    string `json:"purchase_order_item_id"`
	SalesOrderItemId       string `json:"sales_order_item_id"`
	SalesInvoiceItemId     string `json:"sales_invoice_item_id"`
	ReceiptOrderItemId     string `json:"receipt_order_item_id"`
	PrimaryItemVariantId   string `json:"primary_item_variant_id" validate:"required"`
	SecondaryItemVariantId string `json:"secondary_item_variant_id"`
	PrimaryItemUnitId      string `json:"primary_item_unit_id" validate:"required"`
	SecondaryItemUnitId    string `json:"secondary_item_unit_id"`
	PrimaryItemUnitValue   string `json:"primary_item_unit_value" validate:"required"`
	SecondaryItemUnitValue string `json:"secondary_item_unit_value"`
	Amount                 string `json:"amount" validate:"required"`
	Price                  string `json:"price" validate:"required"`
}

type UpsertPurchaseInvoiceRequest struct {
	Id                   string                             `json:"id" validate:"required"`
	SalesInvoiceId       string                             `json:"sales_invoice_id"`
	ReceiptOrderId       string                             `json:"receipt_order_id"`
	CompanyId            string                             `json:"company_id" validate:"required"`
	BranchId             string                             `json:"branch_id" validate:"required"`
	TransactionDate      string                             `json:"transaction_date" validate:"required"`
	ContactBookId        string                             `json:"contact_book_id" validate:"required"`
	SecondaryCompanyId   string                             `json:"secondary_company_id"`
	KonekinId            string                             `json:"konekin_id"`
	CurrencyCode         string                             `json:"currency_code" validate:"required"`
	TotalItems           string                             `json:"total_items" validate:"required"`
	Total                string                             `json:"total" validate:"required"`
	Status               string                             `json:"status" validate:"required"`
	PurchaseInvoiceItems []UpsertPurchaseInvoiceItemRequest `json:"purchase_invoice_items" validate:"required"`
}

type UpsertPurchaseInvoiceResponse struct {
	Message string `json:"message"`
}

type GetPurchaseInvoicesRequest struct {
	CompanyId string `json:"company_id" validate:"required"`
	BranchId  string `json:"branch_id" validate:"required"`
	StartDate string `json:"start_date" validate:"required"`
	EndDate   string `json:"end_date" validate:"required"`
}

type PurchaseInvoice struct {
	Id                          string `json:"id" validate:"required"`
	FormNumber                  string `json:"form_number" validate:"required"`
	SalesInvoiceId              string `json:"sales_invoice_id" validate:"required"`
	ReceiptOrderId              string `json:"receipt_order_id" validate:"required"`
	ReceiptOrderFormNumber      string `json:"receipt_order_form_number" validate:"required"`
	ReceiptOrderTransactionDate string `json:"receipt_order_transaction_date" validate:"required"`
	WarehouseId                 string `json:"warehouse_id" validate:"required"`
	WarehouseName               string `json:"warehouse_name" validate:"required"`
	CompanyId                   string `json:"company_id" validate:"required"`
	BranchId                    string `json:"branch_id" validate:"required"`
	TransactionDate             string `json:"transaction_date" validate:"required"`
	ContactBookId               string `json:"contact_book_id" validate:"required"`
	SecondaryCompanyId          string `json:"secondary_company_id" validate:"required"`
	KonekinId                   string `json:"konekin_id" validate:"required"`
	SupplierName                string `json:"supplier_name" validate:"required"`
	CurrencyCode                string `json:"currency_code" validate:"required"`
	TotalItems                  string `json:"total_items" validate:"required"`
	Total                       string `json:"total" validate:"required"`
	Status                      string `json:"status" validate:"required"`
}

type PurchaseInvoiceItem struct {
	Id                     string  `json:"id" validate:"required"`
	PurchaseOrderItemId    string  `json:"purchase_order_item_id" validate:"required"`
	SalesOrderItemId       string  `json:"sales_order_item_id" validate:"required"`
	SalesInvoiceItemId     string  `json:"sales_invoice_item_id" validate:"required"`
	ReceiptOrderItemId     string  `json:"receipt_order_item_id" validate:"required"`
	PurchaseInvoiceId      string  `json:"purchase_invoice_id" validate:"required"`
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
	Price                  string  `json:"price" validate:"required"`
}

type GetPurchaseInvoicesResponse struct {
	PurchaseInvoices []PurchaseInvoice
}

type GetPurchaseInvoiceItemsRequest struct {
	PurchaseInvoiceId string `json:"purchase_invoice_id" validate:"required"`
}

type GetPurchaseInvoiceItemsResponse struct {
	PurchaseInvoiceItems []PurchaseInvoiceItem
}
