package config

type Components int

const (
	SecretManager Components = iota
)

type StorageConfig struct {
	Type     string `yaml:"type,omitempty"`
	Host     string `yaml:"host,omitempty"`
	Username string `yaml:"username,omitempty"`
	Password string `yaml:"password,omitempty"`
	Database string `yaml:"database,omitempty"`
	Port     int    `yaml:"port,omitempty"`
}
