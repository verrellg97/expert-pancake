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
	AddItemBrandPath = "/inventory/item/brand/add"
	UpdateItemBrandPath = "/inventory/item/brand/update"
	GetItemBrandsPath = "/inventory/item/brands"

	AddItemGroupPath = "/inventory/item/group/add"
	UpdateItemGroupPath = "/inventory/item/group/update"
	GetItemGroupsPath = "/inventory/item/groups"

	AddItemUnitPath = "/inventory/item/unit/add"
	UpdateItemUnitPath = "/inventory/item/unit/update"
	GetItemUnitsPath = "/inventory/item/units"
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

	mux.Method("POST", AddItemBrandPath, httpHandler.New(inventoryService.AddItemBrand))
	mux.Method("POST", UpdateItemBrandPath, httpHandler.New(inventoryService.UpdateItemBrand))
	mux.Method("POST", GetItemBrandsPath, httpHandler.New(inventoryService.GetItemBrands))

	mux.Method("POST", AddItemGroupPath, httpHandler.New(inventoryService.AddItemGroup))
	mux.Method("POST", UpdateItemGroupPath, httpHandler.New(inventoryService.UpdateItemGroup))
	mux.Method("POST", GetItemGroupsPath, httpHandler.New(inventoryService.GetItemGroups))

	mux.Method("POST", AddItemUnitPath, httpHandler.New(inventoryService.AddItemUnit))
	mux.Method("POST", UpdateItemUnitPath, httpHandler.New(inventoryService.UpdateItemUnit))
	mux.Method("POST", GetItemUnitsPath, httpHandler.New(inventoryService.GetItemUnits))

	return mux
}
