package main

import (
	"DemoHttpMock2/handlers"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	router.GET("/", handlers.RequestWrapper(handlers.GetCustomData))

	router.POST("/user", handlers.RequestWrapper(handlers.CreateUser))

	router.Run(":8080")
}
