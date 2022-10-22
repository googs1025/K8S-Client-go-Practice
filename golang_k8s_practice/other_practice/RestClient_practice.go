package main

import (
	"context"
	"fmt"
	"golang_k8s_practice/client"

	//apiv1 "k8s.io/api/core/v1"
	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes/scheme"
	"k8s.io/client-go/rest"
	"k8s.io/client-go/tools/clientcmd"
)


func main() {
	config, err := clientcmd.BuildConfigFromFlags(
		"", "/root/.kube/config")
	if err != nil {
		panic(err)
	}

	config.APIPath = "api"
	config.GroupVersion = &corev1.SchemeGroupVersion
	config.NegotiatedSerializer = scheme.Codecs



	result := &corev1.PodList{}

	err = client.RestClient.Get().Namespace("default").Resource("pods").VersionedParams(&metav1.ListOptions{}, scheme.ParameterCodec).Do(context.TODO()).Into(result)
	if err != nil {
		panic(err)
	}

	for _, d := range result.Items {
		fmt.Printf("Namespace: %v\t Name: %v\t Status: %+v\n", d.Namespace, d.Name, d.Status.Phase)
	}



}
