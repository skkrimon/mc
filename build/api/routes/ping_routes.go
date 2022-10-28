package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/skkrimon/mc/api/controller"
)

func PingRoutes(superRoute *gin.RouterGroup) {
	ping := superRoute.Group("/ping")
	{
		c := new(controller.PingController)

		ping.GET("/", c.Index)
	}
}
