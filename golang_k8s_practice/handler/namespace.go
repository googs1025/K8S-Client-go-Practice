package handler

import (
	"github.com/gin-gonic/gin"
	"golang_k8s_practice/namespace"
)

func GetNamespace(c *gin.Context) {

}


func ListNamespace(c *gin.Context) {
	nsList, err := namespace.ListNamespace()
	if err != nil {
		Fail(c, err)
		return
	}
	Response(c, nsList, nil)
}

func CreateNamespace(c *gin.Context) {

}


func DeleteNamespace(c *gin.Context) {

}

func UpdateNamespace(c *gin.Context) {

}
