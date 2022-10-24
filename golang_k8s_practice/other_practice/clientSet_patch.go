package main

import (
	"context"
	_ "embed"
	"fmt"
	client2 "golang_k8s_practice/client"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime/schema"
	"k8s.io/apimachinery/pkg/types"
	"k8s.io/apimachinery/pkg/util/json"
	"log"
)

func main() {
	//PatchPractice()
	//PatchPractice2()
	//PatchPractice3()
	PatchPracticeWithDynamicClient()
}


func PatchPractice() {
	ctx := context.Background()
	clientSet := client2.K8sClient
	//动态客户端
	//dynamicClient := client2.DynamicClient
	var mgo, err = clientSet.AppsV1().Deployments("default").
		Get(ctx, "nginx-deployment-dynamic-create", metav1.GetOptions{})
	if err != nil {
		log.Fatalln(err)
	}
	//js 代码  {spec:{replicas:3}}
	frontPost := map[string]interface{} {
		"spec":map[string]interface{}{
			"replicas": 1,
		},
	}

	b,_ := json.Marshal(frontPost)

	// patch替换
	_, err = clientSet.AppsV1().Deployments(mgo.Namespace).
		Patch(ctx, mgo.Name, types.StrategicMergePatchType, b, metav1.PatchOptions{})

	if err != nil {
		log.Fatalln(err)
	}
	fmt.Println("patch替换成功")
}


func PatchPractice2() {
	ctx := context.Background()
	clientSet := client2.K8sClient
	//动态客户端
	//dynamicClient := client2.DynamicClient

	var mgo, err = clientSet.AppsV1().Deployments("default").
		Get(ctx, "nginx-deployment-dynamic-create", metav1.GetOptions{})
	if err != nil {
		log.Fatalln(err)
	}

	// 加入一个镜像
	addContainer := map[string]interface{}{
		"spec": map[string]interface{}{
			"template": map[string]interface{}{
				"spec": map[string]interface{}{
					"containers": []map[string]interface{}{
						{
							"name": "redis",
							//"image": "redis:5-alpine", // 加入操作
							//"$patch":"delete", 		 // 删除操作
						},
					},
				},
			},
		},
	}

	b,_ := json.Marshal(addContainer)

	// patch替换
	_, err = clientSet.AppsV1().Deployments(mgo.Namespace).
		Patch(ctx, mgo.Name, types.StrategicMergePatchType, b, metav1.PatchOptions{})

	if err != nil {
		log.Fatalln(err)
	}
	fmt.Println("patch 合并成功")
}



//   go:embed tpls/deploy.yaml
//var deploy_tpl string

/*
	json模式：patch有三种模式 add,replace,remove
	字符串格式
	[
		{"op":"replace", "path":"/xxx/xxx", "value":"xxxx"}
	]
 */

type JSONPatch struct {
	Op    string      `json:"op"`
	Path  string      `json:"path"`
	Value interface{} `json:"value,omitempty"`
}
type JSONPatchList []*JSONPatch

func AddJsonPatch(jps ...*JSONPatch) JSONPatchList {
	list := make([]*JSONPatch,len(jps))
	for index ,jp:=range jps{
		list[index]=jp
	}
	return list
}

func PatchPractice3() {
	ctx := context.Background()
	clientSet := client2.K8sClient

	 //获取 clientset
	var mgo, err = clientSet.AppsV1().Deployments("default").
		Get(ctx, "nginx-deployment-dynamic-create", metav1.GetOptions{})
	if err != nil {
		log.Fatalln(err)
	}
	// v1.Deployment{}

	list := AddJsonPatch( &JSONPatch {
		Op: "add",
		Path: "/spec/template/spec/containers/1", // 后面那个是索引
		Value: map[string]interface{}{
			"name":"redis",
			"image":"redis:5-alpine",
		},
	})


	b, _ := json.Marshal(list)
	fmt.Println(string(b))
	_, err = clientSet.AppsV1().Deployments(mgo.Namespace).
		Patch(ctx,mgo.Name,types.JSONPatchType,b,
			metav1.PatchOptions{})

	if err != nil {
		log.Fatalln(err)
	}
	fmt.Println("patch成功")
}

func PatchPracticeWithDynamicClient() {
	ctx := context.Background()
	//动态客户端
	dynamicClient := client2.DynamicClient
	// 区分下 add与remove的区别
	list := AddJsonPatch( &JSONPatch{
		Op: "remove",
		Path: "/spec/template/spec/containers/1",
		//Value: map[string]interface{}{
		//	"name":"redis",
		//	"image":"redis:5-alpine",
		//},
	})
	GVR := schema.GroupVersionResource{
		Resource: "deployments",
		Version: "v1",
		Group: "apps",
	}
	b, _ := json.Marshal(list)
	_, err := dynamicClient.Resource(GVR).Namespace("default").
		Patch(ctx,"nginx-deployment-dynamic-create",types.JSONPatchType,b,metav1.PatchOptions{})
	if err != nil{
		log.Fatalln(err)
	}
	fmt.Println("动态客户端patch成功")
}