package deployment

import (
	"context"
	"fmt"
	"golang_k8s_practice/client"
	"k8s.io/apimachinery/pkg/api/errors"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"log"
)

func DeleteDeployment(deployname string, namespace string) error {

	ctx := context.Background()
	err := client.K8sClient.AppsV1().Deployments(namespace).Delete(ctx, deployname, metav1.DeleteOptions{})
	if err != nil {
		if errors.IsNotFound(err) {
			log.Println(err)
			fmt.Println("未找到要删除的deployment")
		}

		log.Println(err)

	} else {
		fmt.Printf("deployment: %s 已经删除", deployname)
	}

	return err

}

// TODO: 增加namespace pod service 删除接口

