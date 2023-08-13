package client

type UpsertReceiptOrderItemsRequest struct {
	PurchaseOrderItemId    string `json:"purchase_order_item_id"`
	SalesOrderItemId       string `json:"sales_order_item_id"`
	DeliveryOrderItemId    string `json:"delivery_order_item_id"`
	PrimaryItemVariantId   string `json:"primary_item_variant_id" validate:"required"`
	WarehouseRackId        string `json:"warehouse_rack_id"`
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
	WarehouseId                    string                           `json:"warehouse_id"`
	CompanyId                      string                           `json:"company_id" validate:"required"`
	BranchId                       string                           `json:"branch_id" validate:"required"`
	TransactionDate                string                           `json:"transaction_date" validate:"required"`
	ContactBookId                  string                           `json:"contact_book_id" validate:"required"`
	SecondaryCompanyId             string                           `json:"secondary_company_id"`
	KonekinId                      string                           `json:"konekin_id"`
	TotalItems                     string                           `json:"total_items"`
	UpsertReceiptOrderItemsRequest []UpsertReceiptOrderItemsRequest `json:"receipt_order_items" validate:"required"`
}
