package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"user_service/models"
	"user_service/services"
)

type IUserHandlers interface {
	HandleGetProfile(ctx *gin.Context)
	HandleLogin(ctx *gin.Context)
	HandleSignup(ctx *gin.Context)
	HandleRefreshToken(ctx *gin.Context)
	HandleForgetPassword(ctx *gin.Context)
}

type UserHandlers struct {
	userService services.IUserService
}

func NewUserHandlers(userService services.IUserService) IUserHandlers {
	return &UserHandlers{userService: userService}
}

func (uh *UserHandlers) HandleGetProfile(ctx *gin.Context) {
	var requestBody models.ProfileRequestBody

	if err := ctx.ShouldBindJSON(&requestBody); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{
		"message": "/getprofile Not Implemented yet",
	})
}

func (uh *UserHandlers) HandleLogin(ctx *gin.Context) {
	var requestBody models.LoginRequestBody

	if err := ctx.ShouldBindJSON(&requestBody); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	message := uh.userService.Login(requestBody)

	ctx.JSON(http.StatusOK, gin.H{
		"message": message,
	})
}

func (uh *UserHandlers) HandleSignup(ctx *gin.Context) {
	var requestBody models.SignupRequestBody

	if err := ctx.ShouldBindJSON(&requestBody); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	message := uh.userService.Signup(requestBody)

	ctx.JSON(http.StatusOK, gin.H{
		"message": message,
	})
}

func (uh *UserHandlers) HandleRefreshToken(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{
		"message": "/refresh-token Not Implemented yet",
	})
}

func (uh *UserHandlers) HandleForgetPassword(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{
		"message": "Oh not again! Weak Memory",
	})
}
