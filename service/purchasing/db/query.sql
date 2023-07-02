-- name: UpsertPurchaseOrder :one
INSERT INTO purchasing.purchase_orders(
        id, sales_order_id, company_id, branch_id, form_number, transaction_date,
        contact_book_id, secondary_company_id, konekin_id, currency_code, shipping_date,
        receiving_warehouse_id
    )
VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12) ON CONFLICT (id) DO
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
    shipping_date = EXCLUDED.shipping_date,
    receiving_warehouse_id = EXCLUDED.receiving_warehouse_id,
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

-- name: GetCheckPurchaseOrders :one
SELECT 
    COUNT(id)::bigint AS total_count
FROM purchasing.purchase_orders
WHERE company_id = $1;

-- name: GetPurchaseOrder :one
SELECT 
    *
FROM purchasing.purchase_orders
WHERE id = $1;

-- name: UpdatePurchaseOrderStatus :exec
UPDATE purchasing.purchase_orders
SET status = $2
WHERE id = $1;

-- name: UpdateAcceptedPurchaseOrder :exec
UPDATE purchasing.purchase_orders
SET sales_order_id = $2
WHERE id = $1;

-- name: UpdateAcceptedPurchaseOrderItem :exec
UPDATE purchasing.purchase_order_items
SET sales_order_item_id = $2
WHERE id = $1;

-- name: UpsertReceiptOrder :one
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
RETURNING *;

-- name: GetReceiptOrders :many
SELECT 
    *
FROM purchasing.receipt_orders
WHERE company_id = $1
    AND branch_id = $2
    AND transaction_date BETWEEN @start_date::date AND @end_date::date 
    AND is_deleted = FALSE;

-- name: DeleteReceiptOrderItems :exec
DELETE FROM purchasing.receipt_order_items
WHERE receipt_order_id = $1;

-- name: InsertReceiptOrderItem :exec
INSERT INTO purchasing.receipt_order_items(
    id, purchase_order_item_id, sales_order_item_id, delivery_order_item_id,
    receipt_order_id, primary_item_variant_id, warehouse_rack_id, batch,
    expired_date, item_barcode_id, secondary_item_variant_id,
    primary_item_unit_id, secondary_item_unit_id, primary_item_unit_value,
    secondary_item_unit_value, amount
)
VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12, $13, $14, $15, $16);

-- name: GetReceiptOrderItems :many
SELECT 
    *
FROM purchasing.receipt_order_items
WHERE receipt_order_id = $1 AND is_deleted = FALSE;

-- name: DeleteReceiptOrder :exec
UPDATE purchasing.receipt_orders
SET is_deleted = TRUE
WHERE id = $1;

-- name: UpsertPurchaseInvoice :one
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
RETURNING *;


-- name: DeletePurchaseInvoiceItems :exec
DELETE FROM purchasing.purchase_invoice_items
WHERE purchase_invoice_id = $1;

-- name: InsertPurchaseInvoiceItem :exec
INSERT INTO purchasing.purchase_invoice_items(
    id, purchase_order_item_id, sales_order_item_id, sales_invoice_item_id,
    receipt_order_item_id, purchase_invoice_id, primary_item_variant_id,
    secondary_item_variant_id, primary_item_unit_id, secondary_item_unit_id,
    primary_item_unit_value, secondary_item_unit_value, amount, price
)
VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12, $13, $14);

-- name: GetPurchaseInvoices :many
SELECT 
    a.*,
    b.form_number AS receipt_order_form_number,
    b.transaction_date AS receipt_order_transaction_date,
    b.warehouse_id AS warehouse_id
FROM purchasing.purchase_invoices a
JOIN purchasing.receipt_orders b ON b.id = a.receipt_order_id
WHERE a.company_id = $1
    AND a.branch_id = $2
    AND a.transaction_date BETWEEN @start_date::date AND @end_date::date 
    AND a.is_deleted = FALSE;

-- name: GetPurchaseInvoiceItems :many
SELECT 
    a.*,
    b.warehouse_rack_id AS warehouse_rack_id,
    b.item_barcode_id AS item_barcode_id,
    b.batch AS batch,
    b.expired_date AS expired_date
FROM purchasing.purchase_invoice_items a
JOIN purchasing.receipt_order_items b ON b.id = a.receipt_order_item_id
WHERE a.purchase_invoice_id = $1;
