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
		apiGroup.GET("/ttt", tttController.Get)
		apiGroup.POST("/users", userController.CreateUser)
		apiGroup.POST("/ttt", tttController.Post)
		apiGroup.POST("/callback", userController.CreateUserWithGoogle)
		apiGroup.POST("/sessions", sessionsController.CreateSession)

	}
	return router
}
