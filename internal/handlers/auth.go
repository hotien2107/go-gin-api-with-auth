package handlers

import (
	"net/http"

	"gin-rest-api.com/basic/internal/models"
	"gin-rest-api.com/basic/internal/services"
	"github.com/gin-gonic/gin"
)

type AuthHandler struct {
	services *services.AuthService
}

func NewAuthHandler() *AuthHandler {
	return &AuthHandler{
		services: services.NewAuthService(),
	}
}

func (h *AuthHandler) SignUp(ctx *gin.Context) {
	var newUser models.User
	err := ctx.ShouldBindJSON(&newUser)

	if err != nil {
		ctx.JSON(http.StatusOK, models.Response{
			IsError: true,
			Message: err.Error(),
			Result:  nil,
		})
		return
	}

	newUserId, err := h.services.SignUp(&newUser)
	if err != nil {
		ctx.JSON(http.StatusOK, models.Response{
			IsError: true,
			Message: err.Error(),
			Result:  nil,
		})
		return
	}

	ctx.JSON(http.StatusOK, models.Response{
		IsError: false,
		Message: "Sign up success",
		Result:  newUserId,
	})
}

func (h *AuthHandler) Login(ctx *gin.Context) {
	var user models.User
	err := ctx.ShouldBindJSON(&user)

	if err != nil {
		ctx.JSON(http.StatusUnauthorized, models.Response{
			IsError: true,
			Message: err.Error(),
			Result:  nil,
		})
		return
	}

	err = h.services.Login(&user)
	if err != nil {
		ctx.JSON(http.StatusUnauthorized, models.Response{
			IsError: true,
			Message: err.Error(),
			Result:  nil,
		})
		return
	}

	ctx.JSON(http.StatusOK, models.Response{
		IsError: false,
		Message: "Login success!",
		Result:  nil,
	})
}
