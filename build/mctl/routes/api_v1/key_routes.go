package api_v1

import (
	"github.com/gin-gonic/gin"
	"github.com/skkrimon/mc/mctl/controller"
	"github.com/skkrimon/mc/mctl/middleware"
)

func KeyRoutes(r *gin.RouterGroup) {
	key := r.Group("/key", middleware.AuthMiddleware())
	{
		c := new(controller.KeyController)
		key.POST("/generate", c.Generate)
	}
}
