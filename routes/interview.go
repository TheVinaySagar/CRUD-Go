package routes

import (
	"myproject/controllers"

	"github.com/gin-gonic/gin"
)

// RegisterRoutes initializes API routes
func RegisterInterview(router *gin.Engine) {
	interviewRoutes := router.Group("/interviews")
	{
		interviewRoutes.GET("/", controllers.Getinterviews)
		interviewRoutes.POST("/", controllers.CreateInterview)
	}
}
