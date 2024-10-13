package model

import (
	"encoding/json"
	"fmt"
	"log"
)

type Policy struct {
	Name   string        `json:"name"`
	Effect string        `json:"effect"`
	Target *PolicyTarget `json:"target"`
	Action []string      `json:"action"`
}

func parseToStruct(policy []byte) (*Policy, error) {
	if !json.Valid(policy) {
		log.Fatalf("Invalid policy format!")
		fmt.Println("ERROR")
		return nil, fmt.Errorf("Invalid policy format!")
	}

	tmp := struct {
		Policy *Policy `json:"policy"`
	}{
		Policy: &Policy{},
	}

	if err := json.Unmarshal(policy, &tmp); err != nil {
		fmt.Println("ERROR:", err)
		return nil, fmt.Errorf("Cannot unmarshal policy!")
	}

	return tmp.Policy, nil
}

func GetDefaultPolicy() *Policy {
	p := &Policy{
		Name:   "admin",
		Effect: "Allow",
		Target: &PolicyTarget{
			Deploy: &PolicyTargetDeploy{
				Namespace: []string{"*"},
				GVK:       []string{"*"},
				Resource: &PolicyTargetDeployResource{
					CPU:    "",
					Memory: "",
					Disk:   "",
				},
			},
			Secret: &PolicyTargetSecret{
				Path: []string{"*"},
			},
		},
		Action: []string{"*"},
	}

	StorePolicy(p)
	return p
}

func StorePolicy(p *Policy) error {
	// STORE TO STORAGE
	return nil
}
