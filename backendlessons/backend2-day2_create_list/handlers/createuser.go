package handlers

import (
	"net/http"

	"github.com/Yonis18/users_api/services"
	"github.com/gin-gonic/gin"
)

type CreateUserRequest struct {
	Username string `json:"username"`
	Email    string `json:"email"`
	Phone    int    `json:"phone"`
}

func CreateUser(c *gin.Context) {
	var req CreateUserRequest

	err := c.ShouldBindJSON(&req)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	user := services.CreateUser(req.Username, req.Email, req.Phone)

	c.JSON(http.StatusCreated, gin.H{
		"user": user,
	})
}

func ListUsers(c *gin.Context) {
	users := services.List()

	c.JSON(http.StatusOK, gin.H{
		"users": users,
	})
}
