package main

import (
	"github.com/gin-gonic/gin"
	"k8s_practice_client_go/golang_k8s_practice/handler"
)

func register(router *gin.Engine) {
	api := router.Group("/api")
	{
		k8s_operation := api.Group("/k8s_operation")
		{
			deployment := k8s_operation.Group("deployment")
			{
				deployment.GET("", handler.ListDeployment)
				deployment.GET("/:deploymentName", handler.GetDeployment)
				deployment.POST()
				deployment.DELETE()

			}

			namespace := api.Group("namespace")
			{
				namespace.GET()
				namespace.GET()
				namespace.POST()
				namespace.DELETE()
			}

			service := api.Group("service")
			{
				//service.GET()
				//service.GET()
				//service.POST()
				//service.DELETE()
			}

			pod := api.Group("pod")
			{
				pod.GET()
				pod.GET()
				pod.POST()
				pod.DELETE()
			}


		}


	}

}
