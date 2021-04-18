package app

import (
	"fmt"

	"github.com/bee-well/auth/config"
	"github.com/bee-well/auth/mq"
	"github.com/gin-gonic/gin"
)

func Start() {
	e := gin.Default()
	mapRoutes(e)

	m := mq.NewMq()
	mapMqHandlers(m)

	port := fmt.Sprintf(":%s", config.GetString("PORT"))
	if err := e.Run(port); err != nil {
		panic(err)
	}
}
