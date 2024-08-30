package git

import (
	"github.com/choigonyok/home-idp/pkg/git"
)

type DeployManagerGitClient struct {
	Client *git.GitClient
}

func (c *DeployManagerGitClient) Set(i interface{}) {
	c.Client = parseGitClientFromInterface(i)
}

func parseGitClientFromInterface(i interface{}) *git.GitClient {
	client := i.(*git.GitClient)
	return client
}

func (c *DeployManagerGitClient) GetManifestFromGithub(filepath string) string {
	content := c.Client.GetFilesByPath(filepath)
	return content[0]
}
