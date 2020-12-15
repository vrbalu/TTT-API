package controllers

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type UsersController struct{}

func (*UsersController) CreateChat(c *gin.Context) {
	c.JSON(http.StatusOK, "Hello world result")
}

func (*UsersController) CreateUser(c *gin.Context) {
	c.JSON(http.StatusOK, "Hello world result")
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
