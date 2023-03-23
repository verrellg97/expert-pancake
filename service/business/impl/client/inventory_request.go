package client

type UpsertUnitCategoryRequest struct {
	Id        string `json:"id"`
	CompanyId string `json:"company_id" validate:"required"`
	Name      string `json:"name" validate:"required"`
}

type AddUnitRequest struct {
	CompanyId      string `json:"company_id" validate:"required"`
	UnitCategoryId string `json:"unit_category_id" validate:"required"`
	Name           string `json:"name" validate:"required"`
}

type UpsertPricelistRequest struct {
	PricelistId string `json:"pricelist_id"`
	CompanyId   string `json:"company_id" validate:"required"`
	Name        string `json:"name" validate:"required"`
	StartDate   string `json:"start_date" validate:"required"`
	EndDate     string `json:"end_date"`
	IsDefault   bool   `json:"is_default"`
}
