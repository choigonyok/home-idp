package git

import (
	"encoding/json"
	"fmt"
	"regexp"
	"strings"

	"github.com/choigonyok/home-idp/gateway/pkg/manifest"
	"github.com/choigonyok/home-idp/pkg/env"
	"github.com/choigonyok/home-idp/pkg/git"
	"github.com/choigonyok/home-idp/pkg/model"
	"github.com/google/go-github/v63/github"
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
	c.Client.CreateFilesByFiletype(env.Get("HOME_IDP_GIT_USERNAME"), env.Get("HOME_IDP_GIT_EMAIL"), env.Get("HOME_IDP_NAMESPACE"), ".gitkeep", []byte(""), git.Manifest)
	return nil
}

func (c *GatewayGitClient) CreateGithubWebhook() error {
	apiSchema := "http"
	if env.Get("HOME_IDP_API_TLS_ENABLED") == "true" {
		apiSchema = "https"
	}
	apiHost := env.Get("HOME_IDP_API_HOST")
	apiPort := env.Get("HOME_IDP_API_PORT")
	url := apiSchema + "://" + apiHost + ":" + apiPort
	fmt.Println("TEST WEBHOOK URL : ", url)

	return c.Client.CreateGitWebhook(url + "/webhooks/github")
}

// func (c *GatewayGitClient) CreateManifest(username, email string) error {
// 	return c.Client.CreateFilesByFiletype(username, email, env.Get("HOME_IDP_NAMESPACE"), git.Manifest)
// }

// /github-webhook/"
func (c *GatewayGitClient) GetRepositoryCloneURL() string {
	return c.Client.GetRepositoryCloneURL()
}

func (c *GatewayGitClient) CreateDockerFile(username, project, image, content string) error {
	return c.Client.CreateFile(
		&git.GitDockerFile{
			Username: username,
			Content:  content,
			Image:    image,
		},
		project,
	)
}

func (c *GatewayGitClient) GetDockerFiles(projectName string) []byte {
	fmt.Println("PROJECTNAME:", projectName)

	users := c.Client.GetFilesByPath("docker")
	fmt.Println("USERS:", users)

	datas := []struct {
		Creator    string `json:"creator"`
		Name       string `json:"name"`
		Version    string `json:"version"`
		Repository string `json:"repository"`
	}{}

	for _, user := range users {
		files := c.Client.GetFilesByPath("docker/" + user)
		for _, f := range files {
			content := c.Client.GetFilesByPath("docker/" + user + "/" + f)
			fmt.Println("TEST DOCKERFILE CONTENTS:", content[0])

			img, _ := strings.CutPrefix(f, "Dockerfile.")
			imgName, imgVersion, _ := strings.Cut(img, ":")

			datas = append(datas, struct {
				Creator    string `json:"creator"`
				Name       string `json:"name"`
				Version    string `json:"version"`
				Repository string `json:"repository"`
			}{
				Creator: user,
				Name:    imgName,
				Version: imgVersion,
			})
		}
	}

	// users := c.Client.GetFilesByPath("docker")
	// for _, user := range users {
	// 	files := c.Client.GetFilesByPath("docker/" + user)
	// 	data[user] = make(map[string]string)

	// 	for _, file := range files {
	// 		content := c.Client.GetFilesByPath("docker/" + user + "/" + file)
	// 		fmt.Println("TEST DOCKERFILE CONTENTS:", content)
	// 		fmt.Println("TEST DOCKERFILE CONTENT:", content[0])
	// 		data[user][file] = content[0]
	// 	}
	// }

	b, err := json.Marshal(datas)
	if err != nil {
		fmt.Println("TEST MARSHALED DOCKERFILES ERR:", err)
	}

	fmt.Println("TEST RECIEVED DOCKERFILES DATA:", string(b))
	return b
}

func (c *GatewayGitClient) UpdateDockerFile(username, project, image, content string) error {
	files := c.Client.GetFilesByPath("docker/" + project + "/" + username)
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

func (c *GatewayGitClient) UpdateManifest(image string) error {
	users := c.Client.GetFilesByPath("manifest")
	fmt.Println("TEST USER LIST :", users)
	for _, u := range users {
		fmt.Println("TEST UPDATE MANIFEST USER :", u)
		files := c.Client.Listfile("manifest", u)
		imageName, _, _ := strings.Cut(image, ":")
		r := regexp.MustCompile(imageName + `:\s*[\S]+`)

		m := make(map[git.GitFile]git.GitFile)

		for _, f := range files {
			fmt.Println("TEST FILENAME :", f)
			contents := c.Client.GetFilesByPath("manifest/" + u + "/" + f)
			if strings.Contains(contents[0], imageName+":") {
				output := r.ReplaceAllString(contents[0], image)
				fmt.Println("TEST OUTPUT:", output)
				oldName := r.FindString(contents[0])
				fmt.Println("TEST OLDNAME:", oldName)
				m[&git.GitManifest{
					Username: u,
					Content:  output,
					Filename: f,
				}] = &git.GitManifest{
					Username: u,
					Content:  contents[0],
					Filename: f,
				}
			}
		}

		fmt.Println("TEST FINAL GIT TREE MAP TO MODIFY :", m)
		fmt.Println("TEST FINAL GIT TREE LENGTH TO MODIFY :", len(m))
		if len(m) == 0 {
			fmt.Println("TEST NO MANIFEST TO UPDATE NEW CONTAINER IMAGE")
			return nil
		}

		if err := c.Client.UpdateFile(m); err != nil {
			return err
		}
	}
	return nil
}

func (c *GatewayGitClient) CreatePodManifestFile(username, namespace, email, imageName, imageVersion string, port int, e []*model.EnvVar, f []*model.File) (*github.Response, error) {
	manifest := manifest.GetPodManifest(imageName+"-"+username, namespace, imageName+":"+imageVersion, port, e, f)
	return c.Client.CreateFilesByFiletype(username, email, env.Get("HOME_IDP_NAMESPACE"), "pod-"+imageName+".yaml", []byte(manifest), git.Manifest)
}

func (c *GatewayGitClient) IsDockerfileExist(imagename string) bool {
	projects := c.Client.GetFilesByPath("docker")

	for _, p := range projects {
		users := c.Client.GetFilesByPath("docker/" + p)
		for _, u := range users {
			files := c.Client.GetFilesByPath("docker/" + p + "/" + u)
			for _, f := range files {
				img, _ := strings.CutPrefix(f, "Dockerfile.")
				imgname, _, _ := strings.Cut(img, ":")
				fmt.Println("IMGNAMES:", imgname)
				if imgname == imagename {
					return true
				}
			}
		}
	}
	return false
}
