package model

import (
	"net/http"
)

type BusinessRelationService interface {
	AddContactGroup(w http.ResponseWriter, r *http.Request) error
	UpdateContactGroup(w http.ResponseWriter, r *http.Request) error
}

type ContactGroup struct {
	GroupId   string `json:"group_id" validate:"required"`
	CompanyId string `json:"company_id" validate:"required"`
	Name      string `json:"name" validate:"required"`
}

type AddContactGroupRequest struct {
	CompanyId string `json:"company_id" validate:"required"`
	Name      string `json:"name" validate:"required"`
}

type AddContactGroupResponse struct {
	ContactGroup
}

type UpdateContactGroupRequest struct {
	GroupId string `json:"group_id" validate:"required"`
	Name    string `json:"name" validate:"required"`
}

type UpdateContactGroupResponse struct {
	ContactGroup
}
