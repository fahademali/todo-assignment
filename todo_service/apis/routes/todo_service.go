package routes

import (
	"todo_service/apis/handlers"
	"todo_service/middlewares"

	"github.com/gin-gonic/gin"
)

func AddRoutes(tr *gin.Engine, middleware middlewares.ITodoMiddleware, todoHandlers handlers.ITodoHandlers, listHandlers handlers.IListHandlers, fileHandlers handlers.IFileHandlers) {
	tr.GET("/ping", todoHandlers.Ping)

	tr.GET("/lists/:list_id/todos", middleware.Authenticate, todoHandlers.HandleGetTodosByListID)

	tr.GET("/lists/:list_id/todos/:todo_id", middleware.Authenticate, todoHandlers.HandleGetTodo)

	tr.GET("/lists/:list_id/todos/:todo_id/files/:file_id", middleware.Authenticate, fileHandlers.HandleDownloadFile)

	tr.POST("/lists/:list_id/todos", middleware.Authenticate, todoHandlers.HandleCreateTodo)

	tr.POST("/lists/:list_id/todos/:todo_id/files", middleware.Authenticate, fileHandlers.HandleUploadFile)

	tr.POST("/lists", middleware.Authenticate, listHandlers.HandleCreateList)

	tr.PATCH("/lists/:list_id", middleware.Authenticate, listHandlers.HandleUpdateList)

	tr.PATCH("/lists/:list_id/todos/:todo_id", middleware.Authenticate, todoHandlers.HandleUpdateTodo)

	tr.DELETE("/lists/:list_id", middleware.Authenticate, listHandlers.HandleDeleteList)

	tr.DELETE("/lists/:list_id/todos/:todo_id", middleware.Authenticate, todoHandlers.HandleDeleteTodo)
}
