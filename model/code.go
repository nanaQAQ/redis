package model

import (
	"fmt"
	"math/rand"
	"net/http"
	"redis/utils"
	"time"

	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
)

func CodeService(phone string, c *gin.Context) {

	//判断是否是手机号
	if !utils.IsValidPhoneNumbers(phone) {
		//如果不是则返回错误
		c.JSON(http.StatusOK, gin.H{
			"err":     "2001",
			"message": "Invalid phone number",
			"code":    "",
		})
		return
	}

	//生成验证码
	code := fmt.Sprintf("%06v", rand.New(rand.NewSource(time.Now().UnixNano())).Int31n(1000000))

	//保存验证码
	session := sessions.Default(c)
	session.Set(phone, code)
	session.Save()

	//返回验证码
	c.JSON(http.StatusOK, gin.H{
		"err":     "2000",
		"message": "code",
		"code":    session.Get(phone),
	})
}
