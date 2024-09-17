package git

import (
	"github.com/choigonyok/home-idp/pkg/git"
	"github.com/choigonyok/home-idp/pkg/manifest"
)

type InstallManagerGitClient struct {
	Client *git.GitClient
}

func (c *InstallManagerGitClient) Set(i interface{}) {
	c.Client = parseGitClientFromInterface(i)
}

func parseGitClientFromInterface(i interface{}) *git.GitClient {
	client := i.(*git.GitClient)
	return client
}

func (c *InstallManagerGitClient) CreateArgoCDApplicationManifest(username, email, namespace string) {
	c.Client.CreateFilesByFiletype(username, email, namespace, "app.yaml", manifest.GetArgoCDManifest(username, namespace), git.CD)
}
