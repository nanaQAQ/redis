package redis

import (
	"log"
	"redis/config"

	"github.com/go-redis/redis"
)

var Redis *redis.Client

func init() {
	Redis = redis.NewClient(&redis.Options{
		Addr:     config.ServerConfig.Redis.Ip + ":" + config.ServerConfig.Redis.DbPort,
		Password: "",
		DB:       config.ServerConfig.Redis.DbName,
	})

	_, err := Redis.Ping().Result()
	if err != nil {
		log.Println("ping err :", err)
		return
	}
	log.Println("redis connected")
}
