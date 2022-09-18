package deployment

import (
	"context"
	"fmt"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
	"log"
	"sync"
	"time"
)

// TODO: 提供一个watch接口。

var once sync.Once

// watch接口 并且设置超时退出！
func WatchDeployment(ctx context.Context, client *kubernetes.Clientset, namespce string, d time.Duration) {

	ctx, cancel := context.WithTimeout(ctx, d)
	defer cancel()
	watch, err := client.AppsV1().Deployments(namespce).Watch(ctx, metav1.ListOptions{})
	if err != nil {
		log.Println(err)
	}

	f := func() {
		fmt.Println("时间到了，watch退出")
	}

	for {
		select {

		case e, ok := <-watch.ResultChan():
			if e.Object == nil {
				if !ok {
					return
				}
			} else {
				fmt.Println(e.Type, e.Object)
			}
		case <-ctx.Done():
			once.Do(f)

		}

	}


}

func WatchPod(ctx context.Context, client *kubernetes.Clientset, namespace string) {

}

