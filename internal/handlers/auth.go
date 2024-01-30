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

// SignUp godoc
// @Summary Sign up a new user
// @Description Create a new user account
// @Tags Authentication
// @Accept  json
// @Produce  json
// @Param signUpInput body models.User true "Sign up input"
// @Success 200 {object} models.Response
// @Failure 400 {object} models.Response
// @Router /sign-up [post]
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

// @Summary Login
// @Description Logs user in and returns an access token
// @Tags Authentication
// @Accept json
// @Produce json
// @Param input body models.User true "User credentials"
// @Success 200 {object} models.Response
// @Failure 401 {object} models.Response
// @Router /auth/login [post]
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

	accessToken, err := h.services.Login(&user)
	if err != nil {
		ctx.JSON(http.StatusUnauthorized, models.Response{
			IsError: true,
			Message: err.Error(),
			Result:  nil,
		})
		return
	}

	// Set the token in a HTTP-only cookie
	ctx.SetCookie("token", accessToken, 3600, "/", "localhost", true, true) // HttpOnly set to true

	ctx.JSON(http.StatusOK, models.Response{
		IsError: false,
		Message: "Login success!",
		Result: gin.H{
			"token": accessToken,
		},
	})
}
