package handler

import (
	"github.com/gin-gonic/gin"
	"golang_k8s_practice/pod"
)

func GetPod(c *gin.Context) {

	namespace := c.Param("namespace")
	podName := c.Param("podName")

	res, err := pod.GetPod(namespace, podName)
	if err != nil {
		Fail(c, err)
	}
	Response(c, res, nil)

}

func ListPod(c *gin.Context) {
	namespace := c.Param("namespace")
	resList, err := pod.ListPod(namespace)
	if err != nil {
		Fail(c, err)
	}
	Response(c, resList, nil)
}

func CreatePod(c *gin.Context) {

}


func DeletePod(c *gin.Context) {

}

func UpdatePod(c *gin.Context) {

}
