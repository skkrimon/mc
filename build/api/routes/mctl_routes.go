package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/skkrimon/mc/api/controller"
)

func MctlRoutes(superRoute *gin.RouterGroup) {
	mctl := superRoute.Group("/mctl")
	{
		c := new(controller.MctlController)

		mctl.POST("/start", c.Start)
		mctl.POST("/stop", c.Stop)
		mctl.POST("/restart", c.Restart)
		mctl.POST("/update", c.Update)
	}
}
