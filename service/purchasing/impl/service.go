package impl

import (
	"github.com/calvinkmts/expert-pancake/engine/validator"
	db "github.com/expert-pancake/service/purchasing/db/transaction"
	"github.com/expert-pancake/service/purchasing/model"
	"github.com/expert-pancake/service/purchasing/util"
)

type purchasingService struct {
	config    util.Config
	validator validator.Validator
	dbTrx     db.PurchasingTrx
}

func NewPurchasingService(config util.Config, validator validator.Validator, dbTrx db.PurchasingTrx) model.PurchasingService {
	return &purchasingService{
		config:    config,
		validator: validator,
		dbTrx:     dbTrx,
	}
}
