package routes

import (
	"EventBooking/models"
	"EventBooking/utils"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
	"time"
)

func getEvents(ctx *gin.Context) {
	events, err := models.GetAvailableEvents()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"message": "Could not retrieve events"})
		return
	}
	ctx.JSON(http.StatusOK, events)
}

func createEvents(ctx *gin.Context) {
	token := ctx.Request.Header.Get("Authorization")
	if token == "" {
		ctx.JSON(http.StatusUnauthorized, gin.H{"message": "Unauthorized"})
		return
	}
	err := utils.ValidateToken(token)
	if err != nil {
		ctx.JSON(http.StatusUnauthorized, gin.H{"message": "Unauthorized"})
	}
	var event models.Event
	event.DateTime = time.Now()
	err = ctx.ShouldBindJSON(&event)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": "Could not parse the request data"})
		return
	}

	err = event.Save()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"message": "Could not create the event"})
		return
	}

	ctx.JSON(http.StatusCreated, gin.H{"message": "Event created!", "event": event})
}

func getEvent(ctx *gin.Context) {
	eventId, err := strconv.ParseInt(ctx.Params.ByName("id"), 10, 64)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": "Invalid event ID"})
		return
	}
	event, err := models.GetEventByID(eventId)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"message": "Could not retrieve the event"})
		return
	}
	ctx.JSON(http.StatusOK, event)
}

func updateEvent(ctx *gin.Context) {
	eventId, err := strconv.ParseInt(ctx.Params.ByName("id"), 10, 64)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": "Invalid event ID"})
		return
	}

	_, err = models.GetEventByID(eventId)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"message": "Could not retrieve the event"})
		return
	}

	var updatedEvent models.Event
	err = ctx.ShouldBindJSON(&updatedEvent)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": "Could not parse the request data"})
		return
	}

	updatedEvent.ID = eventId
	err = updatedEvent.Update()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"message": "Could not update the event"})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"message": "Event updated!", "event": updatedEvent})
}

func deleteEvent(ctx *gin.Context) {
	eventID, err := strconv.ParseInt(ctx.Params.ByName("id"), 10, 64)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": "Invalid event ID"})
		return
	}
	event, err := models.GetEventByID(eventID)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"message": "Could not retrieve the event"})
		return
	}
	err = event.DeleteEvent()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"message": "Could not delete the event"})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"message": "Event deleted!"})
}
