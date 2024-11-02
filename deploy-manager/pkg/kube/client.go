package kube

import (
	"github.com/choigonyok/home-idp/deploy-manager/pkg/manifest"
	"github.com/choigonyok/home-idp/pkg/docker"
	"github.com/choigonyok/home-idp/pkg/env"
	"github.com/choigonyok/home-idp/pkg/kube"
	"sigs.k8s.io/yaml"
)

type DeployManagerKubeClient struct {
	Client *kube.KubeClient
}

func (c *DeployManagerKubeClient) Set(i interface{}) {
	c.Client = parseKubeClientFromInterface(i)
}

func parseKubeClientFromInterface(i interface{}) *kube.KubeClient {
	client := i.(*kube.KubeClient)
	return client
}

func (c *DeployManagerKubeClient) ApplyKanikoBuildJob(img, pusher, repo string) error {
	yml := manifest.GetKanikoJobManifest(docker.NewDockerImage(img, pusher), repo)
	b, _ := yaml.Marshal(yml)
	if err := c.Client.ApplyManifest(string(b), "jobs", yml.GetNamespace()); err != nil {
		return err
	}
	return nil
}

func (c *DeployManagerKubeClient) DeployManifest(manifest string) error {
	if err := c.Client.ApplyManifest(manifest, "applications", env.Get("HOME_IDP_NAMESPACE")); err != nil {
		return err
	}
	return nil
}
