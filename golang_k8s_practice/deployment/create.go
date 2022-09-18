package deployment

import (
	"context"
	"fmt"
	"io/ioutil"
	appV1 "k8s.io/api/apps/v1"
	coreV1 "k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/api/errors"
	metaV1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/util/intstr"
	"k8s.io/apimachinery/pkg/util/json"
	"k8s.io/apimachinery/pkg/util/yaml"
	"k8s.io/client-go/kubernetes"
	"log"
)

func CreateNamespace(ctx context.Context, client *kubernetes.Clientset, useNamespace string) {

	namespace := &coreV1.Namespace{
		ObjectMeta: metaV1.ObjectMeta{
			Name: useNamespace,
		},
	}

	if _, err := client.CoreV1().Namespaces().Get(ctx, useNamespace, metaV1.GetOptions{}); err != nil {
		if !errors.IsNotFound(err) {
			log.Println(err)
			return
		}

		res, err := client.CoreV1().Namespaces().Create(ctx, namespace, metaV1.CreateOptions{})
		if err != nil {
			log.Println(err)
			return
		}
		fmt.Println("namespace已经创建完成!", res.Name)
	}



}


func CreateDeploymentService(ctx context.Context, client *kubernetes.Clientset, namespace string, useReplicas int32, deployname string, useTargetport int32) {
	var replicas = useReplicas
	var targetPort = useTargetport
	intString := intstr.IntOrString{
		IntVal: targetPort,
	}

	deployment := &appV1.Deployment{
		ObjectMeta: metaV1.ObjectMeta{
			Name: deployname,
			Labels: map[string]string{
				"app": "nginx",
			},
			Namespace: namespace,
		},
		Spec: appV1.DeploymentSpec{
			Replicas: &replicas,
			Selector: &metaV1.LabelSelector{
				MatchLabels: map[string]string{
					"app": "nginx",
				},
			},
			Template: coreV1.PodTemplateSpec{
				ObjectMeta: metaV1.ObjectMeta{
					Name: "aaaa",
					Labels: map[string]string{
						"app": "nginx",
					},
				},
				Spec: coreV1.PodSpec{
					Containers: []coreV1.Container{
						{
							Name: "nginx",
							Image: "nginx:1.16.1",
							Ports: []coreV1.ContainerPort{
								{
									Name: "http",
									Protocol: coreV1.ProtocolTCP,
									ContainerPort: 80,
								},
							},
						},
					},
				},
			},
		},
	}

	service := &coreV1.Service{
		ObjectMeta: metaV1.ObjectMeta{
			Name: deployname,
			Labels: map[string]string{
				"app": "nginx",
			},
			Namespace: namespace,
		},
		Spec: coreV1.ServiceSpec{
			Type: coreV1.ServiceTypeNodePort,
			Ports: []coreV1.ServicePort{
				{
					Name: "nginx",
					Port: 80,
					TargetPort: intString,
					NodePort: 30088,
					Protocol: coreV1.ProtocolTCP,

				},
			},
			Selector: map[string]string{
				"app": "nginx",
			},

		},
	}

	// TODO: 这里需要做函数重构，预计把deployment service的逻辑搬出去，用一个私有方法写。

	if _, err := client.AppsV1().Deployments(namespace).Get(ctx, deployname, metaV1.GetOptions{}); err != nil {

		if !errors.IsNotFound(err) {
			log.Println(err)
			return
		}

		if deployments, err := client.AppsV1().Deployments(namespace).Create(ctx, deployment, metaV1.CreateOptions{}); err != nil {
			log.Println(err)
		} else {
			fmt.Println("Deployment已成功被建立:", deployments.Name)

		}



	} else {
		fmt.Println("此deployment已经存在，无法建立！")
	}


	if _, err := client.CoreV1().Services(namespace).Get(ctx, deployname, metaV1.GetOptions{}); err != nil {
		if !errors.IsNotFound(err) {
			log.Println(err)
			return
		}

		if service, err := client.CoreV1().Services(namespace).Create(ctx, service, metaV1.CreateOptions{}); err != nil {
			log.Println(err)
		} else {
			fmt.Println("Service已成功建立:", service.Name)
		}
	} else {
		fmt.Println("此service已经存在，无法建立！")
	}



}

func CreateDeploymentFromYaml(ctx context.Context, client *kubernetes.Clientset, namespace string ,yamlfile string) {

	yamlContent, err := ioutil.ReadFile(yamlfile)

	deploymentContainer := &appV1.Deployment{}
	deploymentContentJson, err := yaml.ToJSON(yamlContent)
	if err != nil {
		log.Println("YAML转换JSON格式错误!")
		return
	}

	if err = json.Unmarshal(deploymentContentJson, deploymentContainer); err != nil {
		log.Println("josn Unmarshal时转换错误！")
		return
	}


	if deployments, err := client.AppsV1().Deployments(namespace).Create(ctx, deploymentContainer, metaV1.CreateOptions{}); err != nil {
		log.Println(err)
		return
	} else {
		fmt.Println("deployment建立完成", deployments.Name)
	}


}

func CreateServiceFromYaml() {



}