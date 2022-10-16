package service

import (
	"context"
	"golang_k8s_practice/client"
	"golang_k8s_practice/model"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"log"
	"fmt"
)


func ListService(namespace string) (*model.ServiceList, error) {
	ctx := context.Background()
	services, err := client.K8sClient.CoreV1().Services(namespace).List(ctx, metav1.ListOptions{})
	if err != nil {
		data := &model.ServiceList{
			Err: nil,
		}
		log.Println(err)
		return data, err
	}
	data := &model.ServiceList{
		Services: make([]*model.ServiceGet, 0),
		Err: nil,
	}

	for _, d := range services.Items {
		svc := &model.ServiceGet{
			Name: d.Name,
			Port: d.Spec.Ports[0].Port,
			Type: d.TypeMeta.Kind,
			ClusterIp: d.Spec.ClusterIP,
		}
		data.Services = append(data.Services, svc)
		fmt.Printf("Service Name: %v\t Port: %v\t Type: %+v\t ClusterIp: %v\t", d.Namespace, d.Name, d.Status, d.Spec.ClusterIP)
	}

	return data, nil


}

func GetService() {

}



