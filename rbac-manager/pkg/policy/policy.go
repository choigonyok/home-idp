package policy

type Policy struct {
	Name   string        `json:"name"`
	Effect string        `json:"effect"`
	Target *PolicyTarget `json:"target"`
	Action *PolicyAction `json:"action"`
}
