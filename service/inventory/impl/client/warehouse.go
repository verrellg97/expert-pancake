package client

const (
	WarehouseRootPath     = "http://warehouse-service:4050"
	GetWarehousesPath     = "/warehouses"
	GetWarehouseRacksPath = "/warehouse/racks"
)

type WarehouseService interface {
	GetWarehouses(req GetWarehousesRequest) error
	GetWarehouseRacks(req GetWarehouseRacksRequest) error
}
