package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/skkrimon/mc/api/controller"
)

func PingRoutes(r *gin.Engine) {
	ping := r.Group("/ping")
	{
		pong := new(controller.PingController)
		ping.GET("/", pong.Index)
	}
}
