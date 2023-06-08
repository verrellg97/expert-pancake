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

type GetCheckStockHistoryResponse struct {
	Result struct {
		Status bool `json:"status" validate:"required"`
	} `json:"result"`
}
