package controllers

import (
	"context"
	"log"
	"net/http"
	"time"

	"github.com/OSCode-Community/oscode-app-backend/database"
	"github.com/OSCode-Community/oscode-app-backend/models"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

// func GetEvents() gin.HandlerFunc {}

func GetEvent() gin.HandlerFunc {
	return func(c *gin.Context) {
		eventId := c.Param("event_id")

		objId, err := primitive.ObjectIDFromHex(eventId)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "invalid event id"})
			return
		}

		if err = database.StartMongoDB(); err != nil {
			log.Fatal("Unable to Start a New MongoDB server")
		}
		collection := database.GetCollection("events")
		defer database.CloseMongoDB()

		event := models.Event{}
		err = collection.FindOne(context.Background(), bson.M{"_id": objId}).Decode(&event)
		if err != nil {
			if err == mongo.ErrNoDocuments {
				c.JSON(http.StatusNotFound, gin.H{"error": "event not found"})
				return
			}

			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		c.JSON(http.StatusOK, event)
	}
}

func NewEvent() gin.HandlerFunc {
	return func(c *gin.Context) {
		eventName := c.Request.Header.Get("name")
		startAt := c.Request.Header.Get("start_at")
		endAt := c.Request.Header.Get("end_at")

		if eventName == "" {
			c.JSON(http.StatusBadRequest, gin.H{"error": "event name must not be empty"})
			return
		}

		if startAt == "" {
			c.JSON(http.StatusBadRequest, gin.H{"error": "start date must not be empty"})
			return
		}

		if endAt == "" {
			c.JSON(http.StatusBadRequest, gin.H{"error": "end date must not be empty"})
			return
		}

		startTime, err := time.Parse(time.RFC3339, startAt)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "invalid start date"})
			return
		}

		endTime, err := time.Parse(time.RFC3339, endAt)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "invalid end date"})
			return
		}

		if err = database.StartMongoDB(); err != nil {
			log.Fatal("Unable to Start a New MongoDB server")
		}
		collection := database.GetCollection("events")
		defer database.CloseMongoDB()

		event := models.Event{
			Name:    eventName,
			StartAt: startTime,
			EndAt:   endTime,
		}
		event.ID = primitive.NewObjectID()

		result, err := collection.InsertOne(c, event)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		c.JSON(http.StatusOK, gin.H{"id": result.InsertedID})
	}
}
