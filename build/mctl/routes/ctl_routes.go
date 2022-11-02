package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/skkrimon/mc/mctl/controller"
)

func CtlRoutes(r *gin.Engine) {
	ctl := r.Group("/ctl")
	{
		c := new(controller.CtlController)
		ctl.POST("/start", c.Start)
		ctl.POST("/stop", c.Stop)
		ctl.POST("/update", c.Update)
	}
}
