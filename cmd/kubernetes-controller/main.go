package main

import (
	"log"

	"github.com/jeffpignataro/golang/pkg/kubernetes-client"
)

func main() {
	client, err := kubernetes.Init()
	if err != nil {
		panic(err.Error())
	}
	p, err := client.GetPods("actions-runner-system")
	if err != nil {
		panic(err.Error())
	}
	log.Printf("There are %d pods in the cluster\n", len(p.Items))
}
