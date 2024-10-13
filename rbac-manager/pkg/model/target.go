package model

// type PolicyTarget interface {
// 	StoreTarget()
// }

type PolicyTarget struct {
	Deploy *PolicyTargetDeploy `json:"deploy"`
	Secret *PolicyTargetSecret `json:"secret"`
}

type PolicyTargetDeploy struct {
	Namespace []string                    `json:"namespace"`
	Resource  *PolicyTargetDeployResource `json:"resource"`
	GVK       []string                    `json:"gvk"`
}

type PolicyTargetDeployResource struct {
	CPU    string `json:"cpu"`
	Memory string `json:"memory"`
	Disk   string `json:"disk"`
}

type PolicyTargetSecret struct {
	Path []string `json:"path"`
}

func (s *PolicyTargetSecret) StoreTarget() {}
func (s *PolicyTargetDeploy) StoreTarget() {}
