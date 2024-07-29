package policy

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

func ParseToStruct(policy []byte) (*Policy, error) {
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
