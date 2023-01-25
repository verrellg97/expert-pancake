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

	return mux
}
