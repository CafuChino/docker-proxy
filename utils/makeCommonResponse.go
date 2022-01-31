package utils

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func MakeCommonRespose(ctx *gin.Context, errCode int, errMsg string, data interface{}) {
	ctx.JSON(http.StatusOK, gin.H{
		"code": errCode,
		"msg": errMsg,
		"data": data,
	})
}