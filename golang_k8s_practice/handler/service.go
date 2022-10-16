package handler

import (
	"github.com/gin-gonic/gin"
	"golang_k8s_practice/service"
)

func GetService(c *gin.Context) {

}

func ListService(c *gin.Context) {
	namespace := c.Param("namespace")
	serviceList, err := service.ListService(namespace)
	if err != nil {
		Fail(c, err)
	}
	Response(c, serviceList, nil)
}

func CreateService(c *gin.Context) {

}


func DeleteService(c *gin.Context) {

}

func UpdateService(c *gin.Context) {

}
