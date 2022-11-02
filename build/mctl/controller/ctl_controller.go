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
	out, err := exec.Command("systemctl", "start", "minecraft@server").Output()

	if err != nil {
		util.ErrorResponse(c, err.Error())
		return
	}

	handleSuccess(c, string(out))
}

func (h *CtlController) Stop(c *gin.Context) {
	out, err := exec.Command("systemctl", "stop", "minecraft@server").Output()

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
