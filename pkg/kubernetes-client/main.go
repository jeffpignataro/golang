package kubernetes

import (
	"context"

	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/rest"
)

type Client struct {
	name      string
	ClientSet *kubernetes.Clientset
}

func Init() (*Client, error) {
	// creates the in-cluster config
	config, err := rest.InClusterConfig()
	if err != nil {
		return nil, err
	}
	// creates the clientset
	clientset, err := kubernetes.NewForConfig(config)
	if err != nil {
		return nil, err
	}
	c := Client{
		name:      config.ServerName,
		ClientSet: clientset,
	}
	return &c, nil
}

func (client *Client) GetPods(namespace string) (*corev1.PodList, error) {
	cs := client.ClientSet
	pods, err := cs.CoreV1().Pods(namespace).List(context.TODO(), metav1.ListOptions{})
	if err != nil {
		return nil, err
	}
	return pods, nil
}
