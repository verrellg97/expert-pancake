package model

import (
	"net/http"
	"time"
)

type NotificationService interface {
	InsertNotification(w http.ResponseWriter, r *http.Request) error
	ReadNotification(w http.ResponseWriter, r *http.Request) error
	GetNotifications(w http.ResponseWriter, r *http.Request) error
	DeleteNotification(w http.ResponseWriter, r *http.Request) error
}

type Notification struct {
	NotificationId string    `json:"notification_id" validate:"required"`
	CompanyId      string    `json:"company_id" validate:"required"`
	BranchId       string    `json:"branch_id" validate:"required"`
	Title          string    `json:"title" validate:"required"`
	Content        string    `json:"content" validate:"required"`
	Type           string    `json:"type" validate:"required"`
	CreatedAt      time.Time `json:"created_at" validate:"required"`
}

type InsertNotificationRequest struct {
	CompanyId string `json:"company_id" validate:"required"`
	BranchId  string `json:"branch_id" validate:"required"`
	Title     string `json:"title" validate:"required"`
	Content   string `json:"content" validate:"required"`
	Type      string `json:"type" validate:"required"`
}

type InsertNotificationResponse struct {
	Notification
}

type GetNotificationsRequest struct {
	CompanyId    string `json:"company_id" validate:"required"`
	BranchId     string `json:"branch_id"`
	IsReadFilter *bool  `json:"is_read_filter"`
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
