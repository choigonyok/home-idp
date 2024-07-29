package policy

type PolicyTarget interface {
	StoreTarget()
}

type PolicyTargetDeploy struct {
	Namespace []string                    `json:"namespace"`
	Resource  *PolicyTargetDeployResource `json:"resource"`
	GVKs      []string                    `json:"gvk"`
}

type PolicyTargetDeployResource struct {
	CPU    string `json:"cpu"`
	Memory string `json:"memory"`
	Disk   string `json:"disk"`
}

type PolicyTargetSecret struct {
	Path string `json:"path"`
}

func (s *PolicyTargetSecret) StoreTarget() {}
func (s *PolicyTargetDeploy) StoreTarget() {}
