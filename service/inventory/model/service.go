package model

import (
	"net/http"
)

type InventoryService interface {
	AddItemBrand(w http.ResponseWriter, r *http.Request) error
	UpdateItemBrand(w http.ResponseWriter, r *http.Request) error
	GetItemBrands(w http.ResponseWriter, r *http.Request) error

	AddItemGroup(w http.ResponseWriter, r *http.Request) error
	UpdateItemGroup(w http.ResponseWriter, r *http.Request) error
	GetItemGroups(w http.ResponseWriter, r *http.Request) error

	AddItemUnit(w http.ResponseWriter, r *http.Request) error
	UpdateItemUnit(w http.ResponseWriter, r *http.Request) error
	GetItemUnits(w http.ResponseWriter, r *http.Request) error
}

type Brand struct {
	CompanyId   string `json:"company_id" validate:"required"`
	ItemBrandId string `json:"item_brand_id" validate:"required"`
	Name        string `json:"name" validate:"required"`
}

type AddItemBrandRequest struct {
	CompanyId string `json:"company_id" validate:"required"`
	Name      string `json:"name" validate:"required"`
}

type AddItemBrandResponse struct {
	Brand
}

type UpdateItemBrandRequest struct {
	Id   string `json:"id" validate:"required"`
	Name string `json:"name" validate:"required"`
}

type UpdateItemBrandResponse struct {
	Brand
}

type GetItemBrandsRequest struct {
	CompanyId string `json:"company_id" validate:"required"`
	Keyword   string `json:"keyword"`
}

type Group struct {
	CompanyId   string `json:"company_id" validate:"required"`
	ItemGroupId string `json:"item_group_id" validate:"required"`
	Name        string `json:"name" validate:"required"`
}

type AddItemGroupRequest struct {
	CompanyId string `json:"company_id" validate:"required"`
	Name      string `json:"name" validate:"required"`
}

type AddItemGroupResponse struct {
	Group
}

type UpdateItemGroupRequest struct {
	Id   string `json:"id" validate:"required"`
	Name string `json:"name" validate:"required"`
}

type UpdateItemGroupResponse struct {
	Group
}

type GetItemGroupsRequest struct {
	CompanyId string `json:"company_id" validate:"required"`
	Keyword   string `json:"keyword"`
}

type Unit struct {
	CompanyId  string `json:"company_id" validate:"required"`
	ItemUnitId string `json:"item_unit_id" validate:"required"`
	Name       string `json:"name" validate:"required"`
}

type AddItemUnitRequest struct {
	CompanyId string `json:"company_id" validate:"required"`
	Name      string `json:"name" validate:"required"`
}

type AddItemUnitResponse struct {
	Unit
}

type UpdateItemUnitRequest struct {
	Id   string `json:"id" validate:"required"`
	Name string `json:"name" validate:"required"`
}

type UpdateItemUnitResponse struct {
	Unit
}

type GetItemUnitsRequest struct {
	CompanyId 		string `json:"company_id" validate:"required"`
	Keyword 	string `json:"keyword"`
}
