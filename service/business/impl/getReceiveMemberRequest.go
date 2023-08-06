package impl

import (
	"context"
	"net/http"

	"github.com/calvinkmts/expert-pancake/engine/errors"
	"github.com/calvinkmts/expert-pancake/engine/httpHandler"
	"github.com/expert-pancake/service/business/model"
)

func (a businessService) GetReceiveMemberRequests(w http.ResponseWriter, r *http.Request) error {

	var req model.GetReceiveMemberRequestsRequest
	httpHandler.ParseHTTPRequest(r, &req)

	errMapRequest := a.validator.Validate(req)
	if errMapRequest != nil {
		return errors.NewClientError().WithDataMap(errMapRequest)
	}

	result, err := a.dbTrx.GetReceiveMemberRequests(context.Background(), req.CompanyId)
	if err != nil {
		return errors.NewServerError(model.GetReceiveMemberRequestsError, err.Error())
	}

	var memberRequests = make([]model.MemberRequest, 0)

	for _, d := range result {
		var memberRequest = model.MemberRequest{
			Id:     d.ID,
			Status: d.Status,
		}
		memberRequests = append(memberRequests, memberRequest)
	}

	res := memberRequests
	httpHandler.WriteResponse(w, res)

	return nil
}
