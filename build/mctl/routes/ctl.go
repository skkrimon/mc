package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/skkrimon/mctl/controller"
)

func CtlRoutes(r *gin.Engine)  {
	ctl := r.Group("/ctl") 
	{
		start := new(controller.StartController)
		ctl.POST("/start", start.Index)

		stop := new(controller.StopController)
		ctl.POST("/stop", stop.Index)

		update := new(controller.UpdateController)
		ctl.POST("update", update.Index)
	}
}
