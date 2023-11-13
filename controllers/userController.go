package controllers

import (
	"context"
	"log"
	"net/http"

	"github.com/OSCode-Community/oscode-app-backend/database"
	"github.com/OSCode-Community/oscode-app-backend/models"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

func GetUser() gin.HandlerFunc {
	return func(c *gin.Context) {
		userId := c.Param("user_id")

		objId, err := primitive.ObjectIDFromHex(userId)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "invalid user id"})
			return
		}

		if err = database.StartMongoDB(); err != nil {
			log.Fatal("Unable to Start a New MongoDB server")
		}
		collection := database.GetCollection("users")
		defer database.CloseMongoDB()

		user := models.User{}
		err = collection.FindOne(context.Background(), bson.M{"_id": objId}).Decode(&user)
		if err != nil {
			if err == mongo.ErrNoDocuments {
				c.JSON(http.StatusNotFound, gin.H{"error": "user not found"})
				return
			}

			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		c.JSON(http.StatusOK, user)
	}
}

func NewUser() gin.HandlerFunc {
	return func(c *gin.Context) {
		userName := c.Request.Header.Get("first_name")
		email := c.Request.Header.Get("email")

		if userName == "" {
			c.JSON(http.StatusBadRequest, gin.H{"error": "user name must not be empty"})
			return
		}

		if email == "" {
			c.JSON(http.StatusBadRequest, gin.H{"error": "email must not be empty"})
			return
		}

		if err := database.StartMongoDB(); err != nil {
			log.Fatal("Unable to Start a New MongoDB server")
		}
		collection := database.GetCollection("users")
		defer database.CloseMongoDB()

		// Create a new event
		user := models.User{
			FirstName: userName,
			Email:     email,
		}
		user.ID = primitive.NewObjectID()

		result, err := collection.InsertOne(context.Background(), user)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		c.JSON(http.StatusOK, gin.H{"id": result.InsertedID})
	}
}

func UpdateUser() gin.HandlerFunc {
	return func(c *gin.Context) {
		userId := c.Param("user_id")
		firstName := c.Request.Header.Get("first_name")
		email := c.Request.Header.Get("email")

		if firstName == "" {
			c.JSON(http.StatusBadRequest, gin.H{"error": "first name must not be empty"})
			return
		}

		if email == "" {
			c.JSON(http.StatusBadRequest, gin.H{"error": "email must not be empty"})
			return
		}

		objId, err := primitive.ObjectIDFromHex(userId)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"invalid input": err.Error()})
			return
		}

		// database
		if err = database.StartMongoDB(); err != nil {
			log.Fatal("Unable to Start a New MongoDB server")
		}
		collection := database.GetCollection("users")
		defer database.CloseMongoDB()

		filter := bson.M{"_id": objId}
		update := bson.M{
			"$set": bson.M{
				"first_name": firstName,
				"email":      email,
			},
		}

		result, err := collection.UpdateOne(context.Background(), filter, update)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		c.JSON(http.StatusOK, gin.H{"updated_count": result.ModifiedCount})
	}
}
