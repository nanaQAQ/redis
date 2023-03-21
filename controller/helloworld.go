package controller

import (
	"net/http"

	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
)

func HelloWorld(c *gin.Context) {

	session := sessions.Default(c)

	if session.Get("hello") != "world" {
		session.Set("hello", "world")
		session.Save()
	}
	c.String(http.StatusOK, session.Get("hello").(string))
}
