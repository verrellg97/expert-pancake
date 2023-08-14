package client

type ItemGroup struct {
	Id   string `json:"id" validate:"required"`
	Name string `json:"name" validate:"required"`
}

type Item struct {
	CompanyId   string      `json:"company_id" validate:"required"`
	ItemId      string      `json:"item_id" validate:"required"`
	VariantId   string      `json:"variant_id" validate:"required"`
	ImageUrl    string      `json:"image_url" validate:"required"`
	Code        string      `json:"code" validate:"required"`
	Barcode     string      `json:"barcode" validate:"required"`
	Name        string      `json:"name" validate:"required"`
	VariantName string      `json:"variant_name" validate:"required"`
	BrandId     string      `json:"brand_id" validate:"required"`
	BrandName   string      `json:"brand_name" validate:"required"`
	Groups      []ItemGroup `json:"groups" validate:"required"`
	Tag         []string    `json:"tag" validate:"required"`
	Description string      `json:"description" validate:"required"`
	IsDefault   bool        `json:"is_default" validate:"required"`
	Price       string      `json:"price" validate:"required"`
}

type GetItemVariantsResponse struct {
	Result struct {
		ItemVariants []Item `json:"item_variants" validate:"required"`
	} `json:"result"`
}

type ItemUnit struct {
	ItemUnitId string `json:"item_unit_id" validate:"required"`
	ItemId     string `json:"item_id" validate:"required"`
	UnitId     string `json:"unit_id" validate:"required"`
	UnitName   string `json:"unit_name" validate:"required"`
	Value      string `json:"value" validate:"required"`
	IsDefault  bool   `json:"is_default" validate:"required"`
}

type GetItemUnitsResponse struct {
	Result struct {
		ItemUnits []ItemUnit `json:"item_units" validate:"required"`
	}
}

type InsertStockMovementResponse struct {
	Result struct {
		Message string `json:"message" validate:"required"`
	} `json:"result"`
}

type DeleteStockMovementResponse struct {
	Result struct {
		Message string `json:"message" validate:"required"`
	} `json:"result"`
}
