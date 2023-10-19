package main

import (
	"github.com/gin-gonic/gin"

	"user_service/apis/handlers"
	"user_service/apis/routes"
	"user_service/config"
	"user_service/infra"
	"user_service/repo"
	"user_service/services"
)

func main() {
	r := gin.Default()

	config.Init()

	db := infra.DbConnection()

	var userRepo repo.IUserRepo
	var emailService services.IEmailService
	var tokenService services.ITokenService
	var cryptService services.ICryptService
	var userService services.IUserService
	var userHandlers handlers.IUserHandlers

	userRepo = repo.NewUserRepo(db)

	cryptService = services.NewCryptService()
	tokenService = services.NewTokenService()
	emailService = services.NewEmailService()
	userService = services.NewUserService(userRepo, cryptService, tokenService)

	userHandlers = handlers.NewUserHandlers(userService, emailService)

	routes.AddUserRoutes(r, userHandlers)

	r.Run()
}
