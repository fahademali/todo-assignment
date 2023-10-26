package routes

import (
	"todo_service/apis/handlers"

	"github.com/gin-gonic/gin"
)

func AddRoutes(tr *gin.Engine, handlers handlers.ITodoHandlers) {
	tr.GET("/ping", handlers.Ping)

	tr.POST("/todos", handlers.HandleCreateTodo)

	tr.GET("/todos/:id", handlers.HandleGetTodo)

	tr.DELETE("/todos/:id", handlers.HandleDeleteTodo)

	tr.PATCH("/todos/:id", handlers.HandleUpdateTodo)
}
