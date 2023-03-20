package main

import (
	_ "redis/dao/mysql"
	_ "redis/dao/redis"
	"redis/server"
)

func main() {
	server.Init()
}
