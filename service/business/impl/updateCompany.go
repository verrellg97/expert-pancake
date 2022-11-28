package impl

import (
	"context"
	"net/http"

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

	result, err := a.dbTrx.UpdateCompany(context.Background(), db.UpdateCompanyParams{
		ID:                req.CompanyId,
		Name:              req.Name,
		InitialName:       req.InitialName,
		Type:              req.Type,
		ResponsiblePerson: req.ResponsiblePerson,
		IsDeleted:         req.IsDeleted,
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
