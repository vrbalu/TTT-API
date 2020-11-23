package helpers

import (
	"TTT/mod/controllers"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/toorop/gin-logrus"
	"time"
)

func SetupRouter() *gin.Engine {
	router := gin.New()
	router.Use(ginlogrus.Logger(Log), gin.Recovery()) //Setup logging and panic recovery
	// CORS setup
	router.Use(cors.New(cors.Config{
		AllowAllOrigins:  true,
		AllowMethods:     []string{"GET", "POST"},
		AllowHeaders:     []string{"Origin", "Authorization"},
		AllowCredentials: false,
		MaxAge:           72 * time.Hour,
	}))

	// API routes
	apiGroup := router.Group("/api")
	{
		tttController := controllers.TttController{}

		apiGroup.GET("/ttt", tttController.Get)
	}
	return router
}
