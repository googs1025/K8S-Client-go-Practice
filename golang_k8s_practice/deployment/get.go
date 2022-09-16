package deployment

import (
	"context"
	"fmt"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
	"log"
)

func GetPodsList(ctx context.Context, client *kubernetes.Clientset, namespace string) {

	pods, err := client.CoreV1().Pods(namespace).List(ctx, metav1.ListOptions{})
	if err != nil {
		log.Println(err)
	}

	for _, pod := range pods.Items {
		fmt.Printf("Namespace: %v\t Name: %v\t Status: %+v\n", pod.Namespace, pod.Name, pod.Status.Phase)
	}

}

func GetDeploymentList(ctx context.Context, client *kubernetes.Clientset, namespace string) {

	deployments, err := client.AppsV1().Deployments(namespace).List(ctx, metav1.ListOptions{})
	if err != nil {
		log.Println(err)
	}

	for _, d := range deployments.Items {
		fmt.Printf("Namespace: %v\t Name: %v\t Status: %+v\n", d.Namespace, d.Name, d.Status)
	}

}

func GetPod(ctx context.Context, client *kubernetes.Clientset, namespace string, podname string) {
	pod, err := client.CoreV1().Pods(namespace).Get(ctx, podname, metav1.GetOptions{})
	if err != nil {
		log.Println(err)
	}
	fmt.Println("取到pod", pod.Name)
}

func GetDeployment(ctx context.Context, client *kubernetes.Clientset, namespace string, deployname string) {

	deployment, err := client.AppsV1().Deployments(namespace).Get(ctx, deployname, metav1.GetOptions{})
	if err != nil {
		log.Println(err)
	}

	fmt.Println("取到deployment", deployment.Name)

}


