package model

type PolicyJson struct {
	Name   string   `json:"name"`
	Effect string   `json:"effect"`
	Target []string `json:"target"`
	Action []string `json:"action"`
}
