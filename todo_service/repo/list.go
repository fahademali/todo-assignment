package repo

import (
	"context"
	"database/sql"
	"fmt"
	"todo_service/models"
)

type IListRepo interface {
	Get(id int64) error
	InsertForUser(list models.List, userID int64) error
	DeleteForUser(id int64, userID int64, ctx context.Context, tx *sql.Tx) error
	UpdateForUser(id int64, userID int64, listUpdates models.List) error
	ExecTx(ctx context.Context) (*sql.Tx, error)
}

type ListRepo struct {
	db *sql.DB
}

func NewListRepo(db *sql.DB) IListRepo {
	return &ListRepo{db: db}
}

func (lr *ListRepo) Get(id int64) error {
	if _, err := lr.db.Exec("SELECT * from list where id = $1", id); err != nil {
		return fmt.Errorf("Get List Repo: %v", err)
	}
	return nil
}

func (lr *ListRepo) InsertForUser(list models.List, userID int64) error {
	if _, err := lr.db.Exec("INSERT INTO list (name, user_id) VALUES ($1, $2)", list.Name, userID); err != nil {
		return fmt.Errorf("insert List Repo: %v", err)
	}
	return nil
}

func (lr *ListRepo) DeleteForUser(id int64, userID int64, ctx context.Context, tx *sql.Tx) error {
	result, err := tx.ExecContext(ctx, "DELETE FROM list where id = $1 AND user_id = $2", id, userID)
	if err != nil {
		return fmt.Errorf("delete List Repo: %v", err)
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return fmt.Errorf("error getting RowsAffected: %v", err)
	}

	if rowsAffected == 0 {
		// No rows were deleted, which means the record didn't exist
		return fmt.Errorf("record with id %d and user_id %d not found", id, userID)
	}

	return nil
}

func (lr *ListRepo) UpdateForUser(id int64, userID int64, listUpdates models.List) error {
	result, err := lr.db.Exec("UPDATE list SET name = $1 where id = $2 AND user_id = $3", listUpdates.Name, id, userID)
	if err != nil {
		return fmt.Errorf("update Repo: %v", err)
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return fmt.Errorf("error getting RowsAffected: %v", err)
	}

	if rowsAffected == 0 {
		// No rows were deleted, which means the record didn't exist
		return fmt.Errorf("record with id %d and user_id %d not found", id, userID)
	}

	return nil
}

func (ur *ListRepo) ExecTx(ctx context.Context) (*sql.Tx, error) {
	return ur.db.BeginTx(ctx, nil)
}
