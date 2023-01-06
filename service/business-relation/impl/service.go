package impl

import (
	"github.com/calvinkmts/expert-pancake/engine/validator"
	db "github.com/expert-pancake/service/business-relation/db/transaction"
	"github.com/expert-pancake/service/business-relation/model"
	"github.com/expert-pancake/service/business-relation/util"
)

type businessRelationService struct {
	config    util.Config
	validator validator.Validator
	dbTrx     db.BusinessRelationTrx
}

func NewBusinessRelationService(config util.Config, validator validator.Validator, dbTrx db.BusinessRelationTrx) model.BusinessRelationService {
	return &businessRelationService{
		config:    config,
		validator: validator,
		dbTrx:     dbTrx,
	}
}
