package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/skkrimon/mc/mctl/models"
	"github.com/skkrimon/mc/mctl/util"
	"net/http"
)

type KeyController struct{}

func (h *KeyController) Generate(c *gin.Context) {
	var req models.GenerateRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"success": false,
			"message": err.Error(),
		})
		return
	}

	apiKey, hash, err := util.GenerateKey(18)
	if err != nil {
		util.ErrorResponse(c, "error generating api key")
		return
	}

	var config util.ConfigYaml
	err = config.LoadConfig()
	if err != nil {
		util.ErrorResponse(c, "error generating api key")
		return
	}

	config.AddUser(util.User{
		Username: req.Username,
		ApiKey:   hash,
	})

	err = config.WriteConfig()
	if err != nil {
		util.ErrorResponse(c, "error generating api key")
		return
	}

	var res models.GenerateResponse
	res.Username = req.Username
	res.ApiKey = apiKey

	c.JSON(http.StatusOK, res)
}
