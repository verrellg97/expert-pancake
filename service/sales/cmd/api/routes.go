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
	GetPOSCOASettingPath    = "/sales/pos/coa/setting"

	UpdatePOSCustomerSettingPath = "/sales/pos/customer/setting/update"
	GetPOSCustomerSettingPath    = "/sales/pos/customer/setting"

	UpsertPOSPaymentMethodPath = "/sales/pos/payments/upsert"
	DeletePOSPaymentMethodPath = "/sales/pos/payments/delete"
	GetPOSPaymentMethodPath    = "/sales/pos/payments"

	GetCheckPOSPath = "/sales/pos/check"

	UpsertSalesOrderPath       = "/sales/order/upsert"
	UpdateSalesOrderItemsPath  = "/sales/order/items/update"
	GetSalesOrdersPath         = "/sales/orders"
	GetSalesOrderItemsPath     = "/sales/order/items"
	UpdateSalesOrderStatusPath = "/sales/order/status"

	UpsertDeliveryOrderPath        = "/sales/delivery/order/upsert"
	GetDeliveryOrdersPath          = "/sales/delivery/orders"
	GetSalesOrderDeliveryItemsPath = "/sales/order/delivery-items"
	UpdateDeliveryOrderItemsPath   = "/sales/delivery/order/items/update"
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
	mux.Method("POST", GetPOSCOASettingPath, httpHandler.New(salesService.GetPOSCOASetting))

	mux.Method("POST", UpdatePOSCustomerSettingPath, httpHandler.New(salesService.UpdatePOSCustomerSetting))
	mux.Method("POST", GetPOSCustomerSettingPath, httpHandler.New(salesService.GetPOSCustomerSetting))

	mux.Method("POST", UpsertPOSPaymentMethodPath, httpHandler.New(salesService.UpsertPOSPaymentMethod))
	mux.Method("POST", DeletePOSPaymentMethodPath, httpHandler.New(salesService.DeletePOSPaymentMethod))
	mux.Method("POST", GetPOSPaymentMethodPath, httpHandler.New(salesService.GetPOSPaymentMethod))

	mux.Method("POST", GetCheckPOSPath, httpHandler.New(salesService.GetCheckPOS))

	mux.Method("POST", UpsertSalesOrderPath, httpHandler.New(salesService.UpsertSalesOrder))
	mux.Method("POST", UpdateSalesOrderItemsPath, httpHandler.New(salesService.UpdateSalesOrderItems))
	mux.Method("POST", GetSalesOrdersPath, httpHandler.New(salesService.GetSalesOrders))
	mux.Method("POST", GetSalesOrderItemsPath, httpHandler.New(salesService.GetSalesOrderItems))
	mux.Method("POST", UpdateSalesOrderStatusPath, httpHandler.New(salesService.UpdateSalesOrderStatus))

	mux.Method("POST", UpsertDeliveryOrderPath, httpHandler.New(salesService.UpsertDeliveryOrder))
	mux.Method("POST", GetDeliveryOrdersPath, httpHandler.New(salesService.GetDeliveryOrders))
	mux.Method("POST", GetSalesOrderDeliveryItemsPath, httpHandler.New(salesService.GetSalesOrderDeliveryItems))
	mux.Method("POST", UpdateDeliveryOrderItemsPath, httpHandler.New(salesService.UpdateDeliveryOrderItems))

	return mux
}
