package client

import (
	"flag"
	"k8s.io/client-go/dynamic"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/rest"
	"k8s.io/client-go/tools/clientcmd"
	"log"
	"os"
	"path/filepath"
)


var K8sClient *kubernetes.Clientset
var DynamicClient dynamic.Interface
var RestClient *rest.RESTClient

func init() {
	var kubeConfig *string

	if home := HomeDir(); home != "" {
		kubeConfig = flag.String("kubeconfig", filepath.Join(home, ".kube", "config"), "")
	} else {
		kubeConfig = flag.String("kubeconfig", "", "")
	}
	flag.Parse()

	config, err := clientcmd.BuildConfigFromFlags("", *kubeConfig)
	if err != nil {
		log.Panic(err.Error())
	}

	// clientSet客户端：只能取得k8s内置资源
	clientSet, err := kubernetes.NewForConfig(config)
	if err != nil {
		return
	}
	K8sClient = clientSet

	// dynamicClient客户端：除内置资源以外，还可以取得外部或自定义资源
	dynamicClient, err := dynamic.NewForConfig(config)
	if err != nil {
		return
	}
	DynamicClient = dynamicClient

	restClient, err := rest.RESTClientFor(config)
	if err != nil {
		return
	}
	RestClient = restClient

}

func HomeDir() string {
	if h := os.Getenv("HOME"); h != "" {
		return h
	}

	return os.Getenv("USERPROFILE")

}

