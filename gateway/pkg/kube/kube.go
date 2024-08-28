package kube

import "github.com/choigonyok/home-idp/pkg/kube"

type GatewayKubeClient struct {
	Client *kube.KubeClient
}

func (c *GatewayKubeClient) Set(i interface{}) {
	c.Client = parseKubeClientFromInterface(i)
}

func (c *GatewayKubeClient) IsGatewayHealthy(namespace string) bool {
	if c.Client.IsServiceHealthy("home-idp-gateway", namespace) {
		return true
	}

	return false
}

func parseKubeClientFromInterface(i interface{}) *kube.KubeClient {
	client := i.(*kube.KubeClient)
	return client
}
