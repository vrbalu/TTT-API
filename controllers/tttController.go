package controllers

import (
	"TTT/mod/services"
	"github.com/gin-gonic/gin"
	"net/http"
)

type TttController struct{}

func (*TttController) Get(c *gin.Context) {
	res := services.DbService.Func()
	c.JSON(http.StatusOK, res)

}
func (*TttController) Post(c *gin.Context) {
	services.TttService.DoSmth()
	c.JSON(http.StatusOK, "Hello world result")
}
