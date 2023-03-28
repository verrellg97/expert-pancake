package client

const (
	InventoryRootPath      = "http://inventory-service:4040"
	UpsertUnitCategoryPath = "/inventory/unit-category/upsert"
	AddUnitPath            = "/inventory/unit/add"
	UpsertPricelistPath    = "/inventory/pricelist/upsert"
)

type InventoryService interface {
	UpsertUnitCategory(req UpsertUnitCategoryRequest) error
	AddUnit(req AddUnitRequest) error
	UpsertPricelist(req UpsertPricelistRequest) error
}
