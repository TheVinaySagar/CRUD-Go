package routes

import (
	"myproject/controllers"

	"github.com/gin-gonic/gin"
)

func CloudinaryRoute(router *gin.Engine) {
	imageRoutes := router.Group("/image")
	{
		imageRoutes.POST("/uploadimage", controllers.Uploadimage)
	}
}
