package kube

import (
	"fmt"
	"strconv"
	"time"
)

type Secret struct {
	Name string `json:"name"`
	Type string `json:"type"`
	Data int    `json:"data"`
	Age  string `json:"age"`
}

func (c *GatewayKubeClient) GetSecrets(namespace string) *[]Secret {
	secret, err := c.Client.GetSecrets(namespace)
	if err != nil {
		fmt.Println("TEST GET SECRETS FOR NAMESPACE "+namespace+" ERR:", err)
		return nil
	}

	secrets := []Secret{}

	for _, s := range *secret {
		age := ""
		t := time.Since(s.GetCreationTimestamp().Time)
		if t.Hours() > 24 {
			age = strconv.Itoa(int(t.Hours()/24)) + "d"
		} else if int(60*t.Hours())+int(t.Minutes())%60 > 120 {
			age = strconv.Itoa(int(t.Hours())) + "h"
		} else if int(60*t.Hours())+int(t.Minutes())%60 < 10 {
			age = strconv.Itoa(int(t.Minutes())) + "m" + strconv.Itoa(int(t.Seconds())%60) + "s"
		} else {
			age = strconv.Itoa(int(60*t.Hours()+t.Minutes())) + "m"
		}

		secrets = append(secrets, Secret{
			Name: s.Name,
			Type: string(s.Type),
			Data: len(s.Data),
			Age:  age,
		})
	}

	return &secrets
}
