package impl

import (
	"github.com/calvinkmts/expert-pancake/engine/validator"
	db "github.com/expert-pancake/service/account/db/transaction"
	"github.com/expert-pancake/service/account/model"
	"github.com/expert-pancake/service/account/token"
	"github.com/expert-pancake/service/account/util"
)

type accountService struct {
	config     util.Config
	validator  validator.Validator
	dbTrx      db.AccountTrx
	tokenMaker token.Maker
}

func NewAccountService(config util.Config, validator validator.Validator, dbTrx db.AccountTrx, tokenMaker token.Maker) model.AccountService {
	return &accountService{
		config:     config,
		validator:  validator,
		dbTrx:      dbTrx,
		tokenMaker: tokenMaker,
	}
}
