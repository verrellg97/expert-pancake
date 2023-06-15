// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.17.2

package db

import (
	"database/sql"
	"time"
)

type SalesDeliveryOrder struct {
	ID                 string       `db:"id"`
	ReceiptOrderID     string       `db:"receipt_order_id"`
	CompanyID          string       `db:"company_id"`
	BranchID           string       `db:"branch_id"`
	FormNumber         string       `db:"form_number"`
	TransactionDate    time.Time    `db:"transaction_date"`
	ContactBookID      string       `db:"contact_book_id"`
	SecondaryCompanyID string       `db:"secondary_company_id"`
	KonekinID          string       `db:"konekin_id"`
	SecondaryBranchID  string       `db:"secondary_branch_id"`
	TotalItems         int64        `db:"total_items"`
	IsDeleted          bool         `db:"is_deleted"`
	Status             string       `db:"status"`
	CreatedAt          sql.NullTime `db:"created_at"`
	UpdatedAt          sql.NullTime `db:"updated_at"`
}

type SalesDeliveryOrderItem struct {
	ID                     string         `db:"id"`
	PurchaseOrderItemID    string         `db:"purchase_order_item_id"`
	SalesOrderItemID       string         `db:"sales_order_item_id"`
	ReceiptOrderItemID     string         `db:"receipt_order_item_id"`
	DeliveryOrderID        string         `db:"delivery_order_id"`
	PrimaryItemVariantID   string         `db:"primary_item_variant_id"`
	WarehouseRackID        string         `db:"warehouse_rack_id"`
	Batch                  sql.NullString `db:"batch"`
	ExpiredDate            sql.NullTime   `db:"expired_date"`
	ItemBarcodeID          string         `db:"item_barcode_id"`
	SecondaryItemVariantID string         `db:"secondary_item_variant_id"`
	PrimaryItemUnitID      string         `db:"primary_item_unit_id"`
	SecondaryItemUnitID    string         `db:"secondary_item_unit_id"`
	PrimaryItemUnitValue   int64          `db:"primary_item_unit_value"`
	SecondaryItemUnitValue int64          `db:"secondary_item_unit_value"`
	Amount                 int64          `db:"amount"`
	IsDeleted              bool           `db:"is_deleted"`
	CreatedAt              sql.NullTime   `db:"created_at"`
	UpdatedAt              sql.NullTime   `db:"updated_at"`
}

type SalesPointOfSale struct {
	ID                 string       `db:"id"`
	CompanyID          string       `db:"company_id"`
	BranchID           string       `db:"branch_id"`
	WarehouseID        string       `db:"warehouse_id"`
	FormNumber         string       `db:"form_number"`
	TransactionDate    time.Time    `db:"transaction_date"`
	ContactBookID      string       `db:"contact_book_id"`
	SecondaryCompanyID string       `db:"secondary_company_id"`
	KonekinID          string       `db:"konekin_id"`
	CurrencyCode       string       `db:"currency_code"`
	PosPaymentMethodID string       `db:"pos_payment_method_id"`
	TotalItems         int64        `db:"total_items"`
	Total              int64        `db:"total"`
	IsDeleted          bool         `db:"is_deleted"`
	CreatedAt          sql.NullTime `db:"created_at"`
	UpdatedAt          sql.NullTime `db:"updated_at"`
}

type SalesPointOfSaleItem struct {
	ID              string         `db:"id"`
	PointOfSaleID   string         `db:"point_of_sale_id"`
	WarehouseRackID string         `db:"warehouse_rack_id"`
	ItemVariantID   string         `db:"item_variant_id"`
	ItemUnitID      string         `db:"item_unit_id"`
	ItemUnitValue   int64          `db:"item_unit_value"`
	Batch           sql.NullString `db:"batch"`
	ExpiredDate     sql.NullTime   `db:"expired_date"`
	ItemBarcodeID   string         `db:"item_barcode_id"`
	Amount          int64          `db:"amount"`
	Price           int64          `db:"price"`
	IsDeleted       bool           `db:"is_deleted"`
	CreatedAt       sql.NullTime   `db:"created_at"`
	UpdatedAt       sql.NullTime   `db:"updated_at"`
}

type SalesPosChartOfAccountSetting struct {
	BranchID         string       `db:"branch_id"`
	ChartOfAccountID string       `db:"chart_of_account_id"`
	CreatedAt        sql.NullTime `db:"created_at"`
	UpdatedAt        sql.NullTime `db:"updated_at"`
}

type SalesPosCustomerSetting struct {
	BranchID      string       `db:"branch_id"`
	ContactBookID string       `db:"contact_book_id"`
	CreatedAt     sql.NullTime `db:"created_at"`
	UpdatedAt     sql.NullTime `db:"updated_at"`
}

type SalesPosPaymentMethod struct {
	ID               string       `db:"id"`
	CompanyID        string       `db:"company_id"`
	ChartOfAccountID string       `db:"chart_of_account_id"`
	Name             string       `db:"name"`
	IsDeleted        bool         `db:"is_deleted"`
	CreatedAt        sql.NullTime `db:"created_at"`
	UpdatedAt        sql.NullTime `db:"updated_at"`
}

type SalesPosUserSetting struct {
	UserID          string       `db:"user_id"`
	BranchID        string       `db:"branch_id"`
	WarehouseID     string       `db:"warehouse_id"`
	WarehouseRackID string       `db:"warehouse_rack_id"`
	CreatedAt       sql.NullTime `db:"created_at"`
	UpdatedAt       sql.NullTime `db:"updated_at"`
}

type SalesSalesOrder struct {
	ID                    string       `db:"id"`
	PurchaseOrderID       string       `db:"purchase_order_id"`
	PurchaseOrderBranchID string       `db:"purchase_order_branch_id"`
	CompanyID             string       `db:"company_id"`
	BranchID              string       `db:"branch_id"`
	FormNumber            string       `db:"form_number"`
	TransactionDate       time.Time    `db:"transaction_date"`
	ContactBookID         string       `db:"contact_book_id"`
	SecondaryCompanyID    string       `db:"secondary_company_id"`
	KonekinID             string       `db:"konekin_id"`
	CurrencyCode          string       `db:"currency_code"`
	TotalItems            int64        `db:"total_items"`
	Total                 int64        `db:"total"`
	IsDeleted             bool         `db:"is_deleted"`
	Status                string       `db:"status"`
	IsAllBranches         bool         `db:"is_all_branches"`
	CreatedAt             sql.NullTime `db:"created_at"`
	UpdatedAt             sql.NullTime `db:"updated_at"`
}

type SalesSalesOrderBranch struct {
	SalesOrderID    string       `db:"sales_order_id"`
	CompanyBranchID string       `db:"company_branch_id"`
	CreatedAt       sql.NullTime `db:"created_at"`
	UpdatedAt       sql.NullTime `db:"updated_at"`
}

type SalesSalesOrderItem struct {
	ID                     string       `db:"id"`
	PurchaseOrderItemID    string       `db:"purchase_order_item_id"`
	SalesOrderID           string       `db:"sales_order_id"`
	PrimaryItemVariantID   string       `db:"primary_item_variant_id"`
	SecondaryItemVariantID string       `db:"secondary_item_variant_id"`
	PrimaryItemUnitID      string       `db:"primary_item_unit_id"`
	SecondaryItemUnitID    string       `db:"secondary_item_unit_id"`
	PrimaryItemUnitValue   int64        `db:"primary_item_unit_value"`
	SecondaryItemUnitValue int64        `db:"secondary_item_unit_value"`
	Amount                 int64        `db:"amount"`
	AmountSent             int64        `db:"amount_sent"`
	Price                  int64        `db:"price"`
	IsDeleted              bool         `db:"is_deleted"`
	CreatedAt              sql.NullTime `db:"created_at"`
	UpdatedAt              sql.NullTime `db:"updated_at"`
}
