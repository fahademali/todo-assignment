package main

import (
	"github.com/gin-gonic/gin"

	"user_service/apis/routes"
	"user_service/infra"
)

func main() {
	r := gin.Default()

	routes.AddUserRoutes(r)
	infra.DbConnection()

	// r.GET("/ping", func(ctx *gin.Context) {
	// 	ctx.JSON(200, gin.H{
	// 		"message": "pong",
	// 	})
	// })
	r.Run()
}
