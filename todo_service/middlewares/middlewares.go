package middlewares

import (
	"fmt"
	"net/http"
	"strings"
	httpResponse "todo_service/http"

	"github.com/gin-gonic/gin"

	"todo_service/services"
)

type ITodoMiddleware interface {
	Authenticate(ctx *gin.Context)
}

type TodoMiddleware struct {
	tokenService services.ITokenService
}

func NewTodoMiddlewares(tokenService services.ITokenService) ITodoMiddleware {
	return &TodoMiddleware{
		tokenService: tokenService,
	}
}

func (tm *TodoMiddleware) Authenticate(ctx *gin.Context) {
	accessTokenHeader := ctx.Request.Header.Get("Authorization")

	token, err := validateAndExtractToken(accessTokenHeader)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, httpResponse.GetErrorResponse(err.Error()))
		ctx.Abort()
		return
	}

	user, err := tm.tokenService.GetInfoFromToken(token)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, httpResponse.GetErrorResponse(err.Error()))
		ctx.Abort()
		return
	}
	ctx.Set("user", user)
	ctx.Next()
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
