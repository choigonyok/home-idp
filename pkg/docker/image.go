package docker

import "strings"

type Image struct {
	Name    string
	Version string
}

func NewDockerImage(tag string) *Image {
	name, version, _ := strings.Cut(tag, ":")
	return &Image{
		Name:    name,
		Version: version,
	}
}
