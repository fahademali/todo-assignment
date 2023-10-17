package models

type ProfileRequestBody struct {
	Token string `json:"token" binding:"required"`
}
type LoginRequestBody struct {
	Email    string `json:"email" binding:"required"`
	Password string `json:"password" binding:"required"`
}

type SignupRequestBody struct {
	UserName   string `json:"username" binding:"required"`
	Email      string `json:"email" binding:"required"`
	Password   string `json:"password" binding:"required"`
	Role       string `json:"role" binding:"required"`
	IsVerified bool   `json:"is_verified"`
}

type User struct {
	Id         int    `db:"id"`
	Username   string `db:"username"`
	Email      string `db:"email"`
	Password   string `db:"password"`
	Role       string `db:"role"`
	IsVerified bool   `db:"is_verified"`
}
