package handlers

import (
	"net/http"
	"strconv"
	"time"

	"gin-rest-api.com/basic/internal/models"
	"gin-rest-api.com/basic/internal/services"
	"github.com/gin-gonic/gin"
)

type EventHandler struct {
	services *services.EventService
}

func NewEventHandler() *EventHandler {
	return &EventHandler{
		services: services.NewEventService(),
	}
}

func (e *EventHandler) GetAllEvents(ctx *gin.Context) {
	allEvents, err := e.services.GetAll()
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

func (e *EventHandler) GetEventById(ctx *gin.Context) {
	eventIdStr := ctx.Param("id")
	eventId, err := strconv.ParseInt(eventIdStr, 10, 64)
	if err != nil {
		ctx.JSON(http.StatusOK, models.Response{
			IsError: true,
			Message: err.Error(),
			Result:  nil,
		})
		return
	}

	event, err := e.services.GetById(eventId)

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
		Result:  event,
	})
}

func (e *EventHandler) CreateNewEvent(ctx *gin.Context) {
	var newEvent models.Event
	err := ctx.ShouldBindJSON(&newEvent)

	if err != nil {
		ctx.JSON(http.StatusOK, models.Response{
			IsError: true,
			Message: err.Error(),
			Result:  nil,
		})
		return
	}

	newEvent.UserId = ctx.GetInt64("userId")
	newEvent.DateTime = time.Now()

	newEventId, err := e.services.Save(&newEvent)
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
		Message: "Create event success",
		Result:  newEventId,
	})
}

func (e *EventHandler) UpdateEvent(ctx *gin.Context) {
	eventIdStr := ctx.Param("id")
	eventId, err := strconv.ParseInt(eventIdStr, 10, 64)
	if err != nil {
		ctx.JSON(http.StatusOK, models.Response{
			IsError: true,
			Message: err.Error(),
			Result:  nil,
		})
		return
	}

	var newEvent models.Event
	err = ctx.ShouldBindJSON(&newEvent)
	if err != nil {
		ctx.JSON(http.StatusOK, models.Response{
			IsError: true,
			Message: err.Error(),
			Result:  nil,
		})
		return
	}
	newEvent.ID = eventId

	err = e.services.Update(&newEvent)
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
		Message: "Update event success",
		Result:  newEvent,
	})
}

func (e *EventHandler) DeleteEventById(ctx *gin.Context) {
	eventIdStr := ctx.Param("id")
	eventId, err := strconv.ParseInt(eventIdStr, 10, 64)
	if err != nil {
		ctx.JSON(http.StatusOK, models.Response{
			IsError: true,
			Message: err.Error(),
			Result:  nil,
		})
		return
	}

	err = e.services.Delete(eventId)
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
		Message: "Delete event success",
		Result:  eventId,
	})
}
