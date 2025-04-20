package controllers

import (
	"context"
	"net/http"
	"time"

	"myproject/database"
	"myproject/models"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
)

func Getinterviews(c *gin.Context) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	cursor, err := database.InterviewCollection.Find(ctx, bson.M{})

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"Error": "Their is some error", "details": err.Error()})
		return
	}

	var interviews []models.Interview

	if err = cursor.All(ctx, &interviews); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"Error": "Failed to read users", "details": err.Error()})
		return
	}

	c.JSON(200, interviews)
}

func CreateInterview(c *gin.Context) {
	var interview models.Interview
	if err := c.ShouldBindJSON(&interview); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"Error": err.Error()})
		return
	}

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	_, err := database.InterviewCollection.InsertOne(ctx, interview)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"Error": err.Error()})
		return
	}

	c.JSON(http.StatusAccepted, interview)
}
