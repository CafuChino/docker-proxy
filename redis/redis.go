package redis

import "github.com/gomodule/redigo/redis"

var RedisCli redis.Conn

func init() {
	var err error
	RedisCli, err = redis.Dial("tcp","127.0.0.1:6379")
	if err != nil {
		panic(err)
	}
}

func close()  {
	defer RedisCli.Close()
}