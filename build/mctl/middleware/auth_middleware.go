package middleware

import (
	"crypto/sha256"
	"encoding/hex"
	"github.com/gin-gonic/gin"
	"github.com/skkrimon/mc/mctl/util"
	"io"
	"net/http"
)

func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		var config util.ConfigYaml
		err := config.LoadConfig()
		if err != nil {
			respondWithError(c, http.StatusInternalServerError, "error validating api key")
			return
		}

		key := c.GetHeader("X-API-KEY")

		if key == "" {
			respondWithError(c, 401, "no api key given")
			return
		}

		h := sha256.New()
		_, err = io.WriteString(h, key)
		if err != nil {
			respondWithError(c, http.StatusInternalServerError, "error validating api key")
			return
		}

		hashedKey := hex.EncodeToString(h.Sum(nil))
		valid := false

		for _, v := range config.Users {
			if v.ApiKey == hashedKey {
				valid = true
			}
		}

		if !valid {
			respondWithError(c, 401, "invalid api key")
			return
		}

		c.Next()
	}
}

func respondWithError(c *gin.Context, code int, message interface{}) {
	c.AbortWithStatusJSON(code, gin.H{
		"success": false,
		"message": message,
	})
}
