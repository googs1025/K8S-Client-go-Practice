package main

import (
	"context"
	"fmt"
	"golang_k8s_practice/client"
	"golang_k8s_practice/model"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"log"
)

func main() {

	ListNamespace()

}


func ListNamespace() (*model.NamespaceList, error) {
	nsList, err := client.K8sClient.CoreV1().Namespaces().List(context.Background(), metav1.ListOptions{})
	if err != nil {
		log.Println(err)
		data := &model.NamespaceList{
			Err: err,
		}
		return data, err
	}

	data := &model.NamespaceList{
		Namespaces: make([]*model.NamespaceGet, 0),
		Err: nil,
	}

	for _, d := range nsList.Items {
		ns := &model.NamespaceGet{
			Name: d.Name,
			Status: model.NamespacePhase(d.Status.Phase),
		}
		data.Namespaces = append(data.Namespaces, ns)

		fmt.Printf("Namespace name: %v\t, Status: %v\t \n", d.Name, d.Status)
	}

	return data, nil

}