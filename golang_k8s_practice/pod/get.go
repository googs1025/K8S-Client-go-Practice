package pod

import (
	"context"
	"fmt"
	"golang_k8s_practice/client"
	"golang_k8s_practice/model"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"log"
)

func GetPod(namespace string, podName string) (*model.PodGet, error) {

	ctx := context.Background()
	data := &model.PodGet{}
	pod, err := client.K8sClient.CoreV1().Pods(namespace).Get(ctx, podName, metav1.GetOptions{})
	if err != nil {
		log.Println(err)
		return data, err
	}
	fmt.Println("取到pod", pod.Name)
	data.Namespace = pod.Namespace
	data.Name = pod.Name
	data.ContainerNum = len(pod.Spec.Containers)
	return data, nil
}

func ListPod(namespace string) (*model.PodList, error) {
	ctx := context.Background()
	pods, err := client.K8sClient.CoreV1().Pods(namespace).List(ctx, metav1.ListOptions{})

	if err != nil {
		data := &model.PodList{
			Err: err,
		}
		log.Println(err)
		return data, err

	}

	data := &model.PodList{
		PodGetlist: make([]*model.PodGet, 0),
		Err: nil,
	}

	for _, d := range pods.Items {
		pod := &model.PodGet{
			Name: d.Name,
			Namespace: d.Namespace,
			ContainerNum: len(d.Spec.Containers),
		}
		data.PodGetlist = append(data.PodGetlist, pod)
		fmt.Printf("Namespace: %v\t Name: %v\t Status: %+v\n", d.Namespace, d.Name, d.Status)
	}
	return data, nil
}
