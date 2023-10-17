package models

type ProfileRequestBody struct {
	Token string `json:"token" binding:"required"`
}
type LoginRequestBody struct {
	Email    string `json:"email" binding:"required"`
	Password string `json:"password" binding:"required"`
}

type SignupRequestBody struct {
	UserName string `json:"username" binding:"required"`
	Email    string `json:"email" binding:"required"`
	Password string `json:"password" binding:"required"`
}

type User struct {
	Id         int    `db:"id"`
	Email      string `db:"email"`
	Username   string `db:"username"`
	Password   string `db:"password"`
	Role       string `db:"role"`
	IsVerified bool   `db:"is_verified"`
}
