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
	middlewares.LoadConfig()

	port := os.Getenv("PORT")
	if port == "" {
		port = "8000"
	}

	uri := os.Getenv("MONGODB_URI")
	if uri == "" {
		log.Fatal("You must set your 'MONGODB_URI' environment variable. See\n\t https://www.mongodb.com/docs/drivers/go/current/usage-examples/#environment-variable")
	}

	router := gin.New()
	router.Use(gin.Logger())

	routes.AuthRoutes(router)

	router.Run(":" + port)

	// // Use the SetServerAPIOptions() method to set the Stable API version to 1
	// serverAPI := options.ServerAPI(options.ServerAPIVersion1)
	// opts := options.Client().ApplyURI(uri).SetServerAPIOptions(serverAPI)

	// // Create a new client and connect to the server
	// client, err := mongo.Connect(context.TODO(), opts)
	// if err != nil {
	// 	panic(err)
	// }
	// defer func() {
	// 	if err = client.Disconnect(context.TODO()); err != nil {
	// 		panic(err)
	// 	}
	// }()

	// coll := client.Database("oscode-app-database").Collection("users")
	// doc := models.User{
	// 	FirstName: "Harsh",
	// 	LastName:  "Singh",
	// 	Email:     "singh.harsh9097@gmail.com",
	// }
	// result, err := coll.InsertOne(context.TODO(), doc)
	// if err != nil {
	// 	panic(err)
	// }

	// fmt.Printf("Inserted document with _id: %v\n", result.InsertedID)
}
