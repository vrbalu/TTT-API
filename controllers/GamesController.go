package controllers

import (
	"TTT/mod/models"
	"TTT/mod/services"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

type GamesController struct{}

func (*GamesController) CreateGame(c *gin.Context) {
	var newGame models.CreateGame
	err := c.BindJSON(&newGame)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, "Error parsing body")
		return
	}
	id, err := services.GameService.CreateGame(&newGame)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, "Error creating friendship"+err.Error())
		return
	}
	notificationHub.Broadcast <- []byte(fmt.Sprintf(`{
		"id": %v,
		"user1": "%s",
		"user2": "%s",
		"isGameOn": false
	}`, id, newGame.User1, newGame.User2))
	c.JSON(http.StatusCreated, "")
}

func (*GamesController) GetAllGames(c *gin.Context) {
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
