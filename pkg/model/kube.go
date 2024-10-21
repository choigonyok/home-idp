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

type Configmap struct {
	Name string `json:"name"`
	Data int    `json:"data"`
	Age  string `json:"age"`
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
