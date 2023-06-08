package client

type GetWarehousesRequest struct {
	Id        string   `json:"id"`
	BranchIds []string `json:"branch_ids"`
	BranchId  string   `json:"branch_id" validate:"required"`
	Keyword   string   `json:"keyword"`
}

type GetWarehouseRacksRequest struct {
	Id          string `json:"id"`
	WarehouseId string `json:"warehouse_id" validate:"required"`
	Keyword     string `json:"keyword"`
}
