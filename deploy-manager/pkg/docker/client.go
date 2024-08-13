package docker

import (
	"context"
	"fmt"
	"io"
	"os"

	"github.com/docker/docker/api/types"
	"github.com/docker/docker/api/types/image"
	"github.com/docker/docker/api/types/registry"
	"github.com/docker/docker/client"
	"github.com/docker/docker/pkg/archive"
)

type DockerClient struct {
	Client         *client.Client
	AuthCredential *string
}

func (c *DockerClient) Set(i interface{}) {
	c.Client = parseDockerClientFromInterface(i).Client
	c.AuthCredential = parseDockerClientFromInterface(i).AuthCredential
}

func parseDockerClientFromInterface(i interface{}) *DockerClient {
	client := i.(*DockerClient)
	return client
}

func (c *DockerClient) Close() error {
	return c.Client.Close()
}

func (c *DockerClient) LoginWithEnv(cfg registry.AuthConfig) error {
	ok, err := c.Client.RegistryLogin(context.TODO(), cfg)
	if ok.Status != "true" {
		return err
	}
	return nil
}

func (c *DockerClient) Build(tag, dockerfile string) error {
	// fullName := tag
	opt := types.ImageBuildOptions{
		Dockerfile: "Dockerfile",
		Tags:       []string{tag},
	}

	r, _ := archive.Generate("Dockerfile", dockerfile)

	optss := image.ListOptions{All: true}
	ss, _ := c.Client.ImageList(context.TODO(), optss)
	for _, v := range ss {
		fmt.Println("ID:", v.ID)
	}

	resp, err := c.Client.ImageBuild(context.TODO(), r, opt)
	fmt.Println("ERRORRR:", err)
	defer resp.Body.Close()

	_, err = io.Copy(os.Stdout, resp.Body)
	if err != nil {
		panic(err)
	}

	optss = image.ListOptions{All: true}
	ss, _ = c.Client.ImageList(context.TODO(), optss)
	for _, v := range ss {
		fmt.Println("ID:", v.ID)
	}
	return err
}

func (c *DockerClient) Push(tag string) error {
	resp, err := c.Client.ImagePush(context.TODO(), tag, image.PushOptions{
		All:          true,
		RegistryAuth: *c.AuthCredential,
	})
	defer resp.Close()

	_, err = io.Copy(os.Stdout, resp)
	if err != nil {
		panic(err)
	}

	return err
}

// dc.Build("achoistic98/ewat", `
// FROM node:18
// WORKDIR /app
// COPY . .
// RUN yarn install
// CMD ["node", "src/index.js"]
// EXPOSE 3500
// `)
// err := dc.Push("achoistic98/ewat")
// fmt.Println("TESHTIUHESO ERROROROROR:", err)
