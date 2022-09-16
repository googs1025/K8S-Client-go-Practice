package deployment

import (
	"context"
	"fmt"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
	"log"
)

func UpdateDeploymentImage(ctx context.Context, client *kubernetes.Clientset, deployname string, namespace string, image string) {

	deployment, err := client.AppsV1().Deployments(namespace).Get(ctx, deployname, metav1.GetOptions{})
	if err != nil {
		log.Println(err)
	}

	deployment.Spec.Template.Spec.Containers[0].Image = image
	deployment, err = client.AppsV1().Deployments(namespace).Update(ctx, deployment, metav1.UpdateOptions{})
	if err != nil {
		log.Println(err)
	}
	fmt.Println("deployment image已更新完成!")
}

func UpdateDeploymentReplica(ctx context.Context, client *kubernetes.Clientset, deployname string, namespace string, replicas int32) {
	deployment, err := client.AppsV1().Deployments(namespace).Get(ctx, deployname, metav1.GetOptions{})
	if err != nil {
		log.Println(err)
	}

	deployment.Spec.Replicas = &replicas
	deployment, err = client.AppsV1().Deployments(namespace).Update(ctx, deployment, metav1.UpdateOptions{})
	if err != nil {
		log.Println(err)
	}
	fmt.Println("deployment 副本数已更新完成")

}
