package model

import (
	"time"
)

type Service struct {
	ID           string    `json:"id"`
	Name         string    `json:"name"`
	Status       string    `json:"status"`
	Version      string    `json:"version"`
	Image        string    `json:"image"`
	Tag          string    `json:"tag"`
	Platform     string    `json:"platform"`
	PlatformName string    `json:"platform_name"`
	Namespace    string    `json:"namespace"`
	Env          string    `json:"env"`
	CreatedAt    time.Time `json:"created_at"`
	Assignee     string    `json:"assignee"`
	Git          string    `json:"git"`
	Metrics      []Metric  `json:"metrics"`
}

type Infra struct {
	ID        string        `json:"id"`
	Name      string        `json:"name"`
	Status    string        `json:"status"`
	Version   string        `json:"version"`
	Platform  InfraPlatform `json:"platform"`
	Env       string        `json:"env"`
	CreatedAt time.Time     `json:"created_at"`
	Assignee  string        `json:"assignee"`
}

type InfraPlatform struct {
	Kind       string   `json:"kind"`
	K8sCluster *Cluster `json:"k8s_cluster"`
}

type Metric struct {
	ID          string `json:"id"`
	ServiceName string `json:"service_name"`
	MetricName  string `json:"metric_name"`
	Query       string `json:"query"`
	Source      string `json:"source"`
	Active      string `json:"active"`
}

type Cluster struct {
	Context                  string `json:"context"`
	Name                     string `json:"name"`
	User                     string `json:"user"`
	APIServerEndpoint        string `json:"apiserver_endpoint"`
	ClientKeyData            string `json:"client_key_data"`
	ClientCertificateData    string `json:"client_certificate_data"`
	CertificateAuthorityData string `json:"certificate_authority_data"`
}
