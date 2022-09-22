package handler

import (
	"github.com/gin-gonic/gin"
	"k8s_practice_client_go/golang_k8s_practice/service"
)

func ListDeployment(c *gin.Context) {

	namespace := c.Param("namespace")
	reslist, err := service.DeploymentList(namespace)
	if err != nil {
		Fail(c, err)
	}
	Response(c, reslist, nil)

}

func GetDeployment(c *gin.Context) {

	namespace := c.Param("namespace")
	deploymentName := c.Param("deploymentName")

	res, err := service.GetDeployment(namespace, deploymentName)
	if err != nil {
		Fail(c, err)
	}
	Response(c, res, nil)

}

func CreateDeployment(c *gin.Context) {

}


func DeleteDeployment(c *gin.Context) {

}

func UpdateDeployment(c *gin.Context) {

}
