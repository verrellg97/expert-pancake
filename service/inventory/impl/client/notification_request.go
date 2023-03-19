package client

type InsertNotificationRequest struct {
	CompanyId string `json:"company_id" validate:"required"`
	BranchId  string `json:"branch_id" validate:"required"`
	Title     string `json:"title" validate:"required"`
	Content   string `json:"content" validate:"required"`
	Type      string `json:"type" validate:"required"`
}
