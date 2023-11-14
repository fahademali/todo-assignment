package repo

import (
	"context"
	"database/sql"
	"fmt"
	"todo_service/models"
)

type ITodoRepo interface {
	InsertForUser(todo models.Todo, userID int64, listID int64) error
	DeleteForUser(id int64, userID int64) error
	DeleteByListIDForUser(listID int64, ctx context.Context, tx *sql.Tx) error
	GetForUser(id int64, userID int64) (models.Todo, error)
	GetByDate(dueDate string) ([]int64, error)
	GetByDateRange(startDate string, endDate string) ([]models.TodoWithUserID, error)
	GetByListIDForUser(listID int64, userID int64, limit int, cursor int64) ([]models.Todo, error)
	UpdateForUser(id int64, userID int64, todoUpdates models.Todo) error
	ExecTx(ctx context.Context) (*sql.Tx, error)
}

type TodoRepo struct {
	db *sql.DB
}

func NewTodoRepo(db *sql.DB) ITodoRepo {
	return &TodoRepo{db: db}
}

func (tr *TodoRepo) InsertForUser(todo models.Todo, userID int64, listID int64) error {
	result, err := tr.db.Exec(`INSERT INTO todos (list_id, title, description, due_date)
	SELECT
		$1 AS list_id,
		$2 AS title,
		$3 AS description,
		$4 AS due_date
	FROM list l
	WHERE
		l.id = $1
		AND l.user_id = $5;
	`, listID, todo.Title, todo.Description, todo.DueDate, userID)

	if err != nil {
		return fmt.Errorf("insert: %v", err)
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return fmt.Errorf("error getting RowsAffected: %v", err)
	}

	if rowsAffected == 0 {
		// No rows were added, which means the list didn't belong to the user
		return fmt.Errorf("record with list_id %d and user_id %d not found", listID, userID)
	}

	return nil
}

func (tr *TodoRepo) DeleteForUser(id int64, userID int64) error {
	result, err := tr.db.Exec(`DELETE FROM todos t
	USING list l
	WHERE
		t.id = $1
		AND t.list_id = l.id
		AND l.user_id = $2`, id, userID)
	if err != nil {
		return fmt.Errorf("delete: %v", err)
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return fmt.Errorf("error getting RowsAffected: %v", err)
	}

	if rowsAffected == 0 {
		// No rows were added, which means the todo didn't belong to the user
		return fmt.Errorf("record with todo_id %d and user_id %d not found", id, userID)
	}
	return nil
}

func (tr *TodoRepo) DeleteByListIDForUser(listID int64, ctx context.Context, tx *sql.Tx) error {
	if _, err := tx.ExecContext(ctx, "DELETE FROM todos where list_id= $1", listID); err != nil {
		return fmt.Errorf("DeleteByListID: %v", err)
	}
	return nil
}

func (tr *TodoRepo) GetForUser(id int64, userID int64) (models.Todo, error) {
	var todo models.Todo
	err := tr.db.QueryRow(`select t.* from todos t 
	INNER JOIN list l ON l.id=t.list_id 
	WHERE t.id=$1 AND user_id = $2;`, id, userID).Scan(&todo.ID, &todo.Title, &todo.Description, &todo.DueDate, &todo.IsComplete, &todo.CompletionDate, &todo.CreationTime, &todo.ListID)

	switch {
	case err == sql.ErrNoRows:
		// ASK: is this right way to make an empty struct to by pass compiler screaming
		return todo, fmt.Errorf("no todo with id %d and user_id %d", id, userID)
	case err != nil:
		return todo, fmt.Errorf("query error: %v", err)
	default:
		return todo, nil
	}
}

func (tr *TodoRepo) GetByDate(dueDate string) ([]int64, error) {
	var userIDCollection []int64
	rows, err := tr.db.Query(`select DISTINCT l.user_id from todos t 
	INNER JOIN list l ON l.id=t.list_id 
	WHERE date_trunc('day',t.due_date) = $1`, dueDate)

	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var userID int64

		err := rows.Scan(&userID)
		if err != nil {
			return nil, err
		}
		userIDCollection = append(userIDCollection, userID)
	}

	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("database iteration error: %v", err)
	}

	if len(userIDCollection) == 0 {
		return nil, fmt.Errorf("no users found with task due %s", dueDate)
	}
	return userIDCollection, nil
}

func (tr *TodoRepo) GetByDateRange(startDate string, endDate string) ([]models.TodoWithUserID, error) {
	var todos []models.TodoWithUserID
	rows, err := tr.db.Query(`select l.user_id, t.title from todos t 
	INNER JOIN list l ON l.id=t.list_id 
	WHERE date_trunc('day',t.due_date) >= $1 AND date_trunc('day',t.due_date) <= $2`, startDate, endDate)

	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var todo models.TodoWithUserID

		err := rows.Scan(&todo.UserID, &todo.Title)
		if err != nil {
			return nil, err
		}
		todos = append(todos, todo)
	}

	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("database iteration error: %v", err)
	}

	if len(todos) == 0 {
		return nil, fmt.Errorf("no todos found with range between %s and %s", startDate, endDate)
	}
	return todos, nil
}

func (tr *TodoRepo) GetByListIDForUser(listID int64, userID int64, limit int, cursor int64) ([]models.Todo, error) {
	var todos []models.Todo

	rows, err := tr.db.Query(`SELECT t.* from todos t 
	INNER JOIN list l ON t.list_id = l.id
	WHERE t.list_id = $1 AND l.user_id = $2 AND t.id < $3 ORDER BY t.id DESC LIMIT $4`, listID, userID, cursor, limit)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var todo models.Todo

		err := rows.Scan(&todo.ID, &todo.Title, &todo.Description, &todo.DueDate, &todo.IsComplete, &todo.CompletionDate, &todo.CreationTime, &todo.ListID)
		if err != nil {
			return nil, err
		}
		todos = append(todos, todo)
	}

	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("database iteration error: %v", err)
	}

	if len(todos) == 0 {
		return nil, fmt.Errorf("no todos with list_id %d and user_id %d", listID, userID)
	}

	return todos, nil
}

func (tr *TodoRepo) UpdateForUser(id int64, userID int64, todoUpdates models.Todo) error {
	result, err := tr.db.Exec(`UPDATE todos t
	SET
		title = $1,
		description = $2,
		due_date = $3,
		is_complete = $4,
		completion_date = $5 
	FROM list l
	WHERE
		t.id = $6
		AND t.list_id = l.id
		AND l.user_id = $7`, todoUpdates.Title, todoUpdates.Description, todoUpdates.DueDate, todoUpdates.IsComplete, todoUpdates.CompletionDate, id, userID)
	if err != nil {
		return fmt.Errorf("update Repo: %v", err)
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return fmt.Errorf("error getting RowsAffected: %v", err)
	}

	if rowsAffected == 0 {
		// No rows were updated, which means the list didn't belong to the user
		return fmt.Errorf("record with todo_id %d and user_id %d not found", id, userID)
	}
	return nil
}

func (ur *TodoRepo) ExecTx(ctx context.Context) (*sql.Tx, error) {
	return ur.db.BeginTx(ctx, nil)
}
