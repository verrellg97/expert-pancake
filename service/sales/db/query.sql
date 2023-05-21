-- name: UpsertPOS :one
INSERT INTO sales.point_of_sales(
  id, company_id, branch_id, warehouse_id, form_number, transaction_date,
  contact_book_id, secondary_company_id, konekin_id, currency_code, chart_of_account_id, total_items, total, updated_at
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
    chart_of_account_id = EXCLUDED.chart_of_account_id,
    total_items = EXCLUDED.total_items,
    total = EXCLUDED.total,
    updated_at = NOW()
RETURNING *;

-- name: InsertPOSItem :one
INSERT INTO sales.point_of_sale_items(
  id, point_of_sale_id, warehouse_rack_id, item_variant_id, item_unit_id, item_unit_value, batch, expired_date, item_barcode_id, amount, price, updated_at
)
VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12)
RETURNING *;

-- name: DeletePOSItemsPOS :exec
DELETE FROM sales.point_of_sale_items WHERE point_of_sale_id = $1;

-- name: DeletePOS :exec
UPDATE sales.point_of_sales SET is_deleted = TRUE WHERE id = $1;

-- name: GetPOS :many
SELECT a.* FROM sales.point_of_sales a 
WHERE company_id LIKE $1
    AND branch_id LIKE $2
    AND transaction_date BETWEEN @start_date::date AND @end_date::date 
    AND is_deleted = FALSE;

-- name: GetPOSItemsByPOSId :many
SELECT a.* FROM sales.point_of_sale_items a WHERE a.point_of_sale_id = $1 ORDER BY a.id;

-- name: GetPOSUserSetting :one
SELECT 
    *
FROM sales.pos_user_settings
WHERE user_id = $1
AND branch_id = $2;

-- name: UpsertPOSUserSetting :one
INSERT INTO sales.pos_user_settings(
  user_id, branch_id, warehouse_id, warehouse_rack_id
)
VALUES ($1, $2, $3, $4) ON CONFLICT (user_id, branch_id) DO
UPDATE
SET warehouse_id = EXCLUDED.warehouse_id,
  warehouse_rack_id = EXCLUDED.warehouse_rack_id,
  updated_at = NOW()
RETURNING *;

-- name: InsertPOSCOASetting :one
INSERT INTO sales.pos_chart_of_account_settings(
  branch_id, chart_of_account_id
)
VALUES ($1, $2)
RETURNING *;

-- name: DeletePOSCOASetting :exec
DELETE FROM sales.pos_chart_of_account_settings WHERE branch_id = $1;

-- name: GetPOSCOASetting :many
SELECT 
    *
FROM sales.pos_chart_of_account_settings
WHERE branch_id = $1;

-- name: InsertPOSCustomerSetting :one
INSERT INTO sales.pos_customer_settings(
  branch_id, contact_book_id
)
VALUES ($1, $2)
RETURNING *;

-- name: DeletePOSCustomerSetting :exec
DELETE FROM sales.pos_customer_settings WHERE branch_id = $1;