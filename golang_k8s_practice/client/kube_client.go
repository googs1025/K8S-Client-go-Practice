package client

import (
	"flag"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/tools/clientcmd"
	"log"
	"os"
	"path/filepath"
)


var K8sClient *kubernetes.Clientset

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

	clientSet, err := kubernetes.NewForConfig(config)
	if err != nil {
		return
	}
	K8sClient = clientSet
}

func HomeDir() string {
	if h := os.Getenv("HOME"); h != "" {
		return h
	}

	return os.Getenv("USERPROFILE")

}

