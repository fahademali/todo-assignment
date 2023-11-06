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

	if accessTokenHeader == "" {
		ctx.JSON(http.StatusBadRequest, httpResponse.GetErrorResponse("Authorization Header is missing"))
		ctx.Abort()
		return
	}

	accessTokenParts := strings.Split(accessTokenHeader, " ")
	fmt.Println(accessTokenParts)
	if len(accessTokenParts) != 2 || accessTokenParts[0] != "Bearer" {
		ctx.JSON(http.StatusBadRequest, httpResponse.GetErrorResponse("Invalid or missing token"))
		ctx.Abort()
		return
	}

	token := accessTokenParts[1]
	user, err := tm.tokenService.GetInfoFromToken(token)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, httpResponse.GetErrorResponse(err.Error()))
		ctx.Abort()
		return
	}
	ctx.Set("user", user)
	ctx.Next()
}
