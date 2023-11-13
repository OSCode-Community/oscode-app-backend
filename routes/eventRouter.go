package routes

import (
	controllers "github.com/OSCode-Community/oscode-app-backend/controllers"

	"github.com/gin-gonic/gin"
)

func EventRoutes(incomingRoutes *gin.Engine) {
	// incomingRoutes.GET("/events", controllers.GetEvents())
	incomingRoutes.GET("/events/:event_id", controllers.GetEvent())
	incomingRoutes.POST("/events", controllers.NewEvent())
	incomingRoutes.PUT("/events/:event_id", controllers.UpdateEvent())
	// UPDATE event participants
	// UPDATE event attendees
	// UPDATE event hosts
}
