package impl

import (
	"net/http"

	"github.com/calvinkmts/expert-pancake/engine/httpHandler"
	"github.com/expert-pancake/service/business/model"
)

var defaultTypes = []string{
	"UMKM",
}

func (a businessService) GetCompanyTypes(w http.ResponseWriter, r *http.Request) error {

	res := model.GetCompanyTypesResponse{Types: defaultTypes}
	httpHandler.WriteResponse(w, res)

	return nil
}
