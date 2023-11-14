package models

type User struct {
	ID    int64  `db:"id" binding:"required"`
	Email string `db:"email" binding:"required"`
}

type GetEmailByUserIDPayload struct {
	UserIDs []int64 `json:"userIDs" binding:"required"`
}
