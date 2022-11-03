package static

import (
	"github.com/gin-gonic/gin"
	"github.com/skkrimon/mc/mctl/controller"
)

func Routes(r *gin.RouterGroup) {
	static := r.Group("/")
	{
		c := new(controller.StaticController)

		static.GET("/docs", c.Swagger)
		static.StaticFile("/swagger.json", "./static/swagger.json")
	}
}
