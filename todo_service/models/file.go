package models

import "time"

type File struct {
	ID         int64     `db:"id"`
	FileName   string    `db:"file_name"`
	FilePath   string    `db:"file_path"`
	UploadTime time.Time `db:"upload_time"`
	TodoID     int64     `db:"todo_id"`
}
