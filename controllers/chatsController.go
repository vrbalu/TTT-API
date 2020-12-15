package controllers

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type ChatsController struct{}

func (*ChatsController) CreateChat(c *gin.Context) {
	c.JSON(http.StatusOK, "Hello world result")
}

func (*ChatsController) GetMessage(c *gin.Context) {
	c.JSON(http.StatusOK, "Hello world result")
}

func (*ChatsController) SendMessage(c *gin.Context) {
	c.JSON(http.StatusOK, "Hello world result")
}
