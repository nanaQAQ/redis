package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func SendCode(c *gin.Context) {
	c.String(http.StatusOK, "Hello world!")
}
