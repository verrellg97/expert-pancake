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

	return mux
}
