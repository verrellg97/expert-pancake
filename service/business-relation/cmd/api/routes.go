package main

import (
	"net/http"

	"github.com/calvinkmts/expert-pancake/engine/httpHandler"
	"github.com/expert-pancake/service/business-relation/model"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/cors"
)

const (
	AddContactGroupPath    = "/business-relation/contact-group/add"
	UpdateContactGroupPath = "/business-relation/contact-group/update"
	GetContactGroupsPath   = "/business-relation/contact-groups"
)

func (c *component) Routes(businessRelationService model.BusinessRelationService) http.Handler {
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

	mux.Method("POST", AddContactGroupPath, httpHandler.New(businessRelationService.AddContactGroup))
	mux.Method("POST", UpdateContactGroupPath, httpHandler.New(businessRelationService.UpdateContactGroup))
	mux.Method("POST", GetContactGroupsPath, httpHandler.New(businessRelationService.GetContactGroups))

	return mux
}
