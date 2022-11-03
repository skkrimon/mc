package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/skkrimon/mc/mctl/routes/static"
)

func AddStaticRoutes(r *gin.RouterGroup) {
	static.Routes(r)
}
