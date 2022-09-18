package deployment

import (
	"context"
	"fmt"
	"log"
	"testing"
	"time"
)

func TestDeployment(t *testing.T) {

	ctx := context.Background()

	client, err := GetClientSet()
	if err != nil {
		log.Println(err)
	}
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
	CreateNamespace(ctx, client, "create-test")
}

func TestWatchDeployment(t *testing.T) {

	ctx := context.Background()
	client, err := GetClientSet()
	if err != nil {
		log.Println(err)
	}

	fmt.Println("-----------------------------------------------")
	WatchDeployment(ctx, client, "default", time.Duration(time.Minute * 1))

}

