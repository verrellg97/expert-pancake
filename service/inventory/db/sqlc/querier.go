// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.16.0

package db

import (
	"context"
)

type Querier interface {
	GetBrandById(ctx context.Context, id string) (GetBrandByIdRow, error)
	GetBrands(ctx context.Context, arg GetBrandsParams) ([]GetBrandsRow, error)
	GetGroupById(ctx context.Context, id string) (GetGroupByIdRow, error)
	GetGroups(ctx context.Context, arg GetGroupsParams) ([]GetGroupsRow, error)
	GetItemUnits(ctx context.Context, arg GetItemUnitsParams) ([]GetItemUnitsRow, error)
	GetItemVariant(ctx context.Context, id string) (GetItemVariantRow, error)
	GetItemVariants(ctx context.Context, arg GetItemVariantsParams) ([]GetItemVariantsRow, error)
	GetItems(ctx context.Context, arg GetItemsParams) ([]GetItemsRow, error)
	GetUnit(ctx context.Context, id string) (GetUnitRow, error)
	GetUnits(ctx context.Context, arg GetUnitsParams) ([]GetUnitsRow, error)
	InsertBrand(ctx context.Context, arg InsertBrandParams) (InventoryBrand, error)
	InsertGroup(ctx context.Context, arg InsertGroupParams) (InventoryGroup, error)
	InsertItem(ctx context.Context, arg InsertItemParams) (InventoryItem, error)
	InsertItemVariant(ctx context.Context, arg InsertItemVariantParams) (InventoryItemVariant, error)
	InsertUnit(ctx context.Context, arg InsertUnitParams) (InventoryUnit, error)
	UpdateBrand(ctx context.Context, arg UpdateBrandParams) (InventoryBrand, error)
	UpdateGroup(ctx context.Context, arg UpdateGroupParams) (InventoryGroup, error)
	UpdateItem(ctx context.Context, arg UpdateItemParams) (InventoryItem, error)
	UpdateItemVariantDefault(ctx context.Context, arg UpdateItemVariantDefaultParams) (InventoryItemVariant, error)
	UpdateUnit(ctx context.Context, arg UpdateUnitParams) (InventoryUnit, error)
	UpsertItemUnit(ctx context.Context, arg UpsertItemUnitParams) (InventoryItemUnit, error)
	UpsertItemVariant(ctx context.Context, arg UpsertItemVariantParams) error
}

var _ Querier = (*Queries)(nil)
