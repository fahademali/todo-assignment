package main

import (
	"github.com/gin-gonic/gin"

	"user_service/apis/handlers"
	"user_service/apis/routes"
	"user_service/infra"
	"user_service/repo"
	"user_service/services"
)

func main() {
	r := gin.Default()

	db := infra.DbConnection()

	var userRepo repo.IUserRepo
	var userService services.IUserService
	var userHandlers handlers.IUserHandlers

	userRepo = repo.NewUserRepo(db)
	userService = services.NewUserService(userRepo)
	userHandlers = handlers.NewUserHandlers(userService)

	routes.AddUserRoutes(r, userHandlers)

	r.Run()
}
