package services

import (
	"todo_service/repo"
)

type ITodoService interface {
}

type TodoService struct {
	todoRepo repo.ITodoRepo
}

func NewTodoService(todoRepo repo.ITodoRepo) ITodoService {
	return &TodoService{todoRepo: todoRepo}
}
