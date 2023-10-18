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
	HandleVerifyUser(ctx *gin.Context)
	HandlePingEmail(ctx *gin.Context)
}

type UserHandlers struct {
	userService  services.IUserService
	emailService services.IEmailService
}

func NewUserHandlers(userService services.IUserService, emailService services.IEmailService) IUserHandlers {
	return &UserHandlers{userService: userService, emailService: emailService}
}

func (uh *UserHandlers) HandleGetProfile(ctx *gin.Context) {
	var requestBody models.ProfileRequest

	if err := ctx.ShouldBindJSON(&requestBody); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{
		"message": "/getprofile Not Implemented yet",
	})
}

func (uh *UserHandlers) HandleLogin(ctx *gin.Context) {
	var requestBody models.LoginRequest

	if err := ctx.ShouldBindJSON(&requestBody); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	message, err := uh.userService.Login(requestBody)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{
		"message": message,
	})
}

func (uh *UserHandlers) HandleSignup(ctx *gin.Context) {
	var requestBody models.SignupRequest

	if err := ctx.ShouldBindJSON(&requestBody); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	token, err := uh.userService.Signup(requestBody)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// send email with token
	err = uh.emailService.SendEmail(token, requestBody.Email)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	//ASK: should i just send the token in return or some message as well if not in this case is it appropriate to make the map on the fly
	ctx.JSON(http.StatusOK, gin.H{
		"message": token,
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

func (uh *UserHandlers) HandleVerifyUser(ctx *gin.Context) {
	accessToken := ctx.Param("token")
	err := uh.userService.VerifyUser(accessToken)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"message": "Your Email has been verified, you can use the app now.",
	})
}

func (uh *UserHandlers) HandlePingEmail(ctx *gin.Context) {
	err := uh.emailService.SendEmail("fahadshykh369@gmail.com", "sadfj")
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{
		"message": "Email sent successfully",
	})
}
