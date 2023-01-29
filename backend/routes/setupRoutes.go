package routes

import (
	"detection-no-helmet-web-application/api/controller"
	"github.com/gin-gonic/gin"
)

func SetupRoutes(router *gin.Engine) {
	router.POST("/car", controller.CreateImage())
	router.GET("/car", controller.GetAllImage())
	router.GET("/car/:id", controller.GetImageById())
}
