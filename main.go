package main

import (
	//_ "redis/dao/mysql"
	//_ "redis/dao/redis"
	"redis/router"
)

func main() {

	//设置路由
	r := router.SetupRouter()

	r.SetTrustedProxies(nil)
	r.Run(":8000")
}
