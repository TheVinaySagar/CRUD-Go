package routes

import (
	"github.com/gin-gonic/gin"
	"myproject/controllers"
)

// RegisterRoutes initializes API routes
func RegisterRoutes(router *gin.Engine) {
	userRoutes := router.Group("/users")
	{
		userRoutes.GET("/", controllers.GetUsers)
		userRoutes.POST("/", controllers.CreateUser)
	}
}
