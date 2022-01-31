package v1

import (
	"docker-controller/redis"
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

func GetDockerActionProgress(ctx *gin.Context)  {
	amount, err := strconv.Atoi(ctx.Query("amount"))
	if err != nil {
		amount = 100
	}
	page, err := strconv.Atoi(ctx.Query("page"))
	if err != nil {
		page = 1
	}
	fmt.Print(amount)
	fmt.Print(page)
	listLen, err := redis.RedisCli.Do("LLEN", ctx.Param("session"))
	if err!= nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"errCode": -1,
			"err": err.Error(),
		})
		return
	}
	// 暂时不做分页
	// TODO: ID肯定是有用的 记得做去重
	items, err := redis.RedisCli.Do("LRANGE", ctx.Param("session"),0, listLen)
	var formatList []interface{}
	for _, value := range items.([]interface {}) {
		tmp := string(value.([]uint8))
		var _tmp interface{}
		err := json.Unmarshal([]byte(tmp), &_tmp)
		if err != nil {
			fmt.Print(err.Error())
		}
		formatList = append(formatList,_tmp)
	}
	ctx.JSON(http.StatusOK, gin.H{
		"errCode": 0,
		"list": formatList,
	})
}
