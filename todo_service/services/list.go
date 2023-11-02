package services

import (
	"context"
	"fmt"
	"todo_service/models"
	"todo_service/repo"
)

type IListService interface {
	CreateList(requestBody models.CreateListRequest, userID int64) error
	DeleteList(id int64, ctx context.Context, userID int64) error
	UpdateList(id int64, userID int64, listUpdates models.UpdateListRequest) error
}

type ListService struct {
	listRepo repo.IListRepo
	todoRepo repo.ITodoRepo
}

func NewListService(listRepo repo.IListRepo, todoRepo repo.ITodoRepo) IListService {
	return &ListService{listRepo: listRepo, todoRepo: todoRepo}
}

func (ts *ListService) CreateList(requestBody models.CreateListRequest, userID int64) error {
	var list = models.List{
		Name: requestBody.Name,
	}
	return ts.listRepo.InsertForUser(list, userID)
}

func (ts *ListService) DeleteList(id int64, ctx context.Context, userID int64) error {
	tx, err := ts.listRepo.ExecTx(ctx)
	if err != nil {
		return fmt.Errorf("DeleteListTx: %v", err)
	}
	// Defer a rollback in case anything fails.
	defer tx.Rollback()

	if err = ts.todoRepo.DeleteByListIDForUser(id, ctx, tx); err != nil {
		return err
	}

	if err = ts.listRepo.DeleteForUser(id, userID, ctx, tx); err != nil {
		return err
	}

	return tx.Commit()
}

func (ts *ListService) UpdateList(id int64, userID int64, listUpdates models.UpdateListRequest) error {
	var list = models.List{
		Name: listUpdates.Name,
	}
	return ts.listRepo.UpdateForUser(id, userID, list)
}
