package main

import (
	"net/http"

	"github.com/calvinkmts/expert-pancake/engine/httpHandler"
	"github.com/expert-pancake/service/purchasing/model"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/cors"
)

const (
	UpsertPurchaseOrderPath     = "/purchasing/order/upsert"
	UpsertPurchaseOrderItemPath = "/purchasing/order/item/upsert"
	GetPurchaseOrdersPath       = "/purchasing/order/get"
	GetPurcaseOrderItemsPath    = "/purchasing/order/item/get"
)

func (c *component) Routes(purchasingService model.PurchasingService) http.Handler {
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

	mux.Method("POST", UpsertPurchaseOrderPath, httpHandler.New(purchasingService.UpsertPurchaseOrder))
	mux.Method("POST", UpsertPurchaseOrderItemPath, httpHandler.New(purchasingService.UpsertPurchaseOrderItem))
	mux.Method("POST", GetPurchaseOrdersPath, httpHandler.New(purchasingService.GetPurchaseOrders))
	mux.Method("POST", GetPurcaseOrderItemsPath, httpHandler.New(purchasingService.GetPurchaseOrderItems))

	return mux
}
