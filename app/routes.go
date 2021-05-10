package app

import (
	"time"

	"github.com/bee-well/auth/controllers"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func mapRoutes(e *gin.Engine) {
	corsConfig := cors.Config{
		AllowMethods:     []string{"GET", "POST", "PUT", "PATCH", "DELETE", "HEAD"},
		AllowHeaders:     []string{"Origin", "Content-Length", "Content-Type", "Authorization"},
		AllowCredentials: false,
		MaxAge:           12 * time.Hour,
	}
	corsConfig.AllowAllOrigins = true
	e.Use(cors.New(corsConfig))

	e.GET("/marco", controllers.Marco)
	e.POST("/sign-in", controllers.SignIn)
	e.POST("/sign-up", controllers.SignUp)
	e.GET("/me", controllers.GetUser)
	e.GET("/users", controllers.CountUsers)
}
