package controller

import (
	"net/http"
	"os/exec"

	"github.com/gin-gonic/gin"
)

type StopController struct{}

func (h *StopController) Index(c *gin.Context) {
	out, err := exec.Command("systemctl", "stop", "minecraft").Output()

	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"success": false,
			"message": err.Error(),
		})

		return
	}

	c.JSON(http.StatusOK, gin.H{
		"success": true,
		"message": string(out),
	})
}