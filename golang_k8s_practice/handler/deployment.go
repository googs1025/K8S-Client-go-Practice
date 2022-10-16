package handler

import (
	"github.com/gin-gonic/gin"
	"golang_k8s_practice/deployment"
	"log"
)


// ListDeployment function
// @Summary 获取deployment列表
// @version 1.0
// @Accept  json
// @Produce  json
// @Success 200  object model.DeploymentList "成功后返回"
// @Router /k8s_operation/deployment [get]


func ListDeployment(c *gin.Context) {

	namespace := c.Param("namespace")
	resList, err := deployment.ListDeployment(namespace)
	if err != nil {
		Fail(c, err)
		return
	}
	Response(c, resList, nil)

}

// GetDeployment function
// @Summary 获取特定deployment
// @version 1.0
// @Success 200  object model.DeploymentGet "成功后返回deployment名"
// @Router /k8s_operation/deployment/:deploymentName [get]

type DeploymentGetRequest struct {
	Namespace  string "json:namespace"
	Deployment string "json:deployment"
}

func GetDeployment(c *gin.Context)  {

	var deploymentGetRequest DeploymentGetRequest
	err := c.ShouldBindJSON(&deploymentGetRequest)
	if err != nil {
		log.Println(err)
	}
	res, err := deployment.GetDeployment(deploymentGetRequest.Namespace, deploymentGetRequest.Deployment)
	if err != nil {
		Fail(c, err)
		return
	}
	Response(c, res, nil)

}

func CreateDeployment(c *gin.Context) {

}


func DeleteDeployment(c *gin.Context) {
	namespace := c.Param("namespace")
	deploymentName := c.Param("deploymentName")
	err := deployment.DeleteDeployment(deploymentName, namespace)
	if err != nil {
		Fail(c, err)
		return
	}
	Response(c, nil, nil)
}

func UpdateDeployment(c *gin.Context) {

}
