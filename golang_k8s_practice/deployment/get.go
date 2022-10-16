package deployment

import (
	"context"
	"fmt"
	"golang_k8s_practice/client"
	"golang_k8s_practice/model"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"log"
)

func ListDeployment(namespace string) (*model.DeploymentList, error) {
	ctx := context.Background()
	deployments, err := client.K8sClient.AppsV1().Deployments(namespace).List(ctx, metav1.ListOptions{})

	if err != nil {
		data := &model.DeploymentList{
			Err: err,
		}
		log.Println(err)
		return data, err

	}

	data := &model.DeploymentList{
		DeploymentGetlist: make([]*model.DeploymentGet, 0),
		Err: nil,
	}

	for _, d := range deployments.Items {
		deploy := &model.DeploymentGet{
			Name: d.Name,
			Namespace: d.Namespace,
			Replicas: *d.Spec.Replicas,

		}
		data.DeploymentGetlist = append(data.DeploymentGetlist, deploy)
		fmt.Printf("Namespace: %v\t Name: %v\t Status: %+v\n", d.Namespace, d.Name, d.Status)
	}
	return data, nil
}

func GetDeployment(namespace string, deployName string) (*model.DeploymentGet, error) {

	ctx := context.Background()
	data := &model.DeploymentGet{}
	deployment, err := client.K8sClient.AppsV1().Deployments(namespace).Get(ctx, deployName, metav1.GetOptions{})
	if err != nil {
		log.Println(err)
		return data, err
	}

	fmt.Println("取到deployment", deployment.Name)
	data.Name = deployment.Name
	data.Namespace = deployment.Namespace
	data.Replicas = *deployment.Spec.Replicas

	return data, nil

}


