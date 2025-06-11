package controllers

import (
	"context"
	"fmt"
	"net/http"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"

	"myproject/database"
	"myproject/models"

	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

func GetUsers(c *gin.Context) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	// Fetch users from MongoDB
	cursor, err := database.UserCollection.Find(ctx, bson.M{})
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch users"})
		return
	}

	var users []models.User
	if err = cursor.All(ctx, &users); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error reading users"})
		return
	}

	c.JSON(http.StatusOK, users)
}

func CreateUser(c *gin.Context) {

	ctx, cancel := context.WithTimeout(context.Background(), 100*time.Second)
	var user models.User
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	count, err := database.UserCollection.CountDocuments(ctx, bson.M{"email": user.Email})
	defer cancel()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"Message": err.Error()})
		return
	}

	if count > 0 {
		c.JSON(http.StatusInternalServerError, gin.H{"Message": "User Exists"})
		return
	}

	// Set ObjectID
	user.ID = primitive.NewObjectID()

	passwordhash, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"Error": err.Error()})
		return
	}

	fmt.Println("Hashed Password is : ", string(passwordhash))

	user.Password = string(passwordhash)
	ctx, cancel = context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	_, err = database.UserCollection.InsertOne(ctx, user)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create user", "details": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, user)
}

func loginuser(c *gin.Context){
	var user models.User
	var foundUser models.User
	if err := c.ShouldBindJSON(&user); err != nil{
		c.JSON(http.StatusInternalServerError, gin.H{"Error":err.Error()})
		return
	}

	ctx, cancel := context.WithTimeout(context.Background(), 100*time.Second)

	err := database.UserCollection.FindOne(ctx, bson.M{"email":user.Email}).Decode(&foundUser)
	defer cancel()

	if err != nil{
		c.JSON(http.StatusInternalServerError, gin.H{"Message": "Email or password is incorrect"})
		return
	}

	// passwordIsValid = d








}
