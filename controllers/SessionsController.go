package controllers

import (
	"TTT/mod/models"
	"TTT/mod/services"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

type SessionsController struct{}

func (*SessionsController) CreateSession(c *gin.Context) {
	var auth *models.Auth
	err := c.BindJSON(&auth)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, "Error getting body")
		return
	}
	log.Print(auth)
	authorized := services.UserService.AuthorizeUser(auth)
	if !authorized {
		c.AbortWithStatusJSON(http.StatusUnauthorized, "Wrong e-mail or password")
	}
	user := services.UserService.GetUserByEmail(auth.Email)
	c.JSON(200, gin.H{
		"username": &user.Username,
		"email":    &user.Email,
		"online":   &user.Online,
	})
}

func (*SessionsController) DeleteSession(c *gin.Context) {
	c.JSON(http.StatusOK, "Hello world result")

}
