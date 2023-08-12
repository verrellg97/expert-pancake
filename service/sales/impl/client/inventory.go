package client

const (
	InventoryRootPath   = "http://inventory-service:4040"
	GetItemVariantsPath = "/inventory/item/variants"
	GetItemUnitsPath    = "/inventory/item/units"

	InsertStockMovementPath = "/inventory/stock-movement/insert"
	DeleteStockMovementPath = "/inventory/stock-movement/delete"
)

type InventoryService interface {
	GetItemVariants(req GetItemVariantsRequest) error
	GetItemUnits(req GetItemUnitsRequest) error

	InsertStockMovement(req InsertStockMovementRequest) error
	DeleteStockMovement(req DeleteStockMovementRequest) error
}
