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
	yml := manifest.GetHarborCredManifest(env.Get("HOME_IDP_ADMIN_PASSWORD"))
	b, _ := yaml.Marshal(yml)
	if err := c.Client.ApplyManifest(string(b), "secrets", yml.GetNamespace()); err != nil {
		return err
	}
	return nil
}

func (c *InstallManagerKubeClient) IsArgoCDRunning(ns string) bool {
	if c.Client.IsServiceHealthy("home-idp-cd-argocd-applicationset-controller", ns) &&
		c.Client.IsServiceHealthy("home-idp-cd-argocd-repo-server", ns) &&
		c.Client.IsServiceHealthy("home-idp-cd-argocd-server", ns) {
		return true
	}
	return false
}

func (c *InstallManagerKubeClient) GetArgoCDPassword(ns string) string {
	b := c.Client.GetSecret("argocd-initial-admin-secret", ns, "password")
	// pw, err := base64.StdEncoding.DecodeString(string(b))
	// fmt.Println("TEST BASE64 DECODE ARGOCD PASSWORD ERR:", err)
	// fmt.Println("TEST ARGOCD PASSWORD:", string(pw))
	return string(b)
}
