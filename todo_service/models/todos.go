package models

import "time"

type Todo struct {
	ID             int        `db:"id"`
	Title          string     `db:"title"`
	Description    string     `db:"description"`
	DueDate        time.Time  `db:"due_date"`
	IsComplete     bool       `db:"is_complete"`
	CompletionDate *time.Time `db:"completion_date"`
	CreationTime   *time.Time `db:"creation_time"`
	ListID         int        `db:"list_id"`
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
