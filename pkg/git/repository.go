package git

import (
	"context"
	"fmt"

	"github.com/google/go-github/v63/github"
)

func CreateRepository(name string, private bool, gc *GitClient) error {
	r := &github.Repository{
		Name:    github.String(name),
		Private: github.Bool(private),
	}
	r, _, err := gc.Client.Repositories.Create(context.TODO(), "", r)
	if err != nil {
		return err
	}

	gc.Repository = r
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

func ConnectRepository(gc *GitClient, repo string) error {
	if !checkRepository(gc) {
		r, resp, err := gc.Client.Repositories.Get(context.TODO(), gc.Owner, repo)
		fmt.Println("REPO GET ERR:", err)
		fmt.Println("REPO GET RESP:", resp.StatusCode)
		gc.Repository = r
		return nil
	}
	return fmt.Errorf("%s", "Already Connected")
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
