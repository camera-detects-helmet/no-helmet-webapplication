package main

import (
	"detection-no-helmet-web-application/api/configs"
	"detection-no-helmet-web-application/api/routes"
	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	//database connection
	configs.ConnectDB()

	routes.SetupRoutes(router)

	err := router.Run(":8000")
	if err != nil {
		return
	}

}
