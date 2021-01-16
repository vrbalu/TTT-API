package controllers

import (
	"TTT/mod/services"
	"fmt"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

type TttController struct{}

func (*TttController) Get(c *gin.Context) {
	//res := services.DbService.Func()
	//c.JSON(http.StatusOK, res)

}
func (*TttController) Post(c *gin.Context) {
	var table = "dev.Users"
	sql := fmt.Sprintf("INSERT INTO %s (%s) %s", table, table, table)
	_, err := services.DbService.Exec(sql)
	if err != nil {
		log.Print(err)
	}
	c.JSON(http.StatusOK, "Hello world result")
}

func (*TttController) Register(c *gin.Context) {
	var table = "dev.Users"
	sql := fmt.Sprintf("INSERT INTO %s (%s) %s", table, table, table)
	_, err := services.DbService.Exec(sql)
	if err != nil {
		log.Print(err)
	}
	c.JSON(http.StatusOK, "Hello world result")
}
