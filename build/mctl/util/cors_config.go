package util

import (
	"github.com/gin-contrib/cors"
	"time"
)

func CorsConfig() cors.Config {
	return cors.Config{
		AllowMethods:     []string{"GET", "POST", "DELETE"},
		AllowHeaders:     []string{"Origin", "Content-Length", "Content-Type", "X-API-KEY"},
		AllowCredentials: false,
		AllowAllOrigins:  true,
		MaxAge:           12 * time.Hour,
	}
}
