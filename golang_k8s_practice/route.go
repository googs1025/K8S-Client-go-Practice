package main

import (
	"github.com/gin-gonic/gin"
	"golang_k8s_practice/handler"
)

func register(router *gin.Engine) {
	api := router.Group("/api")
	{
		k8sOperation := api.Group("/k8s_operation")
		{
			deployment := k8sOperation.Group("deployment")
			{
				deployment.GET("/:namespace", handler.ListDeployment)
				// FIXME: Get 需要重新写
				//deployment.GET("/:deploymentName/:namespace", handler.GetDeployment)
				//deployment.DELETE("/:deploymentName/:namespace", handler.DeleteDeployment)

			}
			namespace := k8sOperation.Group("namespace")
			{
				namespace.GET("", handler.ListNamespace)
				//namespace.GET()
				//namespace.POST()
				//namespace.DELETE()
			}
			service := k8sOperation.Group("service")
			{
				service.GET("/:namespace", handler.ListService)
				//service.GET()
				//service.POST()
				//service.DELETE()
			}
			pod := k8sOperation.Group("pod")
			{
				pod.GET("/:namespace", handler.ListPod)
				//pod.GET("/:podName", handler.GetPod)
				//pod.POST()
				//pod.DELETE()
			}
		}
	}

	api.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})

	//swagger
	//api.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))


}