package service

type GatewaySvcConfig struct {
	Port int    `yaml:"port,omitempty"`
	Type string `yaml:"type,omitempty"`
}
