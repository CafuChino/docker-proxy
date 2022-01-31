package docker

import (
	"context"
	"docker-controller/utils"
	"github.com/docker/docker/api/types"
	"github.com/docker/docker/api/types/container"
	"github.com/docker/docker/api/types/filters"
	"github.com/docker/docker/api/types/network"
	"github.com/docker/docker/client"
	v1 "github.com/opencontainers/image-spec/specs-go/v1"
	"time"
)

var cli *client.Client

func init()  {
	var err error
	cli, err = client.NewClientWithOpts(client.FromEnv)
	if err != nil {
		panic(err)
	}
}

func closeClient() {
	defer cli.Close()
}

func GetImageList()(images []types.ImageSummary) {
	images, err := cli.ImageList(context.Background(), types.ImageListOptions{})
	if err != nil {
		panic(err)
	}
	return
}



func PullNewImage(name string, tag string, sessionId string) {
	retStr := name + ":" + tag
	io, err := cli.ImagePull(context.Background(), retStr, types.ImagePullOptions{})
	if err != nil {
		panic(err)
	}
	go utils.StoreSessionIO2Redis(io, sessionId)
	return
}

func RemoveExistedImage(id string) (deletedItem []types.ImageDeleteResponseItem, err error){
	deletedItem, err = cli.ImageRemove(context.Background(), id, types.ImageRemoveOptions{})
	return
}

func FetchPublicImageTags(name string) (res utils.FriendlyHttpResponse, err error) {
	res, err = utils.HttpGet("https://registry.hub.docker.com/v1/repositories/" + name + "/tags")
	return
}

func GetContainerList() (containers []types.Container, err error) {
	containers, err = cli.ContainerList(context.Background(), types.ContainerListOptions{ All: true })
	return
}

func GetContainerStatus(id string) (containers []types.Container, err error){
	containers, err = cli.ContainerList(context.Background(), types.ContainerListOptions{ All: true, Filters: filters.NewArgs(filters.Arg("id", id))})
	return
}

func CreateContainer(config *container.Config, hostConfig *container.HostConfig, networkingConfig *network.NetworkingConfig, platform *v1.Platform, name string) (createBody container.ContainerCreateCreatedBody, err error) {
	createBody, err = cli.ContainerCreate(context.Background(), config, hostConfig, networkingConfig, platform, name)
	return
}

func StartContainer(id string, options types.ContainerStartOptions) (status bool, err error) {
	err = cli.ContainerStart(context.Background(), id, options)
	status = true
	if err != nil {
		status = false
	}
	return
}

func StopContainer(id string) (status bool, err error) {
	timeOut, _ := time.ParseDuration("10s")
	err = cli.ContainerStop(context.Background(), id, &timeOut)
	status = true
	if err != nil {
		status = false
	}
	return
}

func RenameContainer(id string, name string) (status bool, err error)  {
	err = cli.ContainerRename(context.Background(), id, name)
	status = true
	if err != nil {
		status = false
	}
	return
}

func InspectContainer(id string) (res types.ContainerJSON, err error) {
	res, err = cli.ContainerInspect(context.Background(), id)
	return
}

func RemoveContainer(id string, force bool) (status bool, err error) {
	err = cli.ContainerRemove(context.Background(), id, types.ContainerRemoveOptions{Force: force})
	status = true
	if err != nil {
		status = false
	}
	return
}