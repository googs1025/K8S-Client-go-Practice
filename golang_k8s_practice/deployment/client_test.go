package deployment

import (
	"context"
	"fmt"
	"testing"
	"time"
)

func TestDeployment(t *testing.T) {

	ctx := context.Background()

	client, _ := GetClientSet()
	//client := K8sClient

	GetPodsList(ctx, client, "default")
	fmt.Println("-----------------------------------------------")
	GetPod(ctx, client, "default", "stress-pod")
	fmt.Println("-----------------------------------------------")
	GetDeployment(ctx, client, "default", "webapp")
	fmt.Println("-----------------------------------------------")
	GetDeploymentList(ctx, client, "default")
	fmt.Println("-----------------------------------------------")
	CreateDeploymentService(ctx, client, "default", 3, "nginx333333", 80)
	fmt.Println("-----------------------------------------------")
	//CreateNamespace(ctx, client, "create-test")
	fmt.Println("-----------------------------------------------")

	// TODO: 记录一下，重复调用GetClientSet()好像会报错，需要用func init() 来执行 进行初始化
	WatchDeployment(ctx, client, "default", time.Duration(time.Minute * 1))
}

//func TestWatchDeployment(t *testing.T) {
//
//	//ctx := context.Background()
//	//client, err := GetClientSet()
//	//if err != nil {
//	//	log.Println(err)
//	//}
//	//
//	//fmt.Println("-----------------------------------------------")
//	//WatchDeployment(ctx, client, "default", time.Duration(time.Minute * 1))
//
//}

