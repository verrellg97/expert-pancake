package client

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

type UpsertReceiptOrderResponse struct {
	Result struct {
		ReceiptOrder      ReceiptOrder       `json:"receipt_order" validate:"required"`
		ReceiptOrderItems []ReceiptOrderItem `json:"items" validate:"required"`
	} `json:"result"`
}
