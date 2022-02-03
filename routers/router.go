package routers

import (
	v1 "docker-controller/routers/api/v1"

	"github.com/gin-gonic/gin"
)

func InitRouter() *gin.Engine {
	r := gin.New()
	r.Use(gin.Logger())
	r.Use(gin.Recovery())
	gin.SetMode("debug")
	apiV1 := r.Group("/api/v1")
	{
		apiV1.GET("/images", v1.GetCurrentImages)
		apiV1.PUT("/images/:name/:tag", v1.PullNewImage)
		apiV1.DELETE("/images", v1.RemoveExistedImage)
		apiV1.GET("/images/tags/:name", v1.FetchImageTags)
		apiV1.GET("/progress/docker/:session", v1.GetDockerActionProgress)
		apiV1.GET("/containers", v1.GetCurrentContainers)
		apiV1.PUT("/containers", v1.StartNewContainer)
		apiV1.DELETE("/containers", v1.RemoveExistedContainer)
		apiV1.GET("/containers/status/:id", v1.GetContainersStatus)
		apiV1.POST("/containers/start/:id", v1.StartExistedContainer)
		apiV1.POST("/containers/stop/:id", v1.StopContainer)
		apiV1.POST("/containers/kill/:id", v1.KillExistedContainer)
		apiV1.POST("/containers/rename/:id", v1.RenameExistedContainer)
		apiV1.POST("/containers/restart/:id", v1.RestartExistedContainer)
		apiV1.GET("/containers/inspect/:id", v1.InspectExistedContainer)
		apiV1.GET("/networks", v1.GetNetworkList)
		apiV1.GET("/networks/:id", v1.InspectExistedNetwork)
		apiV1.PUT("/networks", v1.CreateNewNetwork)
		apiV1.DELETE("/networks/:id", v1.RemoveExistedNetwork)
		apiV1.POST("/network/connect", v1.ConnectExistedContainerToNetwork)
		apiV1.POST("/network/disconnect", v1.DisconnectExistedContainerFromNetwork)
	}
	return r
}