package controllers

import (
	"TTT/mod/helpers"
	"TTT/mod/models"
	"TTT/mod/services"
	"fmt"
	"github.com/gin-gonic/gin"
	"log"
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
		Username:            gtrm.UserData.Name + gtrm.UserData.ID[len(gtrm.UserData.ID)-5:],
		Email:               gtrm.UserData.Email,
		ExtID:               gtrm.UserData.ID,
		InGame:              false,
		Online:              true,
		RegisteredViaGoogle: true,
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
	err = services.UserService.UpdateStatus(user.Email, "Online", true)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, "Error updating")
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"username":            user.Username,
		"email":               user.Email,
		"inGame":              user.InGame,
		"online":              user.Online,
		"registeredViaGoogle": user.RegisteredViaGoogle,
	})
}
func (*UsersController) DeleteUser(c *gin.Context) {
	c.JSON(http.StatusOK, "Hello world result")
}

func (*UsersController) GetAllUsers(c *gin.Context) {
	online := c.Query("online")
	if online != "" {
		users, err := services.UserService.GetAllUsers(true)
		if err != nil {
			c.AbortWithStatusJSON(http.StatusInternalServerError, "Error updating")
			return
		}
		c.JSON(http.StatusOK, users)
	} else {
		users, err := services.UserService.GetAllUsers(false)
		if err != nil {
			c.AbortWithStatusJSON(http.StatusInternalServerError, "Error updating")
			return
		}
		c.JSON(http.StatusOK, users)
	}
}

func (*UsersController) GetUserByName(c *gin.Context) {
	c.JSON(http.StatusOK, "Hello world result")
}

func (*UsersController) UpdateUser(c *gin.Context) {
	email := c.Query("email")
	operationType := c.Query("type")
	log.Print(email)
	log.Print(operationType)
	if operationType == "status" {
		var status *models.UpdateStatus
		err := c.BindJSON(&status)
		if err != nil {
			c.AbortWithStatusJSON(http.StatusInternalServerError, "Error getting body")
			return
		}
		log.Print(status.Online)
		err = userService.UpdateStatus(email, "Online", status.Online)
		if err != nil {
			c.AbortWithStatusJSON(http.StatusInternalServerError, err.Error())
			return
		}
		c.JSON(http.StatusOK, "")

	}
	if operationType == "password" {
		var passwordUpdate *models.UpdatePassword
		var authModel models.Auth
		err := c.BindJSON(&passwordUpdate)
		if err != nil {
			c.AbortWithStatusJSON(http.StatusInternalServerError, "Error getting body")
			return
		}
		authModel.Email = email
		authModel.Password = passwordUpdate.OldPassword
		isOldPasswordCorrect := userService.AuthorizeUser(&authModel)
		if isOldPasswordCorrect {
			passwordUpdate.Password, err = helpers.HashPassword(passwordUpdate.Password)
			if err != nil {
				c.AbortWithStatusJSON(500, "Internal server error.")
			}
			err = userService.UpdatePassword(email, passwordUpdate.Password)
			if err != nil {
				c.AbortWithStatusJSON(http.StatusInternalServerError, "Error inserting into DB")
				return
			}
			c.JSON(http.StatusOK, "")
		} else {
			c.AbortWithStatusJSON(http.StatusUnauthorized, "Wrong old password.")
		}

	}
}
