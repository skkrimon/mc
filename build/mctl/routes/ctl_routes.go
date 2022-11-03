package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/skkrimon/mc/mctl/controller"
)

func CtlRoutes(r *gin.Engine) {
	ctl := r.Group("/ctl")
	{
		c := new(controller.CtlController)
		ctl.POST("/start/:server", c.Start)
		ctl.POST("/stop/:server", c.Stop)
		ctl.POST("/update", c.Update)
	}
}
