package main

import (
	"net/http"

	"github.com/calvinkmts/expert-pancake/engine/httpHandler"
	"github.com/expert-pancake/service/notification/model"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/cors"
)

const (
	InsertNotificationPath = "/notification/insert"
	GetNotificationsPath   = "/notifications"
	ReadNotificationPath = "/notification/read"
	DeleteNotificationPath = "/notification/delete"
)

func (c *component) Routes(notificationService model.NotificationService) http.Handler {
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

	mux.Method("POST", InsertNotificationPath, httpHandler.New(notificationService.InsertNotification))
	mux.Method("POST", GetNotificationsPath, httpHandler.New(notificationService.GetNotifications))
	mux.Method("POST", ReadNotificationPath, httpHandler.New(notificationService.ReadNotification))
	mux.Method("POST", DeleteNotificationPath, httpHandler.New(notificationService.DeleteNotification))

	return mux
}
