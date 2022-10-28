package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/skkrimon/mc/api/middleware"
	"github.com/skkrimon/mc/api/routes"
	"github.com/skkrimon/mc/api/util"
	"log"
	"net/http"
)

func main() {
	port := util.GetEnv("API_PORT")
	ginMode := util.GetEnv("GIN_MODE")

	gin.SetMode(ginMode)

	app := gin.Default()
	app.Use(middleware.AuthMiddleware())

	router := app.Group("/api/v1")
	routes.AddRoutes(router)

	err := app.SetTrustedProxies(nil)
	if err != nil {
		log.Fatal(err)
	}

	app.NoRoute(func(c *gin.Context) {
		c.JSON(http.StatusNotFound, gin.H{
			"status":  http.StatusNotFound,
			"message": "not found",
		})
	})

	err = app.Run(fmt.Sprintf(":%s", port))
	if err != nil {
		log.Fatal(err)
	}
}
