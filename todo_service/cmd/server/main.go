package main

import (
	"github.com/gin-gonic/gin"

	"todo_service/apis/handlers"
	"todo_service/apis/routes"
	"todo_service/infra"
	"todo_service/repo"
	"todo_service/services"
)

func main() {
	var todoRepo repo.ITodoRepo
	var todoService services.ITodoService
	var todoHandlers handlers.ITodoHandlers

	r := gin.Default()

	db := infra.DbConnection()

	todoRepo = repo.NewTodoRepo(db)
	todoService = services.NewTodoService(todoRepo)
	todoHandlers = handlers.NewTodoHandlers(todoService)

	routes.AddRoutes(r, todoHandlers)

	r.Run()
}
