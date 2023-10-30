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
	var listRepo repo.IListRepo

	var todoService services.ITodoService
	var listService services.IListService

	var todoHandlers handlers.ITodoHandlers
	var listHandlers handlers.IListHandlers

	r := gin.Default()

	db := infra.DbConnection()

	todoRepo = repo.NewTodoRepo(db)
	listRepo = repo.NewListRepo(db)

	todoService = services.NewTodoService(todoRepo)
	listService = services.NewListService(listRepo)

	todoHandlers = handlers.NewTodoHandlers(todoService)
	listHandlers = handlers.NewListHandlers(listService)

	routes.AddRoutes(r, todoHandlers, listHandlers)

	r.Run()
}
