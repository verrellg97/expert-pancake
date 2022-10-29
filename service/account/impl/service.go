package impl

import (
	"github.com/calvinkmts/expert-pancake/engine/validator"
	db "github.com/expert-pancake/service/account/db/transaction"
	"github.com/expert-pancake/service/account/model"
)

type accountService struct {
	validator validator.Validator
	dbTrx     db.AccountTrx
}

func NewAccountService(validator validator.Validator, dbTrx db.AccountTrx) model.AccountService {
	return &accountService{
		validator: validator,
		dbTrx:     dbTrx,
	}
}
