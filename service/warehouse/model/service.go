package model

import (
	"net/http"
)

type WarehouseService interface {
	UpsertWarehouse(w http.ResponseWriter, r *http.Request) error
	GetWarehouses(w http.ResponseWriter, r *http.Request) error
	DeleteWarehouse(w http.ResponseWriter, r *http.Request) error

	UpsertWarehouseRack(w http.ResponseWriter, r *http.Request) error
	GetWarehouseRacks(w http.ResponseWriter, r *http.Request) error
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

type GetWarehousesRequest struct {
	Id       *string `json:"id"`
	BranchId string  `json:"branch_id" validate:"required"`
	Keyword  string  `json:"keyword"`
}

type GetWarehousesResponse struct {
	Warehouses []Warehouse `json:"warehouses" validate:"required"`
}

type WarehouseRack struct {
	RackId      string `json:"rack_id" validate:"required"`
	WarehouseId string `json:"warehouse_id" validate:"required"`
	Name        string `json:"name" validate:"required"`
}

type UpsertWarehouseRackRequest struct {
	RackId      string `json:"rack_id"`
	WarehouseId string `json:"warehouse_id" validate:"required"`
	Name        string `json:"name" validate:"required"`
}

type UpsertWarehouseRackResponse struct {
	WarehouseRack
}

type GetWarehouseRacksRequest struct {
	WarehouseId string `json:"warehouse_id" validate:"required"`
	Keyword     string `json:"keyword"`
}

type GetWarehouseRacksResponse struct {
	WarehouseRacks []WarehouseRack `json:"warehouse_racks" validate:"required"`
}

type DeleteWarehouseRequest struct {
	WarehouseId string `json:"warehouse_id" validate:"required"`
}

type DeleteWarehouseResponse struct {
	Message string `json:"message"`
}
