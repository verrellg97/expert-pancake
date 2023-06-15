package client

type UpsertSalesOrderRequest struct {
	Id                    string   `json:"id"`
	PurchaseOrderId       string   `json:"purchase_order_id"`
	PurchaseOrderBranchId string   `json:"purchase_order_branch_id"`
	CompanyId             string   `json:"company_id" validate:"required"`
	BranchId              string   `json:"branch_id"`
	TransactionDate       string   `json:"transaction_date" validate:"required"`
	ContactBookId         string   `json:"contact_book_id" validate:"required"`
	SecondaryCompanyId    string   `json:"secondary_company_id"`
	KonekinId             string   `json:"konekin_id"`
	CurrencyCode          string   `json:"currency_code" validate:"required"`
	IsAllBranches         bool     `json:"is_all_branches"`
	Branches              []string `json:"branches"`
}

type UpdateSalesOrderItemsRequest struct {
	SalesOrderId    string                   `json:"sales_order_id" validate:"required"`
	SalesOrderItems []SalesOrderItemsRequest `json:"sales_order_items" validate:"required"`
}
