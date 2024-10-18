package kube

import (
	"fmt"

	"github.com/choigonyok/home-idp/pkg/kube"
)

type GatewayKubeClient struct {
	Client *kube.KubeClient
}

func (c *GatewayKubeClient) Set(i interface{}) {
	c.Client = parseKubeClientFromInterface(i)
}

func (c *GatewayKubeClient) IsGatewayHealthy(namespace string) bool {
	return c.Client.IsServiceHealthy("home-idp-gateway", namespace)
}

func parseKubeClientFromInterface(i interface{}) *kube.KubeClient {
	client := i.(*kube.KubeClient)
	return client
}

func (c *GatewayKubeClient) GetNamespaces() *[]string {
	ns, err := c.Client.GetNamespaces()
	if err != nil {
		fmt.Println("TEST GET NAMESPACES ERR:", err)
		return nil
	}
	namespaces := []string{}

	for _, n := range *ns {
		namespaces = append(namespaces, n.Name)
	}

	return &namespaces
}
