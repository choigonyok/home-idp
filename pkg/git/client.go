package git

import (
	"net/http"

	"github.com/google/go-github/v63/github"
)

type GitClient struct {
	Owner      string
	Repository *github.Repository
	Client     *github.Client
	Token      string
}

func NewClient(owner, token string) *GitClient {
	dc := http.DefaultClient
	return &GitClient{
		Owner:      owner,
		Client:     github.NewClient(dc).WithAuthToken(token),
		Repository: nil,
		Token:      token,
	}
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
