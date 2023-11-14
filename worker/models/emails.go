package models

type RecepietDetails struct {
	UserEmail string `json:"userEmail" binding:"required"`
	Subject   string `json:"subject" binding:"required"`
	Body      string `json:"body" binding:"required"`
}
type SendEmailPayload struct {
	RecepietDetails []RecepietDetails `json:"recepientDetails" binding:"required"`
}
