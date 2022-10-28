package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"github.com/skkrimon/mc/mctl/routes"
)

func main() {
	envErr := godotenv.Load()
	if envErr != nil {
		log.Fatal("Error loading .env file")
	}

	port := os.Getenv("MCTL_PORT")
	ginMode := os.Getenv("GIN_MODE")

	gin.SetMode(ginMode)

	r := gin.Default()
	routes.CtlRoutes(r)

	r.NoRoute(func(c *gin.Context) {
		c.JSON(http.StatusNotFound, gin.H{
			"message": "not found",
		})
	})

	proxyErr := r.SetTrustedProxies([]string{"127.0.0.1"})
	if proxyErr != nil {
		log.Fatal(proxyErr)
	}

	srvErr := r.Run(fmt.Sprintf(":%s", port))
	if srvErr != nil {
		log.Fatal(srvErr)
	}
}
