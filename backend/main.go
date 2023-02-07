package main

import (
	"detection-no-helmet-web-application/api/configs"
	"detection-no-helmet-web-application/api/routes"
	"detection-no-helmet-web-application/api/services"
	"github.com/gin-gonic/gin"
)

func CORSMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {

		c.Header("Access-Control-Allow-Origin", "*")
		c.Header("Access-Control-Allow-Credentials", "true")
		c.Header("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With")
		c.Header("Access-Control-Allow-Methods", "POST,HEAD,PATCH, OPTIONS, GET, PUT")

		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(204)
			return
		}

		c.Next()
	}
}

func main() {

	router := gin.Default()
	router.Use(CORSMiddleware())
	//database connection
	configs.ConnectDB()

	routes.SetupRoutes(router)
	router.Static("/images", "assets/images")

	services.CheckExistDir("assets")
	err := router.Run(configs.GetLocalIP() + ":" + configs.EnvPort())

	if err != nil {
		return
	}

}
