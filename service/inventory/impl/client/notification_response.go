package client

import "time"

type Notification struct {
	NotificationId string    `json:"notification_id" validate:"required"`
	CompanyId      string    `json:"company_id" validate:"required"`
	BranchId       string    `json:"branch_id" validate:"required"`
	Title          string    `json:"title" validate:"required"`
	Content        string    `json:"content" validate:"required"`
	Type           string    `json:"type" validate:"required"`
	CreatedAt      time.Time `json:"created_at" validate:"required"`
}

type InsertNotificationResponse struct {
	Notification
}
