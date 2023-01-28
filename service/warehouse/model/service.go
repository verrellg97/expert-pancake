package model

import (
	"net/http"
)

type WarehouseService interface {
	UpsertRack(w http.ResponseWriter, r *http.Request) error
	GetRacks(w http.ResponseWriter, r *http.Request) error
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
	BranchId string `json:"branch_id" validate:"required"`
	Keyword  string `json:"keyword"`
	Mode     string `json:"mode" validate:"required"`
}

type GetRacksResponse struct {
	Racks []Rack `json:"racks"`
}
