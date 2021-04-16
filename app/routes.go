package app

import (
	"github.com/bee-well/auth/controllers"
	"github.com/gin-gonic/gin"
)

func mapRoutes(e *gin.Engine) {
	e.GET("/marco", controllers.Marco)
	e.POST("/sign-in", controllers.SignIn)
}
