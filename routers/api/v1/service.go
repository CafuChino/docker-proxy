package v1

import (
	"github.com/gin-gonic/gin"
)

// Group Manage

func GetGroupList(){
	//mongo.Database.Collection("group").Find()
}

func GetGroupInfo(){}

func AddGroup(){}

func EditGroup(){}

func DeleteGroup(){}

// Service Manage

func GetServiceList(){}

func GetServiceInfo() {}

func AddService() {}

func EditService() {}

func StopService() {}

func DownGradeService() {}

func DeployService() {}

// Conf Manage

func CreateConf() {}

func EditConf(){}

func GetConfList(ctx *gin.Context){}

func GetConfInfo(){}

func DeleteConf(){}