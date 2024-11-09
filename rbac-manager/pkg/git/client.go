package git

import (
	"context"
	"fmt"

	"github.com/choigonyok/home-idp/pkg/git"
)

type RbacGitClient struct {
	Client *git.GitClient
}

func (c *RbacGitClient) Set(i interface{}) {
	c.Client = parseGitClientFromInterface(i)
}

func parseGitClientFromInterface(i interface{}) *git.GitClient {
	client := i.(*git.GitClient)
	return client
}

func (c *RbacGitClient) GetAdminGithubId() int64 {
	user, resp, err := c.Client.Client.Users.Get(context.TODO(), "")
	if err != nil {
		fmt.Println("ERR GETTING ADMIN USER GITHUB ID:", err)
		return 0
	}
	fmt.Println("ADMIN USER GITHUB ID RESPONSE STATUS CODE:", resp.StatusCode)

	return *user.ID
}
