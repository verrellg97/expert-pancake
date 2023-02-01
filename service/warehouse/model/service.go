package model

import (
	"net/http"
)

type WarehouseService interface {
	UpsertWarehouse(w http.ResponseWriter, r *http.Request) error
}

type Warehouse struct {
	WarehouseId string `json:"warehouse_id" validate:"required"`
	BranchId    string `json:"branch_id" validate:"required"`
	Code        string `json:"code" validate:"required"`
	Name        string `json:"name" validate:"required"`
	Type        string `json:"type" validate:"required"`
	Address     string `json:"address" validate:"required"`
}

type UpsertWarehouseRequest struct {
	WarehouseId string `json:"warehouse_id"`
	BranchId    string `json:"branch_id" validate:"required"`
	Name        string `json:"name" validate:"required"`
	Type        string `json:"type" validate:"required"`
	Address     string `json:"address" validate:"required"`
}

type UpsertWarehouseResponse struct {
	Warehouse
}
