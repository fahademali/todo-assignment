package models

type UserTodos struct {
	UserID     int64    `json:"UserID" binding:"required"`
	UserEmail  string   `json:"UserEmail"`
	TodoTitles []string `json:"TodoTitles" binding:"required"`
}

type User struct {
	ID    int64  `db:"id" binding:"required"`
	Email string `db:"email" binding:"required"`
}

type TodosDueTodaySuccessResponse struct {
	Status  string `json:"status" binding:"required"`
	Message string `json:"message" binding:"required"`

	Data []UserTodos `json:"data" binding:"required"`
}
type GetEmailsByIDSuccessResponse struct {
	Status  string `json:"status" binding:"required"`
	Message string `json:"message" binding:"required"`

	//ASK: should i have returned data as type map key as user_id and value as email would have saved me a loop
	Data []User `json:"data" binding:"required"`
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

type RecepietDetails struct {
	UserEmail string `json:"userEmail" binding:"required"`
	Subject   string `json:"subject" binding:"required"`
	Body      string `json:"body" binding:"required"`
}
type SendEmailPayload struct {
	RecepietDetails []RecepietDetails `json:"recepientDetails" binding:"required"`
}
