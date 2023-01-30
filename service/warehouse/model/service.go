package model

import (
	"net/http"
)

type WarehouseService interface {
	UpsertRack(w http.ResponseWriter, r *http.Request) error
	GetRacks(w http.ResponseWriter, r *http.Request) error

	GetWarehouses(w http.ResponseWriter, r *http.Request) error
}

type Rack struct {
	RackId   string `json:"rack_id" validate:"required"`
	BranchId string `json:"branch_id" validate:"required"`
	Name     string `json:"name" validate:"required"`
}

type UpsertRackRequest struct {
	Id       string `json:"id"`
	BranchId string `json:"branch_id" validate:"required"`
	Name     string `json:"name" validate:"required"`
}

type UpsertRackResponse struct {
	Rack
}

type GetRacksRequest struct {
	BranchId       string `json:"branch_id" validate:"required"`
	Keyword        string `json:"keyword"`
	IsGetAvailable bool   `json:"is_get_available" validate:"boolean"`
}

type GetRacksResponse struct {
	Racks []Rack `json:"racks"`
}

type Warehouse struct {
	WarehouseId string `json:"warehouse_id" validate:"required"`
	BranchId    string `json:"branch_id" validate:"required"`
	Code        string `json:"code" validate:"required"`
	Name        string `json:"name" validate:"required"`
	Type        string `json:"type" validate:"required"`
	Address     string `json:"address" validate:"required"`
}

type GetWarehousesRequest struct {
	BranchId string `json:"branch_id" validate:"required"`
	Keyword  string `json:"keyword"`
	Type     string `json:"type"`
}
