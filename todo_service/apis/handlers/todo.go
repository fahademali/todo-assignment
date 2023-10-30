package handlers

import (
	"net/http"
	"strconv"
	"user_service/log"

	"github.com/gin-gonic/gin"

	httpResponse "todo_service/http"
	"todo_service/models"
	"todo_service/services"
)

type ITodoHandlers interface {
	Ping(ctx *gin.Context)
	HandleCreateTodo(ctx *gin.Context)
	HandleDeleteTodo(ctx *gin.Context)
	HandleGetTodo(ctx *gin.Context)
	HandleUpdateTodo(ctx *gin.Context)
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

func (th *TodoHandlers) HandleCreateTodo(ctx *gin.Context) {
	var requestBody models.TodoInput
	listIDStr := ctx.Param("listID")
	listID, err := strconv.Atoi(listIDStr)

	if err != nil {
		ctx.JSON(http.StatusBadRequest, httpResponse.GetErrorResponse(err.Error()))
		return
	}

	if err := ctx.ShouldBindJSON(&requestBody); err != nil {
		ctx.JSON(http.StatusBadRequest, httpResponse.GetErrorResponse(err.Error()))
		return
	}

	if err := th.todoService.CreateTodo(listID, requestBody); err != nil {
		ctx.JSON(http.StatusBadRequest, httpResponse.GetErrorResponse(err.Error()))
		return
	}

	ctx.JSON(http.StatusOK, httpResponse.GetSuccessResponse("todo has been created"))
}

func (th *TodoHandlers) HandleDeleteTodo(ctx *gin.Context) {
	todoID := ctx.Param("id")
	//TODO: should i keep todoId as a string
	if err := th.todoService.DeleteTodo(todoID); err != nil {
		ctx.JSON(http.StatusBadRequest, httpResponse.GetErrorResponse(err.Error()))
		return
	}

	ctx.JSON(http.StatusOK, httpResponse.GetSuccessResponse("Todo has been deleted"))
}

func (th *TodoHandlers) HandleGetTodo(ctx *gin.Context) {
	todoID := ctx.Param("id")
	//TODO: should i keep todoId as a string
	todo, err := th.todoService.GetTodo(todoID)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, httpResponse.GetErrorResponse(err.Error()))
		return
	}

	ctx.JSON(http.StatusOK, httpResponse.GetSuccessResponse(todo))
}

func (th *TodoHandlers) HandleUpdateTodo(ctx *gin.Context) {
	//TODO: should i keep todoId as a string
	var requestBody models.UpdateTodoRequest

	todoID := ctx.Param("id")

	if err := ctx.ShouldBindJSON(&requestBody); err != nil {
		// errorResponse := httpResponse.GetErrorResponse(err.Error())
		ctx.JSON(http.StatusBadRequest, err.Error())
		return
	}
	log.GetLog().Info(requestBody)

	err := th.todoService.UpdateTodo(todoID, requestBody)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, httpResponse.GetErrorResponse(err.Error()))
		return
	}

	ctx.JSON(http.StatusOK, httpResponse.GetSuccessResponse("todo has been updated"))
}
