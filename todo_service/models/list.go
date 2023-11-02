package models

type List struct {
	ID     int    `db:"id"`
	Name   string `db:"name"`
	UserID int    `db:"user_id"`
}

type UpdateListRequest struct {
	Name string `json:"name" binding:"required"`
}

type CreateListRequest struct {
	Name string `json:"name" binding:"required"`
}
