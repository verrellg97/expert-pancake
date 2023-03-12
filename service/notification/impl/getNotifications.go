package impl

import (
	"context"
	"net/http"

	"github.com/calvinkmts/expert-pancake/engine/errors"
	"github.com/calvinkmts/expert-pancake/engine/httpHandler"
	db "github.com/expert-pancake/service/notification/db/sqlc"
	"github.com/expert-pancake/service/notification/model"
	"github.com/expert-pancake/service/notification/util"
)

func (a notificationService) GetNotifications(w http.ResponseWriter, r *http.Request) error {

	var req model.GetNotificationsRequest
	httpHandler.ParseHTTPRequest(r, &req)

	errMapRequest := a.validator.Validate(req)
	if errMapRequest != nil {
		return errors.NewClientError().WithDataMap(errMapRequest)
	}

	result, err := a.dbTrx.GetNotifications(context.Background(), db.GetNotificationsParams{
		CompanyID: req.CompanyId,
		BranchID:  util.WildCardString(req.BranchId),
		IsUnread:  req.IsUnread,
	})
	if err != nil {
		return errors.NewServerError(model.GetNotificationsError, err.Error())
	}

	var notifications = make([]model.Notification, 0)

	for _, d := range result {
		var notification = model.Notification{
			NotificationId: d.ID,
			CompanyId:      d.CompanyID,
			BranchId:       d.BranchID,
			Type:           d.Type,
			Title:          d.Title,
			Content:        d.Content,
			CreatedAt:      d.CreatedAt,
		}
		notifications = append(notifications, notification)
	}

	res := model.GetNotificationsResponse{
		Notifications: notifications,
	}
	httpHandler.WriteResponse(w, res)

	return nil
}
