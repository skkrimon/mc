package middleware

import (
	"github.com/gin-gonic/gin"
	"github.com/skkrimon/mc/api/util"
	"net/http"
)

func respondWithError(c *gin.Context, code int, message interface{}) {
	c.AbortWithStatusJSON(code, gin.H{
		"success": false,
		"message": message,
	})
}

func AuthMiddleware() gin.HandlerFunc {
	requiredKey := util.GetEnv("API_KEY")

	return func(c *gin.Context) {
		givenKey := c.Query("key")

		if givenKey == "" {
			respondWithError(c, http.StatusUnauthorized, "No API key given")
			return
		}

		if givenKey != requiredKey {
			respondWithError(c, http.StatusUnauthorized, "Invalid API key")
			return
		}

		c.Next()
	}
}
