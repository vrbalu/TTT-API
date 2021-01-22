package controllers

import (
	"TTT/mod/models"
	"TTT/mod/services"
	"fmt"
	"github.com/gin-gonic/gin"
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
	authorized := services.UserService.AuthorizeUser(auth)
	if !authorized {
		c.AbortWithStatusJSON(http.StatusUnauthorized, "Wrong e-mail or password")
		return
	}
	user := services.UserService.GetUserByField("Email", auth.Email)
	err = services.UserService.UpdateStatus(user.Username, false, true)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, "Error updating")
		return
	}
	statusHub.Broadcast <- []byte(fmt.Sprintf(`{
		"username": "%s",
		"online": %v,
		"inGame": %v
		}`, user.Username, user.Online, user.InGame))
	c.JSON(200, gin.H{
		"username": &user.Username,
		"email":    &user.Email,
		"online":   &user.Online,
	})
}
