package object

import (
	"fmt"
	"strings"

	"gopkg.in/yaml.v3"
	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	"k8s.io/apimachinery/pkg/runtime/schema"
)

func ParseObjectsFromManifest(manifest string) (schema.GroupVersionKind, unstructured.Unstructured) {
	manifest = strings.ReplaceAll(manifest, "\t", "  ")

	fmt.Println("TEST MANIFEST:", manifest)
	objects := unstructured.Unstructured{}

	err := yaml.Unmarshal([]byte(manifest), &objects.Object)
	fmt.Println("TEST YAML UNMARSHAL ERR:", err)

	gvk := objects.GetObjectKind().GroupVersionKind()
	fmt.Println("GROUP:", gvk.Group)
	fmt.Println("VERSION:", gvk.Version)
	fmt.Println("KIND:", gvk.Kind)

	for i, annotation := range objects.GetAnnotations() {
		fmt.Println("ANNOTATION", i, ":", annotation)
	}

	return gvk, objects
}
