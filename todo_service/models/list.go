package models

type List struct{}

type UpdateListRequest struct {
	Name string `json:"name" binding:"required"`
}

type CreateListRequest struct {
	Name string `json:"name" binding:"required"`
}
