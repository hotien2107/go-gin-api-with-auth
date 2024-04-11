package handlers

import (
	"net/http"

	"gin-rest-api.com/basic/internal/models"
	"gin-rest-api.com/basic/internal/services"
	"github.com/gin-gonic/gin"
)

type MessageHandler struct {
	services *services.MessageService
}

func NewMessageHandler() *MessageHandler {
	return &MessageHandler{
		services: services.NewMessageService(),
	}
}

func (h *MessageHandler) GetAllMessages(ctx *gin.Context) {
	allEvents, err := h.services.GetAllMessages(ctx)
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
		Result:  allEvents,
	})
}
