package impl

import (
	"context"
	"log"
	"net/http"
	"strconv"

	"github.com/calvinkmts/expert-pancake/engine/errors"
	"github.com/calvinkmts/expert-pancake/engine/httpHandler"
	db "github.com/expert-pancake/service/business/db/transaction"
	"github.com/expert-pancake/service/business/model"
)

func (a businessService) UpdateCompany(w http.ResponseWriter, r *http.Request) error {

	var req model.UpdateCompanyRequest
	httpHandler.ParseHTTPRequest(r, &req)

	errMapRequest := a.validator.Validate(req)
	if errMapRequest != nil {
		return errors.NewClientError().WithDataMap(errMapRequest)
	}

	is_deleted, err := strconv.Atoi(req.IsDeleted)
	if err != nil {
		// ... handle error
		log.Panic(err)
	}

	result, err := a.dbTrx.UpdateCompanyTrx(context.Background(), db.UpdateCompanyTrxParams{
		UserID:            req.AccountId,
		CompanyID:         req.CompanyId,
		Name:              req.Name,
		InitialName:       req.InitialName,
		Type:              req.Type,
		ResponsiblePerson: req.ResponsiblePerson,
		IsDeleted:         is_deleted,
	})
	if err != nil {
		return errors.NewServerError(model.UpdateCompanyError, err.Error())
	}

	res := result
	httpHandler.WriteResponse(w, res)

	return nil
}
