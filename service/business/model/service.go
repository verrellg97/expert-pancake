package model

import (
	"net/http"
)

type BusinessService interface {
	HelloWorld(w http.ResponseWriter, r *http.Request) error
	HelloError(w http.ResponseWriter, r *http.Request) error

	RegisterCompany(w http.ResponseWriter, r *http.Request) error
	GetCompanyTypes(w http.ResponseWriter, r *http.Request) error
}

type Company struct {
	AccountId         string `json:"account_id" validate:"required"`
	CompanyId         string `json:"company_id" validate:"required"`
	Name              string `json:"name" validate:"required"`
	InitialName       string `json:"initial_name" validate:"required"`
	Type              string `json:"type" validate:"required"`
	ResponsiblePerson string `json:"responsible_person" validate:"required"`
}

type CompanyBranch struct {
	AccountId   string `json:"account_id" validate:"required"`
	CompanyId   string `json:"company_id" validate:"required"`
	BranchId    string `json:"branch_id" validate:"required"`
	Name        string `json:"name" validate:"required"`
	Address     string `json:"address" validate:"required"`
	PhoneNumber string `json:"phone_number" validate:"required"`
}

type RegisterCompanyRequest struct {
	AccountId         string `json:"account_id" validate:"required"`
	Name              string `json:"name" validate:"required"`
	InitialName       string `json:"initial_name" validate:"required"`
	Type              string `json:"type" validate:"required"`
	ResponsiblePerson string `json:"responsible_person" validate:"required"`
}

type RegisterCompanyResponse struct {
	Company
}

type GetCompanyTypesResponse struct {
	Types []string `json:"types"`
}
