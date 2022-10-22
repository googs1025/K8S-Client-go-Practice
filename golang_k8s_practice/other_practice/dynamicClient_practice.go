package main

import (
	"context"
	"fmt"
	"golang_k8s_practice/client"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime/schema"
	"log"
)

func main() {

	deployGVR := schema.GroupVersionResource{
		Group: "apps",
		Version: "v1",
		Resource: "deployments",
	}
	deployList, err := client.DynamicClient.Resource(deployGVR).
		List(context.Background(), metav1.ListOptions{})
	if err != nil {
		log.Println(err)
	}
	for _, deploy := range deployList.Items {

		fmt.Println(deploy.GetName())
	}



}
