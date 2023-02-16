-- name: InsertBrand :one
INSERT INTO inventory.brands(id, company_id, name)
VALUES ($1, $2, $3)
RETURNING *;

-- name: UpdateBrand :one
UPDATE inventory.brands
SET name = $2,
    updated_at = NOW()
WHERE id = $1
RETURNING *;

-- name: GetBrands :many
SELECT id,
    company_id,
    name
FROM inventory.brands
WHERE company_id = $1
    AND name LIKE $2;

-- name: InsertGroup :one
INSERT INTO inventory.groups(id, company_id, name)
VALUES ($1, $2, $3)
RETURNING *;

-- name: UpdateGroup :one
UPDATE inventory.groups
SET name = $2,
    updated_at = NOW()
WHERE id = $1
RETURNING *;

-- name: GetGroups :many
SELECT id,
    company_id,
    name
FROM inventory.groups
WHERE company_id = $1
    AND name LIKE $2;

-- name: InsertUnit :one
INSERT INTO inventory.units(id, company_id, unit_category_id, name)
VALUES ($1, $2, $3, $4)
RETURNING *;

-- name: UpdateUnit :one
UPDATE inventory.units
SET unit_category_id = $2,
    name = $3,
    updated_at = NOW()
WHERE id = $1
RETURNING *;

-- name: GetUnits :many
SELECT id,
    company_id,
    unit_category_id,
    name
FROM inventory.units
WHERE company_id = $1
    AND unit_category_id LIKE $2
    AND name LIKE $3;

-- name: InsertItem :one
INSERT INTO inventory.items(
        id,
        company_id,
        image_url,
        code,
        name,
        brand_id,
        group_id,
        tag,
        description
    )
VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9)
RETURNING *;

-- name: InsertItemVariant :one
INSERT INTO inventory.item_variants(
        id,
        item_id,
        image_url,
        barcode,
        name,
        price,
        is_default
    )
VALUES ($1, $2, $3, $4, $5, $6, $7)
RETURNING *;

-- name: GetBrandById :one
SELECT id,
    company_id,
    name
FROM inventory.brands
WHERE id = $1;

-- name: GetGroupById :one
SELECT id,
    company_id,
    name
FROM inventory.groups
WHERE id = $1;

-- name: UpdateItem :one
UPDATE inventory.items
SET image_url = $2,
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
SET image_url = $2,
    barcode = $3,
    price = $4,
    updated_at = NOW()
WHERE item_id = $1
    AND is_default = true
RETURNING *;

-- name: GetItems :many
SELECT a.id,
    b.id AS variant_id,
    a.company_id,
    b.image_url,
    a.code,
    b.barcode,
    a.name,
    b.name AS variant_name,
    a.brand_id,
    c.name AS brand_name,
    a.group_id,
    d.name AS group_name,
    a.tag,
    a.description,
    b.is_default,
    b.price
FROM inventory.items a
    JOIN inventory.item_variants b ON a.id = b.item_id
    JOIN inventory.brands c ON a.brand_id = c.id
    JOIN inventory.groups d ON a.group_id = d.id
WHERE a.company_id = $1
    AND b.name LIKE $2;

-- name: UpsertItemInfo :exec
INSERT INTO inventory.item_infos(
        item_id,
        is_purchase,
        is_sale,
        is_raw_material,
        is_asset,
        purchase_chart_of_account_id,
        sale_chart_of_account_id,
        purchase_item_unit_id
    )
VALUES ($1, $2, $3, $4, $5, $6, $7, $8) ON CONFLICT (item_id) DO
UPDATE
SET is_purchase = EXCLUDED.is_purchase,
    is_sale = EXCLUDED.is_sale,
    is_raw_material = EXCLUDED.is_raw_material,
    is_asset = EXCLUDED.is_asset,
    purchase_chart_of_account_id = EXCLUDED.purchase_chart_of_account_id,
    sale_chart_of_account_id = EXCLUDED.sale_chart_of_account_id,
    purchase_item_unit_id = EXCLUDED.purchase_item_unit_id,
    updated_at = NOW();

-- name: GetItemInfo :one
SELECT a.item_id,
    d.company_id,
    a.is_purchase,
    a.is_sale,
    a.is_raw_material,
    a.is_asset,
    a.purchase_chart_of_account_id,
    a.sale_chart_of_account_id,
    a.purchase_item_unit_id,
    c.name AS purchase_item_unit_name
FROM inventory.item_infos a
    JOIN inventory.item_units b ON a.purchase_item_unit_id = b.id
    JOIN inventory.units c ON b.unit_id = c.id
    JOIN inventory.items d ON a.item_id = d.id
WHERE a.item_id = $1;

-- name: UpsertItemVariant :exec
INSERT INTO inventory.item_variants(id, item_id, image_url, barcode, name, price)
VALUES ($1, $2, $3, $4, $5, $6) ON CONFLICT (id) DO
UPDATE
SET item_id = EXCLUDED.item_id,
    image_url = EXCLUDED.image_url,
    barcode = EXCLUDED.barcode,
    name = EXCLUDED.name,
    price = EXCLUDED.price,
    updated_at = NOW();

-- name: GetItemVariant :one
SELECT b.id,
    a.id AS variant_id,
    b.company_id,
    a.image_url,
    b.code,
    a.barcode,
    b.name,
    a.name AS variant_name,
    b.brand_id,
    c.name AS brand_name,
    b.group_id,
    d.name AS group_name,
    b.tag,
    b.description,
    a.is_default,
    a.price
FROM inventory.item_variants a
    JOIN inventory.items b ON a.item_id = b.id
    JOIN inventory.brands c ON b.brand_id = c.id
    JOIN inventory.groups d ON b.group_id = d.id
WHERE a.id = $1;

-- name: GetItemVariants :many
SELECT b.id,
    a.id AS variant_id,
    b.company_id,
    a.image_url,
    b.code,
    a.barcode,
    b.name,
    a.name AS variant_name,
    b.brand_id,
    c.name AS brand_name,
    b.group_id,
    d.name AS group_name,
    b.tag,
    b.description,
    a.is_default,
    a.price
FROM inventory.item_variants a
    JOIN inventory.items b ON a.item_id = b.id
    JOIN inventory.brands c ON b.brand_id = c.id
    JOIN inventory.groups d ON b.group_id = d.id
WHERE a.item_id = $1
    AND a.name LIKE $2;

-- name: UpsertItemUnit :one
INSERT INTO inventory.item_units(id, item_id, unit_id, value, is_default)
VALUES ($1, $2, $3, $4, $5) ON CONFLICT (id) DO
UPDATE
SET item_id = EXCLUDED.item_id,
    unit_id = EXCLUDED.unit_id,
    value = EXCLUDED.value,
    is_default = EXCLUDED.is_default,
    updated_at = NOW()
RETURNING *;

-- name: GetUnit :one
SELECT id,
    company_id,
    name
FROM inventory.units
WHERE id = $1;

-- name: GetItemUnits :many
SELECT a.id,
    a.item_id,
    a.unit_id,
    b.name AS unit_name,
    a.value,
    a.is_default
FROM inventory.item_units a
    JOIN inventory.units b ON a.unit_id = b.id
WHERE a.item_id = $1
    AND b.name LIKE $2;

-- name: InsertInternalStockTransfer :one
INSERT INTO inventory.internal_stock_transfers(
        id,
        source_warehouse_id,
        destination_warehouse_id,
        form_number,
        transaction_date
    )
VALUES ($1, $2, $3, $4, $5)
RETURNING *;

-- name: GetInternalStockTransfers :many
SELECT id,
    source_warehouse_id,
    destination_warehouse_id,
    form_number,
    transaction_date
FROM inventory.internal_stock_transfers
WHERE is_deleted = false
    AND transaction_date BETWEEN @start_date::date AND @end_date::date
    AND (
        source_warehouse_id = ANY(@warehouse_ids::text [])
        OR destination_warehouse_id = ANY(@warehouse_ids::text [])
    );

-- name: InsertInternalStockTransferItem :exec
INSERT INTO inventory.internal_stock_transfer_items(
        id,
        internal_stock_transfer_id,
        warehouse_rack_id,
        variant_id,
        item_unit_id,
        item_unit_value,
        amount,
        batch,
        expired_date,
        item_barcode_id
    )
VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10);

-- name: GetInternalStockTransferItems :many
SELECT a.id,
    a.warehouse_rack_id,
    e.name AS item_name,
    a.variant_id,
    b.name AS variant_name,
    a.item_unit_id,
    d.name AS item_unit_name,
    a.item_unit_value,
    a.amount,
    a.batch,
    a.expired_date
FROM inventory.internal_stock_transfer_items a
    JOIN inventory.item_variants b ON a.variant_id = b.id
    JOIN inventory.item_units c ON a.item_unit_id = c.id
    JOIN inventory.units d ON c.unit_id = d.id
    JOIN inventory.items e ON b.item_id = e.id
WHERE a.internal_stock_transfer_id = $1
    AND a.is_deleted = false;

-- name: InsertItemBarcode :exec
INSERT INTO inventory.item_barcodes(id, variant_id, batch, expired_date)
VALUES ($1, $2, $3, $4);

-- name: GetItemBarcode :one
SELECT id
FROM inventory.item_barcodes
WHERE variant_id = $1
    AND CASE
        WHEN @is_null_batch::bool THEN batch is null
        ELSE batch = $2
    END
    AND CASE
        WHEN @is_null_expired_date::bool THEN expired_date is null
        ELSE expired_date = $3
    END;

-- name: InsertStockMovement :exec
INSERT INTO inventory.stock_movements(
        id,
        transaction_id,
        transaction_date,
        transaction_reference,
        detail_transaction_id,
        warehouse_id,
        warehouse_rack_id,
        variant_id,
        item_barcode_id,
        amount
    )
VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10);

-- name: InsertUpdateStock :exec
INSERT INTO inventory.update_stocks(id,
form_number, transaction_date, warehouse_id, warehouse_rack_id,
variant_id, item_unit_id, item_unit_value, beginning_stock, ending_stock,
batch, expired_date, item_barcode_id)
VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12, $13);

-- name: GetUpdateStock :one
SELECT a.id, a.form_number, a.transaction_date, a.warehouse_id, a.warehouse_rack_id,
b.item_id, c.name AS item_name, a.variant_id, b.name AS variant_name,
a.item_unit_id, e.name AS item_unit_name, a.item_unit_value,
a.beginning_stock, a.ending_stock, a.batch, a.expired_date
FROM inventory.update_stocks a
JOIN inventory.item_variants b ON a.variant_id = b.id
JOIN inventory.items c ON b.item_id = c.id
JOIN inventory.item_units d ON a.item_unit_id = d.id
JOIN inventory.units e ON d.unit_id = e.id
WHERE a.id = $1;

-- name: GetUpdateStocks :many
SELECT a.id, a.form_number, a.transaction_date, a.warehouse_id, a.warehouse_rack_id,
b.item_id, c.name AS item_name, a.variant_id, b.name AS variant_name,
a.item_unit_id, e.name AS item_unit_name, a.item_unit_value,
a.beginning_stock, a.ending_stock, a.batch, a.expired_date
FROM inventory.update_stocks a
JOIN inventory.item_variants b ON a.variant_id = b.id
JOIN inventory.items c ON b.item_id = c.id
JOIN inventory.item_units d ON a.item_unit_id = d.id
JOIN inventory.units e ON d.unit_id = e.id
WHERE a.is_deleted = false
AND a.transaction_date BETWEEN @start_date::date AND @end_date::date
AND a.warehouse_id = ANY(@warehouse_ids::text[]);

-- name: UpsertItemReorder :one
INSERT INTO inventory.item_reorders(
        id,
        variant_id,
        item_unit_id,
        warehouse_id,
        minimum_stock
    )
VALUES ($1, $2, $3, $4, $5) ON CONFLICT (id) DO
UPDATE
SET variant_id = EXCLUDED.variant_id,
    item_unit_id = EXCLUDED.item_unit_id,
    warehouse_id = EXCLUDED.warehouse_id,
    minimum_stock = EXCLUDED.minimum_stock,
    updated_at = NOW()
RETURNING *;

-- name: GetItemReorder :one
SELECT a.id,
    a.variant_id,
    b.name as variant_name,
    c.id as item_id,
    c.name as item_name,
    d.id as item_unit_id,
    d.name as item_unit_name,
    a.warehouse_id,
    a.minimum_stock
FROM inventory.item_reorders a
    JOIN inventory.item_variants b ON a.variant_id = b.id
    JOIN inventory.items c ON b.item_id = c.id
    JOIN inventory.units d ON a.item_unit_id = d.id
WHERE a.id = $1;

-- name: GetItemReorders :many
SELECT a.id,
    a.variant_id,
    b.name as variant_name,
    d.id as item_unit_id,
    d.name as item_unit_name,
    c.id as item_id,
    c.name as item_name,
    a.warehouse_id,
    a.minimum_stock
FROM inventory.item_reorders a
    JOIN inventory.item_variants b ON a.variant_id = b.id
    JOIN inventory.items c ON b.item_id = c.id
    JOIN inventory.units d ON a.item_unit_id = d.id
WHERE a.warehouse_id LIKE $1
    AND b.item_id LIKE $2;

-- name: UpsertUnitCategory :one
INSERT INTO inventory.unit_categories(id, company_id, name)
VALUES ($1, $2, $3) ON CONFLICT (id) DO
UPDATE
SET company_id = EXCLUDED.company_id,
    name = EXCLUDED.name,
    updated_at = NOW()
RETURNING *;

-- name: GetUnitCategories :many
SELECT id,
    company_id,
    name
FROM inventory.unit_categories
WHERE company_id = $1
    AND name LIKE $2;

-- name: GetVariantWarehouseRacks :many
SELECT DISTINCT a.warehouse_rack_id
FROM inventory.stock_movements a
WHERE a.variant_id = $1
AND a.warehouse_id = $2;

-- name: GetVariantWarehouseRackBatches :many
SELECT DISTINCT b.batch
FROM inventory.stock_movements a
JOIN inventory.item_barcodes b ON a.item_barcode_id = b.id
WHERE a.variant_id = $1
AND a.warehouse_rack_id = $2;

-- name: GetVariantWarehouseRackBatchExpiredDates :many
SELECT DISTINCT b.expired_date
FROM inventory.stock_movements a
JOIN inventory.item_barcodes b ON a.item_barcode_id = b.id
WHERE a.variant_id = $1
AND a.warehouse_rack_id = $2
AND CASE
    WHEN @is_null_batch::bool THEN b.batch is null
    ELSE b.batch = $3
END;

-- name: GetVariantWarehouseRackStock :one
SELECT SUM(a.amount) AS stock
FROM inventory.stock_movements a
JOIN inventory.item_barcodes b ON a.item_barcode_id = b.id
WHERE a.variant_id = $1
AND a.warehouse_rack_id = $2
AND CASE
    WHEN @is_null_batch::bool THEN b.batch is null
    ELSE b.batch = $3
END
AND CASE
    WHEN @is_null_expired_date::bool THEN b.expired_date is null
    ELSE b.expired_date = $4
END;

-- name: GetTransferHistory :many
SELECT b.transaction_date,
    b.form_number,
    item.id as item_id,
    item.name as item_name,
    item.image_url,
    variant.id as variant_id,
    variant.name as variant_name,
    b.source_warehouse_id,
    b.destination_warehouse_id,
    a.amount
FROM inventory.internal_stock_transfer_items a
    JOIN inventory.item_variants variant ON variant.id = a.variant_id
    JOIN inventory.items item ON item.id = variant.item_id
    JOIN inventory.internal_stock_transfers b ON b.id = a.internal_stock_transfer_id
WHERE a.is_deleted = FALSE
    AND b.is_deleted = FALSE
    AND variant.item_id LIKE $1
    AND transaction_date BETWEEN @start_date::date AND @end_date::date
    AND (
        source_warehouse_id LIKE $2
        OR destination_warehouse_id LIKE $3
    )
GROUP BY b.transaction_date,
    a.variant_id,
    item.id,
    variant.id,
    item.image_url,
    b.form_number,
    b.source_warehouse_id,
    b.destination_warehouse_id,
    a.amount
ORDER BY b.transaction_date DESC;