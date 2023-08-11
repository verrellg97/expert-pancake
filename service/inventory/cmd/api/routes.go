package main

import (
	"net/http"

	"github.com/calvinkmts/expert-pancake/engine/httpHandler"
	"github.com/expert-pancake/service/inventory/model"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/cors"
)

const (
	AddBrandPath    = "/inventory/brand/add"
	UpdateBrandPath = "/inventory/brand/update"
	DeleteBrandPath = "/inventory/brand/delete"
	GetBrandsPath   = "/inventory/brands"

	AddGroupPath    = "/inventory/group/add"
	UpdateGroupPath = "/inventory/group/update"
	DeleteGroupPath = "/inventory/group/delete"
	GetGroupsPath   = "/inventory/groups"

	AddUnitPath    = "/inventory/unit/add"
	UpdateUnitPath = "/inventory/unit/update"
	GetUnitsPath   = "/inventory/units"

	AddItemPath    = "/inventory/item/add"
	UpdateItemPath = "/inventory/item/update"
	GetItemsPath   = "/inventory/items"

	UpsertItemInfoPath = "/inventory/item/info/upsert"
	GetItemInfoPath    = "/inventory/item/info"

	UpsertItemVariantPath = "/inventory/item/variant/upsert"
	GetItemVariantsPath   = "/inventory/item/variants"
	DeleteItemVariantPath = "/inventory/item/variant/delete"

	UpsertItemUnitPath = "/inventory/item/unit/upsert"
	DeleteItemUnitPath = "/inventory/item/unit/delete"
	GetItemUnitsPath   = "/inventory/item/units"

	AddInternalStockTransferPath  = "/inventory/internal-stock-transfer/add"
	GetInternalStockTransfersPath = "/inventory/internal-stock-transfers"

	AddUpdateStockPath  = "/inventory/update-stock/add"
	GetUpdateStocksPath = "/inventory/update-stocks"

	UpsertItemReorderPath = "/inventory/item-reorder/upsert"
	GetItemReordersPath   = "/inventory/item-reorders"

	UpsertUnitCategoryPath = "/inventory/unit-category/upsert"
	GetUnitCategoriesPath  = "/inventory/unit-categories"

	GetVariantWarehouseRacksPath                 = "/inventory/item/variant/racks"
	GetVariantWarehouseRackBatchesPath           = "/inventory/item/variant/rack/batches"
	GetVariantWarehouseRackBatchExpiredDatesPath = "/inventory/item/variant/rack/batch/expired-dates"
	GetVariantWarehouseRackStockPath             = "/inventory/item/variant/rack/stock"
	GetVariantWarehouseStocksPath                = "/inventory/item/variant/warehouse/stocks"

	GetTransferHistoryPath = "/inventory/transfer-history"
	GetStockHistoryPath    = "/inventory/stock-history"

	GetItemReorderNotificationsPath = "/inventory/item-reorder/notifications"

	GetSupplierCatalogsPath    = "/inventory/supplier/catalogs"
	GetMappingItemsPath        = "/inventory/mapping/items"
	GetMappingItemVariantsPath = "/inventory/mapping/item/variants"
	GetMappingItemUnitsPath    = "/inventory/mapping/item/units"

	UpsertItemVariantMapPath = "/inventory/item/variant/mapping/upsert"
	GetItemVariantMapsPath   = "/inventory/item/variant/mappings"

	UpsertPricelistPath = "/inventory/pricelist/upsert"
	GetPricelistsPath   = "/inventory/pricelists"

	UpsertPricelistItemsPath = "/inventory/pricelist/items/upsert"
	GetPricelistItemsPath    = "/inventory/pricelist/items"

	GetPurchaseItemsPath            = "/inventory/purchase/items"
	GetPurchaseItemVariantsPath     = "/inventory/purchase/item/variants"
	GetPurchaseItemVariantUnitsPath = "/inventory/purchase/item/variant/units"

	GetPOSItemsPath                      = "/inventory/pos/items"
	GetVariantWarehouseRacksByBranchPath = "/inventory/item/variant/branch/racks"

	GetCheckStockHistoryPath = "/inventory/check-stock-history"

	InsertStockMovementPath = "/inventory/stock-movement/insert"
	DeleteStockMovementPath = "/inventory/stock-movement/delete"

	GetUnderMinimumOrderPath = "/inventory/under-minimum-order"
	GetOutgoingStockPath     = "/inventory/outgoing-stock"
	GetIncomingStockPath     = "/inventory/incoming-stock"

	GetItemHistoryPath = "/inventory/item-history"

	AddOpeningStockPath  = "/inventory/opening-stock/add"
	GetOpeningStocksPath = "/inventory/opening-stocks"
	UpsertOpeningStockPath  = "/inventory/opening-stock/upsert"
	DeleteOpeningStockPath  = "/inventory/opening-stock/delete"
)

func (c *component) Routes(inventoryService model.InventoryService) http.Handler {
	mux := chi.NewRouter()

	mux.Use(middleware.Logger)

	mux.Use(cors.Handler(cors.Options{
		AllowedOrigins:   []string{"https://*", "http://*"},
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"Accept", "Authorization", "Content-Type", "X-CSRF-Token"},
		ExposedHeaders:   []string{"Link"},
		AllowCredentials: true,
		MaxAge:           300,
	}))

	mux.Use(middleware.Heartbeat("/ping"))

	mux.Method("POST", AddBrandPath, httpHandler.New(inventoryService.AddBrand))
	mux.Method("POST", UpdateBrandPath, httpHandler.New(inventoryService.UpdateBrand))
	mux.Method("POST", DeleteBrandPath, httpHandler.New(inventoryService.DeleteBrand))
	mux.Method("POST", GetBrandsPath, httpHandler.New(inventoryService.GetBrands))

	mux.Method("POST", AddGroupPath, httpHandler.New(inventoryService.AddGroup))
	mux.Method("POST", UpdateGroupPath, httpHandler.New(inventoryService.UpdateGroup))
	mux.Method("POST", DeleteGroupPath, httpHandler.New(inventoryService.DeleteGroup))
	mux.Method("POST", GetGroupsPath, httpHandler.New(inventoryService.GetGroups))

	mux.Method("POST", AddUnitPath, httpHandler.New(inventoryService.AddUnit))
	mux.Method("POST", UpdateUnitPath, httpHandler.New(inventoryService.UpdateUnit))
	mux.Method("POST", GetUnitsPath, httpHandler.New(inventoryService.GetUnits))

	mux.Method("POST", AddItemPath, httpHandler.New(inventoryService.AddItem))
	mux.Method("POST", UpdateItemPath, httpHandler.New(inventoryService.UpdateItem))
	mux.Method("POST", GetItemsPath, httpHandler.New(inventoryService.GetItems))

	mux.Method("POST", UpsertItemInfoPath, httpHandler.New(inventoryService.UpsertItemInfo))
	mux.Method("POST", GetItemInfoPath, httpHandler.New(inventoryService.GetItemInfo))

	mux.Method("POST", UpsertItemVariantPath, httpHandler.New(inventoryService.UpsertItemVariant))
	mux.Method("POST", GetItemVariantsPath, httpHandler.New(inventoryService.GetItemVariants))
	mux.Method("POST", DeleteItemVariantPath, httpHandler.New(inventoryService.DeleteItemVariant))

	mux.Method("POST", UpsertItemUnitPath, httpHandler.New(inventoryService.UpsertItemUnit))
	mux.Method("POST", DeleteItemUnitPath, httpHandler.New(inventoryService.DeleteItemUnit))
	mux.Method("POST", GetItemUnitsPath, httpHandler.New(inventoryService.GetItemUnits))

	mux.Method("POST", AddInternalStockTransferPath, httpHandler.New(inventoryService.AddInternalStockTransfer))
	mux.Method("POST", GetInternalStockTransfersPath, httpHandler.New(inventoryService.GetInternalStockTransfers))

	mux.Method("POST", AddUpdateStockPath, httpHandler.New(inventoryService.AddUpdateStock))
	mux.Method("POST", GetUpdateStocksPath, httpHandler.New(inventoryService.GetUpdateStocks))

	mux.Method("POST", UpsertItemReorderPath, httpHandler.New(inventoryService.UpsertItemReorder))
	mux.Method("POST", GetItemReordersPath, httpHandler.New(inventoryService.GetItemReorders))

	mux.Method("POST", UpsertUnitCategoryPath, httpHandler.New(inventoryService.UpsertUnitCategory))
	mux.Method("POST", GetUnitCategoriesPath, httpHandler.New(inventoryService.GetUnitCategories))

	mux.Method("POST", GetVariantWarehouseRacksPath, httpHandler.New(inventoryService.GetVariantWarehouseRacks))
	mux.Method("POST", GetVariantWarehouseRackBatchesPath, httpHandler.New(inventoryService.GetVariantWarehouseRackBatches))
	mux.Method("POST", GetVariantWarehouseRackBatchExpiredDatesPath, httpHandler.New(inventoryService.GetVariantWarehouseRackBatchExpiredDates))
	mux.Method("POST", GetVariantWarehouseRackStockPath, httpHandler.New(inventoryService.GetVariantWarehouseRackStock))
	mux.Method("POST", GetVariantWarehouseStocksPath, httpHandler.New(inventoryService.GetVariantWarehouseStocks))

	mux.Method("POST", GetTransferHistoryPath, httpHandler.New(inventoryService.GetTransferHistory))
	mux.Method("POST", GetStockHistoryPath, httpHandler.New(inventoryService.GetStockHistory))

	mux.Method("POST", GetItemReorderNotificationsPath, httpHandler.New(inventoryService.GetItemReorderNotifications))

	mux.Method("POST", GetSupplierCatalogsPath, httpHandler.New(inventoryService.GetSupplierCatalogs))
	mux.Method("POST", GetMappingItemsPath, httpHandler.New(inventoryService.GetMappingItems))
	mux.Method("POST", GetMappingItemVariantsPath, httpHandler.New(inventoryService.GetMappingItemVariants))
	mux.Method("POST", GetMappingItemUnitsPath, httpHandler.New(inventoryService.GetMappingItemUnits))

	mux.Method("POST", UpsertItemVariantMapPath, httpHandler.New(inventoryService.UpsertItemVariantMap))
	mux.Method("POST", GetItemVariantMapsPath, httpHandler.New(inventoryService.GetItemVariantMaps))

	mux.Method("POST", UpsertPricelistPath, httpHandler.New(inventoryService.UpsertPricelist))
	mux.Method("POST", GetPricelistsPath, httpHandler.New(inventoryService.GetPricelists))

	mux.Method("POST", UpsertPricelistItemsPath, httpHandler.New(inventoryService.UpsertPricelistItems))
	mux.Method("POST", GetPricelistItemsPath, httpHandler.New(inventoryService.GetPricelistItems))

	mux.Method("POST", GetPurchaseItemsPath, httpHandler.New(inventoryService.GetPurchaseItems))
	mux.Method("POST", GetPurchaseItemVariantsPath, httpHandler.New(inventoryService.GetPurchaseItemVariants))
	mux.Method("POST", GetPurchaseItemVariantUnitsPath, httpHandler.New(inventoryService.GetPurchaseItemVariantUnits))

	mux.Method("POST", GetPOSItemsPath, httpHandler.New(inventoryService.GetPOSItems))
	mux.Method("POST", GetVariantWarehouseRacksByBranchPath, httpHandler.New(inventoryService.GetVariantWarehouseRacksByBranch))

	mux.Method("POST", GetCheckStockHistoryPath, httpHandler.New(inventoryService.GetCheckStockHistory))

	mux.Method("POST", InsertStockMovementPath, httpHandler.New(inventoryService.InsertStockMovement))
	mux.Method("POST", DeleteStockMovementPath, httpHandler.New(inventoryService.DeleteStockMovement))

	mux.Method("POST", GetUnderMinimumOrderPath, httpHandler.New(inventoryService.GetUnderMinimumOrder))
	mux.Method("POST", GetOutgoingStockPath, httpHandler.New(inventoryService.GetOutgoingStock))
	mux.Method("POST", GetIncomingStockPath, httpHandler.New(inventoryService.GetIncomingStock))

	mux.Method("POST", GetItemHistoryPath, httpHandler.New(inventoryService.GetItemHistory))

	mux.Method("POST", AddOpeningStockPath, httpHandler.New(inventoryService.AddOpeningStock))
	mux.Method("POST", GetOpeningStocksPath, httpHandler.New(inventoryService.GetOpeningStocks))
	mux.Method("POST", UpsertOpeningStockPath, httpHandler.New(inventoryService.UpsertOpeningStock))
	mux.Method("POST", DeleteOpeningStockPath, httpHandler.New(inventoryService.DeleteOpeningStock))

	return mux
}
