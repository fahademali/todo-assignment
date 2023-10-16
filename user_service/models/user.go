package models

type ProfileRequestBody struct {
	Token string `json:"token" binding:"required"`
}
type LoginRequestBody struct {
	Email    string `json:"email" binding:"required"`
	Password string `json:"password" binding:"required"`
}

type SignupRequestBody struct {
	UserName string `json:"user_name" binding:"required"`
	Email    string `json:"email" binding:"required"`
	Password string `json:"password" binding:"required"`
}

type User struct {
	Id       string
	Email    string
	Username string
	Password string
}
