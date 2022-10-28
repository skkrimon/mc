package main

import (
	"fmt"
	"github.com/skkrimon/mc/mctl/util"
	"log"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/skkrimon/mc/mctl/routes"
)

func main() {
	util.LoadEnv()

	port := os.Getenv("MCTL_PORT")
	ginMode := os.Getenv("GIN_MODE")

	gin.SetMode(ginMode)

	r := gin.Default()
	routes.CtlRoutes(r)

	r.NoRoute(func(c *gin.Context) {
		c.JSON(http.StatusNotFound, gin.H{
			"status":  http.StatusNotFound,
			"message": "not found",
		})
	})

	proxyErr := r.SetTrustedProxies(nil)
	if proxyErr != nil {
		log.Fatal(proxyErr)
	}

	srvErr := r.Run(fmt.Sprintf(":%s", port))
	if srvErr != nil {
		log.Fatal(srvErr)
	}
}
