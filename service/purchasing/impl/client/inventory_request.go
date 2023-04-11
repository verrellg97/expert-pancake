package client

type GetItemVariantsRequest struct {
	Id      string `json:"id"`
	ItemId  string `json:"item_id"`
	Keyword string `json:"keyword"`
}

type GetItemUnitsRequest struct {
	Id      string `json:"id"`
	ItemId  string `json:"item_id" validate:"required"`
	Keyword string `json:"keyword"`
}
