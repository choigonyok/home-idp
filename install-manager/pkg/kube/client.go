package kube

import (
	"github.com/choigonyok/home-idp/install-manager/pkg/manifest"
	"github.com/choigonyok/home-idp/pkg/env"
	"github.com/choigonyok/home-idp/pkg/kube"
	"sigs.k8s.io/yaml"
)

type InstallManagerKubeClient struct {
	Client *kube.KubeClient
}

func (c *InstallManagerKubeClient) Set(i interface{}) {
	c.Client = parseKubeClientFromInterface(i)
}

func parseKubeClientFromInterface(i interface{}) *kube.KubeClient {
	client := i.(*kube.KubeClient)
	return client
}

func (c *InstallManagerKubeClient) ApplyHarborCredentialSecret() error {
	yml := manifest.GetHarborCredManifest(env.Get("DEPLOY_MANAGER_REGISTRY_PASSWORD"))
	b, _ := yaml.Marshal(yml)
	if err := c.Client.ApplyManifest(string(b), "secrets", yml.GetNamespace()); err != nil {
		return err
	}
	return nil
}
