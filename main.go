package main

import (
	"context"
	"flag"
	"log"
	"path/filepath"

	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"

	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/tools/clientcmd"
	"k8s.io/client-go/util/homedir"
)

func GetListOfPodNames(clientset kubernetes.Interface) []string {

	listOptions := metav1.ListOptions{}
	pods, err := clientset.CoreV1().Pods("default").List(context.TODO(), listOptions)

	if err != nil {
		log.Println("Error fetching pods")
	}

	var podNameList []string

	for _, pod := range pods.Items {
		// log.Println("Pod Name:", pod.GetName())
		if podNameList != nil {
			podNameList = append(podNameList, pod.GetName())
		} else {
			podNameList = []string{pod.GetName()}
		}
	}

	return podNameList
}

func main() {
	// Build default
	home := homedir.HomeDir()
	defaultConfig := filepath.Join(home, ".kube", "config")

	// Get from flags or default
	kubeconfig := flag.String("kubeconfig", defaultConfig, "Path to .kube/config")
	flag.Parse()

	// Build from flags
	config, err := clientcmd.BuildConfigFromFlags("", *kubeconfig)
	if err != nil {
		panic(err)
	}

	// Create clientset from config
	clientset, err := kubernetes.NewForConfig(config)
	if err != nil {
		panic(err)
	}

	// Get Pods
	var result []string = GetListOfPodNames(clientset)

	for _, name := range result {
		log.Println(name)
	}
}
