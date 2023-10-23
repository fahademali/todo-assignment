package handlers

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"

	"user_service/models"
	"user_service/services"
)

type IUserHandlers interface {
	HandleGetProfile(ctx *gin.Context)
	HandleLogin(ctx *gin.Context)
	HandleSignup(ctx *gin.Context)
	HandleSignupTx(ctx *gin.Context)
	HandleRefreshToken(ctx *gin.Context)
	HandleForgetPassword(ctx *gin.Context)
	HandleVerifyUser(ctx *gin.Context)
}

type UserHandlers struct {
	userService  services.IUserService
	emailService services.IEmailService
	tokenService services.ITokenService
}

func NewUserHandlers(userService services.IUserService, emailService services.IEmailService, tokenService services.ITokenService) IUserHandlers {
	return &UserHandlers{userService: userService, emailService: emailService, tokenService: tokenService}
}

func (uh *UserHandlers) HandleGetProfile(ctx *gin.Context) {
	var requestBody models.ProfileRequest

	if err := ctx.ShouldBindJSON(&requestBody); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{
		"data": "/getprofile Not Implemented yet",
	})
}

func (uh *UserHandlers) HandleLogin(ctx *gin.Context) {
	fmt.Println("Running Handle Login.....")
	var requestBody models.LoginRequest

	if err := ctx.ShouldBindJSON(&requestBody); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	token, err := uh.userService.Login(requestBody)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{
		"data": token,
	})
}

func (uh *UserHandlers) HandleSignup(ctx *gin.Context) {
	var requestBody models.SignupRequest

	if err := ctx.ShouldBindJSON(&requestBody); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	_, err := uh.userService.Signup(requestBody)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"data": "Verfication email has been sent to your email address, please verify.",
	})
}

func (uh *UserHandlers) HandleSignupTx(ctx *gin.Context) {
	var requestBody models.SignupRequest

	if err := ctx.ShouldBindJSON(&requestBody); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	err := uh.userService.SignupTx(ctx, requestBody)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"data": "Verfication email has been sent to your email address, please verify.",
	})
}

func (uh *UserHandlers) HandleRefreshToken(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{
		"data": "/refresh-token Not Implemented yet",
	})
}

func (uh *UserHandlers) HandleForgetPassword(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{
		"data": "Oh not again! Weak Memory",
	})
}

func (uh *UserHandlers) HandleVerifyUser(ctx *gin.Context) {
	fmt.Println("running HandleVerifyUser......................")
	fmt.Println(uh)
	fmt.Println(uh.userService)
	accessToken := ctx.Param("token")

	email, err := uh.tokenService.GetEmailFromAccessToken(accessToken)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	err = uh.userService.VerifyUser(email)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"data": "Your Email has been verified, you can use the app now.",
	})
}
