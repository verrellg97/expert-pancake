// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.17.2

package db

import (
	"context"
	"database/sql"
)

type Querier interface {
	DeleteBrand(ctx context.Context, id string) error
	DeleteGroup(ctx context.Context, id string) error
	GetBrandById(ctx context.Context, id string) (GetBrandByIdRow, error)
	GetBrands(ctx context.Context, arg GetBrandsParams) ([]GetBrandsRow, error)
	GetGroupById(ctx context.Context, id string) (GetGroupByIdRow, error)
	GetGroups(ctx context.Context, arg GetGroupsParams) ([]GetGroupsRow, error)
	GetInternalStockTransferItems(ctx context.Context, internalStockTransferID string) ([]GetInternalStockTransferItemsRow, error)
	GetInternalStockTransfers(ctx context.Context, arg GetInternalStockTransfersParams) ([]GetInternalStockTransfersRow, error)
	GetItemBarcode(ctx context.Context, arg GetItemBarcodeParams) (string, error)
	GetItemInfo(ctx context.Context, itemID string) (GetItemInfoRow, error)
	GetItemReorder(ctx context.Context, id string) (GetItemReorderRow, error)
	GetItemReorderNotifications(ctx context.Context, arg GetItemReorderNotificationsParams) ([]GetItemReorderNotificationsRow, error)
	GetItemReorders(ctx context.Context, arg GetItemReordersParams) ([]GetItemReordersRow, error)
	GetItemUnits(ctx context.Context, arg GetItemUnitsParams) ([]GetItemUnitsRow, error)
	GetItemVariant(ctx context.Context, id string) (GetItemVariantRow, error)
	GetItemVariantMap(ctx context.Context, id string) (GetItemVariantMapRow, error)
	GetItemVariants(ctx context.Context, arg GetItemVariantsParams) ([]GetItemVariantsRow, error)
	GetItems(ctx context.Context, arg GetItemsParams) ([]GetItemsRow, error)
	GetStockHistory(ctx context.Context, arg GetStockHistoryParams) ([]GetStockHistoryRow, error)
	GetTransferHistory(ctx context.Context, arg GetTransferHistoryParams) ([]GetTransferHistoryRow, error)
	GetUnit(ctx context.Context, id string) (GetUnitRow, error)
	GetUnitCategories(ctx context.Context, arg GetUnitCategoriesParams) ([]GetUnitCategoriesRow, error)
	GetUnits(ctx context.Context, arg GetUnitsParams) ([]GetUnitsRow, error)
	GetUpdateStock(ctx context.Context, id string) (GetUpdateStockRow, error)
	GetUpdateStocks(ctx context.Context, arg GetUpdateStocksParams) ([]GetUpdateStocksRow, error)
	GetVariantWarehouseRackBatchExpiredDates(ctx context.Context, arg GetVariantWarehouseRackBatchExpiredDatesParams) ([]sql.NullTime, error)
	GetVariantWarehouseRackBatches(ctx context.Context, arg GetVariantWarehouseRackBatchesParams) ([]sql.NullString, error)
	GetVariantWarehouseRackStock(ctx context.Context, arg GetVariantWarehouseRackStockParams) (int64, error)
	GetVariantWarehouseRacks(ctx context.Context, arg GetVariantWarehouseRacksParams) ([]string, error)
	GetVariantWarehouseStocks(ctx context.Context, warehouseID string) ([]GetVariantWarehouseStocksRow, error)
	InsertBrand(ctx context.Context, arg InsertBrandParams) (InventoryBrand, error)
	InsertGroup(ctx context.Context, arg InsertGroupParams) (InventoryGroup, error)
	InsertInternalStockTransfer(ctx context.Context, arg InsertInternalStockTransferParams) (InventoryInternalStockTransfer, error)
	InsertInternalStockTransferItem(ctx context.Context, arg InsertInternalStockTransferItemParams) error
	InsertItem(ctx context.Context, arg InsertItemParams) (InventoryItem, error)
	InsertItemBarcode(ctx context.Context, arg InsertItemBarcodeParams) error
	InsertItemVariant(ctx context.Context, arg InsertItemVariantParams) (InventoryItemVariant, error)
	InsertStockMovement(ctx context.Context, arg InsertStockMovementParams) error
	InsertUnit(ctx context.Context, arg InsertUnitParams) (InventoryUnit, error)
	InsertUpdateStock(ctx context.Context, arg InsertUpdateStockParams) error
	UpdateBrand(ctx context.Context, arg UpdateBrandParams) (InventoryBrand, error)
	UpdateGroup(ctx context.Context, arg UpdateGroupParams) (InventoryGroup, error)
	UpdateItem(ctx context.Context, arg UpdateItemParams) (InventoryItem, error)
	UpdateItemUnitIsDefaultToFalse(ctx context.Context, arg UpdateItemUnitIsDefaultToFalseParams) error
	UpdateItemVariantDefault(ctx context.Context, arg UpdateItemVariantDefaultParams) (InventoryItemVariant, error)
	UpdateUnit(ctx context.Context, arg UpdateUnitParams) (InventoryUnit, error)
	UpsertItemInfo(ctx context.Context, arg UpsertItemInfoParams) error
	UpsertItemReorder(ctx context.Context, arg UpsertItemReorderParams) (InventoryItemReorder, error)
	UpsertItemUnit(ctx context.Context, arg UpsertItemUnitParams) (InventoryItemUnit, error)
	UpsertItemVariant(ctx context.Context, arg UpsertItemVariantParams) error
	UpsertItemVariantMap(ctx context.Context, arg UpsertItemVariantMapParams) error
	UpsertUnitCategory(ctx context.Context, arg UpsertUnitCategoryParams) (InventoryUnitCategory, error)
}

var _ Querier = (*Queries)(nil)
