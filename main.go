package main

import (
	"log"
	"os"

	"github.com/OSCode-Community/oscode-app-backend/middlewares"
	"github.com/OSCode-Community/oscode-app-backend/routes"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	if err := godotenv.Load(); err != nil {
		log.Println("No .env file found")
	}

	port := os.Getenv("PORT")
	if port == "" {
		port = "8000"
	}

	middlewares.LoadConfig()

	router := gin.New()
	router.Use(gin.Logger())

	routes.AuthRoutes(router)
	routes.EventRoutes(router)
	routes.UserRoutes(router)

	router.Run(":" + port)
}
