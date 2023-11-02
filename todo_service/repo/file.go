package repo

import (
	"database/sql"
	"fmt"
	"todo_service/log"
	"todo_service/models"
)

type IFileRepo interface {
	InsertForUser(files []models.File, userID int64) error
	GetForUser(id int64, userID int64) (models.File, error)
}
type FileRepo struct {
	db *sql.DB
}

func NewFileRepo(db *sql.DB) IFileRepo {
	return &FileRepo{db: db}
}

func (fr *FileRepo) InsertForUser(files []models.File, userID int64) error {
	//TODO: optimize the query
	for _, file := range files {
		log.GetLog().Warn(file.TodoID)
		result, err := fr.db.Exec(`INSERT INTO file (file_name, file_path, todo_id)
		SELECT
			$1 AS file_name,
			$2 AS file_path,
			$3 AS todo_id
		FROM todos t
		JOIN list l ON t.list_id = l.id
		WHERE
			t.id = $3
			AND l.user_id = $4
		`, file.FileName, file.FilePath, file.TodoID, userID)
		if err != nil {
			return err
		}

		rowsAffected, err := result.RowsAffected()
		if err != nil {
			return fmt.Errorf("error getting RowsAffected: %v", err)
		}

		if rowsAffected == 0 {
			// No rows were deleted, which means the record didn't exist
			return fmt.Errorf("record with todo_id %d and user_id %d not found", file.TodoID, userID)
		}
	}
	return nil
}

func (fr *FileRepo) GetForUser(id int64, userID int64) (models.File, error) {
	var file models.File
	err := fr.db.QueryRow(`SELECT f.* FROM file f 
	INNER JOIN todos t ON t.id = f.todo_id
	INNER JOIN list l ON l.id = t.list_id
	WHERE f.id = $1 AND l.user_id = $2`, id, userID).Scan(&file.ID, &file.FileName, &file.FilePath, &file.UploadTime, &file.TodoID)

	switch {
	case err == sql.ErrNoRows:
		return file, fmt.Errorf("no file with id %d and user_id %d", id, userID)
	case err != nil:
		return file, fmt.Errorf("query error: %v", err)
	default:
		return file, nil
	}
}
