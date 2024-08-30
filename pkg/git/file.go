package git

import (
	"context"
	"fmt"
	"strings"
	"time"

	"github.com/google/go-github/v63/github"
)

type GitFileType string

const (
	CI       GitFileType = "ci"
	CD       GitFileType = "cd"
	Docker   GitFileType = "docker"
	Manifest GitFileType = "manifest"
)

func (c *GitClient) isFileExist(filePath string) (bool, error) {
	opts := &github.RepositoryContentGetOptions{
		Ref: "main",
	}

	_, _, resp, _ := c.Client.Repositories.GetContents(context.TODO(), c.Owner, *c.Repository.Name, filePath, opts)
	switch resp.StatusCode {
	case 200:
		return true, nil
	case 404:
		return false, nil
	default:
		fmt.Println("TEST GET GIT CONTENT FAIL WITH STATUSCODE: ", resp.StatusCode)
		return false, fmt.Errorf("ERROR")
	}
}

func (c *GitClient) getFilesByFiletype(filetype GitFileType) (f []*github.RepositoryContent, found bool, err error) {
	opts := &github.RepositoryContentGetOptions{
		Ref: "main",
	}

	_, files, resp, _ := c.Client.Repositories.GetContents(context.TODO(), c.Owner, *c.Repository.Name, string(filetype), opts)

	switch resp.StatusCode {
	case 404:
		return nil, false, nil
	case 200:
		return files, true, nil
	default:
		return nil, false, fmt.Errorf("ERROR")
	}
}

func (c *GitClient) GetFilesByPath(path string) []string {
	file, files, resp, _ := c.Client.Repositories.GetContents(context.TODO(), c.Owner, *c.Repository.Name, path, &github.RepositoryContentGetOptions{Ref: "main"})

	fmt.Println("TEST GET GITHUB FILE STATUS:", resp.StatusCode)

	f := []string{}

	if file.GetType() != "type" {
		content, _ := file.GetContent()
		f = append(f, content)
		return f
	}

	for _, v := range files {
		content, _ := v.GetContent()
		f = append(f, content)
	}

	return f
}

func (c *GitClient) CreateFilesByFiletype(username, email, namespace, filename string, content []byte, filetype GitFileType) error {
	t := github.Timestamp{}
	t.Time = time.Now()

	f := ""
	if filename != "" {
		f = "/" + filename
	}

	c.Client.Repositories.CreateFile(
		context.TODO(),
		c.Owner,
		*c.Repository.Name,
		defaultFilePathByFiletype(username, filetype)+f,
		&github.RepositoryContentFileOptions{
			Message: github.String(`create(` + string(filetype) + `): ` + filename + ` by ` + username),
			Content: content,
			Branch:  github.String("main"),
			Author: &github.CommitAuthor{
				Name:  github.String(username),
				Email: github.String(email),
				Date:  &t,
			},
		},
	)
	return nil
}

func defaultFilePathByFiletype(username string, filetype GitFileType) string {
	switch filetype {
	case CD:
		return "cd/" + username
	case Manifest:
		return "manifest/" + username
	}
	return ""
}

func (c *GitClient) UpdateFilesByFiletype(username, email, before, after string, filetype GitFileType) error {
	// files, found, err := c.getFilesByFiletype(filetype)
	// if err != nil {
	// 	return err
	// }

	// if !found {
	// 	return fmt.Errorf("FILE NOT FOUND")
	// }

	a := github.Timestamp{}
	a.Time = time.Now()

	author := &github.CommitAuthor{
		Name:  github.String(username),
		Email: github.String(email),
		Date:  &a,
	}

	// commiter := &github.CommitAuthor{
	// 	Name:  github.String(c.Owner),
	// 	Email: github.String(c.Email),
	// 	Date:  &a,
	// }

	ref, resp, _ := c.Client.Git.GetRef(context.TODO(), c.Owner, *c.Repository.Name, "refs/heads/main")
	fmt.Println("TEST GETREF STATUS:", resp.StatusCode)

	newTreeEntry := []*github.TreeEntry{}

	tree, resp, _ := c.Client.Git.GetTree(context.TODO(), c.Owner, *c.Repository.Name, *ref.Object.SHA, true)
	fmt.Println("TEST GETTREE STATUS:", resp.StatusCode)

	for _, entry := range tree.Entries {
		if *entry.Type == "blob" && strings.HasPrefix(*entry.Path, string(filetype)+"/"+username) {
			f, _, _, _ := c.Client.Repositories.GetContents(context.TODO(), c.Owner, *c.Repository.Name, *entry.Path, &github.RepositoryContentGetOptions{
				Ref: "main",
			})
			content, _ := f.GetContent()
			fmt.Println("TEST ENTRY CONTENT:", content)
			newContent := strings.ReplaceAll(content, before, after)

			newTreeEntry = append(newTreeEntry, &github.TreeEntry{
				Path:    entry.Path,
				Type:    entry.Type,
				Content: github.String(newContent),
				Mode:    github.String("100644"),
			})
		}
	}

	commit, resp, _ := c.Client.Git.GetCommit(context.TODO(), c.Owner, *c.Repository.Name, *ref.Object.SHA)
	fmt.Println("TEST GETCOMMIT STATUS:", resp.StatusCode)
	fmt.Println("TEST LATEST COMMIT MESSAGE: ", commit.GetMessage())

	newTree, resp, _ := c.Client.Git.CreateTree(context.TODO(), c.Owner, *c.Repository.Name, *commit.Tree.SHA, newTreeEntry)
	fmt.Println("TEST CREATETREE STATUS:", resp.StatusCode)

	opts := &github.CreateCommitOptions{}
	newCommit, resp, err := c.Client.Git.CreateCommit(context.TODO(), c.Owner, *c.Repository.Name, &github.Commit{
		Message: github.String(`update(` + string(filetype) + `): from ` + before + ` to ` + after + ` by ` + username),
		Tree:    newTree,
		Parents: []*github.Commit{
			commit,
		},
		Author: author,
		// Committer: commiter,
	}, opts)
	fmt.Println("TEST CREATECOMMIT STATUS:", resp.StatusCode)
	fmt.Println("TEST CREATECOMMIT ERR:", err)

	_, resp, _ = c.Client.Git.UpdateRef(context.TODO(), c.Owner, *c.Repository.Name, &github.Reference{Ref: github.String("refs/heads/main"), Object: &github.GitObject{SHA: newCommit.SHA}}, false)
	fmt.Println("TEST UPDATEREF STATUS:", resp.StatusCode)

	// opts := &github.RepositoryContentFileOptions{
	// 	Message:   github.String(`update(` + string(fileType) + `): from ` + old + ` to ` + new + ` by ` + username),
	// 	Content:   []byte(newContent),
	// 	Branch:    github.String("main"),
	// Author:    author,
	// Committer: commiter,
	// 	SHA:       file.SHA,
	// }

	// c.Client.Repositories.UpdateFile(context.TODO(), c.Owner, *c.Repository.Name, filePath, opts)

	return nil
}
