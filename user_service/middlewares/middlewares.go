package middlewares

import (
	"fmt"
	"net/http"
	"strings"
	httpResponse "user_service/http"
	"user_service/models"

	"github.com/gin-gonic/gin"

	"user_service/services"
)

type IUserMiddleware interface {
	Authorize(ctx *gin.Context)
}

type UserMiddleware struct {
	tokenService services.ITokenService
}

func NewUserMiddlewares(tokenService services.ITokenService) IUserMiddleware {
	return &UserMiddleware{
		tokenService: tokenService,
	}
}

func (um *UserMiddleware) Authorize(ctx *gin.Context) {
	accessTokenHeader := ctx.Request.Header.Get("Authorization")

	token, err := validateAndExtractToken(accessTokenHeader)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, httpResponse.GetErrorResponse(err.Error()))
		ctx.Abort()
		return
	}

	user, err := um.tokenService.GetInfoFromToken(token)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, httpResponse.GetErrorResponse(err.Error()))
		ctx.Abort()
		return
	}

	if user.Role != models.ADMIN {
		ctx.JSON(http.StatusBadRequest, httpResponse.GetErrorResponse("Admin Permissions required!!!!!"))
		ctx.Abort()
		return
	}
}

func validateAndExtractToken(tokenHeader string) (string, error) {
	if tokenHeader == "" {
		return "", fmt.Errorf("authorization Header is missing")
	}

	accessTokenParts := strings.Split(tokenHeader, " ")

	if len(accessTokenParts) != 2 || accessTokenParts[0] != "Bearer" {
		return "", fmt.Errorf("invalid or missing token")
	}

	return accessTokenParts[1], nil
}
