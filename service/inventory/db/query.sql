-- name: InsertBrand :one
INSERT INTO inventory.brands(id, company_id, name)
VALUES ($1, $2, $3)
RETURNING *;

-- name: UpdateBrand :one
UPDATE inventory.brands
SET 
    name = $2,
    updated_at = NOW()
WHERE id = $1
RETURNING *;

-- name: GetBrands :many
SELECT id, company_id, name FROM inventory.brands
WHERE company_id = $1 AND name LIKE $2;

-- name: InsertGroup :one
INSERT INTO inventory.groups(id, company_id, name)
VALUES ($1, $2, $3)
RETURNING *;

-- name: UpdateGroup :one
UPDATE inventory.groups
SET 
    name = $2,
    updated_at = NOW()
WHERE id = $1
RETURNING *;

-- name: GetGroups :many
SELECT id, company_id, name FROM inventory.groups
WHERE company_id = $1 AND name LIKE $2;

-- name: InsertUnit :one
INSERT INTO inventory.units(id, company_id, name)
VALUES ($1, $2, $3)
RETURNING *;

-- name: UpdateUnit :one
UPDATE inventory.units
SET 
    name = $2,
    updated_at = NOW()
WHERE id = $1
RETURNING *;

-- name: GetUnits :many
SELECT id, company_id, name FROM inventory.units
WHERE company_id = $1 AND name LIKE $2;

-- name: InsertItem :one
INSERT INTO inventory.items(id, company_id, image_url,
code, name, brand_id, group_id, tag, description)
VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9)
RETURNING *;

-- name: InsertItemVariant :one
INSERT INTO inventory.item_variants(id, item_id, image_url,
name, price, stock, is_default)
VALUES ($1, $2, $3, $4, $5, $6, $7)
RETURNING *;

-- name: GetBrandById :one
SELECT id, company_id, name FROM inventory.brands
WHERE id = $1;

-- name: GetGroupById :one
SELECT id, company_id, name FROM inventory.groups
WHERE id = $1;

-- name: UpdateItem :one
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
RETURNING *;

-- name: UpdateItemVariantDefault :one
UPDATE inventory.item_variants
SET 
    image_url = $2,
    name = $3,
    updated_at = NOW()
WHERE item_id = $1
AND is_default = true
RETURNING *;

-- name: GetItems :many
SELECT a.id, b.id AS variant_id, a.company_id, b.image_url, a.code, a.name, b.name AS variant_name,
a.brand_id, c.name AS brand_name, a.group_id, d.name AS group_name,
a.tag, a.description, b.is_default, b.price, b.stock
FROM inventory.items a
JOIN inventory.item_variants b ON a.id = b.item_id
JOIN inventory.brands c ON a.brand_id = c.id
JOIN inventory.groups d ON a.group_id = d.id
WHERE a.company_id = $1 AND b.name LIKE $2;

-- name: UpsertItemVariant :exec
INSERT INTO inventory.item_variants(id, item_id, image_url, name, price, stock)
VALUES ($1, $2, $3, $4, $5, $6)
ON CONFLICT (id)
DO UPDATE SET
    item_id = EXCLUDED.item_id,
    image_url = EXCLUDED.image_url,
    name = EXCLUDED.name,
    price = EXCLUDED.price,
    stock = EXCLUDED.stock,
    updated_at = NOW();

-- name: GetItemVariant :one
SELECT b.id, a.id AS variant_id, b.company_id, a.image_url, b.code, b.name, a.name AS variant_name,
b.brand_id, c.name AS brand_name, b.group_id, d.name AS group_name,
b.tag, b.description, a.is_default, a.price, a.stock
FROM inventory.item_variants a
JOIN inventory.items b ON a.item_id = b.id
JOIN inventory.brands c ON b.brand_id = c.id
JOIN inventory.groups d ON b.group_id = d.id
WHERE a.id = $1;

-- name: GetItemVariants :many
SELECT b.id, a.id AS variant_id, b.company_id, a.image_url, b.code, b.name, a.name AS variant_name,
b.brand_id, c.name AS brand_name, b.group_id, d.name AS group_name,
b.tag, b.description, a.is_default, a.price, a.stock
FROM inventory.item_variants a
JOIN inventory.items b ON a.item_id = b.id
JOIN inventory.brands c ON b.brand_id = c.id
JOIN inventory.groups d ON b.group_id = d.id
WHERE a.item_id = $1 AND a.name LIKE $2
AND a.is_default = false;

-- name: UpsertItemUnit :one
INSERT INTO inventory.item_units(id, item_id, unit_id, value, is_default)
VALUES ($1, $2, $3, $4, $5)
ON CONFLICT (id)
DO UPDATE SET
    item_id = EXCLUDED.item_id,
    unit_id = EXCLUDED.unit_id,
    value = EXCLUDED.value,
    is_default = EXCLUDED.is_default,
    updated_at = NOW()
RETURNING *;

-- name: GetUnit :one
SELECT id, company_id, name
FROM inventory.units
WHERE id = $1;

-- name: GetItemUnits :many
SELECT a.id, a.item_id, a.unit_id, b.name AS unit_name,
a.value, a.is_default
FROM inventory.item_units a
JOIN inventory.units b ON a.unit_id = b.id
WHERE a.item_id = $1 AND b.name LIKE $2;

-- name: InsertInternalStockTransfer :one
INSERT INTO inventory.internal_stock_transfers(id,
source_warehouse_id, destination_warehouse_id, form_number, transaction_date)
VALUES ($1, $2, $3, $4, $5)
RETURNING *;

-- name: GetInternalStockTransfers :many
SELECT id, source_warehouse_id, destination_warehouse_id,
form_number, transaction_date
FROM inventory.internal_stock_transfers
WHERE is_deleted = false
AND transaction_date BETWEEN @start_date::date AND @end_date::date;

-- name: InsertInternalStockTransferItem :exec
INSERT INTO inventory.internal_stock_transfer_items(id,
internal_stock_transfer_id, warehouse_rack_id, variant_id,
item_unit_id, item_unit_value, amount, batch, expired_date,
item_barcode_id)
VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10);

-- name: GetInternalStockTransferItems :many
SELECT a.id, a.warehouse_rack_id, e.name AS item_name, a.variant_id, b.name AS variant_name,
a.item_unit_id, d.name AS item_unit_name, a.item_unit_value, a.amount, a.batch, a.expired_date
FROM inventory.internal_stock_transfer_items a
JOIN inventory.item_variants b ON a.variant_id = b.id
JOIN inventory.item_units c ON a.item_unit_id = c.id
JOIN inventory.units d ON c.unit_id = d.id
JOIN inventory.items e ON b.item_id = e.id
WHERE a.internal_stock_transfer_id = $1 AND a.is_deleted = false;

-- name: InsertItemBarcode :exec
INSERT INTO inventory.item_barcodes(id, variant_id, batch, expired_date)
VALUES ($1, $2, $3, $4);

-- name: GetItemBarcode :one
SELECT id
FROM inventory.item_barcodes
WHERE variant_id = $1
AND CASE WHEN @is_null_batch::bool THEN batch is null
ELSE batch = $2 END
AND CASE WHEN @is_null_expired_date::bool THEN expired_date is null
ELSE expired_date = $3 END;

-- name: InsertStockMovement :exec
INSERT INTO inventory.stock_movements(id, transaction_id, transaction_date,
transaction_reference, detail_transaction_id, warehouse_id, warehouse_rack_id,
variant_id, item_barcode_id, amount)
VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10);

-- name: UpsertItemReorder :one
INSERT INTO inventory.item_reorders(id, variant_id, warehouse_id, minimum_stock)
VALUES ($1, $2, $3, $4)
ON CONFLICT (id)
DO UPDATE SET
    variant_id = EXCLUDED.variant_id,
    warehouse_id = EXCLUDED.warehouse_id,
    minimum_stock = EXCLUDED.minimum_stock,
    updated_at = NOW()
RETURNING *;

-- name: GetItemReorder :one
SELECT a.id, a.variant_id, b.name as variant_name, c.id as item_id, c.name as item_name, a.warehouse_id, a.minimum_stock
FROM inventory.item_reorders a
JOIN inventory.item_variants b ON a.variant_id = b.id
JOIN inventory.items c ON b.item_id = c.id
WHERE a.id = $1;
