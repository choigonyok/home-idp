package model

type Ingress struct {
	Name  string        `json:"name"`
	Rules []IngressRule `json:"rules"`
	Age   string        `json:"age"`
}

type IngressRule struct {
	Host    string `json:"host"`
	Path    string `json:"path"`
	Service string `json:"service"`
	Port    string `json:"port"`
}

type ConfigMap struct {
	Name      string `json:"name"`
	Namespace string `json:"namespace"`
	Files     []File `json:"files"`
	Creator   string `json:"creator"`
}

type File struct {
	Name          string   `json:"name"`
	MountServices []string `json:"mount_services"`
	Content       string   `json:"content"`
}

type Pod struct {
	Name   string `json:"name"`
	Status string `json:"status"`
	Age    string `json:"age"`
	IP     string `json:"ip"`
}

type Secret struct {
	Name string `json:"name"`
	Type string `json:"type"`
	Data int    `json:"data"`
	Age  string `json:"age"`
}

type Service struct {
	Name     string            `json:"name"`
	Type     string            `json:"type"`
	Selector map[string]string `json:"selector"`
	Age      string            `json:"age"`
	Port     []string          `json:"port"`
	IP       string            `json:"ip"`
}

type IngressEdge struct {
	Ingress          *Ingress
	ConnectResources *ConnectResources
}

type ConnectResources struct {
	Pods       []string `json:"conntect_pods"`
	Services   []string `json:"conntect_services"`
	Configmaps []string `json:"conntect_configmaps"`
	Secrets    []string `json:"conntect_secrets"`
	Ingresses  []string `json:"conntect_ingresses"`
}
