package utils

import (
	"docker-controller/redis"
	"io"
)

func StoreSessionIO2Redis(reader io.ReadCloser, sessionId string)  {
	var b = make([]byte, 1)
	var tmp string
	for {
		_, err := reader.Read(b)
		_tmpStr := string(b)
		if err == io.EOF {
			break
		}
		if _tmpStr != "\n" {
			tmp += _tmpStr
		} else {
			_,err :=redis.RedisCli.Do("RPUSH", sessionId, tmp)
			if err != nil {
				panic(err)
			}
			tmp = ""
		}
	}
}