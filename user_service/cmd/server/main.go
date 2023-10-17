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
	var tokenService services.ITokenService
	var cryptService services.ICryptService
	var userHandlers handlers.IUserHandlers

	userRepo = repo.NewUserRepo(db)

	cryptService = services.NewCryptService()
	tokenService = services.NewTokenService()
	userService = services.NewUserService(userRepo, cryptService, tokenService)

	userHandlers = handlers.NewUserHandlers(userService)

	routes.AddUserRoutes(r, userHandlers)

	r.Run()
}
