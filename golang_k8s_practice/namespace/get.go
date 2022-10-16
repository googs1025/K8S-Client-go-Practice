package namespace

import (
	"fmt"
	"golang_k8s_practice/client"
	"golang_k8s_practice/model"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"context"
	"log"
)



func ListNamespace() (*model.NamespaceList, error) {
	nslist, err := client.K8sClient.CoreV1().Namespaces().List(context.Background(), metav1.ListOptions{})
	if err != nil {
		log.Println(err)
		data := &model.NamespaceList{
			Err: err,
		}
		return data, err
	}

	data := &model.NamespaceList{
		Namespaces: make([]*model.NamespaceGet, len(nslist.Items)),
		Err: nil,
	}

	for i, d := range nslist.Items {
		data.Namespaces[i].Name = d.Name
		data.Namespaces[i].Status = model.NamespacePhase(d.Status.Phase)
		fmt.Printf("Namespace name: %v\t, Status: %v\t", d.Name, d.Status)
	}

	return data, nil

}
