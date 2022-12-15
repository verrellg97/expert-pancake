package impl

import (
	"context"
	"net/http"

	"github.com/calvinkmts/expert-pancake/engine/errors"
	"github.com/calvinkmts/expert-pancake/engine/httpHandler"
	db "github.com/expert-pancake/service/business/db/sqlc"
	"github.com/expert-pancake/service/business/model"
	uuid "github.com/satori/go.uuid"
)

func (a businessService) RegisterCompany(w http.ResponseWriter, r *http.Request) error {

	var req model.RegisterCompanyRequest

	httpHandler.ParseHTTPRequest(r, &req)

	errMapRequest := a.validator.Validate(req)
	if errMapRequest != nil {
		return errors.NewClientError().WithDataMap(errMapRequest)
	}

	arg := db.InsertCompanyParams{
		ID:                uuid.NewV4().String(),
		UserID:            req.AccountId,
		Name:              req.Name,
		InitialName:       req.InitialName,
		Type:              req.Type,
		ResponsiblePerson: req.ResponsiblePerson,
	}

	result, err := a.dbTrx.CreateNewCompanyTrx(context.Background(), arg)
	if err != nil {
		return errors.NewServerError(model.CreateNewCompanyError, err.Error())
	}

	res := model.RegisterCompanyResponse{
		Company: model.Company{
			AccountId:         result.AccountId,
			CompanyId:         result.CompanyId,
			Name:              result.Name,
			InitialName:       result.InitialName,
			Type:              result.Type,
			ResponsiblePerson: result.ResponsiblePerson,
			Branches:          result.Branches,
		},
	}

	httpHandler.WriteResponse(w, res)

	return nil
}
