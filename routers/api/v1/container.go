package v1

import (
	"docker-controller/docker"
	"docker-controller/utils"
	"net/http"
	"runtime"

	"github.com/docker/docker/api/types"
	"github.com/docker/docker/api/types/container"
	"github.com/docker/docker/api/types/network"
	"github.com/gin-gonic/gin"
	v1 "github.com/opencontainers/image-spec/specs-go/v1"
)

func GetCurrentContainers(ctx *gin.Context) {
	containers, err := docker.GetContainerList()
	if err != nil {
		utils.MakeCommonRespose(ctx, 500, "Get container list Error", err.Error())
		return
	}
	utils.MakeCommonRespose(ctx, 200, "Get container list Success", gin.H{"containers": containers})
}

func GetContainersStatus(ctx *gin.Context) {
	id := ctx.Param("id")
	stat, err := docker.GetContainerStatus(id)
	if err != nil {
		utils.MakeCommonRespose(ctx, 500, "Get container status Error", err.Error())
		return
	}
	utils.MakeCommonRespose(ctx, 200, "Get container status Success", gin.H{"status": stat})
}

func StartNewContainer(ctx *gin.Context)  {
	body,err := utils.ParseJsonReqBody(ctx.Request.Body)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"errCode": -1,
			"err": err.Error(),
		})
		return
	}
	containerName := body["name"].(string)
	image := body["image"].(string)
	tag, succ := body["tag"].(string)
	if !succ {
		tag = ""
	}
	result, err := docker.CreateContainer(&container.Config{Image: image + ":" + utils.If(tag!="",tag,"latest").(string)}, &container.HostConfig{}, &network.NetworkingConfig{}, &v1.Platform{OS: runtime.GOARCH}, containerName)
	if err != nil {
		utils.MakeCommonRespose(ctx, 500, "Create container Error", err.Error())
		return
	}
	status, err := docker.StartContainer(result.ID, types.ContainerStartOptions{})
	if err != nil {
		utils.MakeCommonRespose(ctx, 500, "Start container Error", err.Error())
		return
	}
	utils.MakeCommonRespose(ctx, 200, "Start container Success", gin.H{"startUp": status, "create": result})
}

func StartExistedContainer(ctx *gin.Context) {
	id := ctx.Param("id")
	status, err := docker.StartContainer(id, types.ContainerStartOptions{})
	if err != nil {
		utils.MakeCommonRespose(ctx, 500, "Start container Error", err.Error())
		return
	}
	utils.MakeCommonRespose(ctx, 200, "Start container Success", gin.H{"startUp": status})
}

func StopContainer(ctx *gin.Context)  {
	id := ctx.Param("id")
	status, err := docker.StopContainer(id)
	if err != nil {
		utils.MakeCommonRespose(ctx, 500, "Stop container Error", err.Error())
		return
	}
	utils.MakeCommonRespose(ctx, 200, "Stop container Success", gin.H{"shutDown": status})
}

func RenameExistedContainer(ctx *gin.Context) {
	body,err := utils.ParseJsonReqBody(ctx.Request.Body)
	if err != nil {
		utils.MakeCommonRespose(ctx, 500, "Parse request body Error", err.Error())
		return
	}
	id := ctx.Param("id")
	containerName, succ := body["name"].(string)
	if !succ {
		utils.MakeCommonRespose(ctx, 500, "Parse request body Error", "Container name is not string")
		return
	}
	status, err := docker.RenameContainer(id, containerName)
	if err != nil {
		utils.MakeCommonRespose(ctx, 500, "Rename container Error", err.Error())

		return
	}
	utils.MakeCommonRespose(ctx, 200, "Rename container Success", gin.H{"renamed": status})
}

func InspectExistedContainer(ctx *gin.Context) {
	id := ctx.Param("id")
	status, err := docker.InspectContainer(id)
	if err != nil {
		utils.MakeCommonRespose(ctx, 500, "Inspect container Error", err.Error())
		return
	}
	utils.MakeCommonRespose(ctx, 200, "Inspect container Success", gin.H{"inspect": status})
}

func RemoveExistedContainer(ctx *gin.Context)  {
	id := ctx.Query("id")
	_force := ctx.Query("force")
	force := utils.If(_force=="true",true,false).(bool)
	status, err := docker.RemoveContainer(id, force)
	if err != nil {
		utils.MakeCommonRespose(ctx, 500, "Remove container Error", err.Error())
		return
	}
	utils.MakeCommonRespose(ctx, 200, "Remove container Success", gin.H{"remove": status})
}

func KillExistedContainer(ctx *gin.Context)  {
	id := ctx.Param("id")
	body, err := utils.ParseJsonReqBody(ctx.Request.Body)
	if err != nil {
		utils.MakeCommonRespose(ctx, 500, "Parse request body Error", err.Error())
		return
	}
	signal := body["signal"].(string)
	_err := docker.KillContainer(id, signal)
	if _err != nil {
		utils.MakeCommonRespose(ctx, 500, "Kill container Error", _err.Error())
		return
	}
	utils.MakeCommonRespose(ctx, 200, "Kill container Success", gin.H{"kill": id})
}