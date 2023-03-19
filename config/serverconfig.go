package config

import (
	"log"
	"os"

	"gopkg.in/yaml.v3"
)

var ServerConfig *serverConfig

//var ServerConfig atomic.Value

type serverConfig struct {
	Mysql struct {
		DbPort string `yaml:"db_port"`
		Ip     string `yaml:"ip"`
		User   string `yaml:"user"`
		Pwd    string `yaml:"pwd"`
		DbName string `yaml:"db_name"`
	} `yaml:"mysql"`
	Redis struct {
		DbPort string `yaml:"db_port"`
		Ip     string `yaml:"ip"`
		DbName int    `yaml:"db_name"`
	} `yaml:"redis"`
	Server struct {
		HttpPort string `yaml:"http_port"`
	} `yaml:"server"`
}

// 热更新server-config
func init() {
	ServerConfig = getServerConfig()
}

func getServerConfig() *serverConfig {
	yamlFile, err := os.ReadFile("config/server-config.yml")
	if err != nil {
		log.Fatal(err)
	}
	c := &serverConfig{}
	err = yaml.Unmarshal(yamlFile, c)
	if err != nil {
		log.Fatal(err)
	}
	return c
}
