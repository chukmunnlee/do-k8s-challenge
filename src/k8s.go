package main

import (
	ctx "context"
	"fmt"
	"os"
	"path/filepath"

	"k8s.io/apimachinery/pkg/apis/meta/v1"
	k8s "k8s.io/client-go/kubernetes"
	"k8s.io/client-go/rest"
	"k8s.io/client-go/tools/clientcmd"
	"k8s.io/client-go/util/homedir"
)

type K8SClient struct {
	KubeConfig   string
	ClientConfig *rest.Config
	Clientset    *k8s.Clientset
}

func (client *K8SClient) GetPods(ns string) ([]string, error) {
	podList, err := client.Clientset.CoreV1().Pods(ns).List(ctx.TODO(), v1.ListOptions{})
	if nil != err {
		return nil, err
	}

	var podNames []string
	for _, p := range podList.Items {
		podNames = append(podNames, p.ObjectMeta.Name)
	}

	return podNames, nil
}

func (client *K8SClient) InClusterConfig() error {
	config, err := rest.InClusterConfig()
	if nil != err {
		return err
	}
	client.ClientConfig = config
	return nil
}

func (client *K8SClient) OutClusterConfig(kpath string) error {
	kubeConfig := kpath

	// Check KUBECONFIG first
	if "" == kubeConfig {
		kubeConfig = os.Getenv("KUBECONFIG")
	}

	// check $HOME/.kube/config
	if "" == kubeConfig {
		kubeConfig = filepath.Join(homedir.HomeDir(), ".kube", "config")
	}

	if _, err := os.Stat(kubeConfig); nil != err {
		return fmt.Errorf("kubeconfig does not exists: %s", kubeConfig)
	}

	client.KubeConfig = kubeConfig

	config, err := clientcmd.BuildConfigFromFlags("", client.KubeConfig)
	if nil != err {
		return err
	}

	client.ClientConfig = config

	return nil
}

func (client *K8SClient) InitializeClient() error {

	cs, err := k8s.NewForConfig(client.ClientConfig)
	if nil != err {
		return err
	}

	client.Clientset = cs

	return nil
}
