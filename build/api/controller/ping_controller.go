package controller

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type PingController struct{}

func (h *PingController) Index(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"ping": "pong",
	})
}
