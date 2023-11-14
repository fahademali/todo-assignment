package models

type TodosDueTodayResponse struct {
	Status  string `json:"status" binding:"required"`
	Message string `json:"message" binding:"required"`

	Data []UserTodos `json:"data" binding:"required"`
}
type GetEmailsByIDResponse struct {
	Status  string `json:"status" binding:"required"`
	Message string `json:"message" binding:"required"`

	//ASK: should i have returned data as type map key as user_id and value as email would have saved me a loop
	Data []User `json:"data" binding:"required"`
}
type SendEmailResponse struct {
	Status  string `json:"status" binding:"required"`
	Message string `json:"message" binding:"required"`

	Data string `json:"data" binding:"required"`
}
type ErrorResponse struct {
	Status  string `json:"status" binding:"required"`
	Message string `json:"message" binding:"required"`

	Error string `json:"error" binding:"required"`
}
