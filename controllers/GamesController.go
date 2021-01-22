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
		c.AbortWithStatusJSON(http.StatusBadRequest, "Error parsing body")
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

func (*GamesController) GetGamesStats(c *gin.Context) {
	var gameStats []models.GameStatsModel
	gameStats, err := services.GameService.GetGameStats()
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, "Error creating friendship"+err.Error())
		return
	}
	c.JSON(http.StatusOK, gameStats)
}

func (*GamesController) UpdateGame(c *gin.Context) {
	var updateGame models.GameUpdate
	err := c.BindJSON(&updateGame)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, "Error parsing body")
		return
	}
	err = services.GameService.UpdateGame(updateGame.Id, updateGame.Winner, updateGame.IsPending, updateGame.IsFinished)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, "Error parsing body")
		return
	}
	c.JSON(http.StatusOK, "")

}
