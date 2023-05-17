package client

const (
	InventoryRootPath   = "http://inventory-service:4040"
	GetItemVariantsPath = "/inventory/item/variants"
	GetItemUnitsPath    = "/inventory/item/units"
)

type InventoryService interface {
	GetItemVariants(req GetItemVariantsRequest) error
	GetItemUnits(req GetItemUnitsRequest) error
}
