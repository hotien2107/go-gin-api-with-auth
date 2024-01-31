package middlewares

import (
	"net/http"
	"strings"

	"gin-rest-api.com/basic/internal/models"
	"gin-rest-api.com/basic/pkg/utils"
	"github.com/gin-gonic/gin"
)

func Authenticate(ctx *gin.Context) {
	token := ctx.Request.Header.Get("Authorization")
	if token == "" {
		ctx.AbortWithStatusJSON(http.StatusUnauthorized, models.Response{
			IsError: true,
			Message: "Missing auth token",
			Result:  nil,
		})
		return
	}

	// remove bearer string
	token = strings.Split(token, " ")[1]

	userId, err := utils.VerifyToken(token)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusUnauthorized, models.Response{
			IsError: true,
			Message: "Token is invalid: " + err.Error(),
			Result:  nil,
		})
		return
	}

	ctx.Set("userId", userId)
	ctx.Next()
}
