package services

import (
	"todo_service/models"
	"todo_service/repo"
)

type ITodoService interface {
	CreateTodo(requestBody models.TodoInput) error
	DeleteTodo(id string) error
	GetTodo(id string) (models.Todo, error)
	UpdateTodo(id string, updatedFields models.TodoInput) (models.Todo, error)
}

type TodoService struct {
	todoRepo repo.ITodoRepo
}

func NewTodoService(todoRepo repo.ITodoRepo) ITodoService {
	return &TodoService{todoRepo: todoRepo}
}

func (ts *TodoService) CreateTodo(requestBody models.TodoInput) error {
	err := ts.todoRepo.Insert(requestBody.Title, requestBody.Description, requestBody.DueDate)
	return err
}

func (ts *TodoService) DeleteTodo(id string) error {
	err := ts.todoRepo.Delete(id)
	return err
}

func (ts *TodoService) GetTodo(id string) (models.Todo, error) {
	todo, err := ts.todoRepo.Get(id)
	if err != nil {
		return models.Todo{}, err
	}
	return todo, nil
}

func (ts *TodoService) UpdateTodo(id string, updatedFields models.TodoInput) (models.Todo, error) {
	todo, err := ts.todoRepo.Get(id)
	if err != nil {
		return models.Todo{}, err
	}

	// values := reflect.ValueOf(todo)
	// types := values.Type()
	// for i := 0; i < values.NumField(); i++ {
	// 	todo[types.Field(i).Name]
	// 	fmt.Println(types.Field(i).Index[0], types.Field(i).Name, values.Field(i))
	// }

	return todo, nil
}
