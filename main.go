package main

import (
	"fmt"
	"log"
	"os"
	"path/filepath"

	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/tools/clientcmd"
)

func main() {
	kubeconfig := filepath.Join(os.Getenv("HOME"), ".kube", "config")
	config, err := clientcmd.BuildConfigFromFlags("", kubeconfig)
	if err != nil {
		log.Fatal(err)
	}
	kc, err := kubernetes.NewForConfig(config)
	if err != nil {
		panic(err)
	}
	pods, err := kc.CoreV1().Pods("kube-system").List(v1.ListOptions{})
	if err != nil {
		log.Fatal(err)
	}
	for _, pod := range pods.Items {
		fmt.Println(pod.Name, "<===>", pod.Namespace)
	}
}
