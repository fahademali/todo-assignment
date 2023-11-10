package services

import (
	"time"
	"todo_service/models"
	"todo_service/repo"
)

type ITodoService interface {
	CreateTodo(listID int64, userID int64, todoInput models.TodoInput) error
	DeleteTodo(id int64, userID int64) error
	GetTodo(id int64, userID int64) (models.Todo, error)
	GetTodosByDate(dueDate string) ([]int64, error)
	GetTodosByListID(listID int64, userID int64, limit int, cursor int64) ([]models.Todo, error)
	UpdateTodo(id int64, userID int64, todoUpdates models.UpdateTodoRequest) error
}

type TodoService struct {
	todoRepo    repo.ITodoRepo
	fileService IFileService
}

func NewTodoService(todoRepo repo.ITodoRepo, fileService IFileService) ITodoService {
	return &TodoService{todoRepo: todoRepo, fileService: fileService}
}

func (ts *TodoService) CreateTodo(listID int64, userID int64, todoInput models.TodoInput) error {
	var todo = models.Todo{
		Title:       todoInput.Title,
		Description: todoInput.Description,
		DueDate:     todoInput.DueDate,
		IsComplete:  *todoInput.IsComplete,
	}
	return ts.todoRepo.InsertForUser(todo, userID, listID)
}

func (ts *TodoService) DeleteTodo(id int64, userID int64) error {
	return ts.todoRepo.DeleteForUser(id, userID)
}

func (ts *TodoService) GetTodo(id int64, userID int64) (models.Todo, error) {
	return ts.todoRepo.GetForUser(id, userID)
}

func (ts *TodoService) GetTodosByDate(dueDate string) ([]int64, error) {
	return ts.todoRepo.GetByDate(dueDate)
}

func (ts *TodoService) GetTodosByListID(listID int64, userID int64, limit int, cursor int64) ([]models.Todo, error) {
	return ts.todoRepo.GetByListIDForUser(listID, userID, limit, cursor)
}

func (ts *TodoService) UpdateTodo(id int64, userID int64, todoUpdates models.UpdateTodoRequest) error {
	var completionDate time.Time

	todo, err := ts.todoRepo.GetForUser(id, userID)
	if err != nil {
		return err
	}

	if *todoUpdates.IsComplete && !todo.IsComplete {
		completionDate = time.Now()
	}

	if todo.IsComplete != *todoUpdates.IsComplete {
		//ASK: This is being assigned on heap is there any way to avoid it.
		todoUpdates.CompletionDate = completionDate
	}
	if todoUpdates.Title != nil {
		todo.Title = *todoUpdates.Title
	}
	if todoUpdates.IsComplete != nil {
		todo.IsComplete = *todoUpdates.IsComplete
	}
	if todoUpdates.DueDate != nil {
		todo.DueDate = *todoUpdates.DueDate
	}
	if todoUpdates.Description != nil {
		todo.Description = *todoUpdates.Description
	}

	return ts.todoRepo.UpdateForUser(id, userID, todo)
}
