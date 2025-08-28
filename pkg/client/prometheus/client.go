package client

import (
	"log"

	"github.com/prometheus/client_golang/api"
	"github.com/spf13/viper"
)

type Prometheus struct {
	Client *api.Client
}

func NewPrometheusClient(v *viper.Viper) *Prometheus {
	client, err := api.NewClient(api.Config{
		Address: v.GetString("prometheus.host"),
	})
	if err != nil {
		log.Fatalf("Error creating client: %v", err)
	}

	return &Prometheus{
		Client: &client,
	}
}
