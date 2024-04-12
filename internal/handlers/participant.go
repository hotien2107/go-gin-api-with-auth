package handlers

import (
	"net/http"

	"gin-rest-api.com/basic/internal/models"
	"gin-rest-api.com/basic/internal/services"
	"github.com/gin-gonic/gin"
)

type ParticipantHandler struct {
	services *services.ParticipantService
}

func NewParticipantHandler() *ParticipantHandler {
	return &ParticipantHandler{
		services: services.NewParticipantService(),
	}
}

func (h *ParticipantHandler) JoinRoom(ctx *gin.Context) {
	roomId := ctx.GetInt64("roomIdId")

	err := h.services.JoinRoom(ctx, roomId)
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
		Message: "Successfully joined the room.",
		Result:  nil,
	})
}

func (h *ParticipantHandler) GetAllParticipantInRoom(ctx *gin.Context) {
	roomId := ctx.GetInt64("roomIdId")

	participants, err := h.services.GetAllParticipantInRoom(roomId)
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
		Message: "",
		Result:  participants,
	})
}
