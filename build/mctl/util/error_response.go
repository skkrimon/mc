package util

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func ErrorResponse(c *gin.Context, error string) {
	c.JSON(http.StatusInternalServerError, gin.H{
		"success": false,
		"message": error,
	})
}
