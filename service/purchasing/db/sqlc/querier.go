// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.16.0

package db

import (
	"context"
)

type Querier interface {
	GetPurchaseOrders(ctx context.Context, arg GetPurchaseOrdersParams) ([]PurchasingPurchaseOrder, error)
	UpdatePurchaseOrderAddItem(ctx context.Context, purchaseOrderID string) error
	UpsertPurchaseOrder(ctx context.Context, arg UpsertPurchaseOrderParams) (PurchasingPurchaseOrder, error)
	UpsertPurchaseOrderItem(ctx context.Context, arg UpsertPurchaseOrderItemParams) (PurchasingPurchaseOrderItem, error)
}

var _ Querier = (*Queries)(nil)