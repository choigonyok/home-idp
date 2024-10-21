package kube

import (
	"fmt"
	"strconv"
	"time"

	"github.com/choigonyok/home-idp/pkg/model"
)

func (c *GatewayKubeClient) GetIngresses(namespace string) *[]model.Ingress {
	ingress, err := c.Client.GetIngresses(namespace)
	if err != nil {
		fmt.Println("TEST GET INGRESSES FOR NAMESPACE "+namespace+" ERR:", err)
		return nil
	}

	ingresses := []model.Ingress{}

	for _, i := range *ingress {
		rules := []model.IngressRule{}
		for _, r := range i.Spec.Rules {
			rules = append(rules, model.IngressRule{
				Host:    r.Host,
				Path:    r.HTTP.Paths[0].Path,
				Service: r.HTTP.Paths[0].Backend.Service.Name,
				Port:    string(r.HTTP.Paths[0].Backend.Service.Port.Number),
			})
		}

		age := ""
		t := time.Since(i.GetCreationTimestamp().Time)
		if t.Hours() > 24 {
			age = strconv.Itoa(int(t.Hours()/24)) + "d"
		} else if int(60*t.Hours())+int(t.Minutes())%60 > 120 {
			age = strconv.Itoa(int(t.Hours())) + "h"
		} else if int(60*t.Hours())+int(t.Minutes())%60 < 10 {
			age = strconv.Itoa(int(t.Minutes())) + "m" + strconv.Itoa(int(t.Seconds())%60) + "s"
		} else {
			age = strconv.Itoa(int(60*t.Hours()+t.Minutes())) + "m"
		}

		ingresses = append(ingresses, model.Ingress{
			Name:  i.Name,
			Rules: rules,
			Age:   age,
		})
	}

	return &ingresses
}
