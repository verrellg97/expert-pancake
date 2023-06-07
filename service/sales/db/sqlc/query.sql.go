// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.18.0
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
  b.chart_of_account_id as chart_of_account_id,
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
	ChartOfAccountID     string    `db:"chart_of_account_id"`
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
			&i.ChartOfAccountID,
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
SELECT  id, company_id, chart_of_account_id, name
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
