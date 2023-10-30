package services

import (
	"todo_service/models"
	"todo_service/repo"
)

type IListService interface {
	CreateList(requestBody models.CreateListRequest) error
	DeleteList(id string) error
	UpdateList(id string, name string) error
}

type ListService struct {
	listRepo repo.IListRepo
}

func NewListService(listRepo repo.IListRepo) IListService {
	return &ListService{listRepo: listRepo}
}

func (ts *ListService) CreateList(requestBody models.CreateListRequest) error {
	err := ts.listRepo.Insert(requestBody.Name)
	return err
}

func (ts *ListService) DeleteList(id string) error {
	err := ts.listRepo.Delete(id)
	return err
}

func (ts *ListService) UpdateList(id string, name string) error {
	err := ts.listRepo.Update(id, name)
	return err
}
