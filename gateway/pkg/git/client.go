package git

import (
	"fmt"
	"regexp"
	"strings"

	"github.com/choigonyok/home-idp/gateway/pkg/manifest"
	"github.com/choigonyok/home-idp/pkg/env"
	"github.com/choigonyok/home-idp/pkg/git"
)

type GatewayGitClient struct {
	Client *git.GitClient
}

func (c *GatewayGitClient) Set(i interface{}) {
	c.Client = parseGitClientFromInterface(i)
}

func parseGitClientFromInterface(i interface{}) *git.GitClient {
	client := i.(*git.GitClient)
	return client
}

func (c *GatewayGitClient) CreateAdminDir() error {
	c.Client.CreateFilesByFiletype(env.Get("HOME_IDP_GIT_USERNAME"), env.Get("HOME_IDP_GIT_EMAIL"), env.Get("HOME_IDP_NAMESPACE"), ". gitkeep", []byte(""), git.Manifest)
	return nil
}

func (c *GatewayGitClient) CreateGithubWebhook() error {
	if env.Get("HOME_IDP_TLS_ENABLED") == "true" {
		return c.Client.CreateGitWebhook("https://" + env.Get("HOME_IDP_HOST") + "/webhooks/github")
	} else {
		return c.Client.CreateGitWebhook("http://" + env.Get("HOME_IDP_HOST") + "/webhooks/github")
	}
}

// func (c *GatewayGitClient) CreateManifest(username, email string) error {
// 	return c.Client.CreateFilesByFiletype(username, email, env.Get("HOME_IDP_NAMESPACE"), git.Manifest)
// }

// /github-webhook/"
func (c *GatewayGitClient) GetRepositoryCloneURL() string {
	return c.Client.GetRepositoryCloneURL()
}

func (c *GatewayGitClient) CreateDockerFile(username, image, content string) error {
	return c.Client.CreateFile(
		&git.GitDockerFile{
			Username: username,
			Content:  content,
			Image:    image,
		},
	)
}

func (c *GatewayGitClient) UpdateDockerFile(username, image, content string) error {
	fmt.Println("TEST START UPDATE FILE!")
	fmt.Println("TEST START UPDATE FILE!")
	fmt.Println("TEST START UPDATE FILE!")
	files := c.Client.GetFilesByPath("docker/" + username)
	imageName, _, _ := strings.Cut(image, ":")

	m := make(map[git.GitFile]git.GitFile)
	fmt.Println("TEST UPDATE IMAGENAME:", imageName)
	for _, f := range files {
		if strings.Contains(f, "Dockerfile."+imageName+":") {
			img, _ := strings.CutPrefix(f, "Dockerfile.")
			fmt.Println("TEST UPDATE IMG:", img)
			m[&git.GitDockerFile{
				Username: username,
				Content:  content,
				Image:    image,
			}] = &git.GitDockerFile{
				Username: username,
				Content:  "",
				Image:    img,
			}
		}
	}
	return c.Client.UpdateFile(m)
}

func (c *GatewayGitClient) UpdateManifest(username, content, image string) error {
	files := c.Client.Listfile("manifest", username)
	imageName, _, _ := strings.Cut(image, ":")
	r := regexp.MustCompile(imageName + `:\s*[\S]+`)

	m := make(map[git.GitFile]git.GitFile)

	for _, f := range files {

		contents := c.Client.GetFilesByPath("manifest/" + username + "/" + f)
		if strings.Contains(contents[0], "image: "+imageName+":") {
			output := r.ReplaceAllString(contents[0], "image: "+image)
			oldName, _ := strings.CutPrefix(r.FindString(contents[0]), "image: ")
			m[&git.GitDockerFile{
				Username: username,
				Content:  output,
				Image:    image,
			}] = &git.GitDockerFile{
				Username: username,
				Content:  contents[0],
				Image:    oldName,
			}
		}
	}
	return c.Client.UpdateFile(m)
}

func (c *GatewayGitClient) CreatePodManifestFile(username, email, image string, port int) error {
	name, _, _ := strings.Cut(image, ":")
	manifest := manifest.GetPodManifest(name+"-"+username, image, port)
	return c.Client.CreateFilesByFiletype(username, email, env.Get("HOME_IDP_NAMESPACE"), "pod.yaml", []byte(manifest), git.Manifest)
}

func (c *GatewayGitClient) IsDockerfileExist(username, imagename string) bool {
	files := c.Client.GetFilesByPath("docker/" + username)
	fmt.Println("FOUND FILES:", files)
	for _, f := range files {
		img, _ := strings.CutPrefix(f, "Dockerfile.")
		fmt.Println("IMGS:", img)
		imgname, _, _ := strings.Cut(img, ":")
		fmt.Println("IMGNAMES:", imgname)
		if imgname == imagename {
			return true
		}
	}
	return false
}

// curl -X POST localhost:5106/api/dockerfile -H "Content-Type: application/json" \
// -d'{
// 	"username": "tester123",
// 	"tag": "testnginx:v1.12.5",
// 	"content": "FROM ubuntu:222.04",
// }'

// 파일의 제목도 함께 바뀌는 경우 : 파일을 삭제하고 생성해야한다.
// 파일의 내용만 바뀌는 경우 : UpdateFile 메소드 지원
