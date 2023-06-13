// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.17.2
// source: query.sql

package db

import (
	"context"
	"database/sql"
	"time"
)

const deletePOS = `-- name: DeletePOS :exec
UPDATE sales.point_of_sales SET is_deleted = TRUE WHERE id = $1
`

func (q *Queries) DeletePOS(ctx context.Context, id string) error {
	_, err := q.db.ExecContext(ctx, deletePOS, id)
	return err
}

const deletePOSCOASetting = `-- name: DeletePOSCOASetting :exec
DELETE FROM sales.pos_chart_of_account_settings WHERE branch_id = $1
`

func (q *Queries) DeletePOSCOASetting(ctx context.Context, branchID string) error {
	_, err := q.db.ExecContext(ctx, deletePOSCOASetting, branchID)
	return err
}

const deletePOSCustomerSetting = `-- name: DeletePOSCustomerSetting :exec
DELETE FROM sales.pos_customer_settings WHERE branch_id = $1
`

func (q *Queries) DeletePOSCustomerSetting(ctx context.Context, branchID string) error {
	_, err := q.db.ExecContext(ctx, deletePOSCustomerSetting, branchID)
	return err
}

const deletePOSItemsPOS = `-- name: DeletePOSItemsPOS :exec
DELETE FROM sales.point_of_sale_items WHERE point_of_sale_id = $1
`

func (q *Queries) DeletePOSItemsPOS(ctx context.Context, pointOfSaleID string) error {
	_, err := q.db.ExecContext(ctx, deletePOSItemsPOS, pointOfSaleID)
	return err
}

const deletePOSPaymentMethod = `-- name: DeletePOSPaymentMethod :exec
UPDATE sales.pos_payment_methods SET is_deleted = TRUE WHERE id = $1
`

func (q *Queries) DeletePOSPaymentMethod(ctx context.Context, id string) error {
	_, err := q.db.ExecContext(ctx, deletePOSPaymentMethod, id)
	return err
}

const deleteSalesOrderItems = `-- name: DeleteSalesOrderItems :exec
DELETE FROM sales.sales_order_items
WHERE sales_order_id = $1
`

func (q *Queries) DeleteSalesOrderItems(ctx context.Context, salesOrderID string) error {
	_, err := q.db.ExecContext(ctx, deleteSalesOrderItems, salesOrderID)
	return err
}

const getCheckPOS = `-- name: GetCheckPOS :one
SELECT 
    COUNT(id)::bigint AS total_count
FROM sales.point_of_sales
WHERE company_id = $1
`

func (q *Queries) GetCheckPOS(ctx context.Context, companyID string) (int64, error) {
	row := q.db.QueryRowContext(ctx, getCheckPOS, companyID)
	var total_count int64
	err := row.Scan(&total_count)
	return total_count, err
}

const getDeliveryOrders = `-- name: GetDeliveryOrders :many
SELECT 
    id, receipt_order_id, company_id, branch_id, form_number, transaction_date, contact_book_id, secondary_company_id, konekin_id, secondary_branch_id, total_items, is_deleted, status, created_at, updated_at
FROM sales.delivery_orders
WHERE company_id = $1
    AND branch_id = $2
    AND transaction_date BETWEEN $3::date AND $4::date 
    AND is_deleted = FALSE
`

type GetDeliveryOrdersParams struct {
	CompanyID string    `db:"company_id"`
	BranchID  string    `db:"branch_id"`
	StartDate time.Time `db:"start_date"`
	EndDate   time.Time `db:"end_date"`
}

func (q *Queries) GetDeliveryOrders(ctx context.Context, arg GetDeliveryOrdersParams) ([]SalesDeliveryOrder, error) {
	rows, err := q.db.QueryContext(ctx, getDeliveryOrders,
		arg.CompanyID,
		arg.BranchID,
		arg.StartDate,
		arg.EndDate,
	)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []SalesDeliveryOrder
	for rows.Next() {
		var i SalesDeliveryOrder
		if err := rows.Scan(
			&i.ID,
			&i.ReceiptOrderID,
			&i.CompanyID,
			&i.BranchID,
			&i.FormNumber,
			&i.TransactionDate,
			&i.ContactBookID,
			&i.SecondaryCompanyID,
			&i.KonekinID,
			&i.SecondaryBranchID,
			&i.TotalItems,
			&i.IsDeleted,
			&i.Status,
			&i.CreatedAt,
			&i.UpdatedAt,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const getPOS = `-- name: GetPOS :many
SELECT 
  a.id as id,
  a.company_id as company_id, a.branch_id as branch_id, a.warehouse_id as warehouse_id,
  a.form_number as form_number,
  a.transaction_date as transaction_date,
  a.contact_book_id as contact_book_id,
  a.secondary_company_id as secondary_company_id,
  a.konekin_id as konekin_id,
  a.currency_code as currency_code,
  a.pos_payment_method_id as pos_payment_method_id,
  b.name as pos_payment_method_name,
  a.total_items as total_items,
  a.total as total
FROM sales.point_of_sales a 
JOIN sales.pos_payment_methods b ON b.id = a.pos_payment_method_id
WHERE a.company_id LIKE $1
    AND a.branch_id LIKE $2
    AND a.transaction_date BETWEEN $3::date AND $4::date 
    AND a.is_deleted = FALSE
`

type GetPOSParams struct {
	CompanyID string    `db:"company_id"`
	BranchID  string    `db:"branch_id"`
	StartDate time.Time `db:"start_date"`
	EndDate   time.Time `db:"end_date"`
}

type GetPOSRow struct {
	ID                   string    `db:"id"`
	CompanyID            string    `db:"company_id"`
	BranchID             string    `db:"branch_id"`
	WarehouseID          string    `db:"warehouse_id"`
	FormNumber           string    `db:"form_number"`
	TransactionDate      time.Time `db:"transaction_date"`
	ContactBookID        string    `db:"contact_book_id"`
	SecondaryCompanyID   string    `db:"secondary_company_id"`
	KonekinID            string    `db:"konekin_id"`
	CurrencyCode         string    `db:"currency_code"`
	PosPaymentMethodID   string    `db:"pos_payment_method_id"`
	PosPaymentMethodName string    `db:"pos_payment_method_name"`
	TotalItems           int64     `db:"total_items"`
	Total                int64     `db:"total"`
}

func (q *Queries) GetPOS(ctx context.Context, arg GetPOSParams) ([]GetPOSRow, error) {
	rows, err := q.db.QueryContext(ctx, getPOS,
		arg.CompanyID,
		arg.BranchID,
		arg.StartDate,
		arg.EndDate,
	)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []GetPOSRow
	for rows.Next() {
		var i GetPOSRow
		if err := rows.Scan(
			&i.ID,
			&i.CompanyID,
			&i.BranchID,
			&i.WarehouseID,
			&i.FormNumber,
			&i.TransactionDate,
			&i.ContactBookID,
			&i.SecondaryCompanyID,
			&i.KonekinID,
			&i.CurrencyCode,
			&i.PosPaymentMethodID,
			&i.PosPaymentMethodName,
			&i.TotalItems,
			&i.Total,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const getPOSCOASetting = `-- name: GetPOSCOASetting :many
SELECT 
    branch_id, chart_of_account_id, created_at, updated_at
FROM sales.pos_chart_of_account_settings
WHERE branch_id = $1
`

func (q *Queries) GetPOSCOASetting(ctx context.Context, branchID string) ([]SalesPosChartOfAccountSetting, error) {
	rows, err := q.db.QueryContext(ctx, getPOSCOASetting, branchID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []SalesPosChartOfAccountSetting
	for rows.Next() {
		var i SalesPosChartOfAccountSetting
		if err := rows.Scan(
			&i.BranchID,
			&i.ChartOfAccountID,
			&i.CreatedAt,
			&i.UpdatedAt,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const getPOSCustomerSetting = `-- name: GetPOSCustomerSetting :many
SELECT 
    branch_id, contact_book_id, created_at, updated_at
FROM sales.pos_customer_settings
WHERE branch_id = $1
`

func (q *Queries) GetPOSCustomerSetting(ctx context.Context, branchID string) ([]SalesPosCustomerSetting, error) {
	rows, err := q.db.QueryContext(ctx, getPOSCustomerSetting, branchID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []SalesPosCustomerSetting
	for rows.Next() {
		var i SalesPosCustomerSetting
		if err := rows.Scan(
			&i.BranchID,
			&i.ContactBookID,
			&i.CreatedAt,
			&i.UpdatedAt,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const getPOSItemsByPOSId = `-- name: GetPOSItemsByPOSId :many
SELECT a.id, a.point_of_sale_id, a.warehouse_rack_id, a.item_variant_id, a.item_unit_id, a.item_unit_value, a.batch, a.expired_date, a.item_barcode_id, a.amount, a.price, a.is_deleted, a.created_at, a.updated_at FROM sales.point_of_sale_items a WHERE a.point_of_sale_id = $1 ORDER BY a.id
`

func (q *Queries) GetPOSItemsByPOSId(ctx context.Context, pointOfSaleID string) ([]SalesPointOfSaleItem, error) {
	rows, err := q.db.QueryContext(ctx, getPOSItemsByPOSId, pointOfSaleID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []SalesPointOfSaleItem
	for rows.Next() {
		var i SalesPointOfSaleItem
		if err := rows.Scan(
			&i.ID,
			&i.PointOfSaleID,
			&i.WarehouseRackID,
			&i.ItemVariantID,
			&i.ItemUnitID,
			&i.ItemUnitValue,
			&i.Batch,
			&i.ExpiredDate,
			&i.ItemBarcodeID,
			&i.Amount,
			&i.Price,
			&i.IsDeleted,
			&i.CreatedAt,
			&i.UpdatedAt,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const getPOSPaymentMethod = `-- name: GetPOSPaymentMethod :many
SELECT id, company_id, chart_of_account_id, name
FROM sales.pos_payment_methods 
WHERE is_deleted = FALSE AND company_id = $1 AND name LIKE $2
`

type GetPOSPaymentMethodParams struct {
	CompanyID string `db:"company_id"`
	Name      string `db:"name"`
}

type GetPOSPaymentMethodRow struct {
	ID               string `db:"id"`
	CompanyID        string `db:"company_id"`
	ChartOfAccountID string `db:"chart_of_account_id"`
	Name             string `db:"name"`
}

func (q *Queries) GetPOSPaymentMethod(ctx context.Context, arg GetPOSPaymentMethodParams) ([]GetPOSPaymentMethodRow, error) {
	rows, err := q.db.QueryContext(ctx, getPOSPaymentMethod, arg.CompanyID, arg.Name)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []GetPOSPaymentMethodRow
	for rows.Next() {
		var i GetPOSPaymentMethodRow
		if err := rows.Scan(
			&i.ID,
			&i.CompanyID,
			&i.ChartOfAccountID,
			&i.Name,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const getPOSUserSetting = `-- name: GetPOSUserSetting :one
SELECT 
    user_id, branch_id, warehouse_id, warehouse_rack_id, created_at, updated_at
FROM sales.pos_user_settings
WHERE user_id = $1
AND branch_id = $2
`

type GetPOSUserSettingParams struct {
	UserID   string `db:"user_id"`
	BranchID string `db:"branch_id"`
}

func (q *Queries) GetPOSUserSetting(ctx context.Context, arg GetPOSUserSettingParams) (SalesPosUserSetting, error) {
	row := q.db.QueryRowContext(ctx, getPOSUserSetting, arg.UserID, arg.BranchID)
	var i SalesPosUserSetting
	err := row.Scan(
		&i.UserID,
		&i.BranchID,
		&i.WarehouseID,
		&i.WarehouseRackID,
		&i.CreatedAt,
		&i.UpdatedAt,
	)
	return i, err
}

const getSalesOrder = `-- name: GetSalesOrder :one
SELECT 
    id, purchase_order_id, purchase_order_branch_id, company_id, branch_id, form_number, transaction_date, contact_book_id, secondary_company_id, konekin_id, currency_code, total_items, total, is_deleted, status, created_at, updated_at
FROM sales.sales_orders
WHERE id = $1
`

func (q *Queries) GetSalesOrder(ctx context.Context, id string) (SalesSalesOrder, error) {
	row := q.db.QueryRowContext(ctx, getSalesOrder, id)
	var i SalesSalesOrder
	err := row.Scan(
		&i.ID,
		&i.PurchaseOrderID,
		&i.PurchaseOrderBranchID,
		&i.CompanyID,
		&i.BranchID,
		&i.FormNumber,
		&i.TransactionDate,
		&i.ContactBookID,
		&i.SecondaryCompanyID,
		&i.KonekinID,
		&i.CurrencyCode,
		&i.TotalItems,
		&i.Total,
		&i.IsDeleted,
		&i.Status,
		&i.CreatedAt,
		&i.UpdatedAt,
	)
	return i, err
}

const getSalesOrderItems = `-- name: GetSalesOrderItems :many
SELECT 
    id, purchase_order_item_id, sales_order_id, primary_item_variant_id, secondary_item_variant_id, primary_item_unit_id, secondary_item_unit_id, primary_item_unit_value, secondary_item_unit_value, amount, amount_sent, price, is_deleted, created_at, updated_at
FROM sales.sales_order_items
WHERE sales_order_id = $1 AND is_deleted = FALSE
`

func (q *Queries) GetSalesOrderItems(ctx context.Context, salesOrderID string) ([]SalesSalesOrderItem, error) {
	rows, err := q.db.QueryContext(ctx, getSalesOrderItems, salesOrderID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []SalesSalesOrderItem
	for rows.Next() {
		var i SalesSalesOrderItem
		if err := rows.Scan(
			&i.ID,
			&i.PurchaseOrderItemID,
			&i.SalesOrderID,
			&i.PrimaryItemVariantID,
			&i.SecondaryItemVariantID,
			&i.PrimaryItemUnitID,
			&i.SecondaryItemUnitID,
			&i.PrimaryItemUnitValue,
			&i.SecondaryItemUnitValue,
			&i.Amount,
			&i.AmountSent,
			&i.Price,
			&i.IsDeleted,
			&i.CreatedAt,
			&i.UpdatedAt,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const getSalesOrders = `-- name: GetSalesOrders :many
SELECT 
    id, purchase_order_id, purchase_order_branch_id, company_id, branch_id, form_number, transaction_date, contact_book_id, secondary_company_id, konekin_id, currency_code, total_items, total, is_deleted, status, created_at, updated_at
FROM sales.sales_orders
WHERE company_id = $1
    AND branch_id = $2
    AND transaction_date BETWEEN $3::date AND $4::date 
    AND is_deleted = FALSE
`

type GetSalesOrdersParams struct {
	CompanyID string    `db:"company_id"`
	BranchID  string    `db:"branch_id"`
	StartDate time.Time `db:"start_date"`
	EndDate   time.Time `db:"end_date"`
}

func (q *Queries) GetSalesOrders(ctx context.Context, arg GetSalesOrdersParams) ([]SalesSalesOrder, error) {
	rows, err := q.db.QueryContext(ctx, getSalesOrders,
		arg.CompanyID,
		arg.BranchID,
		arg.StartDate,
		arg.EndDate,
	)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []SalesSalesOrder
	for rows.Next() {
		var i SalesSalesOrder
		if err := rows.Scan(
			&i.ID,
			&i.PurchaseOrderID,
			&i.PurchaseOrderBranchID,
			&i.CompanyID,
			&i.BranchID,
			&i.FormNumber,
			&i.TransactionDate,
			&i.ContactBookID,
			&i.SecondaryCompanyID,
			&i.KonekinID,
			&i.CurrencyCode,
			&i.TotalItems,
			&i.Total,
			&i.IsDeleted,
			&i.Status,
			&i.CreatedAt,
			&i.UpdatedAt,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const insertPOSCOASetting = `-- name: InsertPOSCOASetting :one
INSERT INTO sales.pos_chart_of_account_settings(
  branch_id, chart_of_account_id
)
VALUES ($1, $2)
RETURNING branch_id, chart_of_account_id, created_at, updated_at
`

type InsertPOSCOASettingParams struct {
	BranchID         string `db:"branch_id"`
	ChartOfAccountID string `db:"chart_of_account_id"`
}

func (q *Queries) InsertPOSCOASetting(ctx context.Context, arg InsertPOSCOASettingParams) (SalesPosChartOfAccountSetting, error) {
	row := q.db.QueryRowContext(ctx, insertPOSCOASetting, arg.BranchID, arg.ChartOfAccountID)
	var i SalesPosChartOfAccountSetting
	err := row.Scan(
		&i.BranchID,
		&i.ChartOfAccountID,
		&i.CreatedAt,
		&i.UpdatedAt,
	)
	return i, err
}

const insertPOSCustomerSetting = `-- name: InsertPOSCustomerSetting :one
INSERT INTO sales.pos_customer_settings(
  branch_id, contact_book_id
)
VALUES ($1, $2)
RETURNING branch_id, contact_book_id, created_at, updated_at
`

type InsertPOSCustomerSettingParams struct {
	BranchID      string `db:"branch_id"`
	ContactBookID string `db:"contact_book_id"`
}

func (q *Queries) InsertPOSCustomerSetting(ctx context.Context, arg InsertPOSCustomerSettingParams) (SalesPosCustomerSetting, error) {
	row := q.db.QueryRowContext(ctx, insertPOSCustomerSetting, arg.BranchID, arg.ContactBookID)
	var i SalesPosCustomerSetting
	err := row.Scan(
		&i.BranchID,
		&i.ContactBookID,
		&i.CreatedAt,
		&i.UpdatedAt,
	)
	return i, err
}

const insertPOSItem = `-- name: InsertPOSItem :one
INSERT INTO sales.point_of_sale_items(
  id, point_of_sale_id, warehouse_rack_id, item_variant_id, item_unit_id, item_unit_value, batch, expired_date, item_barcode_id, amount, price, updated_at
)
VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12)
RETURNING id, point_of_sale_id, warehouse_rack_id, item_variant_id, item_unit_id, item_unit_value, batch, expired_date, item_barcode_id, amount, price, is_deleted, created_at, updated_at
`

type InsertPOSItemParams struct {
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
	UpdatedAt       sql.NullTime   `db:"updated_at"`
}

func (q *Queries) InsertPOSItem(ctx context.Context, arg InsertPOSItemParams) (SalesPointOfSaleItem, error) {
	row := q.db.QueryRowContext(ctx, insertPOSItem,
		arg.ID,
		arg.PointOfSaleID,
		arg.WarehouseRackID,
		arg.ItemVariantID,
		arg.ItemUnitID,
		arg.ItemUnitValue,
		arg.Batch,
		arg.ExpiredDate,
		arg.ItemBarcodeID,
		arg.Amount,
		arg.Price,
		arg.UpdatedAt,
	)
	var i SalesPointOfSaleItem
	err := row.Scan(
		&i.ID,
		&i.PointOfSaleID,
		&i.WarehouseRackID,
		&i.ItemVariantID,
		&i.ItemUnitID,
		&i.ItemUnitValue,
		&i.Batch,
		&i.ExpiredDate,
		&i.ItemBarcodeID,
		&i.Amount,
		&i.Price,
		&i.IsDeleted,
		&i.CreatedAt,
		&i.UpdatedAt,
	)
	return i, err
}

const updateSalesOrderAddItem = `-- name: UpdateSalesOrderAddItem :exec
UPDATE sales.sales_orders
SET total_items=sub.total_items,
    total=sub.total,
    updated_at = NOW()
FROM (SELECT sales_order_id, COUNT(id) AS total_items, SUM(amount*price) AS total
      FROM sales.sales_order_items
      WHERE sales_order_id = $1
      GROUP BY sales_order_id) AS sub
WHERE sales.sales_orders.id = sub.sales_order_id
`

func (q *Queries) UpdateSalesOrderAddItem(ctx context.Context, salesOrderID string) error {
	_, err := q.db.ExecContext(ctx, updateSalesOrderAddItem, salesOrderID)
	return err
}

const updateSalesOrderStatus = `-- name: UpdateSalesOrderStatus :exec
UPDATE sales.sales_orders
SET status = $2
WHERE id = $1
`

type UpdateSalesOrderStatusParams struct {
	ID     string `db:"id"`
	Status string `db:"status"`
}

func (q *Queries) UpdateSalesOrderStatus(ctx context.Context, arg UpdateSalesOrderStatusParams) error {
	_, err := q.db.ExecContext(ctx, updateSalesOrderStatus, arg.ID, arg.Status)
	return err
}

const upsertDeliveryOrder = `-- name: UpsertDeliveryOrder :one
INSERT INTO sales.delivery_orders(
    id, company_id, branch_id,
    form_number, transaction_date,
    contact_book_id, secondary_company_id, konekin_id,
    secondary_branch_id
)
VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9) ON CONFLICT (id) DO
UPDATE
SET company_id = EXCLUDED.company_id,
    branch_id = EXCLUDED.branch_id,
    form_number = EXCLUDED.form_number,
    transaction_date = EXCLUDED.transaction_date,
    contact_book_id = EXCLUDED.contact_book_id,
    secondary_company_id = EXCLUDED.secondary_company_id,
    konekin_id = EXCLUDED.konekin_id,
    secondary_branch_id = EXCLUDED.secondary_branch_id,
    updated_at = NOW()
RETURNING id, receipt_order_id, company_id, branch_id, form_number, transaction_date, contact_book_id, secondary_company_id, konekin_id, secondary_branch_id, total_items, is_deleted, status, created_at, updated_at
`

type UpsertDeliveryOrderParams struct {
	ID                 string    `db:"id"`
	CompanyID          string    `db:"company_id"`
	BranchID           string    `db:"branch_id"`
	FormNumber         string    `db:"form_number"`
	TransactionDate    time.Time `db:"transaction_date"`
	ContactBookID      string    `db:"contact_book_id"`
	SecondaryCompanyID string    `db:"secondary_company_id"`
	KonekinID          string    `db:"konekin_id"`
	SecondaryBranchID  string    `db:"secondary_branch_id"`
}

func (q *Queries) UpsertDeliveryOrder(ctx context.Context, arg UpsertDeliveryOrderParams) (SalesDeliveryOrder, error) {
	row := q.db.QueryRowContext(ctx, upsertDeliveryOrder,
		arg.ID,
		arg.CompanyID,
		arg.BranchID,
		arg.FormNumber,
		arg.TransactionDate,
		arg.ContactBookID,
		arg.SecondaryCompanyID,
		arg.KonekinID,
		arg.SecondaryBranchID,
	)
	var i SalesDeliveryOrder
	err := row.Scan(
		&i.ID,
		&i.ReceiptOrderID,
		&i.CompanyID,
		&i.BranchID,
		&i.FormNumber,
		&i.TransactionDate,
		&i.ContactBookID,
		&i.SecondaryCompanyID,
		&i.KonekinID,
		&i.SecondaryBranchID,
		&i.TotalItems,
		&i.IsDeleted,
		&i.Status,
		&i.CreatedAt,
		&i.UpdatedAt,
	)
	return i, err
}

const upsertPOS = `-- name: UpsertPOS :one
INSERT INTO sales.point_of_sales(
  id, company_id, branch_id, warehouse_id, form_number, transaction_date,
  contact_book_id, secondary_company_id, konekin_id, currency_code, pos_payment_method_id, total_items, total, updated_at
) VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12, $13, $14) ON CONFLICT (id) DO
UPDATE
SET company_id = EXCLUDED.company_id,
    branch_id = EXCLUDED.branch_id,
    warehouse_id = EXCLUDED.warehouse_id,
    form_number = EXCLUDED.form_number,
    transaction_date = EXCLUDED.transaction_date,
    contact_book_id = EXCLUDED.contact_book_id,
    secondary_company_id = EXCLUDED.secondary_company_id,
    konekin_id = EXCLUDED.konekin_id,
    currency_code = EXCLUDED.currency_code,
    pos_payment_method_id = EXCLUDED.pos_payment_method_id,
    total_items = EXCLUDED.total_items,
    total = EXCLUDED.total,
    updated_at = NOW()
RETURNING id, company_id, branch_id, warehouse_id, form_number, transaction_date, contact_book_id, secondary_company_id, konekin_id, currency_code, pos_payment_method_id, total_items, total, is_deleted, created_at, updated_at
`

type UpsertPOSParams struct {
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
	UpdatedAt          sql.NullTime `db:"updated_at"`
}

func (q *Queries) UpsertPOS(ctx context.Context, arg UpsertPOSParams) (SalesPointOfSale, error) {
	row := q.db.QueryRowContext(ctx, upsertPOS,
		arg.ID,
		arg.CompanyID,
		arg.BranchID,
		arg.WarehouseID,
		arg.FormNumber,
		arg.TransactionDate,
		arg.ContactBookID,
		arg.SecondaryCompanyID,
		arg.KonekinID,
		arg.CurrencyCode,
		arg.PosPaymentMethodID,
		arg.TotalItems,
		arg.Total,
		arg.UpdatedAt,
	)
	var i SalesPointOfSale
	err := row.Scan(
		&i.ID,
		&i.CompanyID,
		&i.BranchID,
		&i.WarehouseID,
		&i.FormNumber,
		&i.TransactionDate,
		&i.ContactBookID,
		&i.SecondaryCompanyID,
		&i.KonekinID,
		&i.CurrencyCode,
		&i.PosPaymentMethodID,
		&i.TotalItems,
		&i.Total,
		&i.IsDeleted,
		&i.CreatedAt,
		&i.UpdatedAt,
	)
	return i, err
}

const upsertPOSPaymentMethod = `-- name: UpsertPOSPaymentMethod :exec
INSERT INTO sales.pos_payment_methods(id, company_id, chart_of_account_id, name)
VALUES ($1, $2, $3, $4)
ON CONFLICT (id) DO UPDATE
SET name = EXCLUDED.name,
  chart_of_account_id = EXCLUDED.chart_of_account_id,
  company_id = EXCLUDED.company_id,
  updated_at = NOW()
`

type UpsertPOSPaymentMethodParams struct {
	ID               string `db:"id"`
	CompanyID        string `db:"company_id"`
	ChartOfAccountID string `db:"chart_of_account_id"`
	Name             string `db:"name"`
}

func (q *Queries) UpsertPOSPaymentMethod(ctx context.Context, arg UpsertPOSPaymentMethodParams) error {
	_, err := q.db.ExecContext(ctx, upsertPOSPaymentMethod,
		arg.ID,
		arg.CompanyID,
		arg.ChartOfAccountID,
		arg.Name,
	)
	return err
}

const upsertPOSUserSetting = `-- name: UpsertPOSUserSetting :one
INSERT INTO sales.pos_user_settings(
  user_id, branch_id, warehouse_id, warehouse_rack_id
)
VALUES ($1, $2, $3, $4) ON CONFLICT (user_id, branch_id) DO
UPDATE
SET warehouse_id = EXCLUDED.warehouse_id,
  warehouse_rack_id = EXCLUDED.warehouse_rack_id,
  updated_at = NOW()
RETURNING user_id, branch_id, warehouse_id, warehouse_rack_id, created_at, updated_at
`

type UpsertPOSUserSettingParams struct {
	UserID          string `db:"user_id"`
	BranchID        string `db:"branch_id"`
	WarehouseID     string `db:"warehouse_id"`
	WarehouseRackID string `db:"warehouse_rack_id"`
}

func (q *Queries) UpsertPOSUserSetting(ctx context.Context, arg UpsertPOSUserSettingParams) (SalesPosUserSetting, error) {
	row := q.db.QueryRowContext(ctx, upsertPOSUserSetting,
		arg.UserID,
		arg.BranchID,
		arg.WarehouseID,
		arg.WarehouseRackID,
	)
	var i SalesPosUserSetting
	err := row.Scan(
		&i.UserID,
		&i.BranchID,
		&i.WarehouseID,
		&i.WarehouseRackID,
		&i.CreatedAt,
		&i.UpdatedAt,
	)
	return i, err
}

const upsertSalesOrder = `-- name: UpsertSalesOrder :one
INSERT INTO sales.sales_orders(
    id, purchase_order_id, purchase_order_branch_id, company_id, branch_id,
    form_number, transaction_date,
    contact_book_id, secondary_company_id, konekin_id, currency_code
)
VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11) ON CONFLICT (id) DO
UPDATE
SET purchase_order_id = EXCLUDED.purchase_order_id,
    purchase_order_branch_id = EXCLUDED.purchase_order_branch_id,
    company_id = EXCLUDED.company_id,
    branch_id = EXCLUDED.branch_id,
    form_number = EXCLUDED.form_number,
    transaction_date = EXCLUDED.transaction_date,
    contact_book_id = EXCLUDED.contact_book_id,
    secondary_company_id = EXCLUDED.secondary_company_id,
    konekin_id = EXCLUDED.konekin_id,
    currency_code = EXCLUDED.currency_code,
    updated_at = NOW()
RETURNING id, purchase_order_id, purchase_order_branch_id, company_id, branch_id, form_number, transaction_date, contact_book_id, secondary_company_id, konekin_id, currency_code, total_items, total, is_deleted, status, created_at, updated_at
`

type UpsertSalesOrderParams struct {
	ID                    string    `db:"id"`
	PurchaseOrderID       string    `db:"purchase_order_id"`
	PurchaseOrderBranchID string    `db:"purchase_order_branch_id"`
	CompanyID             string    `db:"company_id"`
	BranchID              string    `db:"branch_id"`
	FormNumber            string    `db:"form_number"`
	TransactionDate       time.Time `db:"transaction_date"`
	ContactBookID         string    `db:"contact_book_id"`
	SecondaryCompanyID    string    `db:"secondary_company_id"`
	KonekinID             string    `db:"konekin_id"`
	CurrencyCode          string    `db:"currency_code"`
}

func (q *Queries) UpsertSalesOrder(ctx context.Context, arg UpsertSalesOrderParams) (SalesSalesOrder, error) {
	row := q.db.QueryRowContext(ctx, upsertSalesOrder,
		arg.ID,
		arg.PurchaseOrderID,
		arg.PurchaseOrderBranchID,
		arg.CompanyID,
		arg.BranchID,
		arg.FormNumber,
		arg.TransactionDate,
		arg.ContactBookID,
		arg.SecondaryCompanyID,
		arg.KonekinID,
		arg.CurrencyCode,
	)
	var i SalesSalesOrder
	err := row.Scan(
		&i.ID,
		&i.PurchaseOrderID,
		&i.PurchaseOrderBranchID,
		&i.CompanyID,
		&i.BranchID,
		&i.FormNumber,
		&i.TransactionDate,
		&i.ContactBookID,
		&i.SecondaryCompanyID,
		&i.KonekinID,
		&i.CurrencyCode,
		&i.TotalItems,
		&i.Total,
		&i.IsDeleted,
		&i.Status,
		&i.CreatedAt,
		&i.UpdatedAt,
	)
	return i, err
}

const upsertSalesOrderItem = `-- name: UpsertSalesOrderItem :one
INSERT INTO sales.sales_order_items(
        id, purchase_order_item_id, sales_order_id,
        primary_item_variant_id, secondary_item_variant_id,
        primary_item_unit_id, secondary_item_unit_id,
        primary_item_unit_value, secondary_item_unit_value,
        amount, price
    )
VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11) ON CONFLICT (id) DO
UPDATE
SET purchase_order_item_id = EXCLUDED.purchase_order_item_id,
    sales_order_id = EXCLUDED.sales_order_id,
    primary_item_variant_id = EXCLUDED.primary_item_variant_id,
    secondary_item_variant_id = EXCLUDED.secondary_item_variant_id,
    primary_item_unit_id = EXCLUDED.primary_item_unit_id,
    secondary_item_unit_id = EXCLUDED.secondary_item_unit_id,
    primary_item_unit_value = EXCLUDED.primary_item_unit_value,
    secondary_item_unit_value = EXCLUDED.secondary_item_unit_value,
    amount = EXCLUDED.amount,
    price = EXCLUDED.price,
    updated_at = NOW()
RETURNING id, purchase_order_item_id, sales_order_id, primary_item_variant_id, secondary_item_variant_id, primary_item_unit_id, secondary_item_unit_id, primary_item_unit_value, secondary_item_unit_value, amount, amount_sent, price, is_deleted, created_at, updated_at
`

type UpsertSalesOrderItemParams struct {
	ID                     string `db:"id"`
	PurchaseOrderItemID    string `db:"purchase_order_item_id"`
	SalesOrderID           string `db:"sales_order_id"`
	PrimaryItemVariantID   string `db:"primary_item_variant_id"`
	SecondaryItemVariantID string `db:"secondary_item_variant_id"`
	PrimaryItemUnitID      string `db:"primary_item_unit_id"`
	SecondaryItemUnitID    string `db:"secondary_item_unit_id"`
	PrimaryItemUnitValue   int64  `db:"primary_item_unit_value"`
	SecondaryItemUnitValue int64  `db:"secondary_item_unit_value"`
	Amount                 int64  `db:"amount"`
	Price                  int64  `db:"price"`
}

func (q *Queries) UpsertSalesOrderItem(ctx context.Context, arg UpsertSalesOrderItemParams) (SalesSalesOrderItem, error) {
	row := q.db.QueryRowContext(ctx, upsertSalesOrderItem,
		arg.ID,
		arg.PurchaseOrderItemID,
		arg.SalesOrderID,
		arg.PrimaryItemVariantID,
		arg.SecondaryItemVariantID,
		arg.PrimaryItemUnitID,
		arg.SecondaryItemUnitID,
		arg.PrimaryItemUnitValue,
		arg.SecondaryItemUnitValue,
		arg.Amount,
		arg.Price,
	)
	var i SalesSalesOrderItem
	err := row.Scan(
		&i.ID,
		&i.PurchaseOrderItemID,
		&i.SalesOrderID,
		&i.PrimaryItemVariantID,
		&i.SecondaryItemVariantID,
		&i.PrimaryItemUnitID,
		&i.SecondaryItemUnitID,
		&i.PrimaryItemUnitValue,
		&i.SecondaryItemUnitValue,
		&i.Amount,
		&i.AmountSent,
		&i.Price,
		&i.IsDeleted,
		&i.CreatedAt,
		&i.UpdatedAt,
	)
	return i, err
}
