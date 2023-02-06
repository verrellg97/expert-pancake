package client

const (
	WarehouseRootPath     = "http://127.0.0.1:4050"
	GetWarehousesPath     = "/warehouses"
	GetWarehouseRacksPath = "/warehouse/racks"
)

type WarehouseService interface {
	GetWarehouses(req GetWarehousesRequest) error
	GetWarehouseRacks(req GetWarehouseRacksRequest) error
}
