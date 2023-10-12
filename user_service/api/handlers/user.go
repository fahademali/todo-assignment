package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type loginRequestBody struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

func HandleGetProfile(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{
		"message": "/getprofile Not Implemented yet",
	})
}

func HandleLogin(ctx *gin.Context) {
	var requestBody loginRequestBody

	if err := ctx.ShouldBindJSON(&requestBody); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"message": "/login Not Implemented yet",
		"body":    ctx.Request.Body,
	})
}

func HandleSignup(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{
		"message": "/signup Not Implemented yet",
	})
}

func HandleRefreshToken(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{
		"message": "/refresh-token Not Implemented yet",
	})
}

func HandleForgetPassword(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{
		"message": "Oh not again! Weak Memory",
	})
}
