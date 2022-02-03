package docker

import (
	"context"
	"docker-controller/utils"
	"time"

	"github.com/docker/docker/api/types"
	"github.com/docker/docker/api/types/container"
	"github.com/docker/docker/api/types/filters"
	"github.com/docker/docker/api/types/network"
	"github.com/docker/docker/client"
	v1 "github.com/opencontainers/image-spec/specs-go/v1"
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

func RestartContainer(id string) (status bool, err error) {
	timeOut, _ := time.ParseDuration("10s")
	err = cli.ContainerRestart(context.Background(), id, &timeOut)
	status = true
	if err != nil {
		status = false
	}
	return
}

func KillContainer(id string, signal string) (err error) {
	err = cli.ContainerKill(context.Background(), id, signal)
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

func GetNetworkList() (networks []types.NetworkResource, err error) {
	networks, err = cli.NetworkList(context.Background(), types.NetworkListOptions{})
	return
}

func InspectNetwork(id string) (network types.NetworkResource, err error) {
	network, err = cli.NetworkInspect(context.Background(), id, types.NetworkInspectOptions{})
	return
}

func CreateNetwork(name string, options types.NetworkCreate) (network types.NetworkCreateResponse, err error) {
	network, err = cli.NetworkCreate(context.Background(), name, options)
	return
}


func ConnectNetwork(id string,containerId string ,network *network.EndpointSettings) (err error) {
	err = cli.NetworkConnect(context.Background(), id, containerId, network)
	return
}

func DisconnectNetwork(id string,containerId string) (err error) {
	err = cli.NetworkDisconnect(context.Background(), id, containerId, true)
	return
}

func RemoveNetwork(id string) (err error) {
	err = cli.NetworkRemove(context.Background(), id)
	return
}