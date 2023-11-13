package routes

import (
	controllers "github.com/OSCode-Community/oscode-app-backend/controllers"

	"github.com/gin-gonic/gin"
)

func EventRoutes(incomingRoutes *gin.Engine) {
	// incomingRoutes.GET("/events", controllers.GetEvents())
	incomingRoutes.GET("/events/:event_id", controllers.GetEvent())
	incomingRoutes.POST("/events", controllers.NewEvent())
	// UPDATE event
	// UPDATE event participants
	// UPDATE event attendees
	// UPDATE event hosts
}
