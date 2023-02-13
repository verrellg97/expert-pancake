// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.16.0
// source: query.sql

package db

import (
	"context"
	"database/sql"
	"time"
)

const getBrandById = `-- name: GetBrandById :one
SELECT id, company_id, name FROM inventory.brands
WHERE id = $1
`

type GetBrandByIdRow struct {
	ID        string `db:"id"`
	CompanyID string `db:"company_id"`
	Name      string `db:"name"`
}

func (q *Queries) GetBrandById(ctx context.Context, id string) (GetBrandByIdRow, error) {
	row := q.db.QueryRowContext(ctx, getBrandById, id)
	var i GetBrandByIdRow
	err := row.Scan(&i.ID, &i.CompanyID, &i.Name)
	return i, err
}

const getBrands = `-- name: GetBrands :many
SELECT id, company_id, name FROM inventory.brands
WHERE company_id = $1 AND name LIKE $2
`

type GetBrandsParams struct {
	CompanyID string `db:"company_id"`
	Name      string `db:"name"`
}

type GetBrandsRow struct {
	ID        string `db:"id"`
	CompanyID string `db:"company_id"`
	Name      string `db:"name"`
}

func (q *Queries) GetBrands(ctx context.Context, arg GetBrandsParams) ([]GetBrandsRow, error) {
	rows, err := q.db.QueryContext(ctx, getBrands, arg.CompanyID, arg.Name)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []GetBrandsRow
	for rows.Next() {
		var i GetBrandsRow
		if err := rows.Scan(&i.ID, &i.CompanyID, &i.Name); err != nil {
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

const getGroupById = `-- name: GetGroupById :one
SELECT id, company_id, name FROM inventory.groups
WHERE id = $1
`

type GetGroupByIdRow struct {
	ID        string `db:"id"`
	CompanyID string `db:"company_id"`
	Name      string `db:"name"`
}

func (q *Queries) GetGroupById(ctx context.Context, id string) (GetGroupByIdRow, error) {
	row := q.db.QueryRowContext(ctx, getGroupById, id)
	var i GetGroupByIdRow
	err := row.Scan(&i.ID, &i.CompanyID, &i.Name)
	return i, err
}

const getGroups = `-- name: GetGroups :many
SELECT id, company_id, name FROM inventory.groups
WHERE company_id = $1 AND name LIKE $2
`

type GetGroupsParams struct {
	CompanyID string `db:"company_id"`
	Name      string `db:"name"`
}

type GetGroupsRow struct {
	ID        string `db:"id"`
	CompanyID string `db:"company_id"`
	Name      string `db:"name"`
}

func (q *Queries) GetGroups(ctx context.Context, arg GetGroupsParams) ([]GetGroupsRow, error) {
	rows, err := q.db.QueryContext(ctx, getGroups, arg.CompanyID, arg.Name)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []GetGroupsRow
	for rows.Next() {
		var i GetGroupsRow
		if err := rows.Scan(&i.ID, &i.CompanyID, &i.Name); err != nil {
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

const getInternalStockTransferItems = `-- name: GetInternalStockTransferItems :many
SELECT a.id, a.warehouse_rack_id, e.name AS item_name, a.variant_id, b.name AS variant_name,
a.item_unit_id, d.name AS item_unit_name, a.item_unit_value, a.amount, a.batch, a.expired_date
FROM inventory.internal_stock_transfer_items a
JOIN inventory.item_variants b ON a.variant_id = b.id
JOIN inventory.item_units c ON a.item_unit_id = c.id
JOIN inventory.units d ON c.unit_id = d.id
JOIN inventory.items e ON b.item_id = e.id
WHERE a.internal_stock_transfer_id = $1 AND a.is_deleted = false
`

type GetInternalStockTransferItemsRow struct {
	ID              string         `db:"id"`
	WarehouseRackID string         `db:"warehouse_rack_id"`
	ItemName        string         `db:"item_name"`
	VariantID       string         `db:"variant_id"`
	VariantName     string         `db:"variant_name"`
	ItemUnitID      string         `db:"item_unit_id"`
	ItemUnitName    string         `db:"item_unit_name"`
	ItemUnitValue   int64          `db:"item_unit_value"`
	Amount          int64          `db:"amount"`
	Batch           sql.NullString `db:"batch"`
	ExpiredDate     sql.NullTime   `db:"expired_date"`
}

func (q *Queries) GetInternalStockTransferItems(ctx context.Context, internalStockTransferID string) ([]GetInternalStockTransferItemsRow, error) {
	rows, err := q.db.QueryContext(ctx, getInternalStockTransferItems, internalStockTransferID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []GetInternalStockTransferItemsRow
	for rows.Next() {
		var i GetInternalStockTransferItemsRow
		if err := rows.Scan(
			&i.ID,
			&i.WarehouseRackID,
			&i.ItemName,
			&i.VariantID,
			&i.VariantName,
			&i.ItemUnitID,
			&i.ItemUnitName,
			&i.ItemUnitValue,
			&i.Amount,
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

const getInternalStockTransfers = `-- name: GetInternalStockTransfers :many
SELECT id, source_warehouse_id, destination_warehouse_id,
form_number, transaction_date
FROM inventory.internal_stock_transfers
WHERE is_deleted = false
AND transaction_date BETWEEN $1::date AND $2::date
`

type GetInternalStockTransfersParams struct {
	StartDate time.Time `db:"start_date"`
	EndDate   time.Time `db:"end_date"`
}

type GetInternalStockTransfersRow struct {
	ID                     string    `db:"id"`
	SourceWarehouseID      string    `db:"source_warehouse_id"`
	DestinationWarehouseID string    `db:"destination_warehouse_id"`
	FormNumber             string    `db:"form_number"`
	TransactionDate        time.Time `db:"transaction_date"`
}

func (q *Queries) GetInternalStockTransfers(ctx context.Context, arg GetInternalStockTransfersParams) ([]GetInternalStockTransfersRow, error) {
	rows, err := q.db.QueryContext(ctx, getInternalStockTransfers, arg.StartDate, arg.EndDate)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []GetInternalStockTransfersRow
	for rows.Next() {
		var i GetInternalStockTransfersRow
		if err := rows.Scan(
			&i.ID,
			&i.SourceWarehouseID,
			&i.DestinationWarehouseID,
			&i.FormNumber,
			&i.TransactionDate,
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

const getItemBarcode = `-- name: GetItemBarcode :one
SELECT id
FROM inventory.item_barcodes
WHERE variant_id = $1
AND CASE WHEN $4::bool THEN batch is null
ELSE batch = $2 END
AND CASE WHEN $5::bool THEN expired_date is null
ELSE expired_date = $3 END
`

type GetItemBarcodeParams struct {
	VariantID         string         `db:"variant_id"`
	Batch             sql.NullString `db:"batch"`
	ExpiredDate       sql.NullTime   `db:"expired_date"`
	IsNullBatch       bool           `db:"is_null_batch"`
	IsNullExpiredDate bool           `db:"is_null_expired_date"`
}

func (q *Queries) GetItemBarcode(ctx context.Context, arg GetItemBarcodeParams) (string, error) {
	row := q.db.QueryRowContext(ctx, getItemBarcode,
		arg.VariantID,
		arg.Batch,
		arg.ExpiredDate,
		arg.IsNullBatch,
		arg.IsNullExpiredDate,
	)
	var id string
	err := row.Scan(&id)
	return id, err
}

const getItemReorder = `-- name: GetItemReorder :one
SELECT
    a.id, 
    a.variant_id,
    b.name as variant_name, 
    c.id as item_id, c.name as item_name,
    d.id as item_unit_id, d.name as item_unit_name, 
    a.warehouse_id, 
    a.minimum_stock
FROM inventory.item_reorders a
JOIN inventory.item_variants b ON a.variant_id = b.id
JOIN inventory.items c ON b.item_id = c.id
JOIN inventory.units d ON a.item_unit_id = d.id
WHERE a.id = $1
`

type GetItemReorderRow struct {
	ID           string `db:"id"`
	VariantID    string `db:"variant_id"`
	VariantName  string `db:"variant_name"`
	ItemID       string `db:"item_id"`
	ItemName     string `db:"item_name"`
	ItemUnitID   string `db:"item_unit_id"`
	ItemUnitName string `db:"item_unit_name"`
	WarehouseID  string `db:"warehouse_id"`
	MinimumStock int64  `db:"minimum_stock"`
}

func (q *Queries) GetItemReorder(ctx context.Context, id string) (GetItemReorderRow, error) {
	row := q.db.QueryRowContext(ctx, getItemReorder, id)
	var i GetItemReorderRow
	err := row.Scan(
		&i.ID,
		&i.VariantID,
		&i.VariantName,
		&i.ItemID,
		&i.ItemName,
		&i.ItemUnitID,
		&i.ItemUnitName,
		&i.WarehouseID,
		&i.MinimumStock,
	)
	return i, err
}

const getItemReorders = `-- name: GetItemReorders :many
SELECT 
    a.id, 
    a.variant_id, 
    b.name as variant_name,
    d.id as item_unit_id, d.name as item_unit_name,
    c.id as item_id, c.name as item_name, 
    a.warehouse_id, 
    a.minimum_stock
FROM inventory.item_reorders a
JOIN inventory.item_variants b ON a.variant_id = b.id
JOIN inventory.items c ON b.item_id = c.id
JOIN inventory.units d ON a.item_unit_id = d.id
WHERE a.warehouse_id LIKE $1 AND b.item_id LIKE $2
`

type GetItemReordersParams struct {
	WarehouseID string `db:"warehouse_id"`
	ItemID      string `db:"item_id"`
}

type GetItemReordersRow struct {
	ID           string `db:"id"`
	VariantID    string `db:"variant_id"`
	VariantName  string `db:"variant_name"`
	ItemUnitID   string `db:"item_unit_id"`
	ItemUnitName string `db:"item_unit_name"`
	ItemID       string `db:"item_id"`
	ItemName     string `db:"item_name"`
	WarehouseID  string `db:"warehouse_id"`
	MinimumStock int64  `db:"minimum_stock"`
}

func (q *Queries) GetItemReorders(ctx context.Context, arg GetItemReordersParams) ([]GetItemReordersRow, error) {
	rows, err := q.db.QueryContext(ctx, getItemReorders, arg.WarehouseID, arg.ItemID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []GetItemReordersRow
	for rows.Next() {
		var i GetItemReordersRow
		if err := rows.Scan(
			&i.ID,
			&i.VariantID,
			&i.VariantName,
			&i.ItemUnitID,
			&i.ItemUnitName,
			&i.ItemID,
			&i.ItemName,
			&i.WarehouseID,
			&i.MinimumStock,
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

const getItemUnits = `-- name: GetItemUnits :many
SELECT a.id, a.item_id, a.unit_id, b.name AS unit_name,
a.value, a.is_default
FROM inventory.item_units a
JOIN inventory.units b ON a.unit_id = b.id
WHERE a.item_id = $1 AND b.name LIKE $2
`

type GetItemUnitsParams struct {
	ItemID string `db:"item_id"`
	Name   string `db:"name"`
}

type GetItemUnitsRow struct {
	ID        string `db:"id"`
	ItemID    string `db:"item_id"`
	UnitID    string `db:"unit_id"`
	UnitName  string `db:"unit_name"`
	Value     int64  `db:"value"`
	IsDefault bool   `db:"is_default"`
}

func (q *Queries) GetItemUnits(ctx context.Context, arg GetItemUnitsParams) ([]GetItemUnitsRow, error) {
	rows, err := q.db.QueryContext(ctx, getItemUnits, arg.ItemID, arg.Name)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []GetItemUnitsRow
	for rows.Next() {
		var i GetItemUnitsRow
		if err := rows.Scan(
			&i.ID,
			&i.ItemID,
			&i.UnitID,
			&i.UnitName,
			&i.Value,
			&i.IsDefault,
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

const getItemVariant = `-- name: GetItemVariant :one
SELECT b.id, a.id AS variant_id, b.company_id, a.image_url, b.code, a.barcode, b.name, a.name AS variant_name,
b.brand_id, c.name AS brand_name, b.group_id, d.name AS group_name,
b.tag, b.description, a.is_default, a.price
FROM inventory.item_variants a
JOIN inventory.items b ON a.item_id = b.id
JOIN inventory.brands c ON b.brand_id = c.id
JOIN inventory.groups d ON b.group_id = d.id
WHERE a.id = $1
`

type GetItemVariantRow struct {
	ID          string `db:"id"`
	VariantID   string `db:"variant_id"`
	CompanyID   string `db:"company_id"`
	ImageUrl    string `db:"image_url"`
	Code        string `db:"code"`
	Barcode     string `db:"barcode"`
	Name        string `db:"name"`
	VariantName string `db:"variant_name"`
	BrandID     string `db:"brand_id"`
	BrandName   string `db:"brand_name"`
	GroupID     string `db:"group_id"`
	GroupName   string `db:"group_name"`
	Tag         string `db:"tag"`
	Description string `db:"description"`
	IsDefault   bool   `db:"is_default"`
	Price       int64  `db:"price"`
}

func (q *Queries) GetItemVariant(ctx context.Context, id string) (GetItemVariantRow, error) {
	row := q.db.QueryRowContext(ctx, getItemVariant, id)
	var i GetItemVariantRow
	err := row.Scan(
		&i.ID,
		&i.VariantID,
		&i.CompanyID,
		&i.ImageUrl,
		&i.Code,
		&i.Barcode,
		&i.Name,
		&i.VariantName,
		&i.BrandID,
		&i.BrandName,
		&i.GroupID,
		&i.GroupName,
		&i.Tag,
		&i.Description,
		&i.IsDefault,
		&i.Price,
	)
	return i, err
}

const getItemVariants = `-- name: GetItemVariants :many
SELECT b.id, a.id AS variant_id, b.company_id, a.image_url, b.code, a.barcode, b.name, a.name AS variant_name,
b.brand_id, c.name AS brand_name, b.group_id, d.name AS group_name,
b.tag, b.description, a.is_default, a.price
FROM inventory.item_variants a
JOIN inventory.items b ON a.item_id = b.id
JOIN inventory.brands c ON b.brand_id = c.id
JOIN inventory.groups d ON b.group_id = d.id
WHERE a.item_id = $1 AND a.name LIKE $2
`

type GetItemVariantsParams struct {
	ItemID string `db:"item_id"`
	Name   string `db:"name"`
}

type GetItemVariantsRow struct {
	ID          string `db:"id"`
	VariantID   string `db:"variant_id"`
	CompanyID   string `db:"company_id"`
	ImageUrl    string `db:"image_url"`
	Code        string `db:"code"`
	Barcode     string `db:"barcode"`
	Name        string `db:"name"`
	VariantName string `db:"variant_name"`
	BrandID     string `db:"brand_id"`
	BrandName   string `db:"brand_name"`
	GroupID     string `db:"group_id"`
	GroupName   string `db:"group_name"`
	Tag         string `db:"tag"`
	Description string `db:"description"`
	IsDefault   bool   `db:"is_default"`
	Price       int64  `db:"price"`
}

func (q *Queries) GetItemVariants(ctx context.Context, arg GetItemVariantsParams) ([]GetItemVariantsRow, error) {
	rows, err := q.db.QueryContext(ctx, getItemVariants, arg.ItemID, arg.Name)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []GetItemVariantsRow
	for rows.Next() {
		var i GetItemVariantsRow
		if err := rows.Scan(
			&i.ID,
			&i.VariantID,
			&i.CompanyID,
			&i.ImageUrl,
			&i.Code,
			&i.Barcode,
			&i.Name,
			&i.VariantName,
			&i.BrandID,
			&i.BrandName,
			&i.GroupID,
			&i.GroupName,
			&i.Tag,
			&i.Description,
			&i.IsDefault,
			&i.Price,
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

const getItems = `-- name: GetItems :many
SELECT a.id, b.id AS variant_id, a.company_id, b.image_url, a.code, b.barcode, a.name, b.name AS variant_name,
a.brand_id, c.name AS brand_name, a.group_id, d.name AS group_name,
a.tag, a.description, b.is_default, b.price
FROM inventory.items a
JOIN inventory.item_variants b ON a.id = b.item_id
JOIN inventory.brands c ON a.brand_id = c.id
JOIN inventory.groups d ON a.group_id = d.id
WHERE a.company_id = $1 AND b.name LIKE $2
`

type GetItemsParams struct {
	CompanyID string `db:"company_id"`
	Name      string `db:"name"`
}

type GetItemsRow struct {
	ID          string `db:"id"`
	VariantID   string `db:"variant_id"`
	CompanyID   string `db:"company_id"`
	ImageUrl    string `db:"image_url"`
	Code        string `db:"code"`
	Barcode     string `db:"barcode"`
	Name        string `db:"name"`
	VariantName string `db:"variant_name"`
	BrandID     string `db:"brand_id"`
	BrandName   string `db:"brand_name"`
	GroupID     string `db:"group_id"`
	GroupName   string `db:"group_name"`
	Tag         string `db:"tag"`
	Description string `db:"description"`
	IsDefault   bool   `db:"is_default"`
	Price       int64  `db:"price"`
}

func (q *Queries) GetItems(ctx context.Context, arg GetItemsParams) ([]GetItemsRow, error) {
	rows, err := q.db.QueryContext(ctx, getItems, arg.CompanyID, arg.Name)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []GetItemsRow
	for rows.Next() {
		var i GetItemsRow
		if err := rows.Scan(
			&i.ID,
			&i.VariantID,
			&i.CompanyID,
			&i.ImageUrl,
			&i.Code,
			&i.Barcode,
			&i.Name,
			&i.VariantName,
			&i.BrandID,
			&i.BrandName,
			&i.GroupID,
			&i.GroupName,
			&i.Tag,
			&i.Description,
			&i.IsDefault,
			&i.Price,
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

const getUnit = `-- name: GetUnit :one
SELECT id, company_id, name
FROM inventory.units
WHERE id = $1
`

type GetUnitRow struct {
	ID        string `db:"id"`
	CompanyID string `db:"company_id"`
	Name      string `db:"name"`
}

func (q *Queries) GetUnit(ctx context.Context, id string) (GetUnitRow, error) {
	row := q.db.QueryRowContext(ctx, getUnit, id)
	var i GetUnitRow
	err := row.Scan(&i.ID, &i.CompanyID, &i.Name)
	return i, err
}

const getUnitCategories = `-- name: GetUnitCategories :many
SELECT id, company_id, name
FROM inventory.unit_categories
WHERE company_id = $1 AND name LIKE $2
`

type GetUnitCategoriesParams struct {
	CompanyID string `db:"company_id"`
	Name      string `db:"name"`
}

type GetUnitCategoriesRow struct {
	ID        string `db:"id"`
	CompanyID string `db:"company_id"`
	Name      string `db:"name"`
}

func (q *Queries) GetUnitCategories(ctx context.Context, arg GetUnitCategoriesParams) ([]GetUnitCategoriesRow, error) {
	rows, err := q.db.QueryContext(ctx, getUnitCategories, arg.CompanyID, arg.Name)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []GetUnitCategoriesRow
	for rows.Next() {
		var i GetUnitCategoriesRow
		if err := rows.Scan(&i.ID, &i.CompanyID, &i.Name); err != nil {
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

const getUnits = `-- name: GetUnits :many
SELECT id, company_id, unit_category_id, name FROM inventory.units
WHERE company_id = $1 AND unit_category_id LIKE $2 AND name LIKE $3
`

type GetUnitsParams struct {
	CompanyID      string `db:"company_id"`
	UnitCategoryID string `db:"unit_category_id"`
	Name           string `db:"name"`
}

type GetUnitsRow struct {
	ID             string `db:"id"`
	CompanyID      string `db:"company_id"`
	UnitCategoryID string `db:"unit_category_id"`
	Name           string `db:"name"`
}

func (q *Queries) GetUnits(ctx context.Context, arg GetUnitsParams) ([]GetUnitsRow, error) {
	rows, err := q.db.QueryContext(ctx, getUnits, arg.CompanyID, arg.UnitCategoryID, arg.Name)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []GetUnitsRow
	for rows.Next() {
		var i GetUnitsRow
		if err := rows.Scan(
			&i.ID,
			&i.CompanyID,
			&i.UnitCategoryID,
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

const insertBrand = `-- name: InsertBrand :one
INSERT INTO inventory.brands(id, company_id, name)
VALUES ($1, $2, $3)
RETURNING id, company_id, name, created_at, updated_at
`

type InsertBrandParams struct {
	ID        string `db:"id"`
	CompanyID string `db:"company_id"`
	Name      string `db:"name"`
}

func (q *Queries) InsertBrand(ctx context.Context, arg InsertBrandParams) (InventoryBrand, error) {
	row := q.db.QueryRowContext(ctx, insertBrand, arg.ID, arg.CompanyID, arg.Name)
	var i InventoryBrand
	err := row.Scan(
		&i.ID,
		&i.CompanyID,
		&i.Name,
		&i.CreatedAt,
		&i.UpdatedAt,
	)
	return i, err
}

const insertGroup = `-- name: InsertGroup :one
INSERT INTO inventory.groups(id, company_id, name)
VALUES ($1, $2, $3)
RETURNING id, company_id, name, created_at, updated_at
`

type InsertGroupParams struct {
	ID        string `db:"id"`
	CompanyID string `db:"company_id"`
	Name      string `db:"name"`
}

func (q *Queries) InsertGroup(ctx context.Context, arg InsertGroupParams) (InventoryGroup, error) {
	row := q.db.QueryRowContext(ctx, insertGroup, arg.ID, arg.CompanyID, arg.Name)
	var i InventoryGroup
	err := row.Scan(
		&i.ID,
		&i.CompanyID,
		&i.Name,
		&i.CreatedAt,
		&i.UpdatedAt,
	)
	return i, err
}

const insertInternalStockTransfer = `-- name: InsertInternalStockTransfer :one
INSERT INTO inventory.internal_stock_transfers(id,
source_warehouse_id, destination_warehouse_id, form_number, transaction_date)
VALUES ($1, $2, $3, $4, $5)
RETURNING id, source_warehouse_id, destination_warehouse_id, form_number, transaction_date, is_deleted, created_at, updated_at
`

type InsertInternalStockTransferParams struct {
	ID                     string    `db:"id"`
	SourceWarehouseID      string    `db:"source_warehouse_id"`
	DestinationWarehouseID string    `db:"destination_warehouse_id"`
	FormNumber             string    `db:"form_number"`
	TransactionDate        time.Time `db:"transaction_date"`
}

func (q *Queries) InsertInternalStockTransfer(ctx context.Context, arg InsertInternalStockTransferParams) (InventoryInternalStockTransfer, error) {
	row := q.db.QueryRowContext(ctx, insertInternalStockTransfer,
		arg.ID,
		arg.SourceWarehouseID,
		arg.DestinationWarehouseID,
		arg.FormNumber,
		arg.TransactionDate,
	)
	var i InventoryInternalStockTransfer
	err := row.Scan(
		&i.ID,
		&i.SourceWarehouseID,
		&i.DestinationWarehouseID,
		&i.FormNumber,
		&i.TransactionDate,
		&i.IsDeleted,
		&i.CreatedAt,
		&i.UpdatedAt,
	)
	return i, err
}

const insertInternalStockTransferItem = `-- name: InsertInternalStockTransferItem :exec
INSERT INTO inventory.internal_stock_transfer_items(id,
internal_stock_transfer_id, warehouse_rack_id, variant_id,
item_unit_id, item_unit_value, amount, batch, expired_date,
item_barcode_id)
VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10)
`

type InsertInternalStockTransferItemParams struct {
	ID                      string         `db:"id"`
	InternalStockTransferID string         `db:"internal_stock_transfer_id"`
	WarehouseRackID         string         `db:"warehouse_rack_id"`
	VariantID               string         `db:"variant_id"`
	ItemUnitID              string         `db:"item_unit_id"`
	ItemUnitValue           int64          `db:"item_unit_value"`
	Amount                  int64          `db:"amount"`
	Batch                   sql.NullString `db:"batch"`
	ExpiredDate             sql.NullTime   `db:"expired_date"`
	ItemBarcodeID           string         `db:"item_barcode_id"`
}

func (q *Queries) InsertInternalStockTransferItem(ctx context.Context, arg InsertInternalStockTransferItemParams) error {
	_, err := q.db.ExecContext(ctx, insertInternalStockTransferItem,
		arg.ID,
		arg.InternalStockTransferID,
		arg.WarehouseRackID,
		arg.VariantID,
		arg.ItemUnitID,
		arg.ItemUnitValue,
		arg.Amount,
		arg.Batch,
		arg.ExpiredDate,
		arg.ItemBarcodeID,
	)
	return err
}

const insertItem = `-- name: InsertItem :one
INSERT INTO inventory.items(id, company_id, image_url,
code, name, brand_id, group_id, tag, description)
VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9)
RETURNING id, company_id, image_url, code, name, brand_id, group_id, tag, description, created_at, updated_at
`

type InsertItemParams struct {
	ID          string `db:"id"`
	CompanyID   string `db:"company_id"`
	ImageUrl    string `db:"image_url"`
	Code        string `db:"code"`
	Name        string `db:"name"`
	BrandID     string `db:"brand_id"`
	GroupID     string `db:"group_id"`
	Tag         string `db:"tag"`
	Description string `db:"description"`
}

func (q *Queries) InsertItem(ctx context.Context, arg InsertItemParams) (InventoryItem, error) {
	row := q.db.QueryRowContext(ctx, insertItem,
		arg.ID,
		arg.CompanyID,
		arg.ImageUrl,
		arg.Code,
		arg.Name,
		arg.BrandID,
		arg.GroupID,
		arg.Tag,
		arg.Description,
	)
	var i InventoryItem
	err := row.Scan(
		&i.ID,
		&i.CompanyID,
		&i.ImageUrl,
		&i.Code,
		&i.Name,
		&i.BrandID,
		&i.GroupID,
		&i.Tag,
		&i.Description,
		&i.CreatedAt,
		&i.UpdatedAt,
	)
	return i, err
}

const insertItemBarcode = `-- name: InsertItemBarcode :exec
INSERT INTO inventory.item_barcodes(id, variant_id, batch, expired_date)
VALUES ($1, $2, $3, $4)
`

type InsertItemBarcodeParams struct {
	ID          string         `db:"id"`
	VariantID   string         `db:"variant_id"`
	Batch       sql.NullString `db:"batch"`
	ExpiredDate sql.NullTime   `db:"expired_date"`
}

func (q *Queries) InsertItemBarcode(ctx context.Context, arg InsertItemBarcodeParams) error {
	_, err := q.db.ExecContext(ctx, insertItemBarcode,
		arg.ID,
		arg.VariantID,
		arg.Batch,
		arg.ExpiredDate,
	)
	return err
}

const insertItemVariant = `-- name: InsertItemVariant :one
INSERT INTO inventory.item_variants(id, item_id, image_url,
barcode, name, price, is_default)
VALUES ($1, $2, $3, $4, $5, $6, $7)
RETURNING id, item_id, image_url, barcode, name, price, is_default, created_at, updated_at
`

type InsertItemVariantParams struct {
	ID        string `db:"id"`
	ItemID    string `db:"item_id"`
	ImageUrl  string `db:"image_url"`
	Barcode   string `db:"barcode"`
	Name      string `db:"name"`
	Price     int64  `db:"price"`
	IsDefault bool   `db:"is_default"`
}

func (q *Queries) InsertItemVariant(ctx context.Context, arg InsertItemVariantParams) (InventoryItemVariant, error) {
	row := q.db.QueryRowContext(ctx, insertItemVariant,
		arg.ID,
		arg.ItemID,
		arg.ImageUrl,
		arg.Barcode,
		arg.Name,
		arg.Price,
		arg.IsDefault,
	)
	var i InventoryItemVariant
	err := row.Scan(
		&i.ID,
		&i.ItemID,
		&i.ImageUrl,
		&i.Barcode,
		&i.Name,
		&i.Price,
		&i.IsDefault,
		&i.CreatedAt,
		&i.UpdatedAt,
	)
	return i, err
}

const insertStockMovement = `-- name: InsertStockMovement :exec
INSERT INTO inventory.stock_movements(id, transaction_id, transaction_date,
transaction_reference, detail_transaction_id, warehouse_id, warehouse_rack_id,
variant_id, item_barcode_id, amount)
VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10)
`

type InsertStockMovementParams struct {
	ID                   string    `db:"id"`
	TransactionID        string    `db:"transaction_id"`
	TransactionDate      time.Time `db:"transaction_date"`
	TransactionReference string    `db:"transaction_reference"`
	DetailTransactionID  string    `db:"detail_transaction_id"`
	WarehouseID          string    `db:"warehouse_id"`
	WarehouseRackID      string    `db:"warehouse_rack_id"`
	VariantID            string    `db:"variant_id"`
	ItemBarcodeID        string    `db:"item_barcode_id"`
	Amount               int64     `db:"amount"`
}

func (q *Queries) InsertStockMovement(ctx context.Context, arg InsertStockMovementParams) error {
	_, err := q.db.ExecContext(ctx, insertStockMovement,
		arg.ID,
		arg.TransactionID,
		arg.TransactionDate,
		arg.TransactionReference,
		arg.DetailTransactionID,
		arg.WarehouseID,
		arg.WarehouseRackID,
		arg.VariantID,
		arg.ItemBarcodeID,
		arg.Amount,
	)
	return err
}

const insertUnit = `-- name: InsertUnit :one
INSERT INTO inventory.units(id, company_id, unit_category_id, name)
VALUES ($1, $2, $3, $4)
RETURNING id, unit_category_id, company_id, name, created_at, updated_at
`

type InsertUnitParams struct {
	ID             string `db:"id"`
	CompanyID      string `db:"company_id"`
	UnitCategoryID string `db:"unit_category_id"`
	Name           string `db:"name"`
}

func (q *Queries) InsertUnit(ctx context.Context, arg InsertUnitParams) (InventoryUnit, error) {
	row := q.db.QueryRowContext(ctx, insertUnit,
		arg.ID,
		arg.CompanyID,
		arg.UnitCategoryID,
		arg.Name,
	)
	var i InventoryUnit
	err := row.Scan(
		&i.ID,
		&i.UnitCategoryID,
		&i.CompanyID,
		&i.Name,
		&i.CreatedAt,
		&i.UpdatedAt,
	)
	return i, err
}

const updateBrand = `-- name: UpdateBrand :one
UPDATE inventory.brands
SET 
    name = $2,
    updated_at = NOW()
WHERE id = $1
RETURNING id, company_id, name, created_at, updated_at
`

type UpdateBrandParams struct {
	ID   string `db:"id"`
	Name string `db:"name"`
}

func (q *Queries) UpdateBrand(ctx context.Context, arg UpdateBrandParams) (InventoryBrand, error) {
	row := q.db.QueryRowContext(ctx, updateBrand, arg.ID, arg.Name)
	var i InventoryBrand
	err := row.Scan(
		&i.ID,
		&i.CompanyID,
		&i.Name,
		&i.CreatedAt,
		&i.UpdatedAt,
	)
	return i, err
}

const updateGroup = `-- name: UpdateGroup :one
UPDATE inventory.groups
SET 
    name = $2,
    updated_at = NOW()
WHERE id = $1
RETURNING id, company_id, name, created_at, updated_at
`

type UpdateGroupParams struct {
	ID   string `db:"id"`
	Name string `db:"name"`
}

func (q *Queries) UpdateGroup(ctx context.Context, arg UpdateGroupParams) (InventoryGroup, error) {
	row := q.db.QueryRowContext(ctx, updateGroup, arg.ID, arg.Name)
	var i InventoryGroup
	err := row.Scan(
		&i.ID,
		&i.CompanyID,
		&i.Name,
		&i.CreatedAt,
		&i.UpdatedAt,
	)
	return i, err
}

const updateItem = `-- name: UpdateItem :one
UPDATE inventory.items
SET 
    image_url = $2,
    name = $3,
    brand_id = $4,
    group_id = $5,
    tag = $6,
    description = $7,
    updated_at = NOW()
WHERE id = $1
RETURNING id, company_id, image_url, code, name, brand_id, group_id, tag, description, created_at, updated_at
`

type UpdateItemParams struct {
	ID          string `db:"id"`
	ImageUrl    string `db:"image_url"`
	Name        string `db:"name"`
	BrandID     string `db:"brand_id"`
	GroupID     string `db:"group_id"`
	Tag         string `db:"tag"`
	Description string `db:"description"`
}

func (q *Queries) UpdateItem(ctx context.Context, arg UpdateItemParams) (InventoryItem, error) {
	row := q.db.QueryRowContext(ctx, updateItem,
		arg.ID,
		arg.ImageUrl,
		arg.Name,
		arg.BrandID,
		arg.GroupID,
		arg.Tag,
		arg.Description,
	)
	var i InventoryItem
	err := row.Scan(
		&i.ID,
		&i.CompanyID,
		&i.ImageUrl,
		&i.Code,
		&i.Name,
		&i.BrandID,
		&i.GroupID,
		&i.Tag,
		&i.Description,
		&i.CreatedAt,
		&i.UpdatedAt,
	)
	return i, err
}

const updateItemVariantDefault = `-- name: UpdateItemVariantDefault :one
UPDATE inventory.item_variants
SET 
    image_url = $2,
    barcode = $3,
    price = $4,
    updated_at = NOW()
WHERE item_id = $1
AND is_default = true
RETURNING id, item_id, image_url, barcode, name, price, is_default, created_at, updated_at
`

type UpdateItemVariantDefaultParams struct {
	ItemID   string `db:"item_id"`
	ImageUrl string `db:"image_url"`
	Barcode  string `db:"barcode"`
	Price    int64  `db:"price"`
}

func (q *Queries) UpdateItemVariantDefault(ctx context.Context, arg UpdateItemVariantDefaultParams) (InventoryItemVariant, error) {
	row := q.db.QueryRowContext(ctx, updateItemVariantDefault,
		arg.ItemID,
		arg.ImageUrl,
		arg.Barcode,
		arg.Price,
	)
	var i InventoryItemVariant
	err := row.Scan(
		&i.ID,
		&i.ItemID,
		&i.ImageUrl,
		&i.Barcode,
		&i.Name,
		&i.Price,
		&i.IsDefault,
		&i.CreatedAt,
		&i.UpdatedAt,
	)
	return i, err
}

const updateUnit = `-- name: UpdateUnit :one
UPDATE inventory.units
SET 
    unit_category_id = $2,
    name = $3,
    updated_at = NOW()
WHERE id = $1
RETURNING id, unit_category_id, company_id, name, created_at, updated_at
`

type UpdateUnitParams struct {
	ID             string `db:"id"`
	UnitCategoryID string `db:"unit_category_id"`
	Name           string `db:"name"`
}

func (q *Queries) UpdateUnit(ctx context.Context, arg UpdateUnitParams) (InventoryUnit, error) {
	row := q.db.QueryRowContext(ctx, updateUnit, arg.ID, arg.UnitCategoryID, arg.Name)
	var i InventoryUnit
	err := row.Scan(
		&i.ID,
		&i.UnitCategoryID,
		&i.CompanyID,
		&i.Name,
		&i.CreatedAt,
		&i.UpdatedAt,
	)
	return i, err
}

const upsertItemReorder = `-- name: UpsertItemReorder :one
INSERT INTO inventory.item_reorders(id, variant_id, item_unit_id, warehouse_id, minimum_stock)
VALUES ($1, $2, $3, $4, $5)
ON CONFLICT (id)
DO UPDATE SET
    variant_id = EXCLUDED.variant_id,
    item_unit_id = EXCLUDED.item_unit_id,
    warehouse_id = EXCLUDED.warehouse_id,
    minimum_stock = EXCLUDED.minimum_stock,
    updated_at = NOW()
RETURNING id, warehouse_id, item_unit_id, variant_id, minimum_stock, created_at, updated_at
`

type UpsertItemReorderParams struct {
	ID           string `db:"id"`
	VariantID    string `db:"variant_id"`
	ItemUnitID   string `db:"item_unit_id"`
	WarehouseID  string `db:"warehouse_id"`
	MinimumStock int64  `db:"minimum_stock"`
}

func (q *Queries) UpsertItemReorder(ctx context.Context, arg UpsertItemReorderParams) (InventoryItemReorder, error) {
	row := q.db.QueryRowContext(ctx, upsertItemReorder,
		arg.ID,
		arg.VariantID,
		arg.ItemUnitID,
		arg.WarehouseID,
		arg.MinimumStock,
	)
	var i InventoryItemReorder
	err := row.Scan(
		&i.ID,
		&i.WarehouseID,
		&i.ItemUnitID,
		&i.VariantID,
		&i.MinimumStock,
		&i.CreatedAt,
		&i.UpdatedAt,
	)
	return i, err
}

const upsertItemUnit = `-- name: UpsertItemUnit :one
INSERT INTO inventory.item_units(id, item_id, unit_id, value, is_default)
VALUES ($1, $2, $3, $4, $5)
ON CONFLICT (id)
DO UPDATE SET
    item_id = EXCLUDED.item_id,
    unit_id = EXCLUDED.unit_id,
    value = EXCLUDED.value,
    is_default = EXCLUDED.is_default,
    updated_at = NOW()
RETURNING id, item_id, unit_id, value, is_default, created_at, updated_at
`

type UpsertItemUnitParams struct {
	ID        string `db:"id"`
	ItemID    string `db:"item_id"`
	UnitID    string `db:"unit_id"`
	Value     int64  `db:"value"`
	IsDefault bool   `db:"is_default"`
}

func (q *Queries) UpsertItemUnit(ctx context.Context, arg UpsertItemUnitParams) (InventoryItemUnit, error) {
	row := q.db.QueryRowContext(ctx, upsertItemUnit,
		arg.ID,
		arg.ItemID,
		arg.UnitID,
		arg.Value,
		arg.IsDefault,
	)
	var i InventoryItemUnit
	err := row.Scan(
		&i.ID,
		&i.ItemID,
		&i.UnitID,
		&i.Value,
		&i.IsDefault,
		&i.CreatedAt,
		&i.UpdatedAt,
	)
	return i, err
}

const upsertItemVariant = `-- name: UpsertItemVariant :exec
INSERT INTO inventory.item_variants(id, item_id, image_url, barcode, name, price)
VALUES ($1, $2, $3, $4, $5, $6)
ON CONFLICT (id)
DO UPDATE SET
    item_id = EXCLUDED.item_id,
    image_url = EXCLUDED.image_url,
    barcode = EXCLUDED.barcode,
    name = EXCLUDED.name,
    price = EXCLUDED.price,
    updated_at = NOW()
`

type UpsertItemVariantParams struct {
	ID       string `db:"id"`
	ItemID   string `db:"item_id"`
	ImageUrl string `db:"image_url"`
	Barcode  string `db:"barcode"`
	Name     string `db:"name"`
	Price    int64  `db:"price"`
}

func (q *Queries) UpsertItemVariant(ctx context.Context, arg UpsertItemVariantParams) error {
	_, err := q.db.ExecContext(ctx, upsertItemVariant,
		arg.ID,
		arg.ItemID,
		arg.ImageUrl,
		arg.Barcode,
		arg.Name,
		arg.Price,
	)
	return err
}

const upsertUnitCategory = `-- name: UpsertUnitCategory :one
INSERT INTO inventory.unit_categories(id, company_id, name)
VALUES ($1, $2, $3)
ON CONFLICT (id)
DO UPDATE SET
    company_id = EXCLUDED.company_id,
    name = EXCLUDED.name,
    updated_at = NOW()
RETURNING id, company_id, name, created_at, updated_at
`

type UpsertUnitCategoryParams struct {
	ID        string `db:"id"`
	CompanyID string `db:"company_id"`
	Name      string `db:"name"`
}

func (q *Queries) UpsertUnitCategory(ctx context.Context, arg UpsertUnitCategoryParams) (InventoryUnitCategory, error) {
	row := q.db.QueryRowContext(ctx, upsertUnitCategory, arg.ID, arg.CompanyID, arg.Name)
	var i InventoryUnitCategory
	err := row.Scan(
		&i.ID,
		&i.CompanyID,
		&i.Name,
		&i.CreatedAt,
		&i.UpdatedAt,
	)
	return i, err
}
