package main

import (
	"context"
	"worker/utils"

	"github.com/gin-gonic/gin"
	"go.uber.org/cadence/.gen/go/cadence/workflowserviceclient"
	"go.uber.org/cadence/client"
	"go.uber.org/yarpc"
	"go.uber.org/yarpc/transport/tchannel"

	"user_service/apis/handlers"
	"user_service/apis/routes"
	"user_service/config"
	"user_service/infra"
	"user_service/middlewares"
	"user_service/repo"
	"user_service/services"
)

// var HostPort = "127.0.0.1:7933"
// var Domain = "test-domain"
// var TaskListName = "test-list"

// var ClientName = "test-client"
// var CadenceService = "cadence-frontend"

func main() {
	var userRepo repo.IUserRepo
	var emailService services.IEmailService
	var tokenService services.ITokenService
	var cryptService services.ICryptService
	var userService services.IUserService
	var userHandlers handlers.IUserHandlers
	var userMiddleware middlewares.IUserMiddleware
	var dailyMidnightUTC = "0 0 * * *"

	ctx := context.Background()
	cadenceClient := client.NewClient(buildCadenceClient(), config.AppConfig.CADENCE_DOMAIN, &client.Options{})
	startWorkflowOptions := utils.GetDefaultStartWorkflowOptions(dailyMidnightUTC)

	cadenceClient.StartWorkflow(ctx, startWorkflowOptions, "worker/workflows.RemindUsersForDueDateWorkflow")

	r := gin.Default()

	db := infra.DbConnection()

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

func buildCadenceClient() workflowserviceclient.Interface {
	ch, err := tchannel.NewChannelTransport(tchannel.ServiceName(config.AppConfig.CADENCE_CLIENT_NAME))
	if err != nil {
		panic("Failed to setup tchannel")
	}
	dispatcher := yarpc.NewDispatcher(yarpc.Config{
		Name: config.AppConfig.CADENCE_CLIENT_NAME,
		Outbounds: yarpc.Outbounds{
			config.AppConfig.CADENCE_SERVICE: {Unary: ch.NewSingleOutbound(config.AppConfig.CADENCE_HOST_PORT)},
		},
	})
	if err := dispatcher.Start(); err != nil {
		panic("Failed to start dispatcher")
	}

	return workflowserviceclient.New(dispatcher.ClientConfig(config.AppConfig.CADENCE_SERVICE))
}
