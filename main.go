package main

import (
	"github.com/gin-gonic/gin"
	"number-classifier/controller"
)

func main() {
	router := gin.Default()
	router.GET("/api/classify-number", controller.ClassifyNumber)
	router.Run(":8080")
}
