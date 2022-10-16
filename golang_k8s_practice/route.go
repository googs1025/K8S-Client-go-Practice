package main

import (
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"golang_k8s_practice/handler"
)

func register(router *gin.Engine) {
	api := router.Group("/api")
	{
		k8s_operation := api.Group("/k8s_operation")
		{
			deployment := k8s_operation.Group("deployment")
			{
				deployment.GET("", handler.ListDeployment)
				deployment.GET("/:deploymentName/:namespace", handler.GetDeployment)
				deployment.DELETE("/:deploymentName/:namespace", handler.DeleteDeployment)
				//deployment.POST()
				//deployment.DELETE()

			}

			namespace := api.Group("namespace")
			{
				namespace.GET("", handler.ListNamespace)
				//namespace.GET()
				//namespace.POST()
				//namespace.DELETE()
			}
		//
			service := api.Group("service")
			{
				service.GET("", handler.ListService)
				//service.GET()
				//service.POST()
				//service.DELETE()
			}
		//
			pod := api.Group("pod")
			{
				pod.GET("", handler.ListPod)
				pod.GET("/:podName", handler.GetPod)
				//pod.POST()
				//pod.DELETE()
			}
		//
		//
		}


	}

	api.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})

	//swagger
	api.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))


}