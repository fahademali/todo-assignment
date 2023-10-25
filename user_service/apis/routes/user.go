package routes

import (
	"net/http"
	"user_service/middlewares"

	"github.com/gin-gonic/gin"

	"user_service/apis/handlers"
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

	ur.PATCH("/grant-admin-role", middleware.Authorize, handlers.HandleGrantAdminRole)
}
