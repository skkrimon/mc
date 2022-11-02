package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/skkrimon/mc/mctl/controller"
)

func KeyRoutes(r *gin.Engine) {
	key := r.Group("/key")
	{
		c := new(controller.KeyController)
		key.POST("/generate", c.Generate)
	}
}
