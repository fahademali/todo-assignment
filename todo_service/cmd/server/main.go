package main

import (
	"github.com/gin-gonic/gin"

	"todo_service/apis/handlers"
	"todo_service/apis/routes"
	"todo_service/config"
	"todo_service/infra"
	"todo_service/middlewares"
	"todo_service/repo"
	"todo_service/services"
)

func main() {
	var todoRepo repo.ITodoRepo
	var listRepo repo.IListRepo
	var fileRepo repo.IFileRepo

	var todoService services.ITodoService
	var listService services.IListService
	var fileService services.IFileService
	var tokenService services.ITokenService

	var todoHandlers handlers.ITodoHandlers
	var listHandlers handlers.IListHandlers
	var fileHandlers handlers.IFileHandlers

	var todoMiddlewares middlewares.ITodoMiddleware

	r := gin.Default()

	db := infra.DbConnection()

	todoRepo = repo.NewTodoRepo(db)
	listRepo = repo.NewListRepo(db)
	fileRepo = repo.NewFileRepo(db)

	fileService = services.NewFileService(fileRepo)
	todoService = services.NewTodoService(todoRepo, fileService)
	listService = services.NewListService(listRepo, todoRepo)
	tokenService = services.NewTokenService(config.AppConfig.SECRET_KEY)

	todoMiddlewares = middlewares.NewTodoMiddlewares(tokenService)

	todoHandlers = handlers.NewTodoHandlers(todoService)
	listHandlers = handlers.NewListHandlers(listService)
	fileHandlers = handlers.NewFileHandlers(fileService)

	routes.AddRoutes(r, todoMiddlewares, todoHandlers, listHandlers, fileHandlers)

	r.Run()
}
