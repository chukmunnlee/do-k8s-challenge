package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {

	opt := CLIOptions{ Port: 3000 }
	opt.ParseOptions()

	k8sClient := K8SClient{}
	k8sClient.SetupKubeConfig(opt.KubeConfig)

	http.HandleFunc("/", HandleRoot)
	http.HandleFunc(opt.Path, HandleMutate)

	log.Printf(fmt.Sprintf("Listening on port %d, path: %s\n", opt.Port, opt.Path))
	if err := http.ListenAndServe(fmt.Sprintf(":%d", opt.Port), nil); nil != err {
		log.Panicf("Cannot start dns-update: %v\n", err)
	}
}

func HandleMutate(w http.ResponseWriter, r *http.Request) {
}

func HandleRoot(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("dns-edit admission controller"))
}
