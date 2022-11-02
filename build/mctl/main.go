package main

import (
	"fmt"
	"github.com/gin-gonic/autotls"
	"github.com/gin-gonic/gin"
	"github.com/skkrimon/mc/mctl/middleware"
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
	r.Use(middleware.AuthMiddleware())
	routes.CtlRoutes(r)
	routes.KeyRoutes(r)

	proxyErr := r.SetTrustedProxies(nil)
	if proxyErr != nil {
		log.Fatal(proxyErr)
	}

	if config.GinMode == "debug" {
		log.Fatal(r.Run(fmt.Sprintf(":%s", config.Port)))
	} else {
		log.Fatal(autotls.Run(r, "og-in-nbg.de", "api.og-in-nbg.de"))
	}
}
