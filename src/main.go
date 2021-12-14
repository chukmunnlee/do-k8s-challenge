package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {

	opt := CLIOptions{Port: 3000}
	opt.ParseOptions()

	k8sClient := K8SClient{}

	log.Println("Trying in-cluster initialization")
	if err := k8sClient.InClusterConfig(); nil != err {
		log.Println("Trying out-cluster initialization")
		if err := k8sClient.OutClusterConfig(opt.KubeConfig); nil != err {
			log.Panicf("Configuration error: %v\n", err)
		}
	}

	log.Println("Initiazing client")
	if err := k8sClient.InitializeClient(); nil != err {
		log.Panicf("Cannot create clientset: %v\n", err)
	}

	http.HandleFunc("/", HandleRoot)
	http.HandleFunc(opt.Path, HandleMutate(k8sClient))

	log.Printf(fmt.Sprintf("Listening on port %d, path: %s\n", opt.Port, opt.Path))
	if err := http.ListenAndServe(fmt.Sprintf(":%d", opt.Port), nil); nil != err {
		log.Panicf("Cannot start dns-update: %v\n", err)
	}
}

func HandleRoot(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("dns-edit admission controller"))
}
