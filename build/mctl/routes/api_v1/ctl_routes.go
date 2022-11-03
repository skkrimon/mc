package api_v1

import (
	"github.com/gin-gonic/gin"
	"github.com/skkrimon/mc/mctl/controller"
	"github.com/skkrimon/mc/mctl/middleware"
)

func CtlRoutes(r *gin.RouterGroup) {
	ctl := r.Group("/ctl", middleware.AuthMiddleware())
	{
		c := new(controller.CtlController)
		ctl.POST("/start/:server", c.Start)
		ctl.POST("/stop/:server", c.Stop)
		ctl.POST("/update", c.Update)
	}
}
