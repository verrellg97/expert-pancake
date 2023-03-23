package client

type UnitCategory struct {
	Id        string `json:"id" validate:"required"`
	CompanyId string `json:"company_id" validate:"required"`
	Name      string `json:"name" validate:"required"`
}

type UpsertUnitCategoryResponse struct {
	Result struct {
		UnitCategory
	} `json:"result"`
}
