package handlers

import (
	"net/http"
	"todo_service/services"

	"github.com/gin-gonic/gin"
)

type ITodoHandlers interface {
	Ping(ctx *gin.Context)
}

type TodoHandlers struct {
	todoService services.ITodoService
}

func NewTodoHandlers(todoService services.ITodoService) ITodoHandlers {
	return &TodoHandlers{todoService: todoService}
}

func (th *TodoHandlers) Ping(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{
		"message": "pong from admin",
	})
}
