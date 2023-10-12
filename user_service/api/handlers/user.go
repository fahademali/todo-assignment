package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type profileRequestBody struct {
	Token string `json:"token" binding:"required"`
}
type loginRequestBody struct {
	Email    string `json:"email" binding:"required"`
	Password string `json:"password" binding:"required"`
}

type signupRequestBody struct {
	UserName string `json:"user_name" binding:"required"`
	Email    string `json:"email" binding:"required"`
	Password string `json:"password" binding:"required"`
}

func HandleGetProfile(ctx *gin.Context) {
	var requestBody profileRequestBody
	// TODO: what does this shouldbingjson do when produce error
	if err := ctx.ShouldBindJSON(&requestBody); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{
		"message": "/getprofile Not Implemented yet",
	})
}

func HandleLogin(ctx *gin.Context) {
	var requestBody loginRequestBody
	// TODO: what does this shouldbingjson do when produce error
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
	var requestBody signupRequestBody
	// TODO: what does this shouldbingjson do when produce error
	if err := ctx.ShouldBindJSON(&requestBody); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"message": "/signup Not Implemented yet",
	})
}

func HandleRefreshToken(ctx *gin.Context) {
	var requestBody loginRequestBody
	// TODO: what does this shouldbingjson do when produce error
	if err := ctx.ShouldBindJSON(&requestBody); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"message": "/refresh-token Not Implemented yet",
	})
}

func HandleForgetPassword(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{
		"message": "Oh not again! Weak Memory",
	})
}
