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
	GetBrandsPath   = "/inventory/brands"

	AddGroupPath    = "/inventory/group/add"
	UpdateGroupPath = "/inventory/group/update"
	GetGroupsPath   = "/inventory/groups"

	AddUnitPath    = "/inventory/unit/add"
	UpdateUnitPath = "/inventory/unit/update"
	GetUnitsPath   = "/inventory/units"

	AddItemPath    = "/inventory/item/add"
	UpdateItemPath = "/inventory/item/update"
	GetItemsPath   = "/inventory/items"

	UpsertItemInfoPath = "/inventory/item/info/upsert"

	UpsertItemVariantPath = "/inventory/item/variant/upsert"
	GetItemVariantsPath   = "/inventory/item/variants"

	UpsertItemUnitPath = "/inventory/item/unit/upsert"
	GetItemUnitsPath   = "/inventory/item/units"

	AddInternalStockTransferPath  = "/inventory/internal-stock-transfer/add"
	GetInternalStockTransfersPath = "/inventory/internal-stock-transfers"

	UpsertItemReorderPath = "/inventory/item-reorder/upsert"
	GetItemReordersPath   = "/inventory/item-reorders"

	UpsertUnitCategoryPath = "/inventory/unit-category/upsert"
	GetUnitCategoriesPath  = "/inventory/unit-categories"
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
	mux.Method("POST", GetBrandsPath, httpHandler.New(inventoryService.GetBrands))

	mux.Method("POST", AddGroupPath, httpHandler.New(inventoryService.AddGroup))
	mux.Method("POST", UpdateGroupPath, httpHandler.New(inventoryService.UpdateGroup))
	mux.Method("POST", GetGroupsPath, httpHandler.New(inventoryService.GetGroups))

	mux.Method("POST", AddUnitPath, httpHandler.New(inventoryService.AddUnit))
	mux.Method("POST", UpdateUnitPath, httpHandler.New(inventoryService.UpdateUnit))
	mux.Method("POST", GetUnitsPath, httpHandler.New(inventoryService.GetUnits))

	mux.Method("POST", AddItemPath, httpHandler.New(inventoryService.AddItem))
	mux.Method("POST", UpdateItemPath, httpHandler.New(inventoryService.UpdateItem))
	mux.Method("POST", GetItemsPath, httpHandler.New(inventoryService.GetItems))

	mux.Method("POST", UpsertItemInfoPath, httpHandler.New(inventoryService.UpsertItemInfo))

	mux.Method("POST", UpsertItemVariantPath, httpHandler.New(inventoryService.UpsertItemVariant))
	mux.Method("POST", GetItemVariantsPath, httpHandler.New(inventoryService.GetItemVariants))

	mux.Method("POST", UpsertItemUnitPath, httpHandler.New(inventoryService.UpsertItemUnit))
	mux.Method("POST", GetItemUnitsPath, httpHandler.New(inventoryService.GetItemUnits))

	mux.Method("POST", AddInternalStockTransferPath, httpHandler.New(inventoryService.AddInternalStockTransfer))
	mux.Method("POST", GetInternalStockTransfersPath, httpHandler.New(inventoryService.GetInternalStockTransfers))

	mux.Method("POST", UpsertItemReorderPath, httpHandler.New(inventoryService.UpsertItemReorder))
	mux.Method("POST", GetItemReordersPath, httpHandler.New(inventoryService.GetItemReorders))

	mux.Method("POST", UpsertUnitCategoryPath, httpHandler.New(inventoryService.UpsertUnitCategory))
	mux.Method("POST", GetUnitCategoriesPath, httpHandler.New(inventoryService.GetUnitCategories))

	return mux
}
