package main

import (
	"detection-no-helmet-web-application/api/configs"
	"detection-no-helmet-web-application/api/routes"
	"detection-no-helmet-web-application/api/services"
	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	//database connection
	configs.ConnectDB()

	routes.SetupRoutes(router)
	router.Static("/images", "assets/images")

	services.CheckExistDir("assets")
	err := router.Run(configs.EnvIP() + ":" + configs.EnvPort())

	if err != nil {
		return
	}

}
