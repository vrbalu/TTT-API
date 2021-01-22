package controllers

import (
	"TTT/mod/helpers"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/toorop/gin-logrus"
	"time"
)

var notificationHub = helpers.NewHub()
var gameHub = helpers.NewHub()
var statusHub = helpers.NewHub()

func SetupRouter() *gin.Engine {
	router := gin.New()
	// Go routines for websockets
	go notificationHub.Run()
	go gameHub.Run()
	go statusHub.Run()
	router.Use(ginlogrus.Logger(helpers.Log), gin.Recovery()) //Setup logging and panic recovery
	// CORS setup
	router.Use(cors.New(cors.Config{
		AllowAllOrigins:  true,
		AllowMethods:     []string{"GET", "POST", "OPTIONS", "PUT", "PATCH", "DELETE"},
		AllowHeaders:     []string{"Origin", "Authorization", "content-type"},
		AllowCredentials: false,
		MaxAge:           72 * time.Hour,
	}))

	// API routes
	apiGroup := router.Group("/api")
	{
		userController := UsersController{}
		sessionsController := SessionsController{}
		friendshipsController := FriendshipsController{}
		gamesController := GamesController{}
		// Users endpoints
		apiGroup.GET("/users", userController.GetAllUsers)
		apiGroup.POST("/users", userController.CreateUser)
		apiGroup.PUT("/users", userController.UpdateUser)
		// Login endpoints
		apiGroup.POST("/sessions", sessionsController.CreateSession)
		apiGroup.POST("/callback", userController.CreateUserWithGoogle)
		//Friendships endpoints
		apiGroup.GET("/friendships", friendshipsController.GetFriendships)
		apiGroup.POST("/friendships", friendshipsController.CreateFriendship)
		apiGroup.PUT("/friendships/:id", friendshipsController.UpdateFriendship)
		apiGroup.DELETE("/friendships/:id", friendshipsController.DeleteFriendship)
		apiGroup.GET("/friendships/:id", friendshipsController.GetFriendshipById)
		// Games endpoints
		apiGroup.POST("/games", gamesController.CreateGame)
		apiGroup.GET("/games/stats", gamesController.GetGamesStats)
		apiGroup.PUT("/games", gamesController.UpdateGame)
		// Websockets
		apiGroup.GET("/notify", func(c *gin.Context) {
			helpers.ServeWs(notificationHub, c.Writer, c.Request)
		})
		apiGroup.GET("/gaming", func(c *gin.Context) {
			helpers.ServeWs(gameHub, c.Writer, c.Request)
		})
		apiGroup.GET("/status", func(c *gin.Context) {
			helpers.ServeWs(statusHub, c.Writer, c.Request)
		})
	}
	return router
}
