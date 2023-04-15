package impl

import (
	"github.com/calvinkmts/expert-pancake/engine/validator"
	db "github.com/expert-pancake/service/sales/db/transaction"
	"github.com/expert-pancake/service/sales/model"
	"github.com/expert-pancake/service/sales/util"
)

type salesService struct {
	config    util.Config
	validator validator.Validator
	dbTrx     db.SalesTrx
}

func NewSalesService(config util.Config, validator validator.Validator, dbTrx db.SalesTrx) model.SalesService {
	return &salesService{
		config:    config,
		validator: validator,
		dbTrx:     dbTrx,
	}
}
