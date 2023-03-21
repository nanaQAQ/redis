package controller

import (
	"redis/model"

	"github.com/gin-gonic/gin"
)

// /user

// /user/code
func SendCode(c *gin.Context) {
	phone := c.PostForm("phone")
	model.CodeService(phone, c)
}
