-- name: UpsertPurchaseOrder :one
INSERT INTO purchasing.purchase_orders(
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
RETURNING *;

-- name: GetPurchaseOrders :many
SELECT 
    *
FROM purchasing.purchase_orders
WHERE company_id = $1
    AND branch_id = $2
    AND transaction_date BETWEEN @start_date::date AND @end_date::date 
    AND is_deleted = FALSE;

-- name: GetPurchaseOrderItems :many
SELECT 
    *
FROM purchasing.purchase_order_items
WHERE purchase_order_id = $1 AND is_deleted = FALSE;

-- name: UpdatePurchaseOrderAddItem :exec
UPDATE purchasing.purchase_orders
SET total_items=sub.total_items,
    total=sub.total,
    updated_at = NOW()
FROM (SELECT purchase_order_id, COUNT(id) AS total_items, SUM(amount*price) AS total
      FROM purchasing.purchase_order_items
      WHERE purchase_order_id = @purchase_order_id
      GROUP BY purchase_order_id) AS sub
WHERE purchasing.purchase_orders.id = sub.purchase_order_id;

-- name: UpsertPurchaseOrderItem :one
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
RETURNING *;

-- name: DeletePurchaseOrderItems :exec
DELETE FROM purchasing.purchase_order_items
WHERE purchase_order_id = $1;

-- name: GetPurchaseSetting :one
SELECT 
    *
FROM purchasing.purchase_settings
WHERE company_id = $1;

-- name: UpsertPurchaseSetting :one
INSERT INTO purchasing.purchase_settings(
        company_id, is_auto_approve_order
    )
VALUES ($1, $2) ON CONFLICT (company_id) DO
UPDATE
SET is_auto_approve_order = EXCLUDED.is_auto_approve_order,
    updated_at = NOW()
RETURNING *;