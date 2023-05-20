package main

import (
	"net/http"

	"github.com/calvinkmts/expert-pancake/engine/httpHandler"
	"github.com/expert-pancake/service/sales/model"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/cors"
)

const (
	UpsertPOSPath   = "/sales/pos/upsert"
	DeletePOSPath   = "/sales/pos/delete"
	GetPOSPath      = "/sales/pos"
	GetPOSItemsPath = "/sales/pos/items"

	GetPOSUserSettingPath    = "/sales/pos/user/setting"
	UpdatePOSUserSettingPath = "/sales/pos/user/setting/update"

	UpdatePOSCOASettingPath = "/sales/pos/coa/setting/update"
)

func (c *component) Routes(salesService model.SalesService) http.Handler {
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
	mux.Method("POST", UpsertPOSPath, httpHandler.New(salesService.UpsertPOS))
	mux.Method("POST", DeletePOSPath, httpHandler.New(salesService.DeletePOS))
	mux.Method("POST", GetPOSPath, httpHandler.New(salesService.GetPOS))
	mux.Method("POST", GetPOSItemsPath, httpHandler.New(salesService.GetPOSItems))

	mux.Method("POST", GetPOSUserSettingPath, httpHandler.New(salesService.GetPOSUserSetting))
	mux.Method("POST", UpdatePOSUserSettingPath, httpHandler.New(salesService.UpdatePOSUserSetting))

	mux.Method("POST", UpdatePOSCOASettingPath, httpHandler.New(salesService.UpdatePOSCOASetting))

	return mux
}
