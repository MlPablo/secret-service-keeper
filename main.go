package main

import (
	"github.com/gin-gonic/gin"
	"secretservice/handlers"
)

func getRouter() *gin.Engine {
	router := gin.Default()
	router.LoadHTMLGlob("templates/*")
	router.GET("/", handlers.IndexView)
	router.POST("/", handlers.SaveMessageView)
	router.GET("/:key", handlers.ReadMessageHandler)
	return router
}

func main() {
	router := getRouter()
	router.Run(":8080")
}
