package deployment

import (
	"context"
	"log"
	"testing"
)

func TestDeployment(t *testing.T) {

	ctx := context.Background()

	client, err := GetClientSet()
	if err != nil {
		log.Println(err)
	}
	GetPodsList(ctx, client, "default")
	GetPod(ctx, client, "default", "stress-pod")
	GetDeployment(ctx, client, "default", "webapp")
	GetDeploymentList(ctx, client, "default")


}

