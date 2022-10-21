package namespace

import (
	"context"
	"golang_k8s_practice/client"
	coreV1 "k8s.io/api/core/v1"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"log"
)

func CreatNamespace(namespace string) error {
	ctx := context.Background()
	ns, _ := client.K8sClient.CoreV1().Namespaces().Get(ctx, namespace, v1.GetOptions{})
	if ns != nil {
		// 代表需要创建
		nsObj := &coreV1.Namespace{
			ObjectMeta: v1.ObjectMeta{
				Name: namespace,
			},
		}
		_, err := client.K8sClient.CoreV1().Namespaces().Create(ctx, nsObj, v1.CreateOptions{})
		if err != nil {
			log.Println(err)
			return err
		}

	}

	return nil


}
