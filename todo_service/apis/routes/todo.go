package routes

import (
	"todo_service/apis/handlers"

	"github.com/gin-gonic/gin"
)

func AddRoutes(ur *gin.Engine, handlers handlers.ITodoHandlers) {
	ur.GET("/ping", handlers.Ping)
}
