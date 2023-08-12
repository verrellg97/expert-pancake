package client

type GetItemVariantsRequest struct {
	Id      string `json:"id"`
	ItemId  string `json:"item_id"`
	Keyword string `json:"keyword"`
}

type GetItemUnitsRequest struct {
	Id      string `json:"id"`
	ItemId  string `json:"item_id" validate:"required"`
	Keyword string `json:"keyword"`
}

type InsertStockMovementRequest struct {
	TransactionId        string `json:"transaction_id" validate:"required"`
	CompanyId            string `json:"company_id" validate:"required"`
	BranchId             string `json:"branch_id" validate:"required"`
	TransactionCode      string `json:"transaction_code" validate:"required"`
	TransactionDate      string `json:"transaction_date" validate:"required"`
	TransactionReference string `json:"transaction_reference" validate:"required"`
	DetailTransactionId  string `json:"detail_transaction_id" validate:"required"`
	WarehouseId          string `json:"warehouse_id" validate:"required"`
	WarehouseRackId      string `json:"warehouse_rack_id" validate:"required"`
	VariantId            string `json:"variant_id" validate:"required"`
	ItemBarcodeId        string `json:"item_barcode_id" validate:"required"`
	Amount               string `json:"amount" validate:"required"`
}

type DeleteStockMovementRequest struct {
	TransactionId        string `json:"transaction_id" validate:"required"`
	TransactionReference string `json:"transaction_reference" validate:"required"`
}
