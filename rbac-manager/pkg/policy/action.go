package policy

type PolicyAction struct {
	Action []string `json:"action"`
}

func (a *PolicyAction) StoreAction() {}
