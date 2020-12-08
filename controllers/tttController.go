package controllers

import (
	"TTT/mod/services"
	"github.com/gin-gonic/gin"
)

type TttController struct{}

func (*TttController) Get(c *gin.Context) {
	res := services.DbService.Func()
	c.JSON(200, res)
}
func (*TttController) Post(c *gin.Context) {
	services.TttService.DoSmth()
}
