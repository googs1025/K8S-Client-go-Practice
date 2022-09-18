package deployment

import (
	"context"
	"fmt"
	"k8s.io/apimachinery/pkg/api/errors"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
	"log"
)

func DeleteDeployment(ctx context.Context, client *kubernetes.Clientset, deployname string, namespace string) {

	err := client.AppsV1().Deployments(namespace).Delete(ctx, deployname, metav1.DeleteOptions{})
	if err != nil {
		if errors.IsNotFound(err) {
			log.Println(err)
			fmt.Println("未找到要删除的deployment")
		}

		log.Println(err)

	} else {
		fmt.Printf("deployment: %s 已经删除", deployname)
	}


}

// TODO: 增加namespace pod service 删除接口

