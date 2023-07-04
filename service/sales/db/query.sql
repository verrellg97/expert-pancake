-- name: UpsertPOS :one
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
    AND a.transaction_date BETWEEN @start_date::date AND @end_date::date 
    AND a.is_deleted = FALSE;

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

-- name: GetPOSCustomerSetting :many
SELECT 
    *
FROM sales.pos_customer_settings
WHERE branch_id = $1;


-- name: UpsertPOSPaymentMethod :exec
INSERT INTO sales.pos_payment_methods(id, company_id, chart_of_account_id, name)
VALUES ($1, $2, $3, $4)
ON CONFLICT (id) DO UPDATE
SET name = EXCLUDED.name,
  chart_of_account_id = EXCLUDED.chart_of_account_id,
  company_id = EXCLUDED.company_id,
  updated_at = NOW();

-- name: DeletePOSPaymentMethod :exec
UPDATE sales.pos_payment_methods SET is_deleted = TRUE WHERE id = $1;

-- name: GetPOSPaymentMethod :many
SELECT id, company_id, chart_of_account_id, name
FROM sales.pos_payment_methods 
WHERE is_deleted = FALSE AND company_id = $1 AND name LIKE $2;

-- name: GetCheckPOS :one
SELECT 
    COUNT(id)::bigint AS total_count
FROM sales.point_of_sales
WHERE company_id = $1;

-- name: UpsertSalesOrder :one
INSERT INTO sales.sales_orders(
    id, purchase_order_id, purchase_order_branch_id, purchase_order_receiving_warehouse_id,
    company_id, branch_id, form_number, transaction_date,
    contact_book_id, secondary_company_id, konekin_id, currency_code,
    is_all_branches
)
VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12, $13) ON CONFLICT (id) DO
UPDATE
SET purchase_order_id = EXCLUDED.purchase_order_id,
    purchase_order_branch_id = EXCLUDED.purchase_order_branch_id,
    purchase_order_receiving_warehouse_id = EXCLUDED.purchase_order_receiving_warehouse_id,
    company_id = EXCLUDED.company_id,
    branch_id = EXCLUDED.branch_id,
    form_number = EXCLUDED.form_number,
    transaction_date = EXCLUDED.transaction_date,
    contact_book_id = EXCLUDED.contact_book_id,
    secondary_company_id = EXCLUDED.secondary_company_id,
    konekin_id = EXCLUDED.konekin_id,
    currency_code = EXCLUDED.currency_code,
    is_all_branches = EXCLUDED.is_all_branches,
    updated_at = NOW()
RETURNING *;

-- name: UpdateSalesOrderAddItem :exec
UPDATE sales.sales_orders
SET total_items=sub.total_items,
    total=sub.total,
    updated_at = NOW()
FROM (SELECT sales_order_id, COUNT(id) AS total_items, SUM(amount*price) AS total
      FROM sales.sales_order_items
      WHERE sales_order_id = @sales_order_id
      GROUP BY sales_order_id) AS sub
WHERE sales.sales_orders.id = sub.sales_order_id;

-- name: UpsertSalesOrderItem :one
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
RETURNING *;

-- name: DeleteSalesOrderItems :exec
DELETE FROM sales.sales_order_items
WHERE sales_order_id = $1;

-- name: GetSalesOrders :many
SELECT 
    *
FROM sales.sales_orders
WHERE company_id = @company_id::text
AND branch_id = @branch_id::text
AND transaction_date BETWEEN @start_date::date AND @end_date::date 
AND is_deleted = FALSE
UNION ALL
SELECT 
    *
FROM sales.sales_orders
WHERE company_id = @company_id::text
AND branch_id = '' AND is_all_branches = TRUE
AND transaction_date BETWEEN @start_date::date AND @end_date::date 
AND is_deleted = FALSE
UNION ALL
SELECT 
    a.*
FROM sales.sales_orders a
JOIN sales.sales_order_branches b ON a.id = b.sales_order_id
AND b.company_branch_id = @branch_id::text
WHERE a.company_id = @company_id::text
AND a.branch_id = '' AND a.is_all_branches = FALSE
AND a.transaction_date BETWEEN @start_date::date AND @end_date::date 
AND a.is_deleted = FALSE;

-- name: GetSalesOrderItems :many
SELECT 
    *
FROM sales.sales_order_items
WHERE sales_order_id = $1 AND is_deleted = FALSE;

-- name: GetSalesOrder :one
SELECT 
    *
FROM sales.sales_orders
WHERE id = $1;

-- name: UpdateSalesOrderStatus :exec
UPDATE sales.sales_orders
SET status = $2, branch_id = $3
WHERE id = $1;

-- name: UpsertDeliveryOrder :one
INSERT INTO sales.delivery_orders(
    id, sales_order_id, company_id, branch_id,
    form_number, transaction_date,
    contact_book_id, secondary_company_id, konekin_id,
    total_items
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
    total_items = EXCLUDED.total_items,
    updated_at = NOW()
RETURNING *;

-- name: GetDeliveryOrders :many
SELECT 
    a.*, b.form_number AS sales_order_form_number
FROM sales.delivery_orders a
JOIN sales.sales_orders b ON a.sales_order_id = b.id
WHERE a.company_id = $1
    AND a.branch_id = $2
    AND a.transaction_date BETWEEN @start_date::date AND @end_date::date 
    AND a.is_deleted = FALSE;

-- name: GetSalesOrderDeliveryItems :many
SELECT 
    b.*
FROM sales.sales_orders a
JOIN sales.sales_order_items b ON b.sales_order_id = a.id
AND b.is_deleted = FALSE AND b.amount > b.amount_sent
WHERE a.branch_id = $1 AND a.secondary_company_id = $2
AND a.purchase_order_branch_id = $3
AND a.is_deleted = FALSE AND a.status = 'accepted';

-- name: UpsertDeliveryOrderItem :one
INSERT INTO sales.delivery_order_items(
        id, purchase_order_item_id, sales_order_item_id, delivery_order_id,
        primary_item_variant_id, warehouse_rack_id, batch, expired_date, item_barcode_id,
        secondary_item_variant_id, primary_item_unit_id, secondary_item_unit_id,
        primary_item_unit_value, secondary_item_unit_value, amount
    )
VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12, $13, $14, $15) ON CONFLICT (id) DO
UPDATE
SET purchase_order_item_id = EXCLUDED.purchase_order_item_id,
    sales_order_item_id = EXCLUDED.sales_order_item_id,
    delivery_order_id = EXCLUDED.delivery_order_id,
    primary_item_variant_id = EXCLUDED.primary_item_variant_id,
    warehouse_rack_id = EXCLUDED.warehouse_rack_id,
    batch = EXCLUDED.batch,
    expired_date = EXCLUDED.expired_date,
    item_barcode_id = EXCLUDED.item_barcode_id,
    secondary_item_variant_id = EXCLUDED.secondary_item_variant_id,
    primary_item_unit_id = EXCLUDED.primary_item_unit_id,
    secondary_item_unit_id = EXCLUDED.secondary_item_unit_id,
    primary_item_unit_value = EXCLUDED.primary_item_unit_value,
    secondary_item_unit_value = EXCLUDED.secondary_item_unit_value,
    amount = EXCLUDED.amount,
    updated_at = NOW()
RETURNING *;

-- name: DeleteDeliveryOrderItems :exec
DELETE FROM sales.delivery_order_items
WHERE delivery_order_id = $1;

-- name: UpdateDeliveryOrderTotalItems :exec
UPDATE sales.delivery_orders
SET total_items = $2
WHERE id = $1;

-- name: UpdateSalesOrderItemAmountSent :one
UPDATE sales.sales_order_items
SET amount_sent = amount_sent+$2
WHERE id = $1
RETURNING *;

-- name: InsertSalesOrderBranch :exec
INSERT INTO sales.sales_order_branches(sales_order_id, company_branch_id)
VALUES ($1, $2);

-- name: UpsertSalesInvoice :one
INSERT INTO sales.sales_invoices(
    id, sales_order_id, company_id, branch_id,
    form_number, transaction_date,
    contact_book_id, secondary_company_id, konekin_id,
    currency_code
)
VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10) ON CONFLICT (id) DO
UPDATE
SET 
    sales_order_id = EXCLUDED.sales_order_id,
    company_id = EXCLUDED.company_id,
    branch_id = EXCLUDED.branch_id,
    form_number = EXCLUDED.form_number,
    transaction_date = EXCLUDED.transaction_date,
    contact_book_id = EXCLUDED.contact_book_id,
    secondary_company_id = EXCLUDED.secondary_company_id,
    konekin_id = EXCLUDED.konekin_id,
    currency_code = EXCLUDED.currency_code,
    status = EXCLUDED.status
RETURNING *;

-- name: UpdateSalesInvoiceAddItem :exec
UPDATE sales.sales_invoices
SET total_items=sub.total_items,
    total=sub.total,
    updated_at = NOW()
FROM (SELECT sales_invoice_id, COUNT(id) AS total_items, SUM(amount*price) AS total
      FROM sales.sales_invoice_items
      WHERE sales_invoice_id = @sales_invoice_id
      GROUP BY sales_invoice_id) AS sub
WHERE sales.sales_invoices.id = sub.sales_invoice_id;

-- name: DeleteSalesInvoiceItems :exec
DELETE FROM sales.sales_invoice_items
WHERE sales_invoice_id = $1;

-- name: InsertSalesInvoiceItem :exec
INSERT INTO sales.sales_invoice_items(
    id, purchase_order_item_id, sales_order_item_id,
    sales_invoice_id, primary_item_variant_id, secondary_item_variant_id,
    primary_item_unit_id, secondary_item_unit_id,
    primary_item_unit_value, secondary_item_unit_value, amount, price
)
VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12);