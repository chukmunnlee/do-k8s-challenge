package main

import (
	"fmt"
	"log"
	"os"
	"path/filepath"

	"k8s.io/client-go/util/homedir"
)

type K8SClient struct {
}

func (k8s *K8SClient) SetupKubeConfig(kpath string) {
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
		log.Panicf("kubeconfig does not exists: %s\n", kubeConfig)
	}

	fmt.Printf(">>> kubeconfig: %s\n", kubeConfig)

}
