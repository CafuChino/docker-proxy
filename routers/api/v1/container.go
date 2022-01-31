package v1

import (
	"docker-controller/docker"
	"docker-controller/utils"
	"github.com/docker/docker/api/types"
	"github.com/docker/docker/api/types/container"
	"github.com/docker/docker/api/types/network"
	"github.com/gin-gonic/gin"
	v1 "github.com/opencontainers/image-spec/specs-go/v1"
	"net/http"
	"runtime"
)

func GetCurrentContainers(ctx *gin.Context) {
	containers, err := docker.GetContainerList()
	if err != nil {
		ctx.JSON(http.StatusOK, gin.H{
			"errCode": -1,
			"err": err.Error(),
		})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{
		"errCode": 0,
		"data": containers,
	})
}

func GetContainersStatus(ctx *gin.Context) {
	id := ctx.Param("id")
	stat, err := docker.GetContainerStatus(id)
	if err != nil {
		ctx.JSON(http.StatusOK, gin.H{
			"errCode": -1,
			"err": err.Error(),
		})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{
		"errCode": 0,
		"data": stat,
	})

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
		ctx.JSON(http.StatusOK, gin.H{
			"errCode": -1,
			"err": err.Error(),
		})
		return
	}
	status, err := docker.StartContainer(result.ID, types.ContainerStartOptions{})
	if err != nil {
		ctx.JSON(http.StatusOK, gin.H{
			"errCode": -1,
			"err": err.Error(),
		})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{
		"errCode": 0,
		"create": result,
		"startUp": status,
	})
}

func StartExistedContainer(ctx *gin.Context) {
	id := ctx.Param("id")
	status, err := docker.StartContainer(id, types.ContainerStartOptions{})
	if err != nil {
		ctx.JSON(http.StatusOK, gin.H{
			"errCode": -1,
			"err": err.Error(),
		})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{
		"errCode": 0,
		"startUp": status,
	})
}

func StopContainer(ctx *gin.Context)  {
	id := ctx.Param("id")
	status, err := docker.StopContainer(id)
	if err != nil {
		ctx.JSON(http.StatusOK, gin.H{
			"errCode": -1,
			"err": err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"errCode": 0,
		"shutDown": status,
	})
}

func RenameExistedContainer(ctx *gin.Context) {
	body,err := utils.ParseJsonReqBody(ctx.Request.Body)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"errCode": -1,
			"err": err.Error(),
		})
		return
	}
	id := ctx.Param("id")
	containerName, succ := body["name"].(string)
	if !succ {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"errCode": -1,
			"err": err.Error(),
		})
		return
	}
	status, err := docker.RenameContainer(id, containerName)
	if err != nil {
		ctx.JSON(http.StatusOK, gin.H{
			"errCode": -1,
			"err": err.Error(),
		})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{
		"errCode": 0,
		"rename": status,
	})
}

func InspectExistedContainer(ctx *gin.Context) {
	id := ctx.Param("id")
	status, err := docker.InspectContainer(id)
	if err != nil {
		ctx.JSON(http.StatusOK, gin.H{
			"errCode": -1,
			"err": err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"errCode": 0,
		"details": status,
	})
}

func RemoveExistedContainer(ctx *gin.Context)  {
	id := ctx.Query("id")
	_force := ctx.Query("force")
	force := utils.If(_force=="true",true,false).(bool)
	status, err := docker.RemoveContainer(id, force)
	if err != nil {
		ctx.JSON(http.StatusOK, gin.H{
			"errCode": -1,
			"err": err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"errCode": 0,
		"deleted": status,
	})
}