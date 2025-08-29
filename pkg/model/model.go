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
	PlatformID   string    `json:"platform_id"`
	PlatformKind string    `json:"platform_kind"`
	PlatformName string    `json:"platform_name"`
	Namespace    string    `json:"namespace"`
	Env          string    `json:"env"`
	CreatedAt    time.Time `json:"created_at"`
	Assignee     string    `json:"assignee"`
	Git          string    `json:"git"`
	Metrics      []Metric  `json:"metrics"`
}

type Infra struct {
	ID        string    `json:"id"`
	Name      string    `json:"name"`
	Status    string    `json:"status"`
	Version   string    `json:"version"`
	Platform  Platform  `json:"platform"`
	Env       string    `json:"env"`
	CreatedAt time.Time `json:"created_at"`
	Assignee  string    `json:"assignee"`
}

type Platform struct {
	Kind       string  `json:"kind"`
	K8sCluster Cluster `json:"k8s_cluster"`
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

type System struct {
	ID          string        `json:"id"`
	Name        string        `json:"name"`
	Team        string        `json:"team"`
	Role        string        `json:"role"`
	Type        string        `json:"type"`
	LeaderID    string        `json:"leader_id"`
	JSON        string        `json:"json"`
	Description string        `json:"description"`
	Count       int           `json:"count"`
	Permissions *[]Permission `json:"permissions"`
}

type JSONFormat struct {
	Obj string `json:"obj"`
	Act string `json:"act"`
}

type Permission struct {
	ID string `json:"id"`
}
