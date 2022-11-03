package controller

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/skkrimon/mc/mctl/util"
	"net/http"
	"os/exec"
)

type CtlController struct{}

func (h *CtlController) Start(c *gin.Context) {
	server := c.Param("server")
	isValid := validateServer(server)

	if !isValid {
		c.JSON(http.StatusNotFound, gin.H{
			"success": false,
			"message": fmt.Sprintf("%s server not found", server),
		})
		return
	}

	out, err := exec.Command("systemctl", "start", fmt.Sprintf("minecraft@%s", server)).Output()

	if err != nil {
		util.ErrorResponse(c, err.Error())
		return
	}

	handleSuccess(c, string(out))
}

func (h *CtlController) Stop(c *gin.Context) {
	server := c.Param("server")
	out, err := exec.Command("systemctl", "stop", fmt.Sprintf("minecraft@%s", server)).Output()
	isValid := validateServer(server)

	if !isValid {
		c.JSON(http.StatusNotFound, gin.H{
			"success": false,
			"message": fmt.Sprintf("%s server not found", server),
		})
		return
	}

	if err != nil {
		util.ErrorResponse(c, err.Error())
		return
	}

	handleSuccess(c, string(out))
}

func (h *CtlController) Update(c *gin.Context) {
	go updateServer()

	handleSuccess(c, "server update was triggered")
}

func handleSuccess(c *gin.Context, message string) {
	c.JSON(http.StatusOK, gin.H{
		"success": true,
		"message": message,
	})
}

func updateServer() {
	var config util.ConfigYaml
	err := config.LoadConfig()
	if err != nil {
		return
	}

	cmd := exec.Command("python3", "update")
	cmd.Dir = config.UpdatePath
	err = cmd.Run()
	if err != nil {
		fmt.Println("Could not update server")
	}
}

func validateServer(server string) bool {
	var config util.ConfigYaml
	err := config.LoadConfig()
	if err != nil {
		return false
	}

	for _, v := range config.Servers {
		if v.Name == server {
			return true
		}
	}

	return false
}
