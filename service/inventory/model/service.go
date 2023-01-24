package model

import (
	"net/http"
)

type InventoryService interface {
	AddItemBrand(w http.ResponseWriter, r *http.Request) error
}

type Brand struct {
	CompanyId   string `json:"company_id" validate:"required"`
	ItemBrandId string `json:"item_brand_id" validate:"required"`
	Name        string `json:"name" validate:"required"`
}

<<<<<<< HEAD
type AddItemBrandRequest struct {
=======
type AddBrandRequest struct {
>>>>>>> e67e22d (feat(service inventory) new endpoint add item brand)
	CompanyId string `json:"company_id" validate:"required"`
	Name      string `json:"name" validate:"required"`
}

<<<<<<< HEAD
type AddItemBrandResponse struct {
=======
type AddBrandResponse struct {
>>>>>>> e67e22d (feat(service inventory) new endpoint add item brand)
	Brand
}
