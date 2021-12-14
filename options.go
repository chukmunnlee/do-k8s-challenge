package main

import "flag"

type CLIOptions struct {
	Port int
	Path string
	KubeConfig string
}

func (o *CLIOptions)ParseOptions() {

	flag.IntVar(&o.Port, "port", 8443, "port number")
	flag.StringVar(&o.Path, "path", "/mutate", "admission controller path")
	flag.StringVar(&o.KubeConfig, "kubeconfig", "", "kubeconfig path")
	flag.Parse()
}
