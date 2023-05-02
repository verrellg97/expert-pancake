package client

type GetWarehousesRequest struct {
	Id       string `json:"id"`
	BranchId string `json:"branch_id" validate:"required"`
	Keyword  string `json:"keyword"`
}

type GetWarehouseRacksRequest struct {
	Id          string `json:"id"`
	WarehouseId string `json:"warehouse_id" validate:"required"`
	Keyword     string `json:"keyword"`
}
