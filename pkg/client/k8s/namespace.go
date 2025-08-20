package client

import (
	"context"

	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

func (k *K8s) ListNamespaces() ([]string, error) {
	nsList, err := k.clientset.CoreV1().Namespaces().List(context.Background(), metav1.ListOptions{})
	if err != nil {
		return nil, err
	}
	names := []string{}
	for _, n := range nsList.Items {
		names = append(names, n.Name)
	}
	return names, nil
}
