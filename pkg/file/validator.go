package file

import (
	"bufio"

	"k8s.io/apimachinery/pkg/util/yaml"
)

func ValidateYamlFileFormat(content string) error {
	t := bufio.Reader{}
	_, err := yaml.NewYAMLReader(&t).Read()
	if err != nil {
		return err
	}
	return nil
}
