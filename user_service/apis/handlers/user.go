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
	HandleGetUserEmailsByIDs(ctx *gin.Context)
	HandleLogin(ctx *gin.Context)
	HandleSignup(ctx *gin.Context)
	HandleRefreshToken(ctx *gin.Context)
	HandleForgetPassword(ctx *gin.Context)
	HandleSendEmail(ctx *gin.Context)
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

func (uh *UserHandlers) HandleGetUserEmailsByIDs(ctx *gin.Context) {
	var requestBody models.GetUserEmailsByIdsRequest
	log.GetLog().Info("running HandleGetUserEmailsByIDs..........")

	if err := ctx.ShouldBindJSON(&requestBody); err != nil {
		errorResponse := httpResponse.GetErrorResponse(err.Error())
		ctx.JSON(http.StatusBadRequest, errorResponse)
		return
	}

	userEmails, err := uh.userService.GetUserEmailsByIds(requestBody.UserIDs)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, httpResponse.GetErrorResponse(err.Error()))
		return
	}

	log.GetLog().Info(userEmails)
	ctx.JSON(http.StatusOK, httpResponse.GetSuccessResponse(userEmails))
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

func (uh *UserHandlers) HandleSendEmail(ctx *gin.Context) {
	var requestBody models.SendEmailRequest

	if err := ctx.ShouldBindJSON(&requestBody); err != nil {
		errorResponse := httpResponse.GetErrorResponse(err.Error())
		ctx.JSON(http.StatusBadRequest, errorResponse)
		return
	}
	subject := `[Important] Todo Due Today`
	emailBody := `
	<html>
		<body>
			<p>Hello!</p>
			<p>You have pending todos which are due today open your app and get it down. No More procrastination</p>
		</body>
	</html>
	`

	err := uh.emailService.SendEmailToAll(requestBody.UserEmailAddresses, subject, emailBody)
	if err != nil {
		errorResponse := httpResponse.GetErrorResponse(err.Error())
		ctx.JSON(http.StatusBadRequest, errorResponse)
		return
	}

	ctx.JSON(http.StatusOK, httpResponse.GetSuccessResponse("email sent"))
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
