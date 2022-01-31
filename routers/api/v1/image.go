package v1

import (
	"docker-controller/docker"
	"github.com/docker/distribution/uuid"
	"github.com/gin-gonic/gin"
	"net/http"
)

// GetCurrentImages 获取当前的镜像列表
func GetCurrentImages(ctx *gin.Context)  {
	images := docker.GetImageList()
	ctx.JSON(http.StatusOK, gin.H{
		"images": images,
	})
}

func PullNewImage(ctx *gin.Context)  {
	name := ctx.Param("name")
	tag := ctx.Param("tag")
	sessionId := uuid.Generate().String()
	go docker.PullNewImage(name, tag, sessionId)
	ctx.JSON(http.StatusOK, gin.H{
		"session": sessionId,
	})
}

func RemoveExistedImage(ctx *gin.Context)  {
	id := ctx.Query("id")
	items, err := docker.RemoveExistedImage(id)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"errCode": -1,
			"err": err.Error(),
		})
	} else {
		ctx.JSON(http.StatusOK, gin.H{
			"errCode": 0,
			"deleted": items,
		})
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
			ctx.JSON(http.StatusInternalServerError, gin.H{
				"errCode": -1,
				"err": err.Error(),
			})
		} else {
			ctx.JSON(http.StatusOK, gin.H{
				"errCode": 0,
				"tags": res.BodyParsed,
			})
		}
	}
}