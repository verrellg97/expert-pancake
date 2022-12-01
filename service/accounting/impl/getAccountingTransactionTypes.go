package impl

import (
	"net/http"

	"github.com/calvinkmts/expert-pancake/engine/errors"
	"github.com/calvinkmts/expert-pancake/engine/httpHandler"
	"github.com/expert-pancake/service/accounting/model"
	"github.com/gorilla/schema"
)

var defaultTransactionInTypes = []string{
	"PENERIMAAN KAS",
}
var defaultTransactionOutTypes = []string{
	"PENGELUARAN KAS",
}

func (a accountingService) GetAccountingTransactionTypes(w http.ResponseWriter, r *http.Request) error {

	var req model.GetAccountingTransactionTypesRequest
	schema.NewDecoder().Decode(&req, r.URL.Query())

	errMapRequest := a.validator.Validate(req)
	if errMapRequest != nil {
		return errors.NewClientError().WithDataMap(errMapRequest)
	}

	var res = model.GetAccountingTransactionTypesResponse{TransactionTypes: []string{}}

	if req.Type == "In" {
		res = model.GetAccountingTransactionTypesResponse{TransactionTypes: defaultTransactionInTypes}
	} else if req.Type == "Out" {
		res = model.GetAccountingTransactionTypesResponse{TransactionTypes: defaultTransactionOutTypes}
	}

	httpHandler.WriteResponse(w, res)

	return nil
}
