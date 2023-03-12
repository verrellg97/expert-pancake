package impl

import (
	"context"
	"net/http"

	"github.com/calvinkmts/expert-pancake/engine/errors"
	"github.com/calvinkmts/expert-pancake/engine/httpHandler"
	"github.com/expert-pancake/service/notification/model"
)

func (a notificationService) ReadNotification(w http.ResponseWriter, r *http.Request) error {

	var req model.ReadNotificationRequest
	httpHandler.ParseHTTPRequest(r, &req)

	errMapRequest := a.validator.Validate(req)
	if errMapRequest != nil {
		return errors.NewClientError().WithDataMap(errMapRequest)
	}

	err := a.dbTrx.ReadNotification(context.Background(), req.NotificationId)
	if err != nil {
		return errors.NewServerError(model.ReadNotificationError, err.Error())
	}

	res := model.ReadNotificationResponse{
		Message: "OK",
	}
	httpHandler.WriteResponse(w, res)

	return nil
}
