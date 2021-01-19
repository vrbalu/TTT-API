package controllers

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type GamesController struct{}

func (*GamesController) CreateGame(c *gin.Context) {

}

func (*GamesController) DeleteGame(c *gin.Context) {
	c.JSON(http.StatusOK, "Hello world result")
}

func (*GamesController) GetAllGames(c *gin.Context) {
	c.JSON(http.StatusOK, "Hello world result")
}

func (*GamesController) GetGamePlayerInfo(c *gin.Context) {
	c.JSON(http.StatusOK, "Hello world result")
}

func (*GamesController) GetGamePlayers(c *gin.Context) {
	c.JSON(http.StatusOK, "Hello world result")
}

func (*GamesController) GetOneGame(c *gin.Context) {
	c.JSON(http.StatusOK, "Hello world result")
}

func (*GamesController) PlayMove(c *gin.Context) {
	c.JSON(http.StatusOK, "Hello world result")
}

func (*GamesController) UpdateGame(c *gin.Context) {
	c.JSON(http.StatusOK, "Hello world result")
}
