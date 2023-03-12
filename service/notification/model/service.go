package model

import (
	"database/sql"
	"net/http"
)

type NotificationService interface {
	InsertNotification(w http.ResponseWriter, r *http.Request) error
	ReadNotification(w http.ResponseWriter, r *http.Request) error
	GetNotifications(w http.ResponseWriter, r *http.Request) error
	DeleteNotification(w http.ResponseWriter, r *http.Request) error
}

type Notification struct {
	NotificationId string       `json:"notification_id" validate:"required"`
	CompanyId      string       `json:"company_id" validate:"required"`
	BranchId       string       `json:"branch_id" validate:"required"`
	Type           string       `json:"type" validate:"required"`
	Title          string       `json:"Title" validate:"required"`
	Content        string       `json:"Content" validate:"required"`
	CreatedAt      sql.NullTime `json:"created_at" validate:"required"`
}

type InsertNotificationRequest struct {
	CompanyId string `json:"company_id" validate:"required"`
	BranchId  string `json:"branch_id" validate:"required"`
	Title     string `json:"Title" validate:"required"`
	Content   string `json:"Content" validate:"required"`
	Type      string `json:"type" validate:"required"`
}

type InsertNotificationResponse struct {
	Notification
}

type GetNotificationsRequest struct {
	CompanyId string `json:"Company_id" validate:"required"`
	BranchId  string `json:"Branch_id"`
	IsUnread  bool   `json:"Is_unread"`
}

type GetNotificationsResponse struct {
	Notifications []Notification `json:"notifications" validate:"required"`
}

type ReadNotificationRequest struct {
	NotificationId string `json:"notification_id" validate:"required"`
}

type ReadNotificationResponse struct {
	Message string `json:"message"`
}

type DeleteNotificationRequest struct {
	NotificationId string `json:"notification_id" validate:"required"`
}

type DeleteNotificationResponse struct {
	Message string `json:"message"`
}
