package models

import "time"

type Todo struct {
	ID             int        `db:"id"`
	Title          string     `db:"title"`
	Description    string     `db:"description"`
	DueDate        time.Time  `db:"due_date"`
	IsComplete     bool       `db:"is_complete"`
	CompletionDate *time.Time `db:"completion_date"`
}

type TodoInput struct {
	Title       string    `json:"title" binding:"required"`
	Description string    `json:"description" binding:"required"`
	DueDate     time.Time `json:"dueDate"  binding:"required"`
	IsComplete  bool      `json:"isComplete"`
}

type UpdateTodoRequest struct {
	Title       *string    `json:"title,omitempty"`
	Description *string    `json:"description,omitempty"`
	DueDate     *time.Time `json:"dueDate,omitempty"`
	IsComplete  *bool      `json:"isComplete,omitempty"`
}

type UpdateListRequest struct {
	Name string `json:"name" binding:"required"`
}

type CreateListRequest struct {
	Name string `json:"name" binding:"required"`
}
