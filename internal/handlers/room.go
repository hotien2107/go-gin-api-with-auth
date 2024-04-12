package handlers

import (
	"fmt"
	"net/http"

	"gin-rest-api.com/basic/internal/models"
	"gin-rest-api.com/basic/internal/services"
	"github.com/gin-gonic/gin"
)

type RoomHandler struct {
	services *services.RoomService
}

func NewRoomHandler() *RoomHandler {
	return &RoomHandler{
		services: services.NewRoomService(),
	}
}

func (h *RoomHandler) GetByUser(ctx *gin.Context) {
	rooms, err := h.services.GetByUser(ctx)
	if err != nil {
		ctx.JSON(http.StatusOK, models.Response{
			IsError: true,
			Message: err.Error(),
			Result:  nil,
		})
		return
	}

	ctx.JSON(http.StatusOK, models.Response{
		IsError: true,
		Message: "Successfully fetched user's rooms",
		Result:  rooms,
	})
}

func (h *RoomHandler) Create(ctx *gin.Context) {
	var newRoom models.Room
	err := ctx.ShouldBindJSON(&newRoom)
	if err != nil {
		ctx.JSON(http.StatusOK, models.Response{
			IsError: true,
			Message: err.Error(),
			Result:  nil,
		})
		return
	}

	roomId, err := h.services.Create(ctx, &newRoom)
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
		Message: "Create tag success",
		Result:  "Room's id is: " + fmt.Sprintf("%d", roomId),
	})
}
