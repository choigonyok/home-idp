package docker

import "strings"

type Image struct {
	Pusher  string
	Name    string
	Version string
}

func NewDockerImage(img, pusher string) *Image {
	name, version, _ := strings.Cut(img, ":")
	return &Image{
		Name:    name,
		Version: version,
		Pusher:  pusher,
	}
}
