package impl

import (
	"context"
	"net/http"

	"github.com/calvinkmts/expert-pancake/engine/errors"
	"github.com/calvinkmts/expert-pancake/engine/httpHandler"
	"github.com/expert-pancake/service/business/impl/client"
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
		user, err := client.GetUserInformation(
			client.GetUserInformationRequest{
				AccountId: d.UserID,
			})
		if err != nil {
			return errors.NewServerError(model.GetReceiveMemberRequestsError, err.Error())
		}

		var memberRequest = model.MemberRequest{
			Id:          d.ID,
			UserId:      d.UserID,
			ImageUrl:    user.Result.ImageUrl,
			FullName:    user.Result.FullName,
			Nickname:    user.Result.Nickname,
			Email:       user.Result.Email,
			PhoneNumber: user.Result.PhoneNumber,
			Status:      d.Status,
		}
		memberRequests = append(memberRequests, memberRequest)
	}

	res := model.GetReceiveMemberRequestsResponse{
		MemberRequests: memberRequests,
	}
	httpHandler.WriteResponse(w, res)

	return nil
}
