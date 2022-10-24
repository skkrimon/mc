package controller

import (
	"net/http"
	"os"
	"os/exec"

	"github.com/gin-gonic/gin"
)

type UpdateController struct{}

func (h *UpdateController) Index(c *gin.Context) {
	go update()

	c.JSON(http.StatusOK, gin.H{
		"success": true,
		"message": "server update was triggered",
	})
}

func update() {
	updatePath := os.Getenv("UPDATE_PATH")

	cmd := exec.Command("python3", "update")
	cmd.Dir = updatePath
	cmd.Run()
}