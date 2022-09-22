package namespace

import (
	"context"
	"github.com/gin-gonic/gin"
	"k8s_practice_client_go/golang_k8s_practice/deployment"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

func ListNamespace(c *gin.Context) {
	clientSet, err := deployment.GetClientSet()
	if err != nil {
		_ = c.Error(err)
		return
	}
	nslist, err := clientSet.CoreV1().Namespaces().List(context.Background(), metav1.ListOptions{})
	if err != nil {
		_ = c.Error(err)
		return
	}
	c.JSON(200, nslist)
}
