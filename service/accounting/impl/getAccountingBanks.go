package impl

import (
	"net/http"

	"github.com/calvinkmts/expert-pancake/engine/errors"
	"github.com/calvinkmts/expert-pancake/engine/httpHandler"
	"github.com/expert-pancake/service/accounting/model"
	"github.com/gorilla/schema"
)

var defaultIndonesiaBanks = []model.Bank{
	{BankName: "BCA", BankCode: "64"},
	{BankName: "BNI", BankCode: "65"},
	{BankName: "BRI", BankCode: "66"},
}

var defaultWorldWideBanks = []model.Bank{
	{BankName: "PAYPAL", BankCode: "1"},
}

func (a accountingService) GetAccountingBanks(w http.ResponseWriter, r *http.Request) error {

	var req model.GetAccountingBanksRequest
	schema.NewDecoder().Decode(&req, r.URL.Query())

	errMapRequest := a.validator.Validate(req)
	if errMapRequest != nil {
		return errors.NewClientError().WithDataMap(errMapRequest)
	}

	var res = model.GetAccountingBanksResponse{Banks: []model.Bank{}}

	if req.Type == "Indonesia" {
		res = model.GetAccountingBanksResponse{Banks: defaultIndonesiaBanks}
	} else if req.Type == "Worldwide" {
		res = model.GetAccountingBanksResponse{Banks: defaultWorldWideBanks}
	}

	httpHandler.WriteResponse(w, res)

	return nil
}
