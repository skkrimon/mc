package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/skkrimon/mc/mctl/routes/api_v1"
)

func AddApiV1Routes(r *gin.RouterGroup) {
	api_v1.CtlRoutes(r)
	api_v1.ServerRoutes(r)
	api_v1.KeyRoutes(r)
}
