package impl

import (
	"net/http"

	"github.com/calvinkmts/expert-pancake/engine/errors"
	"github.com/calvinkmts/expert-pancake/engine/httpHandler"
	"github.com/expert-pancake/service/accounting/model"
)

var defaultIndonesiaBanks = []string{
	"BCA",
	"BNI",
	"BRI",
	"MANDIRI",
}

var defaultWorldWideBanks = []string{
	"PAYPAL",
}

func (a accountingService) GetAccountingBanks(w http.ResponseWriter, r *http.Request) error {

	var req model.GetAccountingBanksRequest
	httpHandler.ParseHTTPRequest(r, &req)

	errMapRequest := a.validator.Validate(req)
	if errMapRequest != nil {
		return errors.NewClientError().WithDataMap(errMapRequest)
	}

	if req.Type == "Indonesia" {
		res := model.GetAccountingBanksResponse{Banks: defaultIndonesiaBanks}
		httpHandler.WriteResponse(w, res)
	} else if req.Type == "Worldwide" {
		res := model.GetAccountingBanksResponse{Banks: defaultWorldWideBanks}
		httpHandler.WriteResponse(w, res)
	}

	return nil
}
