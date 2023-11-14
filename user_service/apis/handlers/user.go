package handlers

import (
	"net/http"
	"todo_service/log"

	"github.com/gin-gonic/gin"

	httpResponse "user_service/http"
	"user_service/models"
	"user_service/services"
)

type IUserHandlers interface {
	HandleGetProfile(ctx *gin.Context)
	HandleGetUserByIDs(ctx *gin.Context)
	HandleLogin(ctx *gin.Context)
	HandleSignup(ctx *gin.Context)
	HandleRefreshToken(ctx *gin.Context)
	HandleForgetPassword(ctx *gin.Context)
	HandleSendEmails(ctx *gin.Context)
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

func (uh *UserHandlers) HandleGetProfile(ctx *gin.Context) {
	var requestBody models.ProfileRequest

	if err := ctx.ShouldBindJSON(&requestBody); err != nil {
		errorResponse := httpResponse.GetErrorResponse(err.Error())
		ctx.JSON(http.StatusBadRequest, errorResponse)
		return
	}

	successResponse := httpResponse.GetSuccessResponse("/getprofile Not Implemented yet")
	ctx.JSON(http.StatusOK, successResponse)
}

func (uh *UserHandlers) HandleGetUserByIDs(ctx *gin.Context) {
	var requestBody models.GetUserByIDsRequest

	if err := ctx.ShouldBindJSON(&requestBody); err != nil {
		errorResponse := httpResponse.GetErrorResponse(err.Error())
		ctx.JSON(http.StatusBadRequest, errorResponse)
		return
	}

	users, err := uh.userService.GetUserByIds(requestBody.UserIDs)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, httpResponse.GetErrorResponse(err.Error()))
		return
	}

	log.GetLog().Info(users)
	ctx.JSON(http.StatusOK, httpResponse.GetSuccessResponse(users))
}

func (uh *UserHandlers) HandleLogin(ctx *gin.Context) {
	var requestBody models.LoginRequest

	if err := ctx.ShouldBindJSON(&requestBody); err != nil {
		errorResponse := httpResponse.GetErrorResponse(err.Error())
		ctx.JSON(http.StatusBadRequest, errorResponse)
		return
	}

	token, err := uh.userService.Login(requestBody)
	if err != nil {
		errorResponse := httpResponse.GetErrorResponse(err.Error())
		ctx.JSON(http.StatusBadRequest, errorResponse)
		return
	}
	successResponse := httpResponse.GetSuccessResponse(token)
	ctx.JSON(http.StatusOK, successResponse)
}

func (uh *UserHandlers) HandleSignup(ctx *gin.Context) {
	var requestBody models.SignupRequest

	if err := ctx.ShouldBindJSON(&requestBody); err != nil {
		errorResponse := httpResponse.GetErrorResponse(err.Error())
		ctx.JSON(http.StatusBadRequest, errorResponse)
		return
	}

	err := uh.userService.Signup(ctx, requestBody)
	if err != nil {
		errorResponse := httpResponse.GetErrorResponse(err.Error())
		ctx.JSON(http.StatusBadRequest, errorResponse)
		return
	}

	successResponse := httpResponse.GetSuccessResponse("Verfication email has been sent to your email address, please verify.")
	ctx.JSON(http.StatusOK, successResponse)
}

func (uh *UserHandlers) HandleRefreshToken(ctx *gin.Context) {
	successResponse := httpResponse.GetSuccessResponse("/refresh-token Not Implemented yet")
	ctx.JSON(http.StatusOK, successResponse)
}

func (uh *UserHandlers) HandleForgetPassword(ctx *gin.Context) {
	successResponse := httpResponse.GetSuccessResponse("Oh not again! Weak Memory")
	ctx.JSON(http.StatusOK, successResponse)
}

func (uh *UserHandlers) HandleSendEmails(ctx *gin.Context) {
	var requestBody models.SendEmailsRequest

	if err := ctx.ShouldBindJSON(&requestBody); err != nil {
		errorResponse := httpResponse.GetErrorResponse(err.Error())
		ctx.JSON(http.StatusBadRequest, errorResponse)
		return
	}

	err := uh.emailService.SendEmailToAll(requestBody.RecepietDetails)
	if err != nil {
		log.GetLog().Error(err)
		errorResponse := httpResponse.GetErrorResponse(err.Error())
		ctx.JSON(http.StatusBadRequest, errorResponse)
		return
	}

	ctx.JSON(http.StatusOK, httpResponse.GetSuccessResponse("emails have been sent to everybody"))
}

func (uh *UserHandlers) HandleVerifyUser(ctx *gin.Context) {
	accessToken := ctx.Param("token")

	email, err := uh.tokenService.GetEmailFromAccessToken(accessToken)
	if err != nil {
		errorResponse := httpResponse.GetErrorResponse(err.Error())
		ctx.JSON(http.StatusBadRequest, errorResponse)
		return
	}

	err = uh.userService.VerifyUser(email)
	if err != nil {
		errorResponse := httpResponse.GetErrorResponse(err.Error())
		ctx.JSON(http.StatusBadRequest, errorResponse)
		return
	}

	successResponse := httpResponse.GetSuccessResponse("Your Email has been verified, you can use the app now.")
	ctx.JSON(http.StatusOK, successResponse)
}

func (uh *UserHandlers) HandleGrantAdminRole(ctx *gin.Context) {
	var requestBody models.AdminPromotionRequest
	if err := ctx.ShouldBindJSON(&requestBody); err != nil {
		errorResponse := httpResponse.GetErrorResponse(err.Error())
		ctx.JSON(http.StatusBadRequest, errorResponse)
		return
	}

	err := uh.userService.GrantAdminRole(requestBody.Email)
	if err != nil {
		errorResponse := httpResponse.GetErrorResponse(err.Error())
		ctx.JSON(http.StatusBadRequest, errorResponse)
		return
	}

	successResponse := httpResponse.GetSuccessResponse("Admin Role has been granted to the provided email.")
	ctx.JSON(http.StatusOK, successResponse)
}
