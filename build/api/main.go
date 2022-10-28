package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/skkrimon/mc/api/routes"
	"github.com/skkrimon/mc/api/util"
	"log"
	"net/http"
	"os"
)

func main() {
	util.LoadEnv()

	port := os.Getenv("API_PORT")
	ginMode := os.Getenv("GIN_MODE")

	gin.SetMode(ginMode)

	r := gin.Default()
	routes.PingRoutes(r)

	proxyErr := r.SetTrustedProxies(nil)
	if proxyErr != nil {
		log.Fatal(proxyErr)
	}

	r.NoRoute(func(c *gin.Context) {
		c.JSON(http.StatusNotFound, gin.H{
			"status":  http.StatusNotFound,
			"message": "not found",
		})
	})

	srvErr := r.Run(fmt.Sprintf(":%s", port))
	if srvErr != nil {
		log.Fatal(srvErr)
	}
}
