package models

type UserTodos struct {
	UserID     int64    `json:"UserID" binding:"required"`
	UserEmail  string   `json:"UserEmail"`
	TodoTitles []string `json:"TodoTitles" binding:"required"`
}
