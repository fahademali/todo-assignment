package routes

import (
	"todo_service/apis/handlers"

	"github.com/gin-gonic/gin"
)

func AddRoutes(tr *gin.Engine, handlers handlers.ITodoHandlers) {
	tr.GET("/ping", handlers.Ping)
	//TODO: check naming convention for path param
	tr.POST("/lists", handlers.HandleCreateList)

	tr.GET("/todos/:id", handlers.HandleGetTodo)

	tr.POST("/todos", handlers.HandleCreateTodo)

	tr.DELETE("/todos/:id", handlers.HandleDeleteTodo)

	tr.PATCH("/todos/:id", handlers.HandleUpdateTodo)
}
