package file

import (
	"errors"
	"io/fs"
	"log"
	"os"
	"strings"

	"github.com/choigonyok/home-idp/pkg/config"
	"gopkg.in/yaml.v2"
)

type FileReader struct {
	Filepath string
}

func NewReader(filepath string) *FileReader {
	return &FileReader{
		Filepath: filepath,
	}
}

func readFile(filepath string) ([]byte, error) {
	return os.ReadFile(strings.TrimSpace(filepath))
}

func exist(filepath string) bool {
	// We must explicitly check if the error is due to the file not existing (as opposed to a
	// permissions error).
	_, err := os.Stat(filepath)
	if err == fs.ErrNotExist {
		return false
	}
	return true
}

func ParseConfigFile(cfg config.Config, filepath string) error {
	if !exist(filepath) {
		log.Fatalf("Cannot find config file in %s", filepath)
		return errors.New("Cannot find config file")

	}
	bytes, _ := readFile(filepath)

	switch cfg.GetName() {
	case "secret-manager":
		tmp := &struct {
			Config config.Config `yaml:"secret-manager,omitempty"`
		}{
			Config: cfg,
		}
		if err := yaml.Unmarshal(bytes, tmp); err != nil {
			log.Fatalf("Invalid config file format")
			return errors.New("Invalid config file format")
		}
	case "deploy-manager":
		tmp := &struct {
			Config config.Config `yaml:"deploy-manager,omitempty"`
		}{
			Config: cfg,
		}
		if err := yaml.Unmarshal(bytes, tmp); err != nil {
			log.Fatalf("Invalid config file format")
			return errors.New("Invalid config file format")
		}
	case "rbac-manager":
		tmp := &struct {
			Config config.Config `yaml:"rbac-manager,omitempty"`
		}{
			Config: cfg,
		}
		if err := yaml.Unmarshal(bytes, tmp); err != nil {
			log.Fatalf("Invalid config file format")
			return errors.New("Invalid config file format")
		}
	case "gateway":
		tmp := &struct {
			Config config.Config `yaml:"gateway,omitempty"`
		}{
			Config: cfg,
		}
		if err := yaml.Unmarshal(bytes, tmp); err != nil {
			log.Fatalf("Invalid config file format")
			return errors.New("Invalid config file format")
		}
	}

	return nil
}
