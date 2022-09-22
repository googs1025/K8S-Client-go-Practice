package service

import (
	"context"
	"fmt"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s_practice_client_go/golang_k8s_practice/client"
	"k8s_practice_client_go/golang_k8s_practice/model"
	"log"
)

func DeploymentList(namespace string) (*model.DeploymentList, error) {
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
		DeploymentGetlist: make([]*model.DeploymentGet, len(deployments.Items)),
		Err: nil,
	}

	for i, d := range deployments.Items {
		data.DeploymentGetlist[i].Namespace = d.Namespace
		data.DeploymentGetlist[i].Name = d.Name
		data.DeploymentGetlist[i].Replicas = *d.Spec.Replicas
		fmt.Printf("Namespace: %v\t Name: %v\t Status: %+v\n", d.Namespace, d.Name, d.Status)
	}
	return data, nil
}

func GetDeployment(namespace string, deployname string) (string, error) {

	ctx := context.Background()
	deployment, err := client.K8sClient.AppsV1().Deployments(namespace).Get(ctx, deployname, metav1.GetOptions{})
	if err != nil {
		log.Println(err)
		return "", err
	}

	fmt.Println("取到deployment", deployment.Name)

	return deployment.Name, nil

}

func NamespaceList() {

}

func GetNamespace() {

}

func ServiceList() {

}

func GetService() {

}



