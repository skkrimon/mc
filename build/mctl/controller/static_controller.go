package controller

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type StaticController struct{}

func (h *StaticController) Swagger(c *gin.Context) {
	c.HTML(http.StatusOK, "swagger.html", nil)
}
