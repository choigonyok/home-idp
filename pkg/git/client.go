package git

import (
	"context"
	"log"
	"net/http"

	"github.com/choigonyok/home-idp/pkg/env"
	"github.com/google/go-github/v63/github"
)

type GitClient struct {
	Owner      string
	Repository *github.Repository
	Client     *github.Client
	Token      string
}

func NewGitClient(owner, token string) *GitClient {
	dc := http.DefaultClient

	client := &GitClient{
		Owner:  owner,
		Client: github.NewClient(dc).WithAuthToken(token),
		Token:  token,
	}

	if err := client.ConnectRepository(env.Get("HOME_IDP_GIT_REPO")); err != nil {
		log.Fatalln("TEST CONNECT REPO ERR:", err)
	}

	return client
}

func (c *GitClient) CreateGitWebhook(url string) error {
	h := &github.Hook{
		Active: github.Bool(true),
		Config: &github.HookConfig{
			ContentType: github.String("json"),
			URL:         github.String(url),
		},
	}
	_, resp, err := c.Client.Repositories.CreateHook(context.TODO(), c.Owner, c.Repository.GetName(), h)
	defer resp.Body.Close()
	if err != nil {
		return err
	}

	return nil
}

func (c *GitClient) Pushfile(file GitFile) error {
	c.Client.Repositories.CreateFile(context.TODO(), file.getUsername(), env.Get("HOME_IDP_GIT_REPO"), file.getType(), &github.RepositoryContentFileOptions{
		Message: github.String(""),
		Content: []byte(""),
	})

	return nil
}

type GitFile interface {
	getType() string
	getUsername() string
}

type GitDockerFile struct {
	Username string
	Tag      string
}

func (f *GitDockerFile) getUsername() string {
	return f.Username
}
func (f *GitDockerFile) getType() string {
	return "docker"
}

// func ValidateClient(c *GitClient) {
// 	a, r, err := c.Client.Authorizations.Check(context.TODO(), c.Owner, c.Token)
// 	fmt.Println("USER:", a.User)
// 	fmt.Println("STATUSCODE:", r.StatusCode)
// 	fmt.Println("ERROR:", err)
// }

// owner := "choigonyok"
// repo := "argocd-apps"
// path := "file.txt"
// content := []byte("FILE_CONTENT")
// message := "Test: github package"

// options := &github.RepositoryContentFileOptions{
// 	Message: &message,
// 	Content: content,
// 	Branch:  github.String("main"),
// }

// // github.CommitsListOptions
// // github.CreateCommitOptions

// _, _, err := gc.Repositories.CreateFile(context.TODO(), owner, repo, path, options)
// if err != nil {
// 	fmt.Println("PUSH FAIL")
// 	return
// }
// fmt.Println("PUSH SUCCEED")
// return
