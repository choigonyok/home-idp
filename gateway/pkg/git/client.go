package git

import (
	"strings"

	"github.com/choigonyok/home-idp/gateway/pkg/manifest"
	"github.com/choigonyok/home-idp/pkg/env"
	"github.com/choigonyok/home-idp/pkg/git"
)

type GatewayGitClient struct {
	Client *git.GitClient
}

func (c *GatewayGitClient) Set(i interface{}) {
	c.Client = parseGitClientFromInterface(i)
}

func parseGitClientFromInterface(i interface{}) *git.GitClient {
	client := i.(*git.GitClient)
	return client
}

func (c *GatewayGitClient) CreateGithubWebhook() error {
	if env.Get("HOME_IDP_TLS_ENABLED") == "true" {
		return c.Client.CreateGitWebhook("https://" + env.Get("HOME_IDP_HOST") + "/webhooks/github")
	} else {
		return c.Client.CreateGitWebhook("http://" + env.Get("HOME_IDP_HOST") + "/webhooks/github")
	}
}

// func (c *GatewayGitClient) CreateManifest(username, email string) error {
// 	return c.Client.CreateFilesByFiletype(username, email, env.Get("HOME_IDP_NAMESPACE"), git.Manifest)
// }

// /github-webhook/"
func (c *GatewayGitClient) GetRepositoryCloneURL() string {
	return c.Client.GetRepositoryCloneURL()
}

func (c *GatewayGitClient) PushFile(username, tag, content string) error {
	return c.Client.Pushfile(
		&git.GitDockerFile{
			Username: username,
			Tag:      tag,
			Content:  content,
		},
	)
}

func (c *GatewayGitClient) UpdateImageVersion(username, email, before, after string) {
	c.Client.UpdateFilesByFiletype(username, email, before, after, git.Manifest)
}

func (c *GatewayGitClient) CreatePodManifestFile(username, email, image string, port int) error {
	name, _, _ := strings.Cut(image, ":")
	manifest := manifest.GetPodManifest(name+"-"+username, image, port)
	return c.Client.CreateFilesByFiletype(username, email, env.Get("HOME_IDP_NAMESPACE"), "pod.yaml", []byte(manifest), git.Manifest)
}