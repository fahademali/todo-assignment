package main

import (
	"github.com/gin-gonic/gin"

	"user_service/apis/handlers"
	"user_service/apis/routes"
	"user_service/config"
	"user_service/infra"
	"user_service/middlewares"
	"user_service/repo"
	"user_service/services"
)

func main() {
	r := gin.Default()

	db := infra.DbConnection()

	var userRepo repo.IUserRepo
	var emailService services.IEmailService
	var tokenService services.ITokenService
	var cryptService services.ICryptService
	var userService services.IUserService
	var userHandlers handlers.IUserHandlers
	var userMiddleware middlewares.IUserMiddleware

	userRepo = repo.NewUserRepo(db)

	cryptService = services.NewCryptService()
	tokenService = services.NewTokenService(config.AppConfig.SECRET_KEY)
	emailService = services.NewEmailService(config.AppConfig.SENDER_EMAIL, config.AppConfig.SENDER_APP_PASS, config.AppConfig.SMTP_SERVER, config.AppConfig.SMTP_PORT)
	userService = services.NewUserService(userRepo, cryptService, tokenService, emailService)

	userMiddleware = middlewares.NewUserMiddlewares(tokenService)

	userHandlers = handlers.NewUserHandlers(userService, emailService, tokenService)

	routes.AddUserRoutes(r, userHandlers, userMiddleware)

	r.Run()
}
