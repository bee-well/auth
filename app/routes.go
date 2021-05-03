package app

import (
	"github.com/bee-well/auth/controllers"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func mapRoutes(e *gin.Engine) {
	e.Use(cors.Default())

	e.GET("/marco", controllers.Marco)
	e.POST("/sign-in", controllers.SignIn)
	e.POST("/sign-up", controllers.SignUp)
	e.GET("/me", controllers.GetUser)
}
