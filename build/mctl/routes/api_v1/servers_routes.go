package api_v1

import (
	"github.com/gin-gonic/gin"
	"github.com/skkrimon/mc/mctl/controller"
	"github.com/skkrimon/mc/mctl/middleware"
)

func ServerRoutes(r *gin.RouterGroup) {
	servers := r.Group("/servers", middleware.AuthMiddleware())
	{
		c := new(controller.ServersController)
		servers.GET("", c.GetServers)
	}
}
