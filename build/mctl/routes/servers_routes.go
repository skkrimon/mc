package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/skkrimon/mc/mctl/controller"
)

func ServerRoutes(r *gin.Engine) {
	servers := r.Group("/servers")
	{
		c := new(controller.ServersController)
		servers.GET("", c.GetServers)
	}
}
