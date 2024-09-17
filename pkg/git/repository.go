package git

import (
	"context"
	"fmt"

	"github.com/google/go-github/v63/github"
)

func (c *GitClient) GetRepositoryCloneURL() string {
	return c.Repository.GetCloneURL()
}

func (c *GitClient) CreateRepository(name string, private bool) error {
	r := &github.Repository{
		Name:                github.String(name),
		Private:             github.Bool(private),
		SecurityAndAnalysis: nil,
	}
	r, _, err := c.Client.Repositories.Create(context.TODO(), "", r)
	if err != nil {
		return err
	}

	c.Repository = r
	return nil
}

func DeleteRepository(name string, gc *GitClient) error {
	r, err := gc.Client.Repositories.Delete(context.TODO(), gc.Owner, name)
	if err != nil {
		fmt.Println("FAIL TO DELETE")
		fmt.Println(r.StatusCode)
		return err
	}

	fmt.Println("SUCCEED TO DELETE")
	return nil
}

func (c *GitClient) ConnectRepository(repo string) error {
	r, resp, err := c.Client.Repositories.Get(context.TODO(), c.Owner, repo)
	fmt.Println("TEST CONNECT GIT REPO STATUS CODE:", resp.StatusCode)
	if resp.StatusCode != 404 && err != nil {
		return err
	}

	if resp.StatusCode == 404 {
		c.CreateRepository(repo, false)
		return nil
	}

	c.Repository = r
	return nil
}

func DisonnectRepository(gc *GitClient) {
	if !checkRepository(gc) {

	}
}

func checkRepository(gc *GitClient) bool {
	if gc.Repository != nil {
		return true
	}
	return false
}
