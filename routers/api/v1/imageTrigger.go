package v1

import (
	"bytes"
	"context"
	"docker-controller/mongo"
	"encoding/json"
	"github.com/gin-gonic/gin"
	"net/http"
)

func NewTagTrigger(ctx *gin.Context){
	collection := mongo.Client.Database("docker").Collection("triggers")
	var jsonBody interface{}
	resBody := ctx.Request.Body
	bodyBuf := new(bytes.Buffer)
	_, err := bodyBuf.ReadFrom(resBody)
	if err != nil {
		bodyBuf = nil
	}
	bodyString := bodyBuf.String()
	json.Unmarshal([]byte(bodyString), &jsonBody)
	result, err := collection.InsertOne(context.TODO(), jsonBody)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"errCode": -1,
			"err": err.Error(),
		})
	} else {
		ctx.JSON(http.StatusOK, gin.H{
			"errCode": 0,
			"result": result,
		})
	}

}
