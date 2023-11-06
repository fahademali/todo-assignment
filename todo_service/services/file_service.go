package services

import (
	"fmt"
	"mime/multipart"
	"path/filepath"
	"todo_service/models"
	"todo_service/repo"
	"user_service/log"

	"github.com/gin-gonic/gin"
)

type IFileService interface {
	SaveFiles(ctx *gin.Context, files []*multipart.FileHeader, todoID int64, userID int64) error
	GetFilePath(fileID int64, userID int64) (string, error)
}

type FileService struct {
	fileRepo repo.IFileRepo
}

func NewFileService(fileRepo repo.IFileRepo) IFileService {
	return &FileService{fileRepo: fileRepo}
}

func (fs *FileService) SaveFiles(ctx *gin.Context, files []*multipart.FileHeader, todoID int64, userID int64) error {
	var filesMetaData []models.File

	for _, file := range files {
		filePath := filepath.Join("uploads", file.Filename)
		if err := ctx.SaveUploadedFile(file, filePath); err != nil {
			return err
		}
		filesMetaData = append(filesMetaData, models.File{
			FileName: file.Filename,
			FilePath: filePath,
			TodoID:   todoID,
		})
	}

	log.GetLog().Info(files)
	return fs.fileRepo.InsertForUser(filesMetaData, userID)
}

func (fs *FileService) GetFilePath(fileID int64, userID int64) (string, error) {
	file, err := fs.fileRepo.GetForUser(fileID, userID)
	if err != nil {
		return "", fmt.Errorf("failed to fetch file info using id, Err: %v", err)
	}

	filePath := filepath.Join("uploads", file.FileName)

	return filePath, err
}
