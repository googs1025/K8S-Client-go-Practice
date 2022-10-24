package main

import (
	"context"
	_ "embed"
	"fmt"
	"golang_k8s_practice/client"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	"k8s.io/apimachinery/pkg/runtime/schema"
	"k8s.io/apimachinery/pkg/util/yaml"
	"log"
)

/*
	使用yaml文件，创建deployment。
 */

//go:embed deployment.yaml
var deploy_tpl string

func main() {

	deployGVR := schema.GroupVersionResource{
		Group: "apps",
		Version: "v1",
		Resource: "deployments",
	}

	// 一般客户端可以直接用 v1.Deployment{}
	// 动态客户端使用
	deployObj := &unstructured.Unstructured{}
	err := yaml.Unmarshal([]byte(deploy_tpl), deployObj)
	if err != nil {
		log.Println("unmarshal err:", err)
	}
	fmt.Println(deploy_tpl)
	_, err = client.DynamicClient.Resource(deployGVR).
		Namespace(deployObj.GetNamespace()).
		Create(context.Background(), deployObj, metav1.CreateOptions{})

	if err != nil {
		log.Println("create err:", err)
	}



}