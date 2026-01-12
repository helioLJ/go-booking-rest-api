package routes

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/helio-pt/go-booking-rest-api/models"
)

func getEvents(context *gin.Context) {
	events := models.GetAllEvents()
	context.JSON(http.StatusOK, events)
}

func createEvent(context *gin.Context) {
	var event models.Event
	err := context.ShouldBindJSON(&event)

	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "Could not parse request data."})
		return
	}

	event.ID = time.Now().Unix()
	event.UserID = 1 // Hardcoded for Phase 1

	event.Save()
	context.JSON(http.StatusCreated, gin.H{"message": "Event created!", "event": event})
}

func RegisterRoutes(server *gin.Engine) {
	server.GET("/events", getEvents)
	server.POST("/events", createEvent)
}
