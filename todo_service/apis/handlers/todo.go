package handlers

import (
	"fmt"
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
	HandleGetTodosByDate(ctx *gin.Context)
	HandleGetTodosByDateRange(ctx *gin.Context)
	HandleGetTodosByListID(ctx *gin.Context)
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
	user := ctx.MustGet("user").(models.UserInfo)

	listIDStr := ctx.Param("list_id")
	listID, err := strconv.ParseInt(listIDStr, 10, 64)
	if err != nil {
		customError := fmt.Sprintf("error generated in HandleCreateTodo %v", err)
		ctx.JSON(http.StatusBadRequest, httpResponse.GetErrorResponse(customError))
		return
	}

	if err := ctx.ShouldBindJSON(&requestBody); err != nil {
		ctx.JSON(http.StatusBadRequest, httpResponse.GetErrorResponse(err.Error()))
		return
	}

	if err := th.todoService.CreateTodo(listID, user.ID, requestBody); err != nil {
		ctx.JSON(http.StatusBadRequest, httpResponse.GetErrorResponse(err.Error()))
		return
	}

	ctx.JSON(http.StatusOK, httpResponse.GetSuccessResponse("todo has been created"))
}

func (th *TodoHandlers) HandleDeleteTodo(ctx *gin.Context) {
	user := ctx.MustGet("user").(models.UserInfo)

	todoIDStr := ctx.Param("todo_id")
	todoID, err := strconv.ParseInt(todoIDStr, 10, 64)
	if err != nil {
		customError := fmt.Sprintf("error generated in HandleDeleteTodo %v", err)
		ctx.JSON(http.StatusBadRequest, httpResponse.GetErrorResponse(customError))
		return
	}

	if err := th.todoService.DeleteTodo(todoID, user.ID); err != nil {
		ctx.JSON(http.StatusBadRequest, httpResponse.GetErrorResponse(err.Error()))
		return
	}

	ctx.JSON(http.StatusOK, httpResponse.GetSuccessResponse("Todo has been deleted"))
}

func (th *TodoHandlers) HandleGetTodo(ctx *gin.Context) {
	user := ctx.MustGet("user").(models.UserInfo)

	todoIDStr := ctx.Param("todo_id")
	todoID, err := strconv.ParseInt(todoIDStr, 10, 64)
	if err != nil {
		customError := fmt.Sprintf("error generated in HandleGetTodo %v", err)
		ctx.JSON(http.StatusBadRequest, httpResponse.GetErrorResponse(customError))
		return
	}

	todo, err := th.todoService.GetTodo(todoID, user.ID)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, httpResponse.GetErrorResponse(err.Error()))
		return
	}

	ctx.JSON(http.StatusOK, httpResponse.GetSuccessResponse(todo))
}

func (th *TodoHandlers) HandleGetTodosByDate(ctx *gin.Context) {
	dueDate := ctx.Query("due_date")

	todos, err := th.todoService.GetTodosByDate(dueDate)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, httpResponse.GetErrorResponse(err.Error()))
		return
	}

	ctx.JSON(http.StatusOK, httpResponse.GetSuccessResponse(todos))
}

func (th *TodoHandlers) HandleGetTodosByDateRange(ctx *gin.Context) {
	// dueDate := ctx.Query("due_date")
	startDate := ctx.Query("start_date")
	endDate := ctx.Query("end_date")

	todos, err := th.todoService.GetTodosByDateRange(startDate, endDate)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, httpResponse.GetErrorResponse(err.Error()))
		return
	}

	ctx.JSON(http.StatusOK, httpResponse.GetSuccessResponse(todos))
}

func (th *TodoHandlers) HandleGetTodosByListID(ctx *gin.Context) {
	user := ctx.MustGet("user").(models.UserInfo)
	listIDStr := ctx.Param("list_id")
	limitStr := ctx.DefaultQuery("limit", "10")
	cursorStr := ctx.Query("cursor")

	listID, err := strconv.ParseInt(listIDStr, 10, 64)
	if err != nil {
		customError := fmt.Sprintf("error generated in HandleGetTodosByListID %v", err)
		ctx.JSON(http.StatusBadRequest, httpResponse.GetErrorResponse(customError))
	}

	limit, err := strconv.Atoi(limitStr)
	if err != nil && limitStr != "" {
		ctx.JSON(http.StatusBadRequest, httpResponse.GetErrorResponse("limit parameter is invalid"))
		return
	}
	if limit == 0 {
		limit = 10
	}

	cursor, err := strconv.ParseInt(cursorStr, 10, 64)
	if err != nil && cursorStr != "" {
		ctx.JSON(http.StatusBadRequest, httpResponse.GetErrorResponse("cursor parameter is invalid"))
		return
	}
	if cursor == 0 {
		cursor = models.MAX_ID_DB_VALUE
	}

	todos, err := th.todoService.GetTodosByListID(listID, user.ID, limit, cursor)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, httpResponse.GetErrorResponse(err.Error()))
		return
	}

	ctx.JSON(http.StatusOK, httpResponse.GetSuccessResponse(todos))

}

func (th *TodoHandlers) HandleUpdateTodo(ctx *gin.Context) {
	var requestBody models.UpdateTodoRequest
	user := ctx.MustGet("user").(models.UserInfo)

	todoIDStr := ctx.Param("todo_id")
	todoID, err := strconv.ParseInt(todoIDStr, 10, 64)
	if err != nil {
		customError := fmt.Sprintf("error generated in HandleCreateTodo %v", err)
		ctx.JSON(http.StatusBadRequest, httpResponse.GetErrorResponse(customError))
		return
	}

	if err := ctx.ShouldBindJSON(&requestBody); err != nil {
		ctx.JSON(http.StatusBadRequest, httpResponse.GetErrorResponse(err.Error()))
		return
	}
	log.GetLog().Info(requestBody)

	err = th.todoService.UpdateTodo(todoID, user.ID, requestBody)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, httpResponse.GetErrorResponse(err.Error()))
		return
	}

	ctx.JSON(http.StatusOK, httpResponse.GetSuccessResponse("todo has been updated"))
}
