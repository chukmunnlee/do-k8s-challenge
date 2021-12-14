package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

func HandleMutate(k8sClient K8SClient) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		body, err := ioutil.ReadAll(r.Body)
		if nil != err {
			w.WriteHeader(http.StatusBadRequest)
			w.Write([]byte(fmt.Sprintf("Cannot read request body: %v", err)))
			return
		}

		log.Printf(">> body: %s", string(body))

		w.WriteHeader(http.StatusBadRequest)
		w.Write(body)
		/*
		pods, err := k8sClient.GetPods("kube-system")
		if nil != err {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte(fmt.Sprintf("error: %v", err)))
			return
		}

		w.WriteHeader(http.StatusOK)
		w.Write([]byte(strings.Join(pods, "\n")))
		*/
	}
}
