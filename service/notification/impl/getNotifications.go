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

	var isReadFilter = false
	if req.IsReadFilter != nil {
		isReadFilter = true
	} else {
		req.IsReadFilter = &isReadFilter
	}

	result, err := a.dbTrx.GetNotifications(context.Background(), db.GetNotificationsParams{
		CompanyID:    req.CompanyId,
		BranchID:     util.WildCardString(req.BranchId),
		IsReadFilter: isReadFilter,
		IsRead:       *req.IsReadFilter,
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
			Title:          d.Title,
			Content:        d.Content,
			Type:           d.Type,
			CreatedAt:      d.CreatedAt.Time,
		}
		notifications = append(notifications, notification)
	}

	res := model.GetNotificationsResponse{
		Notifications: notifications,
	}
	httpHandler.WriteResponse(w, res)

	return nil
}
