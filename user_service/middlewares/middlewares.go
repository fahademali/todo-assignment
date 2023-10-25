package middlewares

import (
	"net/http"
	"strings"
	httpResponse "user_service/http"

	"github.com/gin-gonic/gin"

	"user_service/log"
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
	log.GetLog().Info("middleware .............")
	token := strings.Split(ctx.Request.Header["Authorization"][0], " ")[1]
	user, err := um.tokenService.GetInfoFromToken(token)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, httpResponse.GetErrorResponse(err.Error()))
		ctx.Abort()
		return
	}
	log.GetLog().Info(user)
	if user.Role != "admin" {
		ctx.JSON(http.StatusBadRequest, httpResponse.GetErrorResponse("Admin Permissions required!!!!!"))
		ctx.Abort()
		return
	}
	log.GetLog().Info(user)
}
