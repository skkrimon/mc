package controller

import (
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/skkrimon/mc/api/model"
	"github.com/skkrimon/mc/api/util"
	"io"
	"net/http"
)

type MctlController struct{}

func (h *MctlController) Start(c *gin.Context) {
	executeCommand(c, "start")
}

func (h *MctlController) Stop(c *gin.Context) {
	executeCommand(c, "stop")
}

func (h *MctlController) Restart(c *gin.Context) {
	executeCommand(c, "restart")
}

func (h *MctlController) Update(c *gin.Context) {
	executeCommand(c, "update")
}

func executeCommand(c *gin.Context, command string) {
	res, err := http.Post(
		fmt.Sprintf("%s/ctl/%s", getBaseUrl(), command),
		"application/json",
		nil,
	)

	if err != nil {
		handleError(c, err.Error())
		return
	}

	handleSuccess(c, res.Body)

	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {
			fmt.Println(err.Error())
		}
	}(res.Body)
}

func handleError(c *gin.Context, message string) {
	c.JSON(http.StatusInternalServerError, gin.H{
		"success": false,
		"message": message,
	})
}

func handleSuccess(c *gin.Context, body io.Reader) {
	var b model.MctlResponse

	decoder := json.NewDecoder(body)
	decoder.DisallowUnknownFields()
	err := decoder.Decode(&b)

	if err != nil {
		handleError(c, err.Error())
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"success": b.Success,
		"message": b.Message,
	})
}

func getBaseUrl() string {
	return util.GetEnv("MCTL_URL")
}
