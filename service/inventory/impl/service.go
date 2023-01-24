package impl

import (
	"github.com/calvinkmts/expert-pancake/engine/validator"
	db "github.com/expert-pancake/service/inventory/db/transaction"
	"github.com/expert-pancake/service/inventory/model"
	"github.com/expert-pancake/service/inventory/util"
)

type inventoryService struct {
	config    util.Config
	validator validator.Validator
	dbTrx     db.InventoryTrx
}

func NewInventoryService(config util.Config, validator validator.Validator, dbTrx db.InventoryTrx) model.InventoryService {
	return &inventoryService{
		config:    config,
		validator: validator,
		dbTrx:     dbTrx,
	}
}
