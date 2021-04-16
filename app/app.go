package app

import (
	"github.com/bee-well/auth/mq"
	"github.com/gin-gonic/gin"
)

func Start() {
	e := gin.Default()
	mapRoutes(e)

	m := mq.NewMq()
	mapMqHandlers(m)

	if err := e.Run(":8080"); err != nil {
		panic(err)
	}
}
