package models

type TodosDueTodaySuccessResponse struct {
	Status  string `json:"status" binding:"required"`
	Message string `json:"message" binding:"required"`

	Data []int64 `json:"data" binding:"required"`
}
type GetEmailsByIDSuccessResponse struct {
	Status  string `json:"status" binding:"required"`
	Message string `json:"message" binding:"required"`

	Data []string `json:"data" binding:"required"`
}
type SendEmailSuccessResponse struct {
	Status  string `json:"status" binding:"required"`
	Message string `json:"message" binding:"required"`

	Data string `json:"data" binding:"required"`
}
type ErrorResponse struct {
	Status  string `json:"status" binding:"required"`
	Message string `json:"message" binding:"required"`

	Error string `json:"error" binding:"required"`
}

type SendEmailPayload struct {
	UserEmailAddresses []string `json:"userEmailAddresses" binding:"required"`
	Subject            string   `json:"subject" binding:"required"`
	Body               string   `json:"body" binding:"required"`
}
