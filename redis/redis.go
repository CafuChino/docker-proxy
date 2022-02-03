package redis

import (
	"docker-controller/conf"
	"fmt"

	"github.com/gomodule/redigo/redis"
)

var RedisCli redis.Conn

func init() {
	conf.LoadConfig("");
	config := conf.Conf.Redis;
	var err error
	RedisCli, err = redis.Dial("tcp",fmt.Sprintf("%s:%d",config.Host, config.Port))
	if err != nil {
		panic(err)
	}
	fmt.Println("Connected to Redis!")
}

func close()  {
	defer RedisCli.Close()
}