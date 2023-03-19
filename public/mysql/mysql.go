package mysql

import (
	"log"
	"redis/config"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var Database *gorm.DB

func getPath() string {
	dbConfig := config.ServerConfig.Mysql.User + ":" + config.ServerConfig.Mysql.Pwd + "@tcp(" + config.ServerConfig.Mysql.Ip + ":" + config.ServerConfig.Mysql.DbPort + ")/" + config.ServerConfig.Mysql.DbName
	return dbConfig
}

func init() {
	var err error
	Database, err = gorm.Open(mysql.Open(getPath()+"?charset=utf8mb4&parseTime=True&loc=Local"), &gorm.Config{})
	if err != nil {
		log.Println(err)
	}
	log.Println("sql connected")
}
