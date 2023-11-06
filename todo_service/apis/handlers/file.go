package handlers

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"

	httpResponse "todo_service/http"
	"todo_service/models"
	"todo_service/services"
)

type IFileHandlers interface {
	HandleDownloadFile(ctx *gin.Context)
	HandleUploadFile(ctx *gin.Context)
}

type FileHandlers struct {
	fileService services.IFileService
}

func NewFileHandlers(fileService services.IFileService) IFileHandlers {
	return &FileHandlers{fileService: fileService}
}

func (fh *FileHandlers) HandleDownloadFile(ctx *gin.Context) {
	user := ctx.MustGet("user").(models.UserInfo)
	fileIDStr := ctx.Param("file_id")
	fileID, err := strconv.ParseInt(fileIDStr, 10, 64)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, httpResponse.GetErrorResponse(err.Error()))
		return
	}

	filePath, err := fh.fileService.GetFilePath(fileID, user.ID)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, httpResponse.GetErrorResponse(err.Error()))
	}

	ctx.File(filePath)
}

func (fh *FileHandlers) HandleUploadFile(ctx *gin.Context) {
	user := ctx.MustGet("user").(models.UserInfo)
	todoIDStr := ctx.Param("todo_id")
	todoID, err := strconv.ParseInt(todoIDStr, 10, 64)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, httpResponse.GetErrorResponse(err.Error()))
		return
	}

	form, err := ctx.MultipartForm()
	if err != nil {
		ctx.JSON(http.StatusBadRequest, httpResponse.GetErrorResponse(err.Error()))
		return
	}
	//TODO: ensure the request have file if not throw error and move out above some code to handlers layer
	files := form.File["files"]

	err = fh.fileService.SaveFiles(ctx, files, todoID, user.ID)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, httpResponse.GetErrorResponse(err.Error()))
		return
	}

	ctx.JSON(http.StatusOK, httpResponse.GetSuccessResponse("file has been uploaded to todo"))
}
