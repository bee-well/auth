package app

import (
	"github.com/bee-well/auth/controllers"
	"github.com/bee-well/auth/mq"
)

func mapMqHandlers(m mq.Mq) {
	m.AttachHandler("users", controllers.OnExternalAuthentication)
}
