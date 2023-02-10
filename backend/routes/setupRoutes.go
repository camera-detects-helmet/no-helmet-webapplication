package routes

import (
	"detection-no-helmet-web-application/api/controller"
	"github.com/gin-gonic/gin"
)

func SetupRoutes(router *gin.Engine) {
	router.POST("/image", controller.CreateImage())
	router.GET("/image", controller.GetAllImage())
	router.GET("/image/:id", controller.GetImageById())
	router.GET("/image/date/:date", controller.GetImageByDate())
	router.GET("/image/dater/:start/:end", controller.GetImageByDateRange())
	router.GET("/image/location/:location", controller.GetImageByLocation())
}
