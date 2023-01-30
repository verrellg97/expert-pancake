// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.16.0

package db

import (
	"context"
)

type Querier interface {
	GetRacks(ctx context.Context, arg GetRacksParams) ([]GetRacksRow, error)
	GetWarehouses(ctx context.Context, arg GetWarehousesParams) ([]WarehouseWarehouse, error)
	UpsertRack(ctx context.Context, arg UpsertRackParams) (WarehouseRack, error)
}

var _ Querier = (*Queries)(nil)
