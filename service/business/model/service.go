package model

import (
	"net/http"
)

type BusinessService interface {
	RegisterCompany(w http.ResponseWriter, r *http.Request) error
	GetCompanyTypes(w http.ResponseWriter, r *http.Request) error
	UpdateCompany(w http.ResponseWriter, r *http.Request) error
	DeleteCompany(w http.ResponseWriter, r *http.Request) error
	GetUserCompanies(w http.ResponseWriter, r *http.Request) error

	RegisterCompanyBranch(w http.ResponseWriter, r *http.Request) error
	UpdateCompanyBranch(w http.ResponseWriter, r *http.Request) error
	DeleteCompanyBranch(w http.ResponseWriter, r *http.Request) error
	GetUserCompanyBranches(w http.ResponseWriter, r *http.Request) error
	GetCompanyBranches(w http.ResponseWriter, r *http.Request) error
}

type Company struct {
	AccountId         string          `json:"account_id" validate:"required"`
	CompanyId         string          `json:"company_id" validate:"required"`
	Name              string          `json:"name" validate:"required"`
	InitialName       string          `json:"initial_name" validate:"required"`
	Type              string          `json:"type" validate:"required"`
	ResponsiblePerson string          `json:"responsible_person" validate:"required"`
	ImageUrl          string          `json:"image_url"`
	Branches          []CompanyBranch `json:"branches" validate:"required"`
}

type CompanyBranch struct {
	AccountId   string `json:"account_id" validate:"required"`
	CompanyId   string `json:"company_id" validate:"required"`
	BranchId    string `json:"branch_id" validate:"required"`
	Name        string `json:"name" validate:"required"`
	Address     string `json:"address" validate:"required"`
	PhoneNumber string `json:"phone_number" validate:"required"`
	IsCentral   bool   `json:"is_central" validate:"required"`
}

type RegisterCompanyRequest struct {
	AccountId         string `json:"account_id" validate:"required"`
	Name              string `json:"name" validate:"required"`
	InitialName       string `json:"initial_name" validate:"required"`
	Type              string `json:"type" validate:"required"`
	ResponsiblePerson string `json:"responsible_person" validate:"required"`
	ImageUrl          string `json:"image_url"`
}

type RegisterCompanyResponse struct {
	Company
}

type GetCompanyTypesResponse struct {
	Types []string `json:"types"`
}

type UpdateCompanyRequest struct {
	AccountId         string `json:"account_id" validate:"required"`
	CompanyId         string `json:"company_id" validate:"required"`
	Name              string `json:"name" validate:"required"`
	InitialName       string `json:"initial_name" validate:"required"`
	Type              string `json:"type" validate:"required"`
	ResponsiblePerson string `json:"responsible_person" validate:"required"`
	ImageUrl          string `json:"image_url"`
}

type DeleteCompanyRequest struct {
	CompanyId string `json:"company_id" validate:"required"`
}

type UserCompaniesRequest struct {
	AccountId string `json:"account_id" validate:"required"`
	Keyword   string `json:"keyword"`
}

type RegisterCompanyBranchRequest struct {
	AccountId   string `json:"account_id" validate:"required"`
	CompanyId   string `json:"company_id" validate:"required"`
	Name        string `json:"name" validate:"required"`
	Address     string `json:"address" validate:"required"`
	PhoneNumber string `json:"phone_number" validate:"required"`
}

type RegisterCompanyBranchResponse struct {
	CompanyBranch
}

type UserCompanyBranchesRequest struct {
	AccountId string `json:"account_id" validate:"required"`
	CompanyId string `json:"company_id" validate:"required"`
	Keyword   string `json:"keyword"`
}

type UpdateCompanyBranchRequest struct {
	AccountId   string `json:"account_id" validate:"required"`
	CompanyId   string `json:"company_id" validate:"required"`
	BranchId    string `json:"branch_id" validate:"required"`
	Name        string `json:"name" validate:"required"`
	Address     string `json:"address" validate:"required"`
	PhoneNumber string `json:"phone_number" validate:"required"`
}

type DeleteCompanyBranchRequest struct {
	BranchId string `json:"branch_id" validate:"required"`
}

type DeleteDataResponse struct {
	Message string `json:"message"`
}

type CompanyBranchesRequest struct {
	CompanyId string `json:"company_id" validate:"required"`
}
