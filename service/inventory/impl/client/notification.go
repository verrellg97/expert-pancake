package client

const (
	NotificationRootPath   = "http://notification-service:4060"
	InsertNotificationPath = "/notification/insert"
)

type NotificationService interface {
	InsertNotification(req InsertNotificationRequest) error
}
