package model

import (
	"net/http"
)

type InventoryService interface {
	AddItemBrand(w http.ResponseWriter, r *http.Request) error
	UpdateItemBrand(w http.ResponseWriter, r *http.Request) error
	GetItemBrands(w http.ResponseWriter, r *http.Request) error

	AddItemUnit(w http.ResponseWriter, r *http.Request) error
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
	Id 		string `json:"id" validate:"required"`
	Name 	string `json:"name" validate:"required"`
}

type UpdateItemBrandResponse struct {
	Brand
}

type GetItemBrandsRequest struct {
	CompanyId 		string `json:"company_id" validate:"required"`
	Keyword 	string `json:"keyword"`
}

type Unit struct {
	CompanyId   string `json:"company_id" validate:"required"`
	ItemUnitId string `json:"item_unit_id" validate:"required"`
	Name        string `json:"name" validate:"required"`
}

type AddItemUnitRequest struct {
	CompanyId string `json:"company_id" validate:"required"`
	Name      string `json:"name" validate:"required"`
}

type AddItemUnitResponse struct {
	Unit
}
