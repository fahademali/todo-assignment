package repo

import (
	"database/sql"
	"fmt"
)

type IListRepo interface {
	Get(id string) error
	Insert(name string) error
	Delete(id string) error
	Update(id string, name string) error
}

type ListRepo struct {
	db *sql.DB
}

func NewListRepo(db *sql.DB) IListRepo {
	return &ListRepo{db: db}
}

func (lr *ListRepo) Get(id string) error {
	if _, err := lr.db.Exec("SELECT * from list where id = $1", id); err != nil {
		return fmt.Errorf("Get List Repo: %v", err)
	}
	return nil
}

func (lr *ListRepo) Insert(name string) error {
	if _, err := lr.db.Exec("INSERT INTO list (name) VALUES ($1)", name); err != nil {
		return fmt.Errorf("Insert List Repo: %v", err)
	}
	return nil
}

func (lr *ListRepo) Delete(id string) error {
	if _, err := lr.db.Exec("DELETE FROM list where id = $1", id); err != nil {
		return fmt.Errorf("Delete List Repo: %v", err)
	}
	return nil
}

func (lr *ListRepo) Update(id string, name string) error {
	_, err := lr.db.Exec("UPDATE list SET name = $1 where id = $2", name, id)
	if err != nil {
		return fmt.Errorf("Update Repo: %v", err)
	}
	return nil
}
