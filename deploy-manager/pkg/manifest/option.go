package manifest

import "k8s.io/apimachinery/pkg/runtime/schema"

type Manifest struct {
	GVK      schema.GroupVersionKind
	Metadata *ManifestMetadata
	Spec     *ManifestSpec
}

type ManifestMetadata struct {
}

type ManifestSpec struct {
	Replicas        int
	Containers      *ManifestMetadataContainers
	PodAntiAffinity bool
}

type ManifestMetadataContainers struct {
	Name      string
	Image     string
	Ports     *ManifestMetadataContainersPorts
	Resources *ManifestMetadataContainersResource
}

type ManifestMetadataContainersResource struct {
	CPU    string
	Memory string
}

type ManifestMetadataContainersPorts struct {
	ContainerPort int
	Protocol      string
	Port          int
	NodePort      int
}
