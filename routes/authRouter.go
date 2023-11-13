package routes

import (
	controllers "github.com/OSCode-Community/oscode-app-backend/controllers"

	"github.com/gin-gonic/gin"
)

func AuthRoutes(incomingRoutes *gin.Engine) {
	incomingRoutes.GET("/google_login", controllers.GoogleLogin)
	incomingRoutes.GET("/google_callback", controllers.GoogleCallback)
}
