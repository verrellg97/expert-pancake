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

	AddItem(w http.ResponseWriter, r *http.Request) error
	UpdateItem(w http.ResponseWriter, r *http.Request) error
	GetItems(w http.ResponseWriter, r *http.Request) error

	UpsertItemVariant(w http.ResponseWriter, r *http.Request) error
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

type Item struct {
	CompanyId   string `json:"company_id" validate:"required"`
	ItemId      string `json:"item_id" validate:"required"`
	VariantId   string `json:"variant_id" validate:"required"`
	ImageUrl    string `json:"image_url" validate:"required"`
	Code        string `json:"code" validate:"required"`
	Name        string `json:"name" validate:"required"`
	VariantName string `json:"variant_name" validate:"required"`
	BrandId     string `json:"brand_id" validate:"required"`
	BrandName   string `json:"brand_name" validate:"required"`
	GroupId     string `json:"group_id" validate:"required"`
	GroupName   string `json:"group_name" validate:"required"`
	Tag         string `json:"tag" validate:"required"`
	Description string `json:"description" validate:"required"`
	IsDefault   bool   `json:"is_default"`
	Price       string `json:"price" validate:"required"`
	Stock       string `json:"stock" validate:"required"`
}

type AddItemRequest struct {
	CompanyId   string `json:"company_id" validate:"required"`
	ImageUrl    string `json:"image_url"`
	Name        string `json:"name" validate:"required"`
	BrandId     string `json:"brand_id" validate:"required"`
	GroupId     string `json:"group_id" validate:"required"`
	Tag         string `json:"tag"`
	Description string `json:"description"`
}

type AddItemResponse struct {
	Item
}

type UpdateItemRequest struct {
	ItemId      string `json:"item_id" validate:"required"`
	ImageUrl    string `json:"image_url"`
	Name        string `json:"name" validate:"required"`
	BrandId     string `json:"brand_id" validate:"required"`
	GroupId     string `json:"group_id" validate:"required"`
	Tag         string `json:"tag"`
	Description string `json:"description"`
}

type UpdateItemResponse struct {
	Item
}

type GetItemsRequest struct {
	CompanyId string `json:"company_id" validate:"required"`
	Keyword   string `json:"keyword"`
}

type GetItemsResponse struct {
	Item []Item `json:"items" validate:"required"`
}

type UpsertItemVariantRequest struct {
	ItemVariantId string `json:"item_variant_id"`
	ItemId        string `json:"item_id" validate:"required"`
	ImageUrl      string `json:"image_url"`
	Name          string `json:"name" validate:"required"`
	Price         string `json:"price" validate:"required"`
	Stock         string `json:"stock" validate:"required"`
}

type UpsertItemVariantResponse struct {
	Item
}
