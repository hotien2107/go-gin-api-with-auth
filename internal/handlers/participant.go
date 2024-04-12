package handlers

import (
	"errors"
	"fmt"
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

type roomRequest struct {
	RoomId int64 `json:"roomId"`
}

func (h *ParticipantHandler) JoinRoom(ctx *gin.Context) {
	var joinRoomReq roomRequest
	err := ctx.ShouldBindJSON(&joinRoomReq)

	if err != nil {
		ctx.JSON(http.StatusOK, models.Response{
			IsError: true,
			Message: err.Error(),
			Result:  nil,
		})
		return
	}
	if joinRoomReq.RoomId == 0 {
		ctx.JSON(http.StatusOK, models.Response{
			IsError: true,
			Message: errors.New("Invalid Room ID: " + fmt.Sprintf("%d", joinRoomReq.RoomId)).Error(),
			Result:  nil,
		})
		return
	}

	err = h.services.JoinRoom(ctx, joinRoomReq.RoomId)
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
	var roomReq roomRequest
	err := ctx.ShouldBindJSON(&roomReq)
	if err != nil {
		ctx.JSON(http.StatusOK, models.Response{
			IsError: true,
			Message: err.Error(),
			Result:  nil,
		})
		return
	}

	participants, err := h.services.GetAllParticipantInRoom(roomReq.RoomId)
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
