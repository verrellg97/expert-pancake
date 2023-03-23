package client

const (
	InventoryRootPath      = "http://localhost:3000"
	UpsertUnitCategoryPath = "/inventory/unit-category/upsert"
	AddUnitPath            = "/inventory/unit/add"
	UpsertPricelistPath    = "/inventory/pricelist/upsert"
)

type InventoryService interface {
	UpsertUnitCategory(req UpsertUnitCategoryRequest) error
	AddUnit(req AddUnitRequest) error
	UpsertPricelist(req UpsertPricelistRequest) error
}
