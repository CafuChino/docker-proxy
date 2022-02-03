package conf

import (
	"fmt"
	"log"

	"github.com/spf13/viper"
)


type serverConfig struct {
	Port int
	Address string
	Mode string
}

type mongoConfig struct {
	Host string
	Port int
	User string
	Password string
}

type redisConfig struct {
	Host string
	Port int
	Password string
}

type Config struct {
	Server serverConfig
	Mongo mongoConfig
	Redis redisConfig
}

var Conf = new(Config)

func LoadConfig(path string) {
	viper.SetDefault("Server", serverConfig{
		Port: 8080,
		Address: "127.0.0.1",
		Mode: "debug",
	})
	viper.SetDefault("Mongo", mongoConfig{
		Host: "127.0.0.1",
		Port: 27017,
		User: "",
		Password: "",
	})
	viper.SetDefault("Redis", redisConfig{
		Host: "127.0.0.1",
		Port: 6379,
		Password: "",
	})
	if path != "" {
		viper.SetConfigFile(path)
	} else {
		viper.SetConfigName("config")
		viper.AddConfigPath("./conf")
	}
	err := viper.ReadInConfig()
	if err != nil {
		log.Fatalf("Fatal error config file: %s \n", err)
		return
	}
	if err := viper.Unmarshal(Conf); err != nil {
		panic(fmt.Errorf("unmarshal conf failed, err:%s", err))
	}
}