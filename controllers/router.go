package controllers

import (
	"TTT/mod/helpers"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/toorop/gin-logrus"
	"time"
)

func SetupRouter() *gin.Engine {
	router := gin.New()
	router.Use(ginlogrus.Logger(helpers.Log), gin.Recovery()) //Setup logging and panic recovery
	// CORS setup
	router.Use(cors.New(cors.Config{
		AllowAllOrigins:  true,
		AllowMethods:     []string{"GET", "POST", "OPTIONS", "PUT", "PATCH"},
		AllowHeaders:     []string{"Origin", "Authorization", "content-type"},
		AllowCredentials: false,
		MaxAge:           72 * time.Hour,
	}))

	// API routes
	apiGroup := router.Group("/api")
	{
		tttController := TttController{}
		userController := UsersController{}
		sessionsController := SessionsController{}
		friendshipsController := FriendshipsController{}
		apiGroup.GET("/users", userController.GetAllUsers)
		apiGroup.POST("/users", userController.CreateUser)
		apiGroup.PUT("/users", userController.UpdateUser)

		apiGroup.POST("/sessions", sessionsController.CreateSession)

		apiGroup.GET("/friendships", friendshipsController.GetFriendships)
		apiGroup.POST("/friendships", friendshipsController.CreateFriendship)
		apiGroup.PUT("/friendships", friendshipsController.UpdateFriendship)
		apiGroup.DELETE("/friendships", friendshipsController.DeleteFriendship)
		apiGroup.GET("/friendships/:id", friendshipsController.GetFriendshipById)

		apiGroup.GET("/ttt", tttController.Get)
		apiGroup.POST("/ttt", tttController.Post)
		apiGroup.POST("/callback", userController.CreateUserWithGoogle)

	}
	return router
}
