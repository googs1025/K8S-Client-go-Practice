package main

import (
	"context"
	"fmt"
	"golang_k8s_practice/client"
	v1 "k8s.io/api/apps/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime/schema"
	"k8s.io/apimachinery/pkg/util/json"
	"log"
)

func main() {

	// k8s中 对象的GVR
	deployGVR := schema.GroupVersionResource{
		Group: "apps",
		Version: "v1",
		Resource: "deployments",
	}

	// 注意：用动态客户端转：会是unstructured对象 需要反序列化对象
	deployListUnstructured, err := client.DynamicClient.Resource(deployGVR).Namespace("default").
		List(context.Background(), metav1.ListOptions{})
	if err != nil {
		log.Println(err)
	}
	//for _, deploy := range deployList_unstructured.Items {
	//	fmt.Println(deploy.GetName())
	//}

	// 需要用marshal转换一下
	b, err := deployListUnstructured.MarshalJSON()	// json字符串

	deployList := v1.DeploymentList{}	// 用来捞结果

	err = json.Unmarshal(b, &deployList)
	if err != nil {
		return
	}
	for _, deploy := range deployList.Items {
		fmt.Println(deploy.Name)
	}




}
