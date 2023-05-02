package client

type Warehouse struct {
	WarehouseId string `json:"warehouse_id" validate:"required"`
	BranchId    string `json:"branch_id" validate:"required"`
	Code        string `json:"code" validate:"required"`
	Name        string `json:"name" validate:"required"`
	Type        string `json:"type" validate:"required"`
	Address     string `json:"address" validate:"required"`
}

type GetWarehousesResponse struct {
	Result struct {
		Warehouses []Warehouse `json:"warehouses" validate:"required"`
	} `json:"result"`
}

type WarehouseRack struct {
	RackId      string `json:"rack_id" validate:"required"`
	WarehouseId string `json:"warehouse_id" validate:"required"`
	Name        string `json:"name" validate:"required"`
}

type GetWarehouseRacksResponse struct {
	Result struct {
		WarehouseRacks []WarehouseRack `json:"warehouse_racks" validate:"required"`
	} `json:"result"`
}
