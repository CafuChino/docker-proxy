package v1

import (
	"docker-controller/docker"
	"docker-controller/utils"

	"net/http"

	"github.com/docker/docker/api/types"
	// "github.com/docker/docker/api/types/network"
	"github.com/gin-gonic/gin"
)

func GetNetworkList(ctx *gin.Context) {
	networks, err := docker.GetNetworkList()
	if err != nil {
		utils.MakeCommonRespose(ctx, 500, "Get network list Error", err.Error())
		return
	}
	utils.MakeCommonRespose(ctx, 200, "Get network list Success", gin.H{"networks": networks})
}

func InspectExistedNetwork(ctx *gin.Context) {
	id := ctx.Param("id")
	network, err := docker.InspectNetwork(id)
	if err != nil {
		utils.MakeCommonRespose(ctx, 500, "Inspect network Error", err.Error())
		return
	}
	utils.MakeCommonRespose(ctx, 200, "Inspect network Success", gin.H{"network": network})
}

func CreateNewNetwork(ctx *gin.Context) {
	body, err := utils.ParseJsonReqBody(ctx.Request.Body)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"errCode": -1,
			"err": err.Error(),
		})
		return
	}
	name := body["name"].(string)
	driver := body["driver"].(string)
	result, err := docker.CreateNetwork(name, types.NetworkCreate{Driver: driver})
	if err != nil {
		utils.MakeCommonRespose(ctx, 500, "Create network Error", err.Error())
		return
	}
	utils.MakeCommonRespose(ctx, 200, "Create network Success", gin.H{"network": result})
}

func ConnectExistedContainerToNetwork(ctx *gin.Context) {
	body, err := utils.ParseJsonReqBody(ctx.Request.Body)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"errCode": -1,
			"err": err.Error(),
		})
		return
	}
	containerId := body["containerId"].(string)
	networkId := body["networkId"].(string)
	err = docker.ConnectNetwork(networkId, containerId, nil)
	if err != nil {
		utils.MakeCommonRespose(ctx, 500, "Connect container to network Error", err.Error())
		return
	}
	utils.MakeCommonRespose(ctx, 200, "Connect container to network Success", nil)
} 

func DisconnectExistedContainerFromNetwork(ctx *gin.Context) {
	body, err := utils.ParseJsonReqBody(ctx.Request.Body)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"errCode": -1,
			"err": err.Error(),
		})
		return
	}
	containerId := body["containerId"].(string)
	networkId := body["networkId"].(string)
	err = docker.DisconnectNetwork(networkId, containerId)
	if err != nil {
		utils.MakeCommonRespose(ctx, 500, "Disconnect container from network Error", err.Error())
		return
	}
	utils.MakeCommonRespose(ctx, 200, "Disconnect container from network Success", nil)
}

func RemoveExistedNetwork(ctx *gin.Context) {
	id := ctx.Param("id")
	err := docker.RemoveNetwork(id)
	if err != nil {
		utils.MakeCommonRespose(ctx, 500, "Remove network Error", err.Error())
		return
	}
	utils.MakeCommonRespose(ctx, 200, "Remove network Success", nil)
}