package impl

import (
	"github.com/expert-pancake/service/account/model"
)

type accountService struct {
}

func NewAccountService() model.AccountService {
	return &accountService{}
}
