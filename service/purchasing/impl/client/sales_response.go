package client

type SalesOrder struct {
	TransactionId      string `json:"transaction_id" validate:"required"`
	CompanyId          string `json:"company_id" validate:"required"`
	BranchId           string `json:"branch_id" validate:"required"`
	FormNumber         string `json:"form_number" validate:"required"`
	TransactionDate    string `json:"transaction_date" validate:"required"`
	ContactBookId      string `json:"contact_book_id" validate:"required"`
	SecondaryCompanyId string `json:"secondary_company_id" validate:"required"`
	KonekinId          string `json:"konekin_id" validate:"required"`
	CustomerName       string `json:"customer_name" validate:"required"`
	SecondaryBranchId  string `json:"secondary_branch_id" validate:"required"`
	CurrencyCode       string `json:"currency_code" validate:"required"`
	TotalItems         string `json:"total_items" validate:"required"`
	Total              string `json:"total" validate:"required"`
	Status             string `json:"status" validate:"required"`
}

type UpsertSalesOrderResponse struct {
	Result struct {
		SalesOrder
	} `json:"result"`
}

type SalesOrderItem struct {
	DetailId               string `json:"detail_id" validate:"required"`
	PurchaseOrderItemId    string `json:"purchase_order_item_id" validate:"required"`
	SalesOrderId           string `json:"sales_order_id" validate:"required"`
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

type SalesOrderItemsRequest struct {
	PurchaseOrderItemId    string `json:"purchase_order_item_id"`
	PrimaryItemVariantId   string `json:"primary_item_variant_id" validate:"required"`
	SecondaryItemVariantId string `json:"secondary_item_variant_id" validate:"required"`
	PrimaryItemUnitId      string `json:"primary_item_unit_id" validate:"required"`
	SecondaryItemUnitId    string `json:"secondary_item_unit_id" validate:"required"`
	PrimaryItemUnitValue   string `json:"primary_item_unit_value" validate:"required"`
	SecondaryItemUnitValue string `json:"secondary_item_unit_value" validate:"required"`
	Amount                 string `json:"amount" validate:"required"`
	Price                  string `json:"price" validate:"required"`
}

type UpdateSalesOrderItemsResponse struct {
	Result struct {
		SalesOrderItems []SalesOrderItem `json:"sales_order_items"`
	} `json:"result"`
}
