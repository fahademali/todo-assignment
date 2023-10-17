package routes

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"user_service/apis/handlers"
)

func AddUserRoutes(ur *gin.Engine, handlers handlers.IUserHandlers) {
	ur.GET("/ping", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})

	ur.GET("/profile", handlers.HandleGetProfile)

	ur.POST("/login", handlers.HandleLogin)

	ur.POST("/signup", handlers.HandleSignup)

	ur.POST("/refresh-token", handlers.HandleRefreshToken)

	ur.POST("/forget-password", handlers.HandleForgetPassword)
}
