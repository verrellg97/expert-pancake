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

	AddMemberRequest(w http.ResponseWriter, r *http.Request) error
	GetReceiveMemberRequests(w http.ResponseWriter, r *http.Request) error
	UpdateMemberRequest(w http.ResponseWriter, r *http.Request) error

	GetPublicCompanies(w http.ResponseWriter, r *http.Request) error
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

type AddMemberRequestRequest struct {
	UserId    string `json:"user_id" validate:"required"`
	CompanyId string `json:"company_id" validate:"required"`
}

type AddMemberRequestResponse struct {
	Message string `json:"message"`
}

type MemberRequest struct {
	Id          string `json:"id" validate:"required"`
	UserId      string `json:"user_id" validate:"required"`
	ImageUrl    string `json:"image_url" validate:"required"`
	FullName    string `json:"full_name" validate:"required"`
	Nickname    string `json:"nickname" validate:"required"`
	Email       string `json:"email" validate:"required"`
	PhoneNumber string `json:"phone_number" validate:"required"`
	Status      string `json:"status" validate:"required"`
}

type GetReceiveMemberRequestsRequest struct {
	CompanyId string `json:"company_id" validate:"required"`
}

type GetReceiveMemberRequestsResponse struct {
	MemberRequests []MemberRequest `json:"member_requests" validate:"required"`
}

type UpdateMemberRequestRequest struct {
	Id     string `json:"id" validate:"required"`
	Status string `json:"status" validate:"required"`
}

type UpdateMemberRequestResponse struct {
	Message string `json:"message"`
}

type GetPublicCompaniesRequest struct {
	AccountId string `json:"account_id" validate:"required"`
	Keyword   string `json:"keyword"`
}

type GetPublicCompaniesResponse struct {
	Companies []Company `json:"companies" validate:"required"`
}
