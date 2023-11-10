package models

type ProfileRequest struct {
	Token string `json:"token" binding:"required"`
}
type LoginRequest struct {
	Email    string `json:"email" binding:"required"`
	Password string `json:"password" binding:"required"`
}

type SignupRequest struct {
	UserName string `json:"username" binding:"required"`
	Email    string `json:"email" binding:"required"`
	Password string `json:"password" binding:"required"`
}

type SendEmailsRequest struct {
	UserEmailAddresses []string `json:"userEmailAddresses" binding:"required"`
	Subject            string   `json:"subject" binding:"required"`
	Body               string   `json:"body" binding:"required"`
}

type GetUserEmailsByIdsRequest struct {
	UserIDs []int64 `json:"user_ids" binding:"required"`
}

type AdminPromotionRequest struct {
	Email string `json:"email" binding:"required"`
}

type User struct {
	ID         int    `db:"id"`
	Username   string `db:"username"`
	Email      string `db:"email"`
	Password   string `db:"password"`
	Role       string `db:"role"`
	IsVerified bool   `db:"is_verified"`
}

type UserInfo struct {
	Email      string  `json:"email"`
	Role       string  `json:"role"`
	IsVerified bool    `json:"isVerified"`
	ValidFrom  float64 `json:"validFrom"`
}
