package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/skkrimon/mc/mctl/models"
	"github.com/skkrimon/mc/mctl/util"
	"net/http"
)

type ServersController struct{}

func (h *ServersController) GetServers(c *gin.Context) {
	var config util.ConfigYaml
	err := config.LoadConfig()
	if err != nil {
		util.ErrorResponse(c, err.Error())
		return
	}

	var res models.GetServersResponse
	res.Servers = config.Servers

	c.JSON(http.StatusOK, res)
}
