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

func (h *MessageHandler) GetAlls(ctx *gin.Context) {
	allEvents, err := h.services.GetAlls(ctx)
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

func (h *MessageHandler) Send(ctx *gin.Context) {
	var newMess models.Message
	err := ctx.ShouldBindJSON(&newMess)
	if err != nil {
		ctx.JSON(http.StatusOK, models.Response{
			IsError: true,
			Message: err.Error(),
			Result:  nil,
		})
		return
	}

	err = h.services.Send(ctx, &newMess)
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
		Message: "Send message successfully",
		Result:  nil,
	})
}
