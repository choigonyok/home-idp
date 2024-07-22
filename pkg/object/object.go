package object

import (
	"fmt"

	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	"k8s.io/apimachinery/pkg/runtime/schema"
	"k8s.io/apimachinery/pkg/util/yaml"
)

func ParseObjectsFromManifest(manifest string) (schema.GroupVersionKind, *unstructured.Unstructured) {
	var objects unstructured.Unstructured

	json, _ := yaml.ToJSON([]byte(manifest))
	objects.UnmarshalJSON(json)
	gvk := objects.GetObjectKind().GroupVersionKind()
	fmt.Println("GROUP:", gvk.Group)
	fmt.Println("VERSION:", gvk.Version)
	fmt.Println("KIND:", gvk.Kind)

	for i, annotation := range objects.GetAnnotations() {
		fmt.Println("ANNOTATION", i, ":", annotation)
	}

	return gvk, &objects
}
