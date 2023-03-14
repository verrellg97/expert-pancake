package impl

import (
	"context"
	"net/http"

	"github.com/calvinkmts/expert-pancake/engine/errors"
	"github.com/calvinkmts/expert-pancake/engine/httpHandler"
	db "github.com/expert-pancake/service/notification/db/sqlc"
	"github.com/expert-pancake/service/notification/model"
	uuid "github.com/satori/go.uuid"
)

func (a notificationService) InsertNotification(w http.ResponseWriter, r *http.Request) error {

	var req model.InsertNotificationRequest

	httpHandler.ParseHTTPRequest(r, &req)

	errMapRequest := a.validator.Validate(req)
	if errMapRequest != nil {
		return errors.NewClientError().WithDataMap(errMapRequest)
	}

	arg := db.InsertNotificationParams{
		ID:        uuid.NewV4().String(),
		CompanyID: req.CompanyId,
		BranchID:  req.BranchId,
		Title:     req.Title,
		Content:   req.Content,
		Type:      req.Type,
	}

	result, err := a.dbTrx.InsertNotification(context.Background(), arg)
	if err != nil {
		return errors.NewServerError(model.InsertNotificationError, err.Error())
	}

	res := model.InsertNotificationResponse{
		Notification: model.Notification{
			NotificationId: result.ID,
			CompanyId:      result.CompanyID,
			BranchId:       result.BranchID,
			Title:          result.Title,
			Content:        result.Content,
			Type:           result.Type,
			CreatedAt:      result.CreatedAt.Time,
		},
	}

	httpHandler.WriteResponse(w, res)

	return nil
}
