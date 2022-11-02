package main

import (
	"fmt"
	"github.com/skkrimon/mc/mctl/util"
	"log"

	"github.com/gin-gonic/gin"
	"github.com/skkrimon/mc/mctl/routes"
)

func main() {
	port := util.GetEnv("MCTL_PORT")
	ginMode := util.GetEnv("GIN_MODE")

	gin.SetMode(ginMode)

	r := gin.Default()
	routes.CtlRoutes(r)

	proxyErr := r.SetTrustedProxies(nil)
	if proxyErr != nil {
		log.Fatal(proxyErr)
	}

	srvErr := r.Run(fmt.Sprintf(":%s", port))
	if srvErr != nil {
		log.Fatal(srvErr)
	}
}
