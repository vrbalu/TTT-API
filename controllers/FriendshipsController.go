package controllers

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type FriendshipsController struct{}

func (*FriendshipsController) CreateFriendship(c *gin.Context) {
	c.JSON(http.StatusOK, "Hello world result")
}

func (*FriendshipsController) DeleteFriendship(c *gin.Context) {
	c.JSON(http.StatusOK, "Hello world result")
}

func (*FriendshipsController) GetAllFriendships(c *gin.Context) {
	c.JSON(http.StatusOK, "Hello world result")
}

func (*FriendshipsController) GetFriendship(c *gin.Context) {
	c.JSON(http.StatusOK, "Hello world result")
}

func (*FriendshipsController) UpdateFriendship(c *gin.Context) {
	c.JSON(http.StatusOK, "Hello world result")
}
