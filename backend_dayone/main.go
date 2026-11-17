package main

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	r.GET("/ping", pingController)

	r.Run(":8080")

	fmt.Println("server is running on port 8080")
}

func pingController(context *gin.Context) {
	context.JSON(200, gin.H{
		"message": "hello",
	})
}
