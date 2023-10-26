package repo

import (
	"database/sql"
)

type ITodoRepo interface {
}

type TodoRepo struct {
	db *sql.DB
}

func NewTodoRepo(db *sql.DB) ITodoRepo {
	return &TodoRepo{db: db}
}
