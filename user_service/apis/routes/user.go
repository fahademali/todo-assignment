package routes

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"user_service/apis/handlers"
	"user_service/middlewares"
)

func AddUserRoutes(ur *gin.Engine, handlers handlers.IUserHandlers, middleware middlewares.IUserMiddleware) {
	ur.GET("/ping", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})

	ur.GET("/profile", handlers.HandleGetProfile)

	ur.GET("/verify-user/:token", handlers.HandleVerifyUser)

	ur.POST("/login", handlers.HandleLogin)

	ur.POST("/signup", handlers.HandleSignup)

	ur.POST("/refresh-token", handlers.HandleRefreshToken)

	ur.POST("/forget-password", handlers.HandleForgetPassword)

	ur.POST("/internal/send-email", handlers.HandleSendEmails)

	ur.POST("/internal/users", handlers.HandleGetUserEmailsByIDs)

	ur.PATCH("/grant-admin-role", middleware.Authorize, handlers.HandleGrantAdminRole)
}
