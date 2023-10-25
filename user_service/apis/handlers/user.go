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
	HandleGrantAdminRole(ctx *gin.Context)
}

type UserHandlers struct {
	userService  services.IUserService
	emailService services.IEmailService
	tokenService services.ITokenService
}

func NewUserHandlers(userService services.IUserService, emailService services.IEmailService, tokenService services.ITokenService) IUserHandlers {
	return &UserHandlers{userService: userService, emailService: emailService, tokenService: tokenService}
}

var successResponse = models.SuccessResponse{
	Status:  "success",
	Message: "Request was successfull",
}

var errorResponse = models.ErrorResponse{
	Status:  "error",
	Message: "Request failed",
}

func (uh *UserHandlers) HandleGetProfile(ctx *gin.Context) {
	var requestBody models.ProfileRequest

	if err := ctx.ShouldBindJSON(&requestBody); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	successResponse.Data = "/getprofile Not Implemented yet"
	ctx.JSON(http.StatusOK, successResponse)
}

func (uh *UserHandlers) HandleLogin(ctx *gin.Context) {
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

	successResponse.Data = token
	ctx.JSON(http.StatusOK, successResponse)
}

func (uh *UserHandlers) HandleSignup(ctx *gin.Context) {
	var requestBody models.SignupRequest

	if err := ctx.ShouldBindJSON(&requestBody); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	err := uh.userService.Signup(ctx, requestBody)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	successResponse.Data = "Verfication email has been sent to your email address, please verify."
	ctx.JSON(http.StatusOK, successResponse)
}

func (uh *UserHandlers) HandleRefreshToken(ctx *gin.Context) {
	successResponse.Data = "/refresh-token Not Implemented yet."
	ctx.JSON(http.StatusOK, successResponse)
}

func (uh *UserHandlers) HandleForgetPassword(ctx *gin.Context) {
	successResponse.Data = "Oh not again! Weak Memory."
	ctx.JSON(http.StatusOK, successResponse)
}

func (uh *UserHandlers) HandleVerifyUser(ctx *gin.Context) {
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

	successResponse.Data = "Your Email has been verified, you can use the app now."
	ctx.JSON(http.StatusOK, successResponse)
}

func (uh *UserHandlers) HandleGrantAdminRole(ctx *gin.Context) {
	var requestBody models.AdminPromotionRequest
	if err := ctx.ShouldBindJSON(&requestBody); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	err := uh.userService.GrantAdminRole(requestBody.Email)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	successResponse.Data = "Admin Role has been granted to the provided email."
	ctx.JSON(http.StatusOK, successResponse)
}
