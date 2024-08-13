package docker

import (
	"context"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"io"
	"os"

	"github.com/docker/docker/api/types"
	"github.com/docker/docker/api/types/container"
	"github.com/docker/docker/api/types/image"
	"github.com/docker/docker/api/types/registry"
	"github.com/docker/docker/client"
	"github.com/docker/docker/pkg/archive"
	"github.com/docker/go-connections/nat"
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

func (c *DockerClient) RunDefaultRegistry() error {
	_, err := c.Client.ImagePull(context.TODO(), "registry:2", image.PullOptions{})

	portBindings := nat.PortMap{
		"5000/tcp": []nat.PortBinding{
			{
				HostIP:   "0.0.0.0",
				HostPort: "5050",
			},
		},
	}
	resp, err := c.Client.ContainerCreate(
		context.TODO(),
		&container.Config{
			Image: "registry:2",
			ExposedPorts: nat.PortSet{
				"5050/tcp": struct{}{},
			},
		},
		&container.HostConfig{
			PortBindings: portBindings,
		},
		nil, nil, "",
	)
	if err != nil {
		panic(err)
	}

	if err := c.Client.ContainerStart(context.TODO(), resp.ID, container.StartOptions{}); err != nil {
		panic(err)
	}

	cfg := registry.AuthConfig{
		ServerAddress: "http://localhost:5050",
	}

	fmt.Println("wait create")
	c.Client.ContainerWait(context.TODO(), resp.ID, "created")
	fmt.Println("created!")

	err = c.LoginWithEnv(cfg)

	err = c.Client.ImageTag(context.TODO(), "nginx:latest", "localhost:5050/nginx:latest")

	_, err = c.Client.ImagePull(context.TODO(), "nginx:latest", image.PullOptions{})

	tmp, _ := json.Marshal(cfg)
	_, err = c.Client.ImagePush(context.TODO(), "localhost:5050/nginx:latest", image.PushOptions{
		RegistryAuth: base64.URLEncoding.EncodeToString(tmp),
	})

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
