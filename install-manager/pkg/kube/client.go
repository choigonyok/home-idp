package kube

import (
	"encoding/base64"
	"fmt"

	"github.com/choigonyok/home-idp/install-manager/pkg/manifest"
	"github.com/choigonyok/home-idp/pkg/env"
	"github.com/choigonyok/home-idp/pkg/kube"
	corev1 "k8s.io/api/core/v1"
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

func (c *InstallManagerKubeClient) IsArgoCDRunning() bool {
	label := c.Client.GetServiceSelectors("testname", "testns")
	fmt.Println("TEST LABEL:", label)

	pods := c.Client.GetPodsByLabel("testns", label)

	for _, pod := range pods {
		if pod.Status.Phase != corev1.PodRunning {
			return false
		}
	}

	return true
}

func (c *InstallManagerKubeClient) GetArgoCDPassword() string {
	b := c.Client.GetSecret("argocd-initial-admin-secret", env.Get("HOME_IDP_NAMESPACE"), "password")
	pw := []byte{}
	base64.StdEncoding.Decode(b, pw)
	return string(pw)
}
