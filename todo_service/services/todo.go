package services

import (
	"fmt"
	"reflect"
	"time"
	"todo_service/models"
	"todo_service/repo"
)

type ITodoService interface {
	CreateTodo(requestBody models.TodoInput) error
	DeleteTodo(id string) error
	GetTodo(id string) (models.Todo, error)
	UpdateTodo(id string, todoUpdates models.UpdateTodoRequest) error
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

func (ts *TodoService) UpdateTodo(id string, todoUpdates models.UpdateTodoRequest) error {
	var completionDate time.Time

	todo, err := ts.todoRepo.Get(id)
	if err != nil {
		return err
	}

	todoValue := reflect.ValueOf(&todo).Elem()
	todoUpdatesType := reflect.TypeOf(todoUpdates)
	todoUpdatesValue := reflect.ValueOf(todoUpdates)

	if todoUpdates.IsComplete != nil && *todoUpdates.IsComplete && !todo.IsComplete {

		completionDate = time.Now()
	}

	if todo.IsComplete != *todoUpdates.IsComplete {
		//ASK: This is being assigned on heap is there any way to avoid it.
		todo.CompletionDate = &completionDate
	}

	for i := 0; i < todoUpdatesType.NumField(); i++ {
		fieldValue := todoUpdatesValue.Field(i)
		if fieldValue.Type().Kind() == reflect.Ptr && fieldValue.IsNil() {
			continue
		}

		fieldToUpdate := todoUpdatesType.Field(i).Name
		field := todoValue.FieldByName(fieldToUpdate)

		if !field.IsValid() {
			return fmt.Errorf("not a field name: %s", fieldToUpdate)
		}

		if !field.CanSet() {
			return fmt.Errorf("cannot set field %s", fieldToUpdate)
		}
		field.Set(fieldValue.Elem())
	}

	return ts.todoRepo.Update(id, todo)
}
