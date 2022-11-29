package impl

import (
	"github.com/calvinkmts/expert-pancake/engine/validator"
	db "github.com/expert-pancake/service/accounting/db/transaction"
	"github.com/expert-pancake/service/accounting/model"
	"github.com/expert-pancake/service/accounting/util"
)

type accountingService struct {
	config    util.Config
	validator validator.Validator
	dbTrx     db.AccountingTrx
}

func NewAccountingService(config util.Config, validator validator.Validator, dbTrx db.AccountingTrx) model.AccountingService {
	return &accountingService{
		config:    config,
		validator: validator,
		dbTrx:     dbTrx,
	}
}
