package controllers

import (
	"TTT/mod/services"
	"github.com/gin-gonic/gin"
)

type TttController struct{}

func (*TttController) Get(c *gin.Context) {
	services.TttService.DoSmth()
}
