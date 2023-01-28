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
	UpsertRackPath = "/warehouse/rack/upsert"

	GetRacksPath = "/warehouse/racks"
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

	mux.Method("POST", UpsertRackPath, httpHandler.New(warehouseService.UpsertRack))
	mux.Method("POST", GetRacksPath, httpHandler.New(warehouseService.GetRacks))

	return mux
}
