package models

import "time"

type Todo struct {
	ID             int        `db:"id"`
	Title          string     `db:"title"`
	Description    string     `db:"description"`
	DueDate        time.Time  `db:"due_date"`
	IsComplete     string     `db:"is_complete"`
	CompletionDate *time.Time `db:"completion_date"`
}

type TodoInput struct {
	Title       string    `json:"title" binding:"required"`
	Description string    `json:"description" binding:"required"`
	DueDate     time.Time `json:"dueDate" binding:"required"`
	IsComplete  string    `json:"isComplete"`
}
