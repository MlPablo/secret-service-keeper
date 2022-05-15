package main

import (
	"github.com/gin-gonic/gin"
	"secretservice/handlers"
)

func getRouter() *gin.Engine {
	router := gin.Default()
	router.LoadHTMLFiles("templates/index.html", "templates/key.html")
	router.GET("/", handlers.IndexView)
	router.POST("/", handlers.SaveMessageView)
	return router
}

func main() {
	router := getRouter()
	router.Run(":8080")
}
