package controllers

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type SessionsController struct{}

func (*SessionsController) CreateSession(c *gin.Context) {
	c.JSON(http.StatusOK, "Hello world result")
}

func (*SessionsController) DeleteSession(c *gin.Context) {
	c.JSON(http.StatusOK, "Hello world result")

}
