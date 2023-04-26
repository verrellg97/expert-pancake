// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.16.0
// source: query.sql

package db

import (
	"context"
	"time"
)

const getPurchaseOrderItems = `-- name: GetPurchaseOrderItems :many
SELECT 
    id, sales_order_item_id, purchase_order_id, primary_item_variant_id, secondary_item_variant_id, primary_item_unit_id, secondary_item_unit_id, primary_item_unit_value, secondary_item_unit_value, amount, price, is_deleted, created_at, updated_at
FROM sales.purchase_order_items
WHERE purchase_order_id = $1 AND is_deleted = FALSE
`

func (q *Queries) GetPurchaseOrderItems(ctx context.Context, purchaseOrderID string) ([]SalesPurchaseOrderItem, error) {
	rows, err := q.db.QueryContext(ctx, getPurchaseOrderItems, purchaseOrderID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []SalesPurchaseOrderItem
	for rows.Next() {
		var i SalesPurchaseOrderItem
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
    id, sales_order_id, company_id, branch_id, form_number, transaction_date, contact_book_id, secondary_company_id, konekin_id, currency_code, total_items, total, is_deleted, status, created_at, updated_at
FROM sales.purchase_orders
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

func (q *Queries) GetPurchaseOrders(ctx context.Context, arg GetPurchaseOrdersParams) ([]SalesPurchaseOrder, error) {
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
	var items []SalesPurchaseOrder
	for rows.Next() {
		var i SalesPurchaseOrder
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

const updatePurchaseOrderAddItem = `-- name: UpdatePurchaseOrderAddItem :exec
UPDATE sales.purchase_orders
SET total_items=sub.total_items,
    total=sub.total
FROM (SELECT purchase_order_id, COUNT(id) AS total_items, SUM(amount*price) AS total
      FROM sales.purchase_order_items
      WHERE purchase_order_id = $1
      GROUP BY purchase_order_id) AS sub
WHERE sales.purchase_orders.id = sub.purchase_order_id
`

func (q *Queries) UpdatePurchaseOrderAddItem(ctx context.Context, purchaseOrderID string) error {
	_, err := q.db.ExecContext(ctx, updatePurchaseOrderAddItem, purchaseOrderID)
	return err
}

const upsertPurchaseOrder = `-- name: UpsertPurchaseOrder :one
INSERT INTO sales.purchase_orders(
        id, sales_order_id, company_id, branch_id, form_number, transaction_date,
        contact_book_id, secondary_company_id, konekin_id, currency_code
    )
VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10) ON CONFLICT (id) DO
UPDATE
SET sales_order_id = EXCLUDED.sales_order_id,
    company_id = EXCLUDED.company_id,
    branch_id = EXCLUDED.branch_id,
    form_number = EXCLUDED.form_number,
    transaction_date = EXCLUDED.transaction_date,
    contact_book_id = EXCLUDED.contact_book_id,
    secondary_company_id = EXCLUDED.secondary_company_id,
    konekin_id = EXCLUDED.konekin_id,
    currency_code = EXCLUDED.currency_code,
    updated_at = NOW()
RETURNING id, sales_order_id, company_id, branch_id, form_number, transaction_date, contact_book_id, secondary_company_id, konekin_id, currency_code, total_items, total, is_deleted, status, created_at, updated_at
`

type UpsertPurchaseOrderParams struct {
	ID                 string    `db:"id"`
	SalesOrderID       string    `db:"sales_order_id"`
	CompanyID          string    `db:"company_id"`
	BranchID           string    `db:"branch_id"`
	FormNumber         string    `db:"form_number"`
	TransactionDate    time.Time `db:"transaction_date"`
	ContactBookID      string    `db:"contact_book_id"`
	SecondaryCompanyID string    `db:"secondary_company_id"`
	KonekinID          string    `db:"konekin_id"`
	CurrencyCode       string    `db:"currency_code"`
}

func (q *Queries) UpsertPurchaseOrder(ctx context.Context, arg UpsertPurchaseOrderParams) (SalesPurchaseOrder, error) {
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
		arg.CurrencyCode,
	)
	var i SalesPurchaseOrder
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

const upsertPurchaseOrderItem = `-- name: UpsertPurchaseOrderItem :one
INSERT INTO sales.purchase_order_items(
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
RETURNING id, sales_order_item_id, purchase_order_id, primary_item_variant_id, secondary_item_variant_id, primary_item_unit_id, secondary_item_unit_id, primary_item_unit_value, secondary_item_unit_value, amount, price, is_deleted, created_at, updated_at
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

func (q *Queries) UpsertPurchaseOrderItem(ctx context.Context, arg UpsertPurchaseOrderItemParams) (SalesPurchaseOrderItem, error) {
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
	var i SalesPurchaseOrderItem
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
		&i.Price,
		&i.IsDeleted,
		&i.CreatedAt,
		&i.UpdatedAt,
	)
	return i, err
}