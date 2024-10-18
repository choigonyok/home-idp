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
	if c.Client.IsServiceHealthy("home-idp-gateway", namespace) {
		return true
	}

	return false
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

func (c *GatewayKubeClient) GetPods(namespace string) *[]string {
	ps, err := c.Client.GetPods(namespace)
	if err != nil {
		fmt.Println("TEST GET PODS FOR NAMESPACE "+namespace+" ERR:", err)
		return nil
	}

	pods := []string{}

	for _, p := range *ps {
		pods = append(pods, p.Name)
	}

	return &pods
}

func (c *GatewayKubeClient) GetServices(namespace string) *[]string {
	svcs, err := c.Client.GetServices(namespace)
	if err != nil {
		fmt.Println("TEST GET SERVICES FOR NAMESPACE "+namespace+" ERR:", err)
		return nil
	}

	services := []string{}

	for _, svc := range *svcs {
		services = append(services, svc.Name)
	}

	return &services
}

func (c *GatewayKubeClient) GetIngresses(namespace string) *[]string {
	i, err := c.Client.GetIngresses(namespace)
	if err != nil {
		fmt.Println("TEST GET INGRESSES FOR NAMESPACE "+namespace+" ERR:", err)
		return nil
	}

	ingresses := []string{}

	for _, ingress := range *i {
		ingresses = append(ingresses, ingress.Name)
	}

	return &ingresses
}

func (c *GatewayKubeClient) GetConfigmaps(namespace string) *[]string {
	cms, err := c.Client.GetConfigmaps(namespace)
	if err != nil {
		fmt.Println("TEST GET CONFIGMAPS FOR NAMESPACE "+namespace+" ERR:", err)
		return nil
	}

	configmaps := []string{}

	for _, configmap := range *cms {
		configmaps = append(configmaps, configmap.Name)
	}

	return &configmaps
}

func (c *GatewayKubeClient) GetSecrets(namespace string) *[]string {
	s, err := c.Client.GetSecrets(namespace)
	if err != nil {
		fmt.Println("TEST GET SECRETS FOR NAMESPACE "+namespace+" ERR:", err)
		return nil
	}

	secrets := []string{}

	for _, secret := range *s {
		secrets = append(secrets, secret.Name)
	}

	return &secrets
}
