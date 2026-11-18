package main

import (
	"fmt"

	"github.com/Yonis18/users_api/handlers"
	"github.com/gin-gonic/gin"
)

func main() {

	r := gin.Default()

	r.POST("/users/create", handlers.CreateUser)
	r.GET("/users/list", handlers.ListUsers)
	r.Run(":8080")
	fmt.Println("server is running on port 8080🌧️")
}
