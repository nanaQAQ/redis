package router

import (
	"redis/controller"

	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
	"github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine {
	r := gin.Default()

	//创建基于cookie的储存引擎，字节数组是密钥
	store := cookie.NewStore([]byte("woce"))

	//设置session中间件
	r.Use(sessions.Sessions("session", store))

	helloGroup := r.Group("/hello")
	{
		helloGroup.GET("", controller.HelloWorld)
	}
	userGroup := r.Group("/user")
	{
		userGroup.POST("/code", controller.SendCode)
	}
	return r
}
