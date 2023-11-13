package routes

import (
	controllers "github.com/OSCode-Community/oscode-app-backend/controllers"

	"github.com/gin-gonic/gin"
)

func EventRoutes(incomingRoutes *gin.Engine) {
	// incomingRoutes.GET("/events", controllers.GetEvents())
	incomingRoutes.GET("/events/:event_id", controllers.GetEvent())
	incomingRoutes.POST("/events", controllers.NewEvent()) // TODO: add hosts in this function
	incomingRoutes.PUT("/events/:event_id", controllers.UpdateEvent())
	incomingRoutes.PUT("/events/:event_id/participants", controllers.UpdateParticipants())
	// incomingRoutes.PUT("/events/:event_id/participants", controllers.UpdateAttendees())
}
