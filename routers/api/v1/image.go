package v1

import (
	"docker-controller/docker"
	"docker-controller/utils"

	"github.com/docker/distribution/uuid"
	"github.com/gin-gonic/gin"
)

// GetCurrentImages 获取当前的镜像列表
func GetCurrentImages(ctx *gin.Context)  {
	images := docker.GetImageList()
	utils.MakeCommonRespose(ctx, 200, "List all images", images)
}

func PullNewImage(ctx *gin.Context)  {
	name := ctx.Param("name")
	tag := ctx.Param("tag")
	sessionId := uuid.Generate().String()
	go docker.PullNewImage(name, tag, sessionId)
	utils.MakeCommonRespose(ctx, 200, "Pulling new image", gin.H{"session": sessionId})
}

func RemoveExistedImage(ctx *gin.Context)  {
	id := ctx.Query("id")
	items, err := docker.RemoveExistedImage(id)
	if err != nil {
		utils.MakeCommonRespose(ctx, 500, "Remove image failed", gin.H{"err": err.Error()})
	} else {
		utils.MakeCommonRespose(ctx, 200, "Remove image success", gin.H{"items": items})
	}
}

func FetchImageTags(ctx *gin.Context) {
	name := ctx.Param("name")
	imageType := ctx.Query("type")
	if imageType == ""  {
		imageType = "public"
	}
	if imageType == "public" {
		res, err := docker.FetchPublicImageTags(name)
		if err != nil {
			utils.MakeCommonRespose(ctx, 500, "Fetch image tags failed", gin.H{"err": err.Error()})
		} else {
			utils.MakeCommonRespose(ctx, 200, "Fetch image tags success", gin.H{"tags": res.BodyParsed})
		}
	}
}