package impl

import (
	"context"
	"net/http"

	"github.com/calvinkmts/expert-pancake/engine/errors"
	"github.com/calvinkmts/expert-pancake/engine/httpHandler"
	"github.com/expert-pancake/service/notification/model"
)

func (a notificationService) DeleteNotification(w http.ResponseWriter, r *http.Request) error {

	var req model.DeleteNotificationRequest
	httpHandler.ParseHTTPRequest(r, &req)

	errMapRequest := a.validator.Validate(req)
	if errMapRequest != nil {
		return errors.NewClientError().WithDataMap(errMapRequest)
	}

	err := a.dbTrx.DeleteNotification(context.Background(), req.NotificationId)
	if err != nil {
		return errors.NewServerError(model.DeleteNotificationError, err.Error())
	}

	res := model.DeleteNotificationResponse{
		Message: "OK",
	}
	httpHandler.WriteResponse(w, res)

	return nil
}
