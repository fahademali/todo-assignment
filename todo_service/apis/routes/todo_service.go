package routes

import (
	"todo_service/apis/handlers"

	"github.com/gin-gonic/gin"
)

func AddRoutes(tr *gin.Engine, todoHandlers handlers.ITodoHandlers, listHandlers handlers.IListHandlers) {
	//TODO: check naming convention for path param
	tr.POST("/lists", listHandlers.HandleCreateList)

	tr.DELETE("/lists/:listID", listHandlers.HandleDeleteList)

	tr.PATCH("/lists/:listID", listHandlers.HandleUpdateListName)

	tr.GET("/ping", todoHandlers.Ping)

	tr.GET("/todos/:id", todoHandlers.HandleGetTodo)

	tr.POST("/todos", todoHandlers.HandleCreateTodo)

	tr.DELETE("/todos/:id", todoHandlers.HandleDeleteTodo)

	tr.PATCH("/todos/:id", todoHandlers.HandleUpdateTodo)
}
