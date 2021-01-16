package controllers

import (
	"TTT/mod/helpers"
	"TTT/mod/models"
	"TTT/mod/services"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

type UsersController struct{}

var userService services.UserServiceType

func (*UsersController) CreateChat(c *gin.Context) {
	c.JSON(http.StatusOK, "Hello world result")
}

func (*UsersController) CreateUser(c *gin.Context) {
	var user *models.User
	err := c.BindJSON(&user)
	if err != nil {
		c.AbortWithStatusJSON(500, "Failed binding.")
	}
	user.Password, err = helpers.HashPassword(user.Password)
	if err != nil {
		c.AbortWithStatusJSON(500, "Internal server error.")
	}
	err = userService.RegisterUserViaWeb(user)
	if err != nil {
		c.AbortWithStatusJSON(500, "Failed upload to DB.")
	}
	c.JSON(http.StatusOK, "")
}
func (*UsersController) CreateUserWithGoogle(c *gin.Context) {
	var gtrm models.GoogleTokenResponseModel
	err := c.BindJSON(&gtrm)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, "Error getting body")
		return
	}
	fmt.Printf("Body %s", gtrm.UserData.ID)
	user := models.User{
		Username: gtrm.UserData.Name + gtrm.UserData.ID[len(gtrm.UserData.ID)-5:],
		Email:    gtrm.UserData.Email,
		ExtID:    gtrm.UserData.ID,
		IDToken:  gtrm.UserData.IDToken,
		Online:   false,
	}
	_, existsUser := userService.CheckUserExists(user.Email)
	if !existsUser {
		err = userService.RegisterUserViaGoogle(&user)
		if err != nil {
			c.AbortWithStatusJSON(http.StatusInternalServerError, fmt.Sprintf("Registration not successfull: %s", err))
			return
		}
	}
	/*valid, err := validateToken(gtrm.UserData.IDToken)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, fmt.Sprintf("Error validating token: %s", err))
		return
	}
	if !valid {
		c.JSON(http.StatusUnauthorized,"Unauthorized. Token not valid.")
	}*/
	c.JSON(http.StatusOK, gin.H{
		"username": user.Username,
		"email":    user.Email,
		"idToken":  user.IDToken,
		"online":   user.Online,
	})
}
func (*UsersController) DeleteUser(c *gin.Context) {
	c.JSON(http.StatusOK, "Hello world result")
}

func (*UsersController) GetAllUsers(c *gin.Context) {
	c.JSON(http.StatusOK, "Hello world result")
}

func (*UsersController) GetUserByName(c *gin.Context) {
	c.JSON(http.StatusOK, "Hello world result")
}

func (*UsersController) UpdateUser(c *gin.Context) {
	c.JSON(http.StatusOK, "Hello world result")
}
