package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"log"
	"net/http"
	"os"
)

func main() {
	envErr := godotenv.Load()
	if envErr != nil {
		log.Fatal("Error loading .env file")
	}

	port := os.Getenv("API_PORT")
	ginMode := os.Getenv("GIN_MODE")

	gin.SetMode(ginMode)

	r := gin.Default()
	proxyErr := r.SetTrustedProxies(nil)
	if proxyErr != nil {
		log.Fatal(proxyErr)
	}

	r.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"ping": "pong",
		})
	})

	srvErr := r.Run(fmt.Sprintf(":%s", port))
	if srvErr != nil {
		log.Fatal(srvErr)
	}
}
