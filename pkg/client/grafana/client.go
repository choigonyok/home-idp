package client

import (
	"net/http"
)

type Grafana struct {
	Client         *http.Client
	Host           string
	Token          string
	DataSourcesUID *map[string]string
}

func NewGrafanaFromAPIToken(host string, token string, m *map[string]string) *Grafana {
	return &Grafana{
		Client:         http.DefaultClient,
		Host:           host,
		Token:          token,
		DataSourcesUID: m,
	}
}
