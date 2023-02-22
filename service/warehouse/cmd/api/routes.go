package main

import (
	"net/http"

	"github.com/calvinkmts/expert-pancake/engine/httpHandler"
	"github.com/expert-pancake/service/warehouse/model"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/cors"
)

const (
	UpsertWarehousePath = "/warehouse/upsert"
	GetWarehousesPath   = "/warehouses"
	DeleteWarehousePath = "/warehouse/delete"

	UpsertWarehouseRackPath = "/warehouse/rack/upsert"
	DeleteWarehouseRackPath = "/warehouse/rack/delete"
	GetWarehouseRacksPath   = "/warehouse/racks"
)

func (c *component) Routes(warehouseService model.WarehouseService) http.Handler {
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

	mux.Method("POST", UpsertWarehousePath, httpHandler.New(warehouseService.UpsertWarehouse))
	mux.Method("POST", GetWarehousesPath, httpHandler.New(warehouseService.GetWarehouses))
	mux.Method("POST", DeleteWarehousePath, httpHandler.New(warehouseService.DeleteWarehouse))

	mux.Method("POST", UpsertWarehouseRackPath, httpHandler.New(warehouseService.UpsertWarehouseRack))
	mux.Method("POST", DeleteWarehouseRackPath, httpHandler.New(warehouseService.DeleteWarehouseRack))
	mux.Method("POST", GetWarehouseRacksPath, httpHandler.New(warehouseService.GetWarehouseRacks))

	return mux
}
