package repo

import (
	"database/sql"
	"fmt"
	"time"
	"todo_service/models"
)

type ITodoRepo interface {
	Insert(title string, description string, dueDate time.Time) error
	Delete(id string) error
	Get(id string) (models.Todo, error)
	Update(id string, updatedTodo models.Todo) error
}

type TodoRepo struct {
	db *sql.DB
}

func NewTodoRepo(db *sql.DB) ITodoRepo {
	return &TodoRepo{db: db}
}

func (tr *TodoRepo) Insert(title string, description string, dueDate time.Time) error {
	if _, err := tr.db.Exec("INSERT INTO todos (title, description, due_date) VALUES ($1, $2, $3)", title, description, dueDate); err != nil {
		return fmt.Errorf("Insert: %v", err)
	}
	return nil
}

func (tr *TodoRepo) Delete(id string) error {
	if _, err := tr.db.Exec("DELETE FROM todos where id= $1", id); err != nil {
		return fmt.Errorf("Delete: %v", err)
	}
	return nil
}

func (tr *TodoRepo) Get(id string) (models.Todo, error) {
	var todo models.Todo
	err := tr.db.QueryRow("SELECT * from todos WHERE id=$1", id).Scan(&todo.ID, &todo.Title, &todo.Description, &todo.DueDate, &todo.IsComplete, &todo.CompletionDate)
	switch {
	case err == sql.ErrNoRows:
		// ASK: is this right way to make an empty struct to by pass compiler screaming
		return models.Todo{}, fmt.Errorf("no user with id %s", id)
	case err != nil:
		return models.Todo{}, fmt.Errorf("query error: %v", err)
	default:
		return todo, nil
	}
}

func (tr *TodoRepo) Update(id string, updatedTodo models.Todo) error {
	_, err := tr.db.Exec("UPDATE todos SET title = $1, description = $2, due_date = $3, is_complete=$4, completion_date=$5 where id = $6", updatedTodo.Title, updatedTodo.Description, updatedTodo.DueDate, updatedTodo.IsComplete, updatedTodo.CompletionDate, updatedTodo.ID)
	if err != nil {
		return fmt.Errorf("Update Repo: %v", err)
	}
	return nil
}
