package git

import (
	"context"
	"fmt"
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

type GitManifest struct {
	Username string `json:"username"`
	Content  string `json:"content"`
	Filename string `json:"filename"`
}

type GitDockerFile struct {
	Username string `json:"username"`
	Content  string `json:"content"`
	Image    string `json:"image"`
}

func (f *GitManifest) getContent() string {
	return f.Content
}

func (f *GitManifest) getUsername() string {
	return f.Username
}

func (f *GitManifest) getType() string {
	return "manifest"
}

func (f *GitManifest) getFilename() string {
	return f.Filename
}

func (f *GitDockerFile) getContent() string {
	return f.Content
}

func (f *GitDockerFile) getUsername() string {
	return f.Username
}

func (f *GitDockerFile) getType() string {
	return "docker"
}

func (f *GitDockerFile) getFilename() string {
	return "Dockerfile." + f.Image
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
	fmt.Println("TEST PATH:", path)
	file, files, resp, _ := c.Client.Repositories.GetContents(context.TODO(), c.Owner, *c.Repository.Name, path, &github.RepositoryContentGetOptions{Ref: "main"})

	fmt.Println("TEST GET GITHUB FILE STATUS:", resp.StatusCode)

	f := []string{}

	fmt.Println("TEST FILE TYPE:", file.GetType())
	if file.GetType() == "file" {
		content, _ := file.GetContent()
		f = append(f, content)
		return f
	}

	for _, v := range files {
		f = append(f, v.GetName())
	}

	return f
}

func (c *GitClient) CreateFilesByFiletype(username, email, namespace, filename string, content []byte, filetype GitFileType) (*github.Response, error) {
	t := github.Timestamp{}
	t.Time = time.Now()

	f := ""
	if filename != "" {
		f = "/" + filename
	}

	_, resp, err := c.Client.Repositories.CreateFile(
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
	if err != nil {
		fmt.Println("ERR CREATE FILE IN GITHUB: ", err)
		return resp, err
	}
	return resp, nil

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
