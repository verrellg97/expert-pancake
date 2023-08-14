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

const deletePurchaseInvoiceItems = `-- name: DeletePurchaseInvoiceItems :exec
DELETE FROM purchasing.purchase_invoice_items
WHERE purchase_invoice_id = $1
`

func (q *Queries) DeletePurchaseInvoiceItems(ctx context.Context, purchaseInvoiceID string) error {
	_, err := q.db.ExecContext(ctx, deletePurchaseInvoiceItems, purchaseInvoiceID)
	return err
}

const deletePurchaseOrderItems = `-- name: DeletePurchaseOrderItems :exec
DELETE FROM purchasing.purchase_order_items
WHERE purchase_order_id = $1
`

func (q *Queries) DeletePurchaseOrderItems(ctx context.Context, purchaseOrderID string) error {
	_, err := q.db.ExecContext(ctx, deletePurchaseOrderItems, purchaseOrderID)
	return err
}

const deleteReceiptOrder = `-- name: DeleteReceiptOrder :exec
UPDATE purchasing.receipt_orders
SET is_deleted = TRUE
WHERE id = $1
`

func (q *Queries) DeleteReceiptOrder(ctx context.Context, id string) error {
	_, err := q.db.ExecContext(ctx, deleteReceiptOrder, id)
	return err
}

const deleteReceiptOrderItems = `-- name: DeleteReceiptOrderItems :exec
DELETE FROM purchasing.receipt_order_items
WHERE receipt_order_id = $1
`

func (q *Queries) DeleteReceiptOrderItems(ctx context.Context, receiptOrderID string) error {
	_, err := q.db.ExecContext(ctx, deleteReceiptOrderItems, receiptOrderID)
	return err
}

const getCheckPurchaseOrders = `-- name: GetCheckPurchaseOrders :one
SELECT 
    COUNT(id)::bigint AS total_count
FROM purchasing.purchase_orders
WHERE company_id = $1
`

func (q *Queries) GetCheckPurchaseOrders(ctx context.Context, companyID string) (int64, error) {
	row := q.db.QueryRowContext(ctx, getCheckPurchaseOrders, companyID)
	var total_count int64
	err := row.Scan(&total_count)
	return total_count, err
}

const getPurchaseInvoiceItems = `-- name: GetPurchaseInvoiceItems :many
SELECT 
    a.id, a.purchase_order_item_id, a.sales_order_item_id, a.sales_invoice_item_id, a.receipt_order_item_id, a.purchase_invoice_id, a.primary_item_variant_id, a.secondary_item_variant_id, a.primary_item_unit_id, a.secondary_item_unit_id, a.primary_item_unit_value, a.secondary_item_unit_value, a.amount, a.price, a.is_deleted, a.created_at, a.updated_at,
    b.warehouse_rack_id AS warehouse_rack_id,
    b.item_barcode_id AS item_barcode_id,
    b.batch AS batch,
    b.expired_date AS expired_date
FROM purchasing.purchase_invoice_items a
JOIN purchasing.receipt_order_items b ON b.id = a.receipt_order_item_id
WHERE a.purchase_invoice_id = $1
`

type GetPurchaseInvoiceItemsRow struct {
	ID                     string         `db:"id"`
	PurchaseOrderItemID    string         `db:"purchase_order_item_id"`
	SalesOrderItemID       string         `db:"sales_order_item_id"`
	SalesInvoiceItemID     string         `db:"sales_invoice_item_id"`
	ReceiptOrderItemID     string         `db:"receipt_order_item_id"`
	PurchaseInvoiceID      string         `db:"purchase_invoice_id"`
	PrimaryItemVariantID   string         `db:"primary_item_variant_id"`
	SecondaryItemVariantID string         `db:"secondary_item_variant_id"`
	PrimaryItemUnitID      string         `db:"primary_item_unit_id"`
	SecondaryItemUnitID    string         `db:"secondary_item_unit_id"`
	PrimaryItemUnitValue   int64          `db:"primary_item_unit_value"`
	SecondaryItemUnitValue int64          `db:"secondary_item_unit_value"`
	Amount                 int64          `db:"amount"`
	Price                  int64          `db:"price"`
	IsDeleted              bool           `db:"is_deleted"`
	CreatedAt              sql.NullTime   `db:"created_at"`
	UpdatedAt              sql.NullTime   `db:"updated_at"`
	WarehouseRackID        string         `db:"warehouse_rack_id"`
	ItemBarcodeID          string         `db:"item_barcode_id"`
	Batch                  sql.NullString `db:"batch"`
	ExpiredDate            sql.NullTime   `db:"expired_date"`
}

func (q *Queries) GetPurchaseInvoiceItems(ctx context.Context, purchaseInvoiceID string) ([]GetPurchaseInvoiceItemsRow, error) {
	rows, err := q.db.QueryContext(ctx, getPurchaseInvoiceItems, purchaseInvoiceID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []GetPurchaseInvoiceItemsRow
	for rows.Next() {
		var i GetPurchaseInvoiceItemsRow
		if err := rows.Scan(
			&i.ID,
			&i.PurchaseOrderItemID,
			&i.SalesOrderItemID,
			&i.SalesInvoiceItemID,
			&i.ReceiptOrderItemID,
			&i.PurchaseInvoiceID,
			&i.PrimaryItemVariantID,
			&i.SecondaryItemVariantID,
			&i.PrimaryItemUnitID,
			&i.SecondaryItemUnitID,
			&i.PrimaryItemUnitValue,
			&i.SecondaryItemUnitValue,
			&i.Amount,
			&i.Price,
			&i.IsDeleted,
			&i.CreatedAt,
			&i.UpdatedAt,
			&i.WarehouseRackID,
			&i.ItemBarcodeID,
			&i.Batch,
			&i.ExpiredDate,
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

const getPurchaseInvoices = `-- name: GetPurchaseInvoices :many
SELECT 
    a.id, a.sales_invoice_id, a.receipt_order_id, a.company_id, a.branch_id, a.form_number, a.transaction_date, a.contact_book_id, a.secondary_company_id, a.konekin_id, a.currency_code, a.total_items, a.total, a.is_deleted, a.status, a.created_at, a.updated_at,
    b.form_number AS receipt_order_form_number,
    b.transaction_date AS receipt_order_transaction_date,
    b.warehouse_id AS warehouse_id
FROM purchasing.purchase_invoices a
JOIN purchasing.receipt_orders b ON b.id = a.receipt_order_id
WHERE a.company_id = $1
    AND a.branch_id = $2
    AND a.transaction_date BETWEEN $3::date AND $4::date 
    AND a.is_deleted = FALSE
`

type GetPurchaseInvoicesParams struct {
	CompanyID string    `db:"company_id"`
	BranchID  string    `db:"branch_id"`
	StartDate time.Time `db:"start_date"`
	EndDate   time.Time `db:"end_date"`
}

type GetPurchaseInvoicesRow struct {
	ID                          string       `db:"id"`
	SalesInvoiceID              string       `db:"sales_invoice_id"`
	ReceiptOrderID              string       `db:"receipt_order_id"`
	CompanyID                   string       `db:"company_id"`
	BranchID                    string       `db:"branch_id"`
	FormNumber                  string       `db:"form_number"`
	TransactionDate             time.Time    `db:"transaction_date"`
	ContactBookID               string       `db:"contact_book_id"`
	SecondaryCompanyID          string       `db:"secondary_company_id"`
	KonekinID                   string       `db:"konekin_id"`
	CurrencyCode                string       `db:"currency_code"`
	TotalItems                  int64        `db:"total_items"`
	Total                       int64        `db:"total"`
	IsDeleted                   bool         `db:"is_deleted"`
	Status                      string       `db:"status"`
	CreatedAt                   sql.NullTime `db:"created_at"`
	UpdatedAt                   sql.NullTime `db:"updated_at"`
	ReceiptOrderFormNumber      string       `db:"receipt_order_form_number"`
	ReceiptOrderTransactionDate time.Time    `db:"receipt_order_transaction_date"`
	WarehouseID                 string       `db:"warehouse_id"`
}

func (q *Queries) GetPurchaseInvoices(ctx context.Context, arg GetPurchaseInvoicesParams) ([]GetPurchaseInvoicesRow, error) {
	rows, err := q.db.QueryContext(ctx, getPurchaseInvoices,
		arg.CompanyID,
		arg.BranchID,
		arg.StartDate,
		arg.EndDate,
	)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []GetPurchaseInvoicesRow
	for rows.Next() {
		var i GetPurchaseInvoicesRow
		if err := rows.Scan(
			&i.ID,
			&i.SalesInvoiceID,
			&i.ReceiptOrderID,
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
			&i.ReceiptOrderFormNumber,
			&i.ReceiptOrderTransactionDate,
			&i.WarehouseID,
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

const getPurchaseOrder = `-- name: GetPurchaseOrder :one
SELECT 
    id, sales_order_id, company_id, branch_id, form_number, transaction_date, contact_book_id, secondary_company_id, konekin_id, payment_term, currency_code, shipping_date, receiving_warehouse_id, total_items, total, is_deleted, status, created_at, updated_at
FROM purchasing.purchase_orders
WHERE id = $1
`

func (q *Queries) GetPurchaseOrder(ctx context.Context, id string) (PurchasingPurchaseOrder, error) {
	row := q.db.QueryRowContext(ctx, getPurchaseOrder, id)
	var i PurchasingPurchaseOrder
	err := row.Scan(
		&i.ID,
		&i.SalesOrderID,
		&i.CompanyID,
		&i.BranchID,
		&i.FormNumber,
		&i.TransactionDate,
		&i.ContactBookID,
		&i.SecondaryCompanyID,
		&i.KonekinID,
		&i.PaymentTerm,
		&i.CurrencyCode,
		&i.ShippingDate,
		&i.ReceivingWarehouseID,
		&i.TotalItems,
		&i.Total,
		&i.IsDeleted,
		&i.Status,
		&i.CreatedAt,
		&i.UpdatedAt,
	)
	return i, err
}

const getPurchaseOrderItems = `-- name: GetPurchaseOrderItems :many
SELECT 
    id, sales_order_item_id, purchase_order_id, primary_item_variant_id, secondary_item_variant_id, primary_item_unit_id, secondary_item_unit_id, primary_item_unit_value, secondary_item_unit_value, amount, amount_received, amount_invoiced, price, is_deleted, created_at, updated_at
FROM purchasing.purchase_order_items
WHERE purchase_order_id = $1 AND is_deleted = FALSE
`

func (q *Queries) GetPurchaseOrderItems(ctx context.Context, purchaseOrderID string) ([]PurchasingPurchaseOrderItem, error) {
	rows, err := q.db.QueryContext(ctx, getPurchaseOrderItems, purchaseOrderID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []PurchasingPurchaseOrderItem
	for rows.Next() {
		var i PurchasingPurchaseOrderItem
		if err := rows.Scan(
			&i.ID,
			&i.SalesOrderItemID,
			&i.PurchaseOrderID,
			&i.PrimaryItemVariantID,
			&i.SecondaryItemVariantID,
			&i.PrimaryItemUnitID,
			&i.SecondaryItemUnitID,
			&i.PrimaryItemUnitValue,
			&i.SecondaryItemUnitValue,
			&i.Amount,
			&i.AmountReceived,
			&i.AmountInvoiced,
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

const getPurchaseOrders = `-- name: GetPurchaseOrders :many
SELECT 
    id, sales_order_id, company_id, branch_id, form_number, transaction_date, contact_book_id, secondary_company_id, konekin_id, payment_term, currency_code, shipping_date, receiving_warehouse_id, total_items, total, is_deleted, status, created_at, updated_at
FROM purchasing.purchase_orders
WHERE company_id = $1
    AND branch_id = $2
    AND transaction_date BETWEEN $3::date AND $4::date 
    AND is_deleted = FALSE
`

type GetPurchaseOrdersParams struct {
	CompanyID string    `db:"company_id"`
	BranchID  string    `db:"branch_id"`
	StartDate time.Time `db:"start_date"`
	EndDate   time.Time `db:"end_date"`
}

func (q *Queries) GetPurchaseOrders(ctx context.Context, arg GetPurchaseOrdersParams) ([]PurchasingPurchaseOrder, error) {
	rows, err := q.db.QueryContext(ctx, getPurchaseOrders,
		arg.CompanyID,
		arg.BranchID,
		arg.StartDate,
		arg.EndDate,
	)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []PurchasingPurchaseOrder
	for rows.Next() {
		var i PurchasingPurchaseOrder
		if err := rows.Scan(
			&i.ID,
			&i.SalesOrderID,
			&i.CompanyID,
			&i.BranchID,
			&i.FormNumber,
			&i.TransactionDate,
			&i.ContactBookID,
			&i.SecondaryCompanyID,
			&i.KonekinID,
			&i.PaymentTerm,
			&i.CurrencyCode,
			&i.ShippingDate,
			&i.ReceivingWarehouseID,
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

const getPurchaseSetting = `-- name: GetPurchaseSetting :one
SELECT 
    company_id, is_auto_approve_order, created_at, updated_at
FROM purchasing.purchase_settings
WHERE company_id = $1
`

func (q *Queries) GetPurchaseSetting(ctx context.Context, companyID string) (PurchasingPurchaseSetting, error) {
	row := q.db.QueryRowContext(ctx, getPurchaseSetting, companyID)
	var i PurchasingPurchaseSetting
	err := row.Scan(
		&i.CompanyID,
		&i.IsAutoApproveOrder,
		&i.CreatedAt,
		&i.UpdatedAt,
	)
	return i, err
}

const getReceiptOrder = `-- name: GetReceiptOrder :one
SELECT 
    id, delivery_order_id, company_id, branch_id, warehouse_id, form_number, transaction_date, contact_book_id, secondary_company_id, konekin_id, total_items, is_deleted, status, created_at, updated_at
FROM purchasing.receipt_orders
WHERE id = $1
`

func (q *Queries) GetReceiptOrder(ctx context.Context, id string) (PurchasingReceiptOrder, error) {
	row := q.db.QueryRowContext(ctx, getReceiptOrder, id)
	var i PurchasingReceiptOrder
	err := row.Scan(
		&i.ID,
		&i.DeliveryOrderID,
		&i.CompanyID,
		&i.BranchID,
		&i.WarehouseID,
		&i.FormNumber,
		&i.TransactionDate,
		&i.ContactBookID,
		&i.SecondaryCompanyID,
		&i.KonekinID,
		&i.TotalItems,
		&i.IsDeleted,
		&i.Status,
		&i.CreatedAt,
		&i.UpdatedAt,
	)
	return i, err
}

const getReceiptOrderItems = `-- name: GetReceiptOrderItems :many
SELECT 
    id, purchase_order_item_id, sales_order_item_id, delivery_order_item_id, receipt_order_id, primary_item_variant_id, warehouse_rack_id, batch, expired_date, item_barcode_id, secondary_item_variant_id, primary_item_unit_id, secondary_item_unit_id, primary_item_unit_value, secondary_item_unit_value, amount, is_deleted, created_at, updated_at
FROM purchasing.receipt_order_items
WHERE receipt_order_id = $1 AND is_deleted = FALSE
`

func (q *Queries) GetReceiptOrderItems(ctx context.Context, receiptOrderID string) ([]PurchasingReceiptOrderItem, error) {
	rows, err := q.db.QueryContext(ctx, getReceiptOrderItems, receiptOrderID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []PurchasingReceiptOrderItem
	for rows.Next() {
		var i PurchasingReceiptOrderItem
		if err := rows.Scan(
			&i.ID,
			&i.PurchaseOrderItemID,
			&i.SalesOrderItemID,
			&i.DeliveryOrderItemID,
			&i.ReceiptOrderID,
			&i.PrimaryItemVariantID,
			&i.WarehouseRackID,
			&i.Batch,
			&i.ExpiredDate,
			&i.ItemBarcodeID,
			&i.SecondaryItemVariantID,
			&i.PrimaryItemUnitID,
			&i.SecondaryItemUnitID,
			&i.PrimaryItemUnitValue,
			&i.SecondaryItemUnitValue,
			&i.Amount,
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

const getReceiptOrders = `-- name: GetReceiptOrders :many
SELECT 
    id, delivery_order_id, company_id, branch_id, warehouse_id, form_number, transaction_date, contact_book_id, secondary_company_id, konekin_id, total_items, is_deleted, status, created_at, updated_at
FROM purchasing.receipt_orders
WHERE company_id = $1
    AND branch_id = $2
    AND transaction_date BETWEEN $3::date AND $4::date 
    AND is_deleted = FALSE
`

type GetReceiptOrdersParams struct {
	CompanyID string    `db:"company_id"`
	BranchID  string    `db:"branch_id"`
	StartDate time.Time `db:"start_date"`
	EndDate   time.Time `db:"end_date"`
}

func (q *Queries) GetReceiptOrders(ctx context.Context, arg GetReceiptOrdersParams) ([]PurchasingReceiptOrder, error) {
	rows, err := q.db.QueryContext(ctx, getReceiptOrders,
		arg.CompanyID,
		arg.BranchID,
		arg.StartDate,
		arg.EndDate,
	)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []PurchasingReceiptOrder
	for rows.Next() {
		var i PurchasingReceiptOrder
		if err := rows.Scan(
			&i.ID,
			&i.DeliveryOrderID,
			&i.CompanyID,
			&i.BranchID,
			&i.WarehouseID,
			&i.FormNumber,
			&i.TransactionDate,
			&i.ContactBookID,
			&i.SecondaryCompanyID,
			&i.KonekinID,
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

const insertPurchaseInvoiceItem = `-- name: InsertPurchaseInvoiceItem :exec
INSERT INTO purchasing.purchase_invoice_items(
    id, purchase_order_item_id, sales_order_item_id, sales_invoice_item_id,
    receipt_order_item_id, purchase_invoice_id, primary_item_variant_id,
    secondary_item_variant_id, primary_item_unit_id, secondary_item_unit_id,
    primary_item_unit_value, secondary_item_unit_value, amount, price
)
VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12, $13, $14)
`

type InsertPurchaseInvoiceItemParams struct {
	ID                     string `db:"id"`
	PurchaseOrderItemID    string `db:"purchase_order_item_id"`
	SalesOrderItemID       string `db:"sales_order_item_id"`
	SalesInvoiceItemID     string `db:"sales_invoice_item_id"`
	ReceiptOrderItemID     string `db:"receipt_order_item_id"`
	PurchaseInvoiceID      string `db:"purchase_invoice_id"`
	PrimaryItemVariantID   string `db:"primary_item_variant_id"`
	SecondaryItemVariantID string `db:"secondary_item_variant_id"`
	PrimaryItemUnitID      string `db:"primary_item_unit_id"`
	SecondaryItemUnitID    string `db:"secondary_item_unit_id"`
	PrimaryItemUnitValue   int64  `db:"primary_item_unit_value"`
	SecondaryItemUnitValue int64  `db:"secondary_item_unit_value"`
	Amount                 int64  `db:"amount"`
	Price                  int64  `db:"price"`
}

func (q *Queries) InsertPurchaseInvoiceItem(ctx context.Context, arg InsertPurchaseInvoiceItemParams) error {
	_, err := q.db.ExecContext(ctx, insertPurchaseInvoiceItem,
		arg.ID,
		arg.PurchaseOrderItemID,
		arg.SalesOrderItemID,
		arg.SalesInvoiceItemID,
		arg.ReceiptOrderItemID,
		arg.PurchaseInvoiceID,
		arg.PrimaryItemVariantID,
		arg.SecondaryItemVariantID,
		arg.PrimaryItemUnitID,
		arg.SecondaryItemUnitID,
		arg.PrimaryItemUnitValue,
		arg.SecondaryItemUnitValue,
		arg.Amount,
		arg.Price,
	)
	return err
}

const insertReceiptOrderItem = `-- name: InsertReceiptOrderItem :one
INSERT INTO purchasing.receipt_order_items(
    id, purchase_order_item_id, sales_order_item_id, delivery_order_item_id,
    receipt_order_id, primary_item_variant_id, warehouse_rack_id, batch,
    expired_date, item_barcode_id, secondary_item_variant_id,
    primary_item_unit_id, secondary_item_unit_id, primary_item_unit_value,
    secondary_item_unit_value, amount
)
VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12, $13, $14, $15, $16)
RETURNING id, purchase_order_item_id, sales_order_item_id, delivery_order_item_id, receipt_order_id, primary_item_variant_id, warehouse_rack_id, batch, expired_date, item_barcode_id, secondary_item_variant_id, primary_item_unit_id, secondary_item_unit_id, primary_item_unit_value, secondary_item_unit_value, amount, is_deleted, created_at, updated_at
`

type InsertReceiptOrderItemParams struct {
	ID                     string         `db:"id"`
	PurchaseOrderItemID    string         `db:"purchase_order_item_id"`
	SalesOrderItemID       string         `db:"sales_order_item_id"`
	DeliveryOrderItemID    string         `db:"delivery_order_item_id"`
	ReceiptOrderID         string         `db:"receipt_order_id"`
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
}

func (q *Queries) InsertReceiptOrderItem(ctx context.Context, arg InsertReceiptOrderItemParams) (PurchasingReceiptOrderItem, error) {
	row := q.db.QueryRowContext(ctx, insertReceiptOrderItem,
		arg.ID,
		arg.PurchaseOrderItemID,
		arg.SalesOrderItemID,
		arg.DeliveryOrderItemID,
		arg.ReceiptOrderID,
		arg.PrimaryItemVariantID,
		arg.WarehouseRackID,
		arg.Batch,
		arg.ExpiredDate,
		arg.ItemBarcodeID,
		arg.SecondaryItemVariantID,
		arg.PrimaryItemUnitID,
		arg.SecondaryItemUnitID,
		arg.PrimaryItemUnitValue,
		arg.SecondaryItemUnitValue,
		arg.Amount,
	)
	var i PurchasingReceiptOrderItem
	err := row.Scan(
		&i.ID,
		&i.PurchaseOrderItemID,
		&i.SalesOrderItemID,
		&i.DeliveryOrderItemID,
		&i.ReceiptOrderID,
		&i.PrimaryItemVariantID,
		&i.WarehouseRackID,
		&i.Batch,
		&i.ExpiredDate,
		&i.ItemBarcodeID,
		&i.SecondaryItemVariantID,
		&i.PrimaryItemUnitID,
		&i.SecondaryItemUnitID,
		&i.PrimaryItemUnitValue,
		&i.SecondaryItemUnitValue,
		&i.Amount,
		&i.IsDeleted,
		&i.CreatedAt,
		&i.UpdatedAt,
	)
	return i, err
}

const updateAcceptedPurchaseOrder = `-- name: UpdateAcceptedPurchaseOrder :exec
UPDATE purchasing.purchase_orders
SET sales_order_id = $2
WHERE id = $1
`

type UpdateAcceptedPurchaseOrderParams struct {
	ID           string `db:"id"`
	SalesOrderID string `db:"sales_order_id"`
}

func (q *Queries) UpdateAcceptedPurchaseOrder(ctx context.Context, arg UpdateAcceptedPurchaseOrderParams) error {
	_, err := q.db.ExecContext(ctx, updateAcceptedPurchaseOrder, arg.ID, arg.SalesOrderID)
	return err
}

const updateAcceptedPurchaseOrderItem = `-- name: UpdateAcceptedPurchaseOrderItem :exec
UPDATE purchasing.purchase_order_items
SET sales_order_item_id = $2
WHERE id = $1
`

type UpdateAcceptedPurchaseOrderItemParams struct {
	ID               string `db:"id"`
	SalesOrderItemID string `db:"sales_order_item_id"`
}

func (q *Queries) UpdateAcceptedPurchaseOrderItem(ctx context.Context, arg UpdateAcceptedPurchaseOrderItemParams) error {
	_, err := q.db.ExecContext(ctx, updateAcceptedPurchaseOrderItem, arg.ID, arg.SalesOrderItemID)
	return err
}

const updatePurchaseOrderAddItem = `-- name: UpdatePurchaseOrderAddItem :exec
UPDATE purchasing.purchase_orders
SET total_items=sub.total_items,
    total=sub.total,
    updated_at = NOW()
FROM (SELECT purchase_order_id, COUNT(id) AS total_items, SUM(amount*price) AS total
      FROM purchasing.purchase_order_items
      WHERE purchase_order_id = $1
      GROUP BY purchase_order_id) AS sub
WHERE purchasing.purchase_orders.id = sub.purchase_order_id
`

func (q *Queries) UpdatePurchaseOrderAddItem(ctx context.Context, purchaseOrderID string) error {
	_, err := q.db.ExecContext(ctx, updatePurchaseOrderAddItem, purchaseOrderID)
	return err
}

const updatePurchaseOrderItemAmountReceived = `-- name: UpdatePurchaseOrderItemAmountReceived :exec
UPDATE purchasing.purchase_order_items
SET amount_received = amount_received+$2
WHERE id = $1
`

type UpdatePurchaseOrderItemAmountReceivedParams struct {
	ID             string `db:"id"`
	AmountReceived int64  `db:"amount_received"`
}

func (q *Queries) UpdatePurchaseOrderItemAmountReceived(ctx context.Context, arg UpdatePurchaseOrderItemAmountReceivedParams) error {
	_, err := q.db.ExecContext(ctx, updatePurchaseOrderItemAmountReceived, arg.ID, arg.AmountReceived)
	return err
}

const updatePurchaseOrderStatus = `-- name: UpdatePurchaseOrderStatus :exec
UPDATE purchasing.purchase_orders
SET status = $2
WHERE id = $1
`

type UpdatePurchaseOrderStatusParams struct {
	ID     string `db:"id"`
	Status string `db:"status"`
}

func (q *Queries) UpdatePurchaseOrderStatus(ctx context.Context, arg UpdatePurchaseOrderStatusParams) error {
	_, err := q.db.ExecContext(ctx, updatePurchaseOrderStatus, arg.ID, arg.Status)
	return err
}

const updateReceiptOrderStatus = `-- name: UpdateReceiptOrderStatus :exec
UPDATE purchasing.receipt_orders
SET status = $2
WHERE id = $1
`

type UpdateReceiptOrderStatusParams struct {
	ID     string `db:"id"`
	Status string `db:"status"`
}

func (q *Queries) UpdateReceiptOrderStatus(ctx context.Context, arg UpdateReceiptOrderStatusParams) error {
	_, err := q.db.ExecContext(ctx, updateReceiptOrderStatus, arg.ID, arg.Status)
	return err
}

const upsertPurchaseInvoice = `-- name: UpsertPurchaseInvoice :one
INSERT INTO purchasing.purchase_invoices(
        id, sales_invoice_id, receipt_order_id, company_id, branch_id, form_number, transaction_date,
        contact_book_id, secondary_company_id, konekin_id, currency_code, total_items, total, status
    ) 
VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12, $13, $14) ON CONFLICT (id) DO
UPDATE
SET 
    sales_invoice_id = EXCLUDED.sales_invoice_id,
    receipt_order_id = EXCLUDED.receipt_order_id,
    company_id = EXCLUDED.company_id,
    branch_id = EXCLUDED.branch_id,
    form_number = EXCLUDED.form_number,
    transaction_date = EXCLUDED.transaction_date,
    contact_book_id = EXCLUDED.contact_book_id,
    secondary_company_id = EXCLUDED.secondary_company_id,
    konekin_id = EXCLUDED.konekin_id,
    currency_code = EXCLUDED.currency_code,
    total_items = EXCLUDED.total_items,
    total = EXCLUDED.total,
    status = EXCLUDED.status
RETURNING id, sales_invoice_id, receipt_order_id, company_id, branch_id, form_number, transaction_date, contact_book_id, secondary_company_id, konekin_id, currency_code, total_items, total, is_deleted, status, created_at, updated_at
`

type UpsertPurchaseInvoiceParams struct {
	ID                 string    `db:"id"`
	SalesInvoiceID     string    `db:"sales_invoice_id"`
	ReceiptOrderID     string    `db:"receipt_order_id"`
	CompanyID          string    `db:"company_id"`
	BranchID           string    `db:"branch_id"`
	FormNumber         string    `db:"form_number"`
	TransactionDate    time.Time `db:"transaction_date"`
	ContactBookID      string    `db:"contact_book_id"`
	SecondaryCompanyID string    `db:"secondary_company_id"`
	KonekinID          string    `db:"konekin_id"`
	CurrencyCode       string    `db:"currency_code"`
	TotalItems         int64     `db:"total_items"`
	Total              int64     `db:"total"`
	Status             string    `db:"status"`
}

func (q *Queries) UpsertPurchaseInvoice(ctx context.Context, arg UpsertPurchaseInvoiceParams) (PurchasingPurchaseInvoice, error) {
	row := q.db.QueryRowContext(ctx, upsertPurchaseInvoice,
		arg.ID,
		arg.SalesInvoiceID,
		arg.ReceiptOrderID,
		arg.CompanyID,
		arg.BranchID,
		arg.FormNumber,
		arg.TransactionDate,
		arg.ContactBookID,
		arg.SecondaryCompanyID,
		arg.KonekinID,
		arg.CurrencyCode,
		arg.TotalItems,
		arg.Total,
		arg.Status,
	)
	var i PurchasingPurchaseInvoice
	err := row.Scan(
		&i.ID,
		&i.SalesInvoiceID,
		&i.ReceiptOrderID,
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

const upsertPurchaseOrder = `-- name: UpsertPurchaseOrder :one
INSERT INTO purchasing.purchase_orders(
        id, sales_order_id, company_id, branch_id, form_number, transaction_date,
        contact_book_id, secondary_company_id, konekin_id, payment_term,
        currency_code, shipping_date, receiving_warehouse_id
    )
VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12, $13) ON CONFLICT (id) DO
UPDATE
SET sales_order_id = EXCLUDED.sales_order_id,
    company_id = EXCLUDED.company_id,
    branch_id = EXCLUDED.branch_id,
    form_number = EXCLUDED.form_number,
    transaction_date = EXCLUDED.transaction_date,
    contact_book_id = EXCLUDED.contact_book_id,
    secondary_company_id = EXCLUDED.secondary_company_id,
    konekin_id = EXCLUDED.konekin_id,
    payment_term = EXCLUDED.payment_term,
    currency_code = EXCLUDED.currency_code,
    shipping_date = EXCLUDED.shipping_date,
    receiving_warehouse_id = EXCLUDED.receiving_warehouse_id,
    updated_at = NOW()
RETURNING id, sales_order_id, company_id, branch_id, form_number, transaction_date, contact_book_id, secondary_company_id, konekin_id, payment_term, currency_code, shipping_date, receiving_warehouse_id, total_items, total, is_deleted, status, created_at, updated_at
`

type UpsertPurchaseOrderParams struct {
	ID                   string    `db:"id"`
	SalesOrderID         string    `db:"sales_order_id"`
	CompanyID            string    `db:"company_id"`
	BranchID             string    `db:"branch_id"`
	FormNumber           string    `db:"form_number"`
	TransactionDate      time.Time `db:"transaction_date"`
	ContactBookID        string    `db:"contact_book_id"`
	SecondaryCompanyID   string    `db:"secondary_company_id"`
	KonekinID            string    `db:"konekin_id"`
	PaymentTerm          int32     `db:"payment_term"`
	CurrencyCode         string    `db:"currency_code"`
	ShippingDate         time.Time `db:"shipping_date"`
	ReceivingWarehouseID string    `db:"receiving_warehouse_id"`
}

func (q *Queries) UpsertPurchaseOrder(ctx context.Context, arg UpsertPurchaseOrderParams) (PurchasingPurchaseOrder, error) {
	row := q.db.QueryRowContext(ctx, upsertPurchaseOrder,
		arg.ID,
		arg.SalesOrderID,
		arg.CompanyID,
		arg.BranchID,
		arg.FormNumber,
		arg.TransactionDate,
		arg.ContactBookID,
		arg.SecondaryCompanyID,
		arg.KonekinID,
		arg.PaymentTerm,
		arg.CurrencyCode,
		arg.ShippingDate,
		arg.ReceivingWarehouseID,
	)
	var i PurchasingPurchaseOrder
	err := row.Scan(
		&i.ID,
		&i.SalesOrderID,
		&i.CompanyID,
		&i.BranchID,
		&i.FormNumber,
		&i.TransactionDate,
		&i.ContactBookID,
		&i.SecondaryCompanyID,
		&i.KonekinID,
		&i.PaymentTerm,
		&i.CurrencyCode,
		&i.ShippingDate,
		&i.ReceivingWarehouseID,
		&i.TotalItems,
		&i.Total,
		&i.IsDeleted,
		&i.Status,
		&i.CreatedAt,
		&i.UpdatedAt,
	)
	return i, err
}

const upsertPurchaseOrderItem = `-- name: UpsertPurchaseOrderItem :one
INSERT INTO purchasing.purchase_order_items(
        id, sales_order_item_id, purchase_order_id,
        primary_item_variant_id, secondary_item_variant_id,
        primary_item_unit_id, secondary_item_unit_id,
        primary_item_unit_value, secondary_item_unit_value,
        amount, price
    )
VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11) ON CONFLICT (id) DO
UPDATE
SET sales_order_item_id = EXCLUDED.sales_order_item_id,
    purchase_order_id = EXCLUDED.purchase_order_id,
    primary_item_variant_id = EXCLUDED.primary_item_variant_id,
    secondary_item_variant_id = EXCLUDED.secondary_item_variant_id,
    primary_item_unit_id = EXCLUDED.primary_item_unit_id,
    secondary_item_unit_id = EXCLUDED.secondary_item_unit_id,
    primary_item_unit_value = EXCLUDED.primary_item_unit_value,
    secondary_item_unit_value = EXCLUDED.secondary_item_unit_value,
    amount = EXCLUDED.amount,
    price = EXCLUDED.price,
    updated_at = NOW()
RETURNING id, sales_order_item_id, purchase_order_id, primary_item_variant_id, secondary_item_variant_id, primary_item_unit_id, secondary_item_unit_id, primary_item_unit_value, secondary_item_unit_value, amount, amount_received, amount_invoiced, price, is_deleted, created_at, updated_at
`

type UpsertPurchaseOrderItemParams struct {
	ID                     string `db:"id"`
	SalesOrderItemID       string `db:"sales_order_item_id"`
	PurchaseOrderID        string `db:"purchase_order_id"`
	PrimaryItemVariantID   string `db:"primary_item_variant_id"`
	SecondaryItemVariantID string `db:"secondary_item_variant_id"`
	PrimaryItemUnitID      string `db:"primary_item_unit_id"`
	SecondaryItemUnitID    string `db:"secondary_item_unit_id"`
	PrimaryItemUnitValue   int64  `db:"primary_item_unit_value"`
	SecondaryItemUnitValue int64  `db:"secondary_item_unit_value"`
	Amount                 int64  `db:"amount"`
	Price                  int64  `db:"price"`
}

func (q *Queries) UpsertPurchaseOrderItem(ctx context.Context, arg UpsertPurchaseOrderItemParams) (PurchasingPurchaseOrderItem, error) {
	row := q.db.QueryRowContext(ctx, upsertPurchaseOrderItem,
		arg.ID,
		arg.SalesOrderItemID,
		arg.PurchaseOrderID,
		arg.PrimaryItemVariantID,
		arg.SecondaryItemVariantID,
		arg.PrimaryItemUnitID,
		arg.SecondaryItemUnitID,
		arg.PrimaryItemUnitValue,
		arg.SecondaryItemUnitValue,
		arg.Amount,
		arg.Price,
	)
	var i PurchasingPurchaseOrderItem
	err := row.Scan(
		&i.ID,
		&i.SalesOrderItemID,
		&i.PurchaseOrderID,
		&i.PrimaryItemVariantID,
		&i.SecondaryItemVariantID,
		&i.PrimaryItemUnitID,
		&i.SecondaryItemUnitID,
		&i.PrimaryItemUnitValue,
		&i.SecondaryItemUnitValue,
		&i.Amount,
		&i.AmountReceived,
		&i.AmountInvoiced,
		&i.Price,
		&i.IsDeleted,
		&i.CreatedAt,
		&i.UpdatedAt,
	)
	return i, err
}

const upsertPurchaseSetting = `-- name: UpsertPurchaseSetting :one
INSERT INTO purchasing.purchase_settings(
        company_id, is_auto_approve_order
    )
VALUES ($1, $2) ON CONFLICT (company_id) DO
UPDATE
SET is_auto_approve_order = EXCLUDED.is_auto_approve_order,
    updated_at = NOW()
RETURNING company_id, is_auto_approve_order, created_at, updated_at
`

type UpsertPurchaseSettingParams struct {
	CompanyID          string `db:"company_id"`
	IsAutoApproveOrder bool   `db:"is_auto_approve_order"`
}

func (q *Queries) UpsertPurchaseSetting(ctx context.Context, arg UpsertPurchaseSettingParams) (PurchasingPurchaseSetting, error) {
	row := q.db.QueryRowContext(ctx, upsertPurchaseSetting, arg.CompanyID, arg.IsAutoApproveOrder)
	var i PurchasingPurchaseSetting
	err := row.Scan(
		&i.CompanyID,
		&i.IsAutoApproveOrder,
		&i.CreatedAt,
		&i.UpdatedAt,
	)
	return i, err
}

const upsertReceiptOrder = `-- name: UpsertReceiptOrder :one
INSERT INTO purchasing.receipt_orders(
        id, delivery_order_id, warehouse_id, company_id, branch_id, form_number, transaction_date,
        contact_book_id, secondary_company_id, konekin_id, total_items
    )
VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11) ON CONFLICT (id) DO
UPDATE
SET delivery_order_id = EXCLUDED.delivery_order_id,
    warehouse_id = EXCLUDED.warehouse_id,
    company_id = EXCLUDED.company_id,
    branch_id = EXCLUDED.branch_id,
    form_number = EXCLUDED.form_number,
    transaction_date = EXCLUDED.transaction_date,
    contact_book_id = EXCLUDED.contact_book_id,
    secondary_company_id = EXCLUDED.secondary_company_id,
    konekin_id = EXCLUDED.konekin_id,
    total_items = EXCLUDED.total_items,
    updated_at = NOW()
RETURNING id, delivery_order_id, company_id, branch_id, warehouse_id, form_number, transaction_date, contact_book_id, secondary_company_id, konekin_id, total_items, is_deleted, status, created_at, updated_at
`

type UpsertReceiptOrderParams struct {
	ID                 string    `db:"id"`
	DeliveryOrderID    string    `db:"delivery_order_id"`
	WarehouseID        string    `db:"warehouse_id"`
	CompanyID          string    `db:"company_id"`
	BranchID           string    `db:"branch_id"`
	FormNumber         string    `db:"form_number"`
	TransactionDate    time.Time `db:"transaction_date"`
	ContactBookID      string    `db:"contact_book_id"`
	SecondaryCompanyID string    `db:"secondary_company_id"`
	KonekinID          string    `db:"konekin_id"`
	TotalItems         int64     `db:"total_items"`
}

func (q *Queries) UpsertReceiptOrder(ctx context.Context, arg UpsertReceiptOrderParams) (PurchasingReceiptOrder, error) {
	row := q.db.QueryRowContext(ctx, upsertReceiptOrder,
		arg.ID,
		arg.DeliveryOrderID,
		arg.WarehouseID,
		arg.CompanyID,
		arg.BranchID,
		arg.FormNumber,
		arg.TransactionDate,
		arg.ContactBookID,
		arg.SecondaryCompanyID,
		arg.KonekinID,
		arg.TotalItems,
	)
	var i PurchasingReceiptOrder
	err := row.Scan(
		&i.ID,
		&i.DeliveryOrderID,
		&i.CompanyID,
		&i.BranchID,
		&i.WarehouseID,
		&i.FormNumber,
		&i.TransactionDate,
		&i.ContactBookID,
		&i.SecondaryCompanyID,
		&i.KonekinID,
		&i.TotalItems,
		&i.IsDeleted,
		&i.Status,
		&i.CreatedAt,
		&i.UpdatedAt,
	)
	return i, err
}
