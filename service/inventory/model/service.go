package model

import (
	"net/http"
)

type InventoryService interface {
	AddBrand(w http.ResponseWriter, r *http.Request) error
	UpdateBrand(w http.ResponseWriter, r *http.Request) error
	GetBrands(w http.ResponseWriter, r *http.Request) error

	AddGroup(w http.ResponseWriter, r *http.Request) error
	UpdateGroup(w http.ResponseWriter, r *http.Request) error
	GetGroups(w http.ResponseWriter, r *http.Request) error

	AddUnit(w http.ResponseWriter, r *http.Request) error
	UpdateUnit(w http.ResponseWriter, r *http.Request) error
	GetUnits(w http.ResponseWriter, r *http.Request) error
}

type Brand struct {
	CompanyId string `json:"company_id" validate:"required"`
	BrandId   string `json:"brand_id" validate:"required"`
	Name      string `json:"name" validate:"required"`
}

type AddBrandRequest struct {
	CompanyId string `json:"company_id" validate:"required"`
	Name      string `json:"name" validate:"required"`
}

type AddBrandResponse struct {
	Brand
}

type UpdateBrandRequest struct {
	Id   string `json:"id" validate:"required"`
	Name string `json:"name" validate:"required"`
}

type UpdateBrandResponse struct {
	Brand
}

type GetBrandsRequest struct {
	CompanyId string `json:"company_id" validate:"required"`
	Keyword   string `json:"keyword"`
}

type Group struct {
	CompanyId string `json:"company_id" validate:"required"`
	GroupId   string `json:"group_id" validate:"required"`
	Name      string `json:"name" validate:"required"`
}

type AddGroupRequest struct {
	CompanyId string `json:"company_id" validate:"required"`
	Name      string `json:"name" validate:"required"`
}

type AddGroupResponse struct {
	Group
}

type UpdateGroupRequest struct {
	Id   string `json:"id" validate:"required"`
	Name string `json:"name" validate:"required"`
}

type UpdateGroupResponse struct {
	Group
}

type GetGroupsRequest struct {
	CompanyId string `json:"company_id" validate:"required"`
	Keyword   string `json:"keyword"`
}

type Unit struct {
	CompanyId string `json:"company_id" validate:"required"`
	UnitId    string `json:"unit_id" validate:"required"`
	Name      string `json:"name" validate:"required"`
}

type AddUnitRequest struct {
	CompanyId string `json:"company_id" validate:"required"`
	Name      string `json:"name" validate:"required"`
}

type AddUnitResponse struct {
	Unit
}

type UpdateUnitRequest struct {
	Id   string `json:"id" validate:"required"`
	Name string `json:"name" validate:"required"`
}

type UpdateUnitResponse struct {
	Unit
}

type GetUnitsRequest struct {
	CompanyId string `json:"company_id" validate:"required"`
	Keyword   string `json:"keyword"`
}
