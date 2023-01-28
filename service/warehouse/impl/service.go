package impl

import (
	"github.com/calvinkmts/expert-pancake/engine/validator"
	db "github.com/expert-pancake/service/warehouse/db/transaction"
	"github.com/expert-pancake/service/warehouse/model"
	"github.com/expert-pancake/service/warehouse/util"
)

type warehouseService struct {
	config    util.Config
	validator validator.Validator
	dbTrx     db.WarehouseTrx
}

func NewWarehouseService(config util.Config, validator validator.Validator, dbTrx db.WarehouseTrx) model.WarehouseService {
	return &warehouseService{
		config:    config,
		validator: validator,
		dbTrx:     dbTrx,
	}
}
