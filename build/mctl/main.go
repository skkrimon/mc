package main

import (
	"fmt"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/autotls"
	"github.com/gin-gonic/gin"
	"github.com/skkrimon/mc/mctl/routes"
	"github.com/skkrimon/mc/mctl/util"
	"log"
)

func main() {
	var config util.ConfigYaml
	err := config.LoadConfig()
	if err != nil {
		log.Fatal(err)
	}

	gin.SetMode(config.GinMode)

	r := gin.Default()
	r.LoadHTMLGlob("./static/*.html")
	r.Use(cors.New(util.CorsConfig()))

	apiRoutes := r.Group("/api/v1")
	staticRoutes := r.Group("/")

	routes.AddApiV1Routes(apiRoutes)
	routes.AddStaticRoutes(staticRoutes)

	proxyErr := r.SetTrustedProxies(nil)
	if proxyErr != nil {
		log.Fatal(proxyErr)
	}

	if config.GinMode == "release" {
		log.Fatal(autotls.Run(r, "og-in-nbg.de", "api.og-in-nbg.de"))
	} else {
		log.Fatal(r.Run(fmt.Sprintf(":%s", config.Port)))
	}
}
