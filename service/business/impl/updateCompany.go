package impl

import (
	"context"
	"log"
	"net/http"
	"strconv"

	"github.com/calvinkmts/expert-pancake/engine/errors"
	"github.com/calvinkmts/expert-pancake/engine/httpHandler"
	db "github.com/expert-pancake/service/business/db/sqlc"
	"github.com/expert-pancake/service/business/model"
)

func (a businessService) UpdateCompany(w http.ResponseWriter, r *http.Request) error {

	var req model.UpdateCompanyRequest
	httpHandler.ParseHTTPRequest(r, &req)

	errMapRequest := a.validator.Validate(req)
	if errMapRequest != nil {
		return errors.NewClientError().WithDataMap(errMapRequest)
	}

	is_deleted, err := strconv.ParseInt(req.IsDeleted, 10, 32)
	if err != nil {
		// ... handle error
		log.Panic(err)
	}

	result, err := a.dbTrx.UpsertCompany(context.Background(), db.UpsertCompanyParams{
		ID:                req.CompanyId,
		UserID:            req.AccountId,
		Name:              req.Name,
		InitialName:       req.InitialName,
		Type:              req.Type,
		ResponsiblePerson: req.ResponsiblePerson,
		IsDeleted:         int32(is_deleted),
	})
	if err != nil {
		return errors.NewServerError(model.UpdateCompanyError, err.Error())
	}

	res := model.RegisterCompanyResponse{
		Company: model.Company{
			AccountId:         result.UserID,
			CompanyId:         result.ID,
			Name:              result.Name,
			InitialName:       result.InitialName,
			Type:              result.Type,
			ResponsiblePerson: result.ResponsiblePerson,
		},
	}
	httpHandler.WriteResponse(w, res)

	return nil
}
