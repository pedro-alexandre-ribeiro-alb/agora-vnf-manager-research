package kubernetes

import (
	log "agora-vnf-manager/core/log"
	"agora-vnf-manager/features/kubernetes/types"
	"context"

	v1 "k8s.io/api/core/v1"
	kubernetes "k8s.io/client-go/kubernetes"
	clientcmd "k8s.io/client-go/tools/clientcmd"
)

var Namespace = "dolt-testbed"

func InitializeKubernetesClientset(configuration_file string) (*kubernetes.Clientset, error) {
	configuration, err := clientcmd.BuildConfigFromFlags("", configuration_file)
	if err != nil {
		log.Errorf("[KubernetesService - InitializeKubernetesClientset]: %s", err.Error())
		return nil, err
	}
	clientset, err := kubernetes.NewForConfig(configuration)
	if err != nil {
		log.Errorf("[KubernetesService - InitializeKubernetesClientset]: %s", err.Error())
		return nil, err
	}
	return clientset, nil
}

func ListKubernetesPods(specification types.Specifications) ([]v1.Pod, error) {
	clientset, err := InitializeKubernetesClientset(specification.ConfigurationFile)
	if err != nil {
		log.Errorf("[KubernetesService - ListKubernetesPods]: %s", err.Error())
		return nil, err
	}
	list_options := MapSpecificationsToListOptions(specification)
	pods, err := clientset.CoreV1().Pods(Namespace).List(context.Background(), list_options)
	if err != nil {
		log.Errorf("[KubernetesService - ListKubernetesPods]: %s", err.Error())
		return nil, err
	}
	return pods.Items, nil
}

func ListContainers(specification types.Specifications) ([]v1.Container, error) {
	containers := []v1.Container{}
	pods, err := ListKubernetesPods(specification)
	if err != nil {
		log.Errorf("[KubernetesService - ListContainers]: %s", err.Error())
		return nil, err
	}
	for _, pod := range pods {
		pod_instances := pod.Spec.Containers
		for _, container := range pod_instances {
			containers = append(containers, container)
		}
	}
	return containers, nil
}
