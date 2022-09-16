package deployment

import (
	"context"
	"fmt"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
	"log"
)

func DeleteDeployment(ctx context.Context, client *kubernetes.Clientset, deployname string, namespace string) {

	err := client.AppsV1().Deployments(namespace).Delete(ctx, deployname, metav1.DeleteOptions{})
	if err != nil {
		log.Println(err)

	}
	fmt.Printf("deployment: %s 已经删除", deployname)


}
