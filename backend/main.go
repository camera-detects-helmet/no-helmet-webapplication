package main

import (
	"detection-no-helmet-web-application/api/configs"
	"detection-no-helmet-web-application/api/routes"
	"detection-no-helmet-web-application/api/services"
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func main() {

	router := gin.Default()
	router.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"*"},
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "PATCH", "OPTIONS"},
		AllowHeaders:     []string{"Origin", "Content-Length", "Content-Type", "Authorization"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
		MaxAge:           12 * time.Hour,
	}))

	//database connection
	configs.ConnectDB()

	routes.SetupRoutes(router)
	router.Static("/images", "assets/images")

	services.CheckExistDir("assets")
	err := router.Run(":8080")

	if err != nil {
		return
	}

}
