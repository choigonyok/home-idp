package git

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"strconv"
	"time"

	"github.com/choigonyok/home-idp/pkg/env"
	"github.com/google/go-github/v63/github"
)

type GitClient struct {
	Owner      string
	Repository *github.Repository
	Client     *github.Client
	Token      string
	Email      string
}

func NewGitClient(owner, email, token string) *GitClient {
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

func (c *GitClient) CreateFile(file GitFile) error {
	fmt.Println("TEST START CREATE FILE!")
	fmt.Println("TEST START CREATE FILE!")
	fmt.Println("TEST START CREATE FILE!")
	_, resp, err := c.Client.Repositories.CreateFile(context.TODO(), env.Get("HOME_IDP_GIT_USERNAME"), env.Get("HOME_IDP_GIT_REPO"), file.getType()+"/"+file.getUsername()+"/"+file.getFilename(), &github.RepositoryContentFileOptions{
		Message: github.String("create(" + file.getType() + "): " + file.getFilename() + " by " + file.getUsername()),
		Content: []byte(file.getContent()),
	})
	fmt.Println("TEST CREATE FILE STATUS CODE:", resp.StatusCode)
	fmt.Println("TEST CREATE FILE ERR:", err)

	return nil
}

func (c *GitClient) DeleteFile(file GitFile) error {
	files, found, err := c.getFilesByFiletype(Docker)
	if !found {
		return fmt.Errorf("NOT FOUND")
	}
	if err != nil {
		return err
	}

	_, resp, err := c.Client.Repositories.DeleteFile(context.TODO(), env.Get("HOME_IDP_GIT_USERNAME"), env.Get("HOME_IDP_GIT_REPO"), file.getType()+"/"+file.getUsername()+"/"+file.getFilename(), &github.RepositoryContentFileOptions{
		Message: github.String("remove(" + file.getType() + "): " + file.getFilename() + " by " + file.getUsername()),
		Content: []byte(file.getContent()),
		SHA:     files[0].SHA,
	})
	fmt.Println("TEST DELETE FILE STATUS CODE:", resp.StatusCode)
	fmt.Println("TEST DELETE FILE ERR:", err)
	return nil
}

func (c *GitClient) Listfile(dir, username string) []string {
	_, files, resp, err := c.Client.Repositories.GetContents(context.TODO(), c.Owner, *c.Repository.Name, dir+"/"+username, &github.RepositoryContentGetOptions{Ref: "main"})
	fmt.Println("TEST LIST FILE STATUS CODE:", resp.StatusCode)
	fmt.Println("TEST LIST FILE ERR:", err)

	list := []string{}

	for _, f := range files {
		list = append(list, f.GetName())
	}

	return list
}

// m is map[NewFile]OldFile
func (c *GitClient) UpdateFile(m map[GitFile]GitFile) error {
	ref, _, err := c.Client.Git.GetRef(context.TODO(), c.Owner, *c.Repository.Name, "refs/heads/main")
	if err != nil {
		return err
	}

	fmt.Println("TEST1")
	entries := []*github.TreeEntry{}

	filetype := ""
	username := ""

	for new, old := range m {
		fmt.Println("TEST GET NEW FILENAME:", new.getFilename())
		fmt.Println("TEST GET OLD FILENAME:", old.getFilename())
		filetype = new.getType()
		username = new.getUsername()

		fmt.Println("TEST OLD PATH:", old.getType()+"/"+old.getUsername()+"/"+old.getFilename())
		fmt.Println("TEST NEW PATH:", new.getType()+"/"+new.getUsername()+"/"+new.getFilename())
		fmt.Println("TEST2")

		// 	{
		// 		Path:    github.String(filePath),
		// 		Type:    github.String("blob"),
		// 		Mode:    github.String("100644"),
		// 		Content: github.String(newContent),  // 수정된 내용
		// },

		if old.getFilename() == new.getFilename() {
			entries = append(entries,
				&github.TreeEntry{
					Path:    github.String(old.getType() + "/" + old.getUsername() + "/" + old.getFilename()),
					Type:    github.String("blob"),
					Mode:    github.String("100644"),
					Content: github.String(new.getContent()),
				},
			)
		} else {
			entries = append(entries,
				&github.TreeEntry{
					Path: github.String(old.getType() + "/" + old.getUsername() + "/" + old.getFilename()),
					Type: github.String("blob"),
					Mode: github.String("100644"),
					SHA:  nil,
				},
				&github.TreeEntry{
					Path:    github.String(new.getType() + "/" + new.getUsername() + "/" + new.getFilename()),
					Type:    github.String("blob"),
					Mode:    github.String("100644"),
					Content: github.String(new.getContent()),
				},
			)
			fmt.Println("TEST3")
		}
	}

	baseTree, _, err := c.Client.Git.GetTree(context.TODO(), c.Owner, *c.Repository.Name, *ref.Object.SHA, true)
	if err != nil {
		return err
	}
	fmt.Println("TEST4")
	fmt.Println("TEST TREE ENTRIES LENGTH: ", len(entries))

	newTree, _, err := c.Client.Git.CreateTree(context.TODO(), c.Owner, *c.Repository.Name, baseTree.GetSHA(), entries)
	if err != nil {
		return err
	}
	fmt.Println(newTree.String())
	fmt.Println("TEST5")
	parent, _, err := c.Client.Repositories.GetCommit(context.TODO(), c.Owner, *c.Repository.Name, *ref.Object.SHA, nil)
	if err != nil {
		return err
	}

	parent.Commit.SHA = parent.SHA

	fmt.Println("TEST6")
	t := &github.Timestamp{}
	t.Time = time.Now()

	newCommit := &github.Commit{
		Message: github.String(`update(` + filetype + `): by ` + username),
		Tree:    newTree,
		Parents: []*github.Commit{parent.Commit},
		Author: &github.CommitAuthor{
			Name:  github.String("choigonyok"),
			Email: github.String("achoistic98@naver.com"),
			Date:  t,
		},
		Committer: &github.CommitAuthor{
			Name:  github.String("choigonyok"),
			Email: github.String("achoistic98@naver.com"),
			Date:  t,
		},
	}
	fmt.Println("TEST7")

	// newCommit.Parents[0].SHA
	fmt.Println("TEST RECENT COMMIT SHA :", parent.Commit.GetSHA())
	fmt.Println("TEST RECENT SHA  :", parent.GetSHA())

	for i, parent := range newCommit.Parents {
		fmt.Println("TEST PARENT COMMIT FIELD "+strconv.Itoa(i)+": ", parent.String())
	}
	fmt.Println("TEST7 TREE.SHA:", newCommit.Tree.SHA)
	// newCommit.Parents[0].SHA
	commit, _, err := c.Client.Git.CreateCommit(context.TODO(), c.Owner, *c.Repository.Name, newCommit, nil)
	if err != nil {
		return err
	}
	fmt.Println("TEST8")
	ref.Object.SHA = commit.SHA

	fmt.Println("TEST NEW COMMIT SHA: ", *commit.SHA)

	_, _, err = c.Client.Git.UpdateRef(context.TODO(), c.Owner, *c.Repository.Name, ref, false)
	if err != nil {
		return err
	}

	return nil
}

type GitFile interface {
	getType() string
	getUsername() string
	getFilename() string
	getContent() string
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
