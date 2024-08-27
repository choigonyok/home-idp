package git

import (
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

func (c *GatewayGitClient) UpdateArgoCDApplicationManifest(username, email, before, after string) {
	c.Client.UpdateFilesByFiletype(username, email, before, after, git.CD)
}

// func (c *GatewayGitClient) CreateArgoCDApplicationManifest(username, email, image string) {
// 	c.Client.CreateFilesByFiletype(username, email, image, git.CD)
// }
