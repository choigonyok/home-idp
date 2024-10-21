package kube

import (
	"fmt"
	"strings"

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

func (c *GatewayKubeClient) DeleteResources(resources, names, namespace string) error {
	list := strings.Split(names, ", ")
	switch resources {
	case "pods":
		if err := c.Client.DeletePods(&list, namespace); err != nil {
			return err
		}
	case "services":
		if err := c.Client.DeleteServices(&list, namespace); err != nil {
			return err
		}
	case "ingresses":
		if err := c.Client.DeleteIngresses(&list, namespace); err != nil {
			return err
		}
	case "secrets":
		if err := c.Client.DeleteSecrets(&list, namespace); err != nil {
			return err
		}
	case "configmaps":
		if err := c.Client.DeleteConfigmaps(&list, namespace); err != nil {
			return err
		}
	}

	return nil
}
