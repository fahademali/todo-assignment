package middlewares

import (
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

	if accessTokenHeader == "" {
		ctx.JSON(http.StatusBadRequest, httpResponse.GetErrorResponse("Authorization Header is missing"))
		ctx.Abort()
		return
	}

	accessTokenParts := strings.Split(accessTokenHeader, " ")

	if len(accessTokenParts) != 2 || accessTokenParts[0] != "Bearer" {
		ctx.JSON(http.StatusBadRequest, httpResponse.GetErrorResponse("Invalid or missing token"))
		ctx.Abort()
		return
	}

	token := accessTokenParts[1]
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
