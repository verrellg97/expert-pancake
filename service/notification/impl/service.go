package impl

import (
	"github.com/calvinkmts/expert-pancake/engine/validator"
	db "github.com/expert-pancake/service/notification/db/transaction"
	"github.com/expert-pancake/service/notification/model"
	"github.com/expert-pancake/service/notification/util"
)

type notificationService struct {
	config    util.Config
	validator validator.Validator
	dbTrx     db.NotificationTrx
}

func NewNotificationService(config util.Config, validator validator.Validator, dbTrx db.NotificationTrx) model.NotificationService {
	return &notificationService{
		config:    config,
		validator: validator,
		dbTrx:     dbTrx,
	}
}
