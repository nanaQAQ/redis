package router

import (
	"redis/controller"

	"github.com/gin-gonic/gin"
)

func InitRouter(r *gin.Engine) {
	helloGroup := r.Group("/hello")
	{
		helloGroup.GET("", controller.HelloWorld)
	}
	userGroup := r.Group("/user")
	{
		userGroup.POST("/code", controller.SendCode)
	}
}
