package server

import (
	"redis/router"

	"github.com/gin-gonic/gin"
)

func Init() {
	r := gin.Default()
	router.InitRouter(r)
	r.SetTrustedProxies(nil)
	r.Run(":8000")
}
